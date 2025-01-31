// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	apiv1alpha1 "github.com/kgateway-dev/kgateway/projects/gateway2/api/applyconfiguration/api/v1alpha1"
	v1alpha1 "github.com/kgateway-dev/kgateway/projects/gateway2/api/v1alpha1"
	scheme "github.com/kgateway-dev/kgateway/projects/gateway2/pkg/client/clientset/versioned/scheme"
)

// DirectResponsesGetter has a method to return a DirectResponseInterface.
// A group's client should implement this interface.
type DirectResponsesGetter interface {
	DirectResponses(namespace string) DirectResponseInterface
}

// DirectResponseInterface has methods to work with DirectResponse resources.
type DirectResponseInterface interface {
	Create(ctx context.Context, directResponse *v1alpha1.DirectResponse, opts v1.CreateOptions) (*v1alpha1.DirectResponse, error)
	Update(ctx context.Context, directResponse *v1alpha1.DirectResponse, opts v1.UpdateOptions) (*v1alpha1.DirectResponse, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, directResponse *v1alpha1.DirectResponse, opts v1.UpdateOptions) (*v1alpha1.DirectResponse, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DirectResponse, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DirectResponseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DirectResponse, err error)
	Apply(ctx context.Context, directResponse *apiv1alpha1.DirectResponseApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DirectResponse, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, directResponse *apiv1alpha1.DirectResponseApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DirectResponse, err error)
	DirectResponseExpansion
}

// directResponses implements DirectResponseInterface
type directResponses struct {
	*gentype.ClientWithListAndApply[*v1alpha1.DirectResponse, *v1alpha1.DirectResponseList, *apiv1alpha1.DirectResponseApplyConfiguration]
}

// newDirectResponses returns a DirectResponses
func newDirectResponses(c *GatewayV1alpha1Client, namespace string) *directResponses {
	return &directResponses{
		gentype.NewClientWithListAndApply[*v1alpha1.DirectResponse, *v1alpha1.DirectResponseList, *apiv1alpha1.DirectResponseApplyConfiguration](
			"directresponses",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.DirectResponse { return &v1alpha1.DirectResponse{} },
			func() *v1alpha1.DirectResponseList { return &v1alpha1.DirectResponseList{} }),
	}
}
