// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
)

// Http1ProtocolOptionsApplyConfiguration represents a declarative configuration of the Http1ProtocolOptions type for use
// with apply.
type Http1ProtocolOptionsApplyConfiguration struct {
	EnableTrailers                          *bool                     `json:"enableTrailers,omitempty"`
	HeaderFormat                            *apiv1alpha1.HeaderFormat `json:"headerFormat,omitempty"`
	OverrideStreamErrorOnInvalidHttpMessage *bool                     `json:"overrideStreamErrorOnInvalidHttpMessage,omitempty"`
}

// Http1ProtocolOptionsApplyConfiguration constructs a declarative configuration of the Http1ProtocolOptions type for use with
// apply.
func Http1ProtocolOptions() *Http1ProtocolOptionsApplyConfiguration {
	return &Http1ProtocolOptionsApplyConfiguration{}
}

// WithEnableTrailers sets the EnableTrailers field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableTrailers field is set to the value of the last call.
func (b *Http1ProtocolOptionsApplyConfiguration) WithEnableTrailers(value bool) *Http1ProtocolOptionsApplyConfiguration {
	b.EnableTrailers = &value
	return b
}

// WithHeaderFormat sets the HeaderFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeaderFormat field is set to the value of the last call.
func (b *Http1ProtocolOptionsApplyConfiguration) WithHeaderFormat(value apiv1alpha1.HeaderFormat) *Http1ProtocolOptionsApplyConfiguration {
	b.HeaderFormat = &value
	return b
}

// WithOverrideStreamErrorOnInvalidHttpMessage sets the OverrideStreamErrorOnInvalidHttpMessage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OverrideStreamErrorOnInvalidHttpMessage field is set to the value of the last call.
func (b *Http1ProtocolOptionsApplyConfiguration) WithOverrideStreamErrorOnInvalidHttpMessage(value bool) *Http1ProtocolOptionsApplyConfiguration {
	b.OverrideStreamErrorOnInvalidHttpMessage = &value
	return b
}
