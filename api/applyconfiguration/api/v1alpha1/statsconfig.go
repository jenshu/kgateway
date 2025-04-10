// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// StatsConfigApplyConfiguration represents a declarative configuration of the StatsConfig type for use
// with apply.
type StatsConfigApplyConfiguration struct {
	Enabled                 *bool   `json:"enabled,omitempty"`
	RoutePrefixRewrite      *string `json:"routePrefixRewrite,omitempty"`
	EnableStatsRoute        *bool   `json:"enableStatsRoute,omitempty"`
	StatsRoutePrefixRewrite *string `json:"statsRoutePrefixRewrite,omitempty"`
}

// StatsConfigApplyConfiguration constructs a declarative configuration of the StatsConfig type for use with
// apply.
func StatsConfig() *StatsConfigApplyConfiguration {
	return &StatsConfigApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *StatsConfigApplyConfiguration) WithEnabled(value bool) *StatsConfigApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithRoutePrefixRewrite sets the RoutePrefixRewrite field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RoutePrefixRewrite field is set to the value of the last call.
func (b *StatsConfigApplyConfiguration) WithRoutePrefixRewrite(value string) *StatsConfigApplyConfiguration {
	b.RoutePrefixRewrite = &value
	return b
}

// WithEnableStatsRoute sets the EnableStatsRoute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableStatsRoute field is set to the value of the last call.
func (b *StatsConfigApplyConfiguration) WithEnableStatsRoute(value bool) *StatsConfigApplyConfiguration {
	b.EnableStatsRoute = &value
	return b
}

// WithStatsRoutePrefixRewrite sets the StatsRoutePrefixRewrite field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StatsRoutePrefixRewrite field is set to the value of the last call.
func (b *StatsConfigApplyConfiguration) WithStatsRoutePrefixRewrite(value string) *StatsConfigApplyConfiguration {
	b.StatsRoutePrefixRewrite = &value
	return b
}
