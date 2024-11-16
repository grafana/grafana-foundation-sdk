// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimePickerConfig] = (*TimePickerBuilder)(nil)

// Time picker configuration
// It defines the default config for the time picker and the refresh picker for the specific dashboard.
type TimePickerBuilder struct {
	internal *TimePickerConfig
	errors   map[string]cog.BuildErrors
}

func NewTimePickerBuilder() *TimePickerBuilder {
	resource := NewTimePickerConfig()
	builder := &TimePickerBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *TimePickerBuilder) Build() (TimePickerConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimePickerConfig{}, err
	}

	return *builder.internal, nil
}

// Whether timepicker is visible or not.
func (builder *TimePickerBuilder) Hidden(hidden bool) *TimePickerBuilder {
	builder.internal.Hidden = hidden

	return builder
}

// Interval options available in the refresh picker dropdown.
func (builder *TimePickerBuilder) RefreshIntervals(refreshIntervals []string) *TimePickerBuilder {
	builder.internal.RefreshIntervals = refreshIntervals

	return builder
}

// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
func (builder *TimePickerBuilder) TimeOptions(timeOptions []string) *TimePickerBuilder {
	builder.internal.TimeOptions = timeOptions

	return builder
}
