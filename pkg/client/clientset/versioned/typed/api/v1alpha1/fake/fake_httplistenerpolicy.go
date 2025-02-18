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

	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/applyconfiguration/api/v1alpha1"
	v1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
)

// FakeHTTPListenerPolicies implements HTTPListenerPolicyInterface
type FakeHTTPListenerPolicies struct {
	Fake *FakeGatewayV1alpha1
	ns   string
}

var httplistenerpoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("httplistenerpolicies")

var httplistenerpoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("HTTPListenerPolicy")

// Get takes name of the hTTPListenerPolicy, and returns the corresponding hTTPListenerPolicy object, and an error if there is any.
func (c *FakeHTTPListenerPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HTTPListenerPolicy, err error) {
	emptyResult := &v1alpha1.HTTPListenerPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(httplistenerpoliciesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.HTTPListenerPolicy), err
}

// List takes label and field selectors, and returns the list of HTTPListenerPolicies that match those selectors.
func (c *FakeHTTPListenerPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HTTPListenerPolicyList, err error) {
	emptyResult := &v1alpha1.HTTPListenerPolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(httplistenerpoliciesResource, httplistenerpoliciesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HTTPListenerPolicyList{ListMeta: obj.(*v1alpha1.HTTPListenerPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.HTTPListenerPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPListenerPolicies.
func (c *FakeHTTPListenerPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(httplistenerpoliciesResource, c.ns, opts))

}

// Create takes the representation of a hTTPListenerPolicy and creates it.  Returns the server's representation of the hTTPListenerPolicy, and an error, if there is any.
func (c *FakeHTTPListenerPolicies) Create(ctx context.Context, hTTPListenerPolicy *v1alpha1.HTTPListenerPolicy, opts v1.CreateOptions) (result *v1alpha1.HTTPListenerPolicy, err error) {
	emptyResult := &v1alpha1.HTTPListenerPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(httplistenerpoliciesResource, c.ns, hTTPListenerPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.HTTPListenerPolicy), err
}

// Update takes the representation of a hTTPListenerPolicy and updates it. Returns the server's representation of the hTTPListenerPolicy, and an error, if there is any.
func (c *FakeHTTPListenerPolicies) Update(ctx context.Context, hTTPListenerPolicy *v1alpha1.HTTPListenerPolicy, opts v1.UpdateOptions) (result *v1alpha1.HTTPListenerPolicy, err error) {
	emptyResult := &v1alpha1.HTTPListenerPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(httplistenerpoliciesResource, c.ns, hTTPListenerPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.HTTPListenerPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHTTPListenerPolicies) UpdateStatus(ctx context.Context, hTTPListenerPolicy *v1alpha1.HTTPListenerPolicy, opts v1.UpdateOptions) (result *v1alpha1.HTTPListenerPolicy, err error) {
	emptyResult := &v1alpha1.HTTPListenerPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(httplistenerpoliciesResource, "status", c.ns, hTTPListenerPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.HTTPListenerPolicy), err
}

// Delete takes name of the hTTPListenerPolicy and deletes it. Returns an error if one occurs.
func (c *FakeHTTPListenerPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(httplistenerpoliciesResource, c.ns, name, opts), &v1alpha1.HTTPListenerPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPListenerPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(httplistenerpoliciesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HTTPListenerPolicyList{})
	return err
}

// Patch applies the patch and returns the patched hTTPListenerPolicy.
func (c *FakeHTTPListenerPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HTTPListenerPolicy, err error) {
	emptyResult := &v1alpha1.HTTPListenerPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(httplistenerpoliciesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.HTTPListenerPolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied hTTPListenerPolicy.
func (c *FakeHTTPListenerPolicies) Apply(ctx context.Context, hTTPListenerPolicy *apiv1alpha1.HTTPListenerPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HTTPListenerPolicy, err error) {
	if hTTPListenerPolicy == nil {
		return nil, fmt.Errorf("hTTPListenerPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(hTTPListenerPolicy)
	if err != nil {
		return nil, err
	}
	name := hTTPListenerPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("hTTPListenerPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.HTTPListenerPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(httplistenerpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.HTTPListenerPolicy), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeHTTPListenerPolicies) ApplyStatus(ctx context.Context, hTTPListenerPolicy *apiv1alpha1.HTTPListenerPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HTTPListenerPolicy, err error) {
	if hTTPListenerPolicy == nil {
		return nil, fmt.Errorf("hTTPListenerPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(hTTPListenerPolicy)
	if err != nil {
		return nil, err
	}
	name := hTTPListenerPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("hTTPListenerPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.HTTPListenerPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(httplistenerpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.HTTPListenerPolicy), err
}
