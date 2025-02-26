// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// AnthropicConfigApplyConfiguration represents a declarative configuration of the AnthropicConfig type for use
// with apply.
type AnthropicConfigApplyConfiguration struct {
	AuthToken *SingleAuthTokenApplyConfiguration `json:"authToken,omitempty"`
	Version   *string                            `json:"apiVersion,omitempty"`
	Model     *string                            `json:"model,omitempty"`
}

// AnthropicConfigApplyConfiguration constructs a declarative configuration of the AnthropicConfig type for use with
// apply.
func AnthropicConfig() *AnthropicConfigApplyConfiguration {
	return &AnthropicConfigApplyConfiguration{}
}

// WithAuthToken sets the AuthToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AuthToken field is set to the value of the last call.
func (b *AnthropicConfigApplyConfiguration) WithAuthToken(value *SingleAuthTokenApplyConfiguration) *AnthropicConfigApplyConfiguration {
	b.AuthToken = value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *AnthropicConfigApplyConfiguration) WithVersion(value string) *AnthropicConfigApplyConfiguration {
	b.Version = &value
	return b
}

// WithModel sets the Model field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Model field is set to the value of the last call.
func (b *AnthropicConfigApplyConfiguration) WithModel(value string) *AnthropicConfigApplyConfiguration {
	b.Model = &value
	return b
}
