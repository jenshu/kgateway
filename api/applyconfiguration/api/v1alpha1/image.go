// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// ImageApplyConfiguration represents a declarative configuration of the Image type for use
// with apply.
type ImageApplyConfiguration struct {
	Registry   *string        `json:"registry,omitempty"`
	Repository *string        `json:"repository,omitempty"`
	Tag        *string        `json:"tag,omitempty"`
	Digest     *string        `json:"digest,omitempty"`
	PullPolicy *v1.PullPolicy `json:"pullPolicy,omitempty"`
}

// ImageApplyConfiguration constructs a declarative configuration of the Image type for use with
// apply.
func Image() *ImageApplyConfiguration {
	return &ImageApplyConfiguration{}
}

// WithRegistry sets the Registry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Registry field is set to the value of the last call.
func (b *ImageApplyConfiguration) WithRegistry(value string) *ImageApplyConfiguration {
	b.Registry = &value
	return b
}

// WithRepository sets the Repository field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Repository field is set to the value of the last call.
func (b *ImageApplyConfiguration) WithRepository(value string) *ImageApplyConfiguration {
	b.Repository = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *ImageApplyConfiguration) WithTag(value string) *ImageApplyConfiguration {
	b.Tag = &value
	return b
}

// WithDigest sets the Digest field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Digest field is set to the value of the last call.
func (b *ImageApplyConfiguration) WithDigest(value string) *ImageApplyConfiguration {
	b.Digest = &value
	return b
}

// WithPullPolicy sets the PullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullPolicy field is set to the value of the last call.
func (b *ImageApplyConfiguration) WithPullPolicy(value v1.PullPolicy) *ImageApplyConfiguration {
	b.PullPolicy = &value
	return b
}
