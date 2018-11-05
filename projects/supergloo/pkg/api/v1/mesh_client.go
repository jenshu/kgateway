package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type MeshClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Mesh, error)
	Write(resource *Mesh, opts clients.WriteOpts) (*Mesh, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (MeshList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan MeshList, <-chan error, error)
}

type meshClient struct {
	rc clients.ResourceClient
}

func NewMeshClient(rcFactory factory.ResourceClientFactory) (MeshClient, error) {
	return NewMeshClientWithToken(rcFactory, "")
}

func NewMeshClientWithToken(rcFactory factory.ResourceClientFactory, token string) (MeshClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Mesh{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Mesh resource client")
	}
	return &meshClient{
		rc: rc,
	}, nil
}

func (client *meshClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *meshClient) Register() error {
	return client.rc.Register()
}

func (client *meshClient) Read(namespace, name string, opts clients.ReadOpts) (*Mesh, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Mesh), nil
}

func (client *meshClient) Write(mesh *Mesh, opts clients.WriteOpts) (*Mesh, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(mesh, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Mesh), nil
}

func (client *meshClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *meshClient) List(namespace string, opts clients.ListOpts) (MeshList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToMesh(resourceList), nil
}

func (client *meshClient) Watch(namespace string, opts clients.WatchOpts) (<-chan MeshList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	meshesChan := make(chan MeshList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				meshesChan <- convertToMesh(resourceList)
			case <-opts.Ctx.Done():
				close(meshesChan)
				return
			}
		}
	}()
	return meshesChan, errs, nil
}

func convertToMesh(resources resources.ResourceList) MeshList {
	var meshList MeshList
	for _, resource := range resources {
		meshList = append(meshList, resource.(*Mesh))
	}
	return meshList
}
