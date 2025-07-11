// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// OpenTelemetryAccessLogServiceApplyConfiguration represents a declarative configuration of the OpenTelemetryAccessLogService type for use
// with apply.
type OpenTelemetryAccessLogServiceApplyConfiguration struct {
	GrpcService          *CommonAccessLogGrpcServiceApplyConfiguration `json:"grpcService,omitempty"`
	Body                 *string                                       `json:"body,omitempty"`
	DisableBuiltinLabels *bool                                         `json:"disableBuiltinLabels,omitempty"`
	Attributes           *KeyAnyValueListApplyConfiguration            `json:"attributes,omitempty"`
}

// OpenTelemetryAccessLogServiceApplyConfiguration constructs a declarative configuration of the OpenTelemetryAccessLogService type for use with
// apply.
func OpenTelemetryAccessLogService() *OpenTelemetryAccessLogServiceApplyConfiguration {
	return &OpenTelemetryAccessLogServiceApplyConfiguration{}
}

// WithGrpcService sets the GrpcService field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GrpcService field is set to the value of the last call.
func (b *OpenTelemetryAccessLogServiceApplyConfiguration) WithGrpcService(value *CommonAccessLogGrpcServiceApplyConfiguration) *OpenTelemetryAccessLogServiceApplyConfiguration {
	b.GrpcService = value
	return b
}

// WithBody sets the Body field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Body field is set to the value of the last call.
func (b *OpenTelemetryAccessLogServiceApplyConfiguration) WithBody(value string) *OpenTelemetryAccessLogServiceApplyConfiguration {
	b.Body = &value
	return b
}

// WithDisableBuiltinLabels sets the DisableBuiltinLabels field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisableBuiltinLabels field is set to the value of the last call.
func (b *OpenTelemetryAccessLogServiceApplyConfiguration) WithDisableBuiltinLabels(value bool) *OpenTelemetryAccessLogServiceApplyConfiguration {
	b.DisableBuiltinLabels = &value
	return b
}

// WithAttributes sets the Attributes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Attributes field is set to the value of the last call.
func (b *OpenTelemetryAccessLogServiceApplyConfiguration) WithAttributes(value *KeyAnyValueListApplyConfiguration) *OpenTelemetryAccessLogServiceApplyConfiguration {
	b.Attributes = value
	return b
}
