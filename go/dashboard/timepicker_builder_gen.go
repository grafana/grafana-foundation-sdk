// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimePicker] = (*TimePickerBuilder)(nil)

type TimePickerBuilder struct {
	internal *TimePicker
	errors   map[string]cog.BuildErrors
}

func NewTimePickerBuilder() *TimePickerBuilder {
	resource := &TimePicker{}
	builder := &TimePickerBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TimePickerBuilder) Build() (TimePicker, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimePicker{}, err
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

// Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
func (builder *TimePickerBuilder) Collapse(collapse bool) *TimePickerBuilder {
	builder.internal.Collapse = collapse

	return builder
}

// Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
func (builder *TimePickerBuilder) Enable(enable bool) *TimePickerBuilder {
	builder.internal.Enable = enable

	return builder
}

// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
func (builder *TimePickerBuilder) TimeOptions(timeOptions []string) *TimePickerBuilder {
	builder.internal.TimeOptions = timeOptions

	return builder
}

func (builder *TimePickerBuilder) applyDefaults() {
	builder.Hidden(false)
	builder.RefreshIntervals([]string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"})
	builder.Collapse(false)
	builder.Enable(true)
	builder.TimeOptions([]string{"5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"})
}
