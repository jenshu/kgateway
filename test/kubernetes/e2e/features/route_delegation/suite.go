package route_delegation

import (
	"context"
	"net/http"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"

	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	testmatchers "github.com/solo-io/gloo/test/gomega/matchers"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/defaults"
	"github.com/solo-io/gloo/test/kubernetes/e2e/utils"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
)

type tsuite struct {
	suite.Suite

	ctx context.Context
	// ti contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	ti *e2e.TestInstallation

	// maps test name to a list of manifests to apply before the test
	manifests map[string][]string

	manifestObjects map[string][]client.Object
}

func NewTestingSuite(ctx context.Context, testInst *e2e.TestInstallation) *tsuite {
	return &tsuite{
		ctx: ctx,
		ti:  testInst,
	}
}

func (s *tsuite) SetupSuite() {
	s.manifests = map[string][]string{
		"TestBasic":            {commonManifest, basicRoutesManifest},
		"TestRecursive":        {commonManifest, recursiveRoutesManifest},
		"TestCyclic":           {commonManifest, cyclicRoutesManifest},
		"TestInvalidChild":     {commonManifest, invalidChildRoutesManifest},
		"TestHeaderQueryMatch": {commonManifest, headerQueryMatchRoutesManifest},
	}
	// Not every resource that is applied needs to be verified. We are not testing `kubectl apply`,
	// but the below code demonstrates how it can be done if necessary
	s.manifestObjects = map[string][]client.Object{
		commonManifest:                 {proxyService, proxyDeployment, defaults.CurlPod, httpbinTeam1, httpbinTeam2, gateway},
		basicRoutesManifest:            {routeRoot, routeTeam1, routeTeam2},
		cyclicRoutesManifest:           {routeRoot, routeTeam1, routeTeam2},
		recursiveRoutesManifest:        {routeRoot, routeTeam1, routeTeam2},
		invalidChildRoutesManifest:     {routeRoot, routeTeam1, routeTeam2},
		headerQueryMatchRoutesManifest: {routeRoot, routeTeam1, routeTeam2},
	}
	clients, err := gloogateway.NewResourceClients(s.ctx, s.ti.TestCluster.ClusterContext)
	s.Require().NoError(err)
	s.ti.ResourceClients = clients
}

func (s *tsuite) TearDownSuite() {
	// nothing at the moment
}

func (s *tsuite) BeforeTest(suiteName, testName string) {
	manifests := s.manifests[testName]
	for _, manifest := range manifests {
		err := s.ti.Actions.Kubectl().ApplyFile(s.ctx, manifest)
		s.Require().NoError(err)
		s.ti.Assertions.EventuallyObjectsExist(s.ctx, s.manifestObjects[manifest]...)
	}
}

func (s *tsuite) AfterTest(suiteName, testName string) {
	manifests := s.manifests[testName]
	for _, manifest := range manifests {
		err := s.ti.Actions.Kubectl().DeleteFileSafe(s.ctx, manifest)
		s.Require().NoError(err)
		s.ti.Assertions.EventuallyObjectsNotExist(s.ctx, s.manifestObjects[manifest]...)
	}
}

func (s *tsuite) TestBasic() {
	// Assert traffic to team1 route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam1)},
		&testmatchers.HttpResponse{StatusCode: http.StatusOK, Body: ContainSubstring("anything")})

	// Assert traffic to team2 route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam2)},
		&testmatchers.HttpResponse{StatusCode: http.StatusOK, Body: ContainSubstring("anything")})
}

func (s *tsuite) TestRecursive() {
	// Assert traffic to team1 route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam1)},
		&testmatchers.HttpResponse{StatusCode: http.StatusOK, Body: ContainSubstring("anything")})

	// Assert traffic to team2 route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam2)},
		&testmatchers.HttpResponse{StatusCode: http.StatusOK, Body: ContainSubstring("anything")})
}

func (s *tsuite) TestCyclic() {
	// Assert traffic to team1 route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam1)},
		&testmatchers.HttpResponse{StatusCode: http.StatusOK, Body: ContainSubstring("anything")})

	// Assert traffic to team2 route fails with HTTP 404 as it is a cyclic route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam2)},
		&testmatchers.HttpResponse{StatusCode: http.StatusNotFound})

	cyclicRoute := &gwv1.HTTPRoute{}
	err := s.ti.TestCluster.ClusterContext.Client.Get(s.ctx,
		types.NamespacedName{Name: routeTeam2.Name, Namespace: routeTeam2.Namespace},
		cyclicRoute)
	s.Require().NoError(err)
	s.Require().Truef(utils.HTTPRouteStatusContainsMsg(cyclicRoute, "cyclic reference detected"), "missing status on cyclic route")
}

func (s *tsuite) TestInvalidChild() {
	// Assert traffic to team1 route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam1)},
		&testmatchers.HttpResponse{StatusCode: http.StatusOK, Body: ContainSubstring("anything")})

	// Assert traffic to team2 route fails with HTTP 404 as the route is invalid due to specifying a hostname on the child route
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt, []curl.Option{curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam2)},
		&testmatchers.HttpResponse{StatusCode: http.StatusNotFound})

	invalidRoute := &gwv1.HTTPRoute{}
	err := s.ti.TestCluster.ClusterContext.Client.Get(s.ctx,
		types.NamespacedName{Name: routeTeam2.Name, Namespace: routeTeam2.Namespace},
		invalidRoute)
	s.Require().NoError(err)
	s.Require().Truef(utils.HTTPRouteStatusContainsMsg(invalidRoute, "spec.hostnames must be unset"), "missing status on invalid route")
}

func (s *tsuite) TestHeaderQueryMatch() {
	// Assert traffic to team1 route with matching header and query parameters
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam1),
			curl.WithHeader("header1", "val1"),
			curl.WithHeader("headerX", "valX"),
			curl.WithQueryParameters(map[string]string{"query1": "val1", "queryX": "valX"}),
		},
		&testmatchers.HttpResponse{StatusCode: http.StatusOK, Body: ContainSubstring("anything")})

	// Assert traffic to team2 route fails with HTTP 404 as it does not match the parent's header and query parameters
	s.ti.Assertions.AssertEventuallyConsistentCurlResponse(s.ctx, defaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHostPort(proxyHostPort), curl.WithPath(pathTeam2),
			curl.WithHeader("headerX", "valX"),
			curl.WithQueryParameters(map[string]string{"queryX": "valX"}),
		},
		&testmatchers.HttpResponse{StatusCode: http.StatusNotFound})
}
