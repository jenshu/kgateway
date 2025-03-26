// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	http "net/http"

	rest "k8s.io/client-go/rest"

	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	scheme "github.com/kgateway-dev/kgateway/v2/pkg/client/clientset/versioned/scheme"
)

type GatewayV1alpha1Interface interface {
	RESTClient() rest.Interface
	BackendsGetter
	DirectResponsesGetter
	GatewayExtensionsGetter
	GatewayParametersesGetter
	HTTPListenerPoliciesGetter
	RoutePoliciesGetter
}

// GatewayV1alpha1Client is used to interact with features provided by the gateway.kgateway.dev group.
type GatewayV1alpha1Client struct {
	restClient rest.Interface
}

func (c *GatewayV1alpha1Client) Backends(namespace string) BackendInterface {
	return newBackends(c, namespace)
}

func (c *GatewayV1alpha1Client) DirectResponses(namespace string) DirectResponseInterface {
	return newDirectResponses(c, namespace)
}

func (c *GatewayV1alpha1Client) GatewayExtensions(namespace string) GatewayExtensionInterface {
	return newGatewayExtensions(c, namespace)
}

func (c *GatewayV1alpha1Client) GatewayParameterses(namespace string) GatewayParametersInterface {
	return newGatewayParameterses(c, namespace)
}

func (c *GatewayV1alpha1Client) HTTPListenerPolicies(namespace string) HTTPListenerPolicyInterface {
	return newHTTPListenerPolicies(c, namespace)
}

func (c *GatewayV1alpha1Client) RoutePolicies(namespace string) RoutePolicyInterface {
	return newRoutePolicies(c, namespace)
}

// NewForConfig creates a new GatewayV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*GatewayV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new GatewayV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*GatewayV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &GatewayV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new GatewayV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GatewayV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GatewayV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *GatewayV1alpha1Client {
	return &GatewayV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := apiv1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *GatewayV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
