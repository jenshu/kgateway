// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kgateway-dev/kgateway/v2/pkg/client/clientset/versioned/typed/api/v1alpha1"
)

type FakeGatewayV1alpha1 struct {
	*testing.Fake
}

func (c *FakeGatewayV1alpha1) Backends(namespace string) v1alpha1.BackendInterface {
	return newFakeBackends(c, namespace)
}

func (c *FakeGatewayV1alpha1) BackendConfigPolicies(namespace string) v1alpha1.BackendConfigPolicyInterface {
	return newFakeBackendConfigPolicies(c, namespace)
}

func (c *FakeGatewayV1alpha1) DirectResponses(namespace string) v1alpha1.DirectResponseInterface {
	return newFakeDirectResponses(c, namespace)
}

func (c *FakeGatewayV1alpha1) GatewayExtensions(namespace string) v1alpha1.GatewayExtensionInterface {
	return newFakeGatewayExtensions(c, namespace)
}

func (c *FakeGatewayV1alpha1) GatewayParameterses(namespace string) v1alpha1.GatewayParametersInterface {
	return newFakeGatewayParameterses(c, namespace)
}

func (c *FakeGatewayV1alpha1) HTTPListenerPolicies(namespace string) v1alpha1.HTTPListenerPolicyInterface {
	return newFakeHTTPListenerPolicies(c, namespace)
}

func (c *FakeGatewayV1alpha1) TrafficPolicies(namespace string) v1alpha1.TrafficPolicyInterface {
	return newFakeTrafficPolicies(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeGatewayV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
