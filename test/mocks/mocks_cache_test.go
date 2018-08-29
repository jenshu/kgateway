package mocks

import (
	"context"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/services"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("MocksCache", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1         string
		namespace2         string
		cfg                *rest.Config
		cache              Cache
		mockResourceClient MockResourceClient
		fakeResourceClient FakeResourceClient
		mockDataClient     MockDataClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		err := services.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = services.SetupKubeForTest(namespace2)
		kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		Expect(err).NotTo(HaveOccurred())

		// MockResource Constructor
		mockResourceClientFactory := factory.NewResourceClientFactory(&factory.KubeResourceClientOpts{
			Crd: MockResourceCrd,
			Cfg: cfg,
		})
		mockResourceClient, err = NewMockResourceClient(mockResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		// FakeResource Constructor
		fakeResourceClientFactory := factory.NewResourceClientFactory(&factory.KubeResourceClientOpts{
			Crd: FakeResourceCrd,
			Cfg: cfg,
		})
		fakeResourceClient, err = NewFakeResourceClient(fakeResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		// MockData Constructor
		mockDataClientFactory := factory.NewResourceClientFactory(&factory.KubeResourceClientOpts{
			Crd: MockDataCrd,
			Cfg: cfg,
		})
		mockDataClient, err = NewMockDataClient(mockDataClientFactory)
		Expect(err).NotTo(HaveOccurred())
		cache = NewCache(mockResourceClient, fakeResourceClient, mockDataClient)
	})
	AfterEach(func() {
		services.TeardownKube(namespace1)
		services.TeardownKube(namespace2)
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := cache.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := cache.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *Snapshot

		/*
			MockResource
		*/

		assertSnapshotMocks := func(expectMocks MockResourceList, unexpectMocks MockResourceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMocks {
						if _, err := snap.Mocks.List().Find(expected.Metadata.ObjectRef()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMocks {
						if _, err := snap.Mocks.List().Find(unexpected.Metadata.ObjectRef()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := mockResourceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := mockResourceClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}

		mockResource1a, err := mockResourceClient.Write(NewMockResource(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockResource1b, err := mockResourceClient.Write(NewMockResource(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b}, nil)

		mockResource2a, err := mockResourceClient.Write(NewMockResource(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockResource2b, err := mockResourceClient.Write(NewMockResource(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b, mockResource2a, mockResource2b}, nil)

		err = mockResourceClient.Delete(mockResource2a.Metadata.Namespace, mockResource2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockResourceClient.Delete(mockResource2b.Metadata.Namespace, mockResource2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b}, MockResourceList{mockResource2a, mockResource2b})

		err = mockResourceClient.Delete(mockResource1a.Metadata.Namespace, mockResource1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockResourceClient.Delete(mockResource1b.Metadata.Namespace, mockResource1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(nil, MockResourceList{mockResource1a, mockResource1b, mockResource2a, mockResource2b})

		/*
			FakeResource
		*/

		assertSnapshotFakes := func(expectFakes FakeResourceList, unexpectFakes FakeResourceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectFakes {
						if _, err := snap.Fakes.List().Find(expected.Metadata.ObjectRef()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectFakes {
						if _, err := snap.Fakes.List().Find(unexpected.Metadata.ObjectRef()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := fakeResourceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := fakeResourceClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}

		fakeResource1a, err := fakeResourceClient.Write(NewFakeResource(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		fakeResource1b, err := fakeResourceClient.Write(NewFakeResource(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(FakeResourceList{fakeResource1a, fakeResource1b}, nil)

		fakeResource2a, err := fakeResourceClient.Write(NewFakeResource(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		fakeResource2b, err := fakeResourceClient.Write(NewFakeResource(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(FakeResourceList{fakeResource1a, fakeResource1b, fakeResource2a, fakeResource2b}, nil)

		err = fakeResourceClient.Delete(fakeResource2a.Metadata.Namespace, fakeResource2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = fakeResourceClient.Delete(fakeResource2b.Metadata.Namespace, fakeResource2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(FakeResourceList{fakeResource1a, fakeResource1b}, FakeResourceList{fakeResource2a, fakeResource2b})

		err = fakeResourceClient.Delete(fakeResource1a.Metadata.Namespace, fakeResource1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = fakeResourceClient.Delete(fakeResource1b.Metadata.Namespace, fakeResource1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(nil, FakeResourceList{fakeResource1a, fakeResource1b, fakeResource2a, fakeResource2b})

		/*
			MockData
		*/

		assertSnapshotMockDatas := func(expectMockDatas MockDataList, unexpectMockDatas MockDataList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMockDatas {
						if _, err := snap.MockDatas.List().Find(expected.Metadata.ObjectRef()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMockDatas {
						if _, err := snap.MockDatas.List().Find(unexpected.Metadata.ObjectRef()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := mockDataClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := mockDataClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}

		mockData1a, err := mockDataClient.Write(NewMockData(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockData1b, err := mockDataClient.Write(NewMockData(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMockDatas(MockDataList{mockData1a, mockData1b}, nil)

		mockData2a, err := mockDataClient.Write(NewMockData(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockData2b, err := mockDataClient.Write(NewMockData(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMockDatas(MockDataList{mockData1a, mockData1b, mockData2a, mockData2b}, nil)

		err = mockDataClient.Delete(mockData2a.Metadata.Namespace, mockData2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockDataClient.Delete(mockData2b.Metadata.Namespace, mockData2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMockDatas(MockDataList{mockData1a, mockData1b}, MockDataList{mockData2a, mockData2b})

		err = mockDataClient.Delete(mockData1a.Metadata.Namespace, mockData1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockDataClient.Delete(mockData1b.Metadata.Namespace, mockData1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMockDatas(nil, MockDataList{mockData1a, mockData1b, mockData2a, mockData2b})
	})
})
