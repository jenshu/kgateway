// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// DynamicForwardProxyBackendApplyConfiguration represents a declarative configuration of the DynamicForwardProxyBackend type for use
// with apply.
type DynamicForwardProxyBackendApplyConfiguration struct {
	EnableTls *bool `json:"enableTls,omitempty"`
}

// DynamicForwardProxyBackendApplyConfiguration constructs a declarative configuration of the DynamicForwardProxyBackend type for use with
// apply.
func DynamicForwardProxyBackend() *DynamicForwardProxyBackendApplyConfiguration {
	return &DynamicForwardProxyBackendApplyConfiguration{}
}

// WithEnableTls sets the EnableTls field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableTls field is set to the value of the last call.
func (b *DynamicForwardProxyBackendApplyConfiguration) WithEnableTls(value bool) *DynamicForwardProxyBackendApplyConfiguration {
	b.EnableTls = &value
	return b
}
