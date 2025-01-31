// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	apiv1alpha1 "github.com/kgateway-dev/kgateway/projects/gateway2/api/applyconfiguration/api/v1alpha1"
	v1alpha1 "github.com/kgateway-dev/kgateway/projects/gateway2/api/v1alpha1"
)

// FakeRoutePolicies implements RoutePolicyInterface
type FakeRoutePolicies struct {
	Fake *FakeGatewayV1alpha1
	ns   string
}

var routepoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("routepolicies")

var routepoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("RoutePolicy")

// Get takes name of the routePolicy, and returns the corresponding routePolicy object, and an error if there is any.
func (c *FakeRoutePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RoutePolicy, err error) {
	emptyResult := &v1alpha1.RoutePolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(routepoliciesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RoutePolicy), err
}

// List takes label and field selectors, and returns the list of RoutePolicies that match those selectors.
func (c *FakeRoutePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RoutePolicyList, err error) {
	emptyResult := &v1alpha1.RoutePolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(routepoliciesResource, routepoliciesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RoutePolicyList{ListMeta: obj.(*v1alpha1.RoutePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.RoutePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routePolicies.
func (c *FakeRoutePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(routepoliciesResource, c.ns, opts))

}

// Create takes the representation of a routePolicy and creates it.  Returns the server's representation of the routePolicy, and an error, if there is any.
func (c *FakeRoutePolicies) Create(ctx context.Context, routePolicy *v1alpha1.RoutePolicy, opts v1.CreateOptions) (result *v1alpha1.RoutePolicy, err error) {
	emptyResult := &v1alpha1.RoutePolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(routepoliciesResource, c.ns, routePolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RoutePolicy), err
}

// Update takes the representation of a routePolicy and updates it. Returns the server's representation of the routePolicy, and an error, if there is any.
func (c *FakeRoutePolicies) Update(ctx context.Context, routePolicy *v1alpha1.RoutePolicy, opts v1.UpdateOptions) (result *v1alpha1.RoutePolicy, err error) {
	emptyResult := &v1alpha1.RoutePolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(routepoliciesResource, c.ns, routePolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RoutePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoutePolicies) UpdateStatus(ctx context.Context, routePolicy *v1alpha1.RoutePolicy, opts v1.UpdateOptions) (result *v1alpha1.RoutePolicy, err error) {
	emptyResult := &v1alpha1.RoutePolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(routepoliciesResource, "status", c.ns, routePolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RoutePolicy), err
}

// Delete takes name of the routePolicy and deletes it. Returns an error if one occurs.
func (c *FakeRoutePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(routepoliciesResource, c.ns, name, opts), &v1alpha1.RoutePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoutePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(routepoliciesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RoutePolicyList{})
	return err
}

// Patch applies the patch and returns the patched routePolicy.
func (c *FakeRoutePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RoutePolicy, err error) {
	emptyResult := &v1alpha1.RoutePolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(routepoliciesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RoutePolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied routePolicy.
func (c *FakeRoutePolicies) Apply(ctx context.Context, routePolicy *apiv1alpha1.RoutePolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.RoutePolicy, err error) {
	if routePolicy == nil {
		return nil, fmt.Errorf("routePolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(routePolicy)
	if err != nil {
		return nil, err
	}
	name := routePolicy.Name
	if name == nil {
		return nil, fmt.Errorf("routePolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.RoutePolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(routepoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RoutePolicy), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeRoutePolicies) ApplyStatus(ctx context.Context, routePolicy *apiv1alpha1.RoutePolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.RoutePolicy, err error) {
	if routePolicy == nil {
		return nil, fmt.Errorf("routePolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(routePolicy)
	if err != nil {
		return nil, err
	}
	name := routePolicy.Name
	if name == nil {
		return nil, fmt.Errorf("routePolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.RoutePolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(routepoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RoutePolicy), err
}
