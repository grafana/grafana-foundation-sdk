// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimePicker] = (*TimePickerBuilder)(nil)

type TimePickerBuilder struct {
	internal *TimePicker
	errors   cog.BuildErrors
}

func NewTimePickerBuilder() *TimePickerBuilder {
	resource := NewTimePicker()
	builder := &TimePickerBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TimePickerBuilder) Build() (TimePicker, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimePicker{}, err
	}

	if len(builder.errors) > 0 {
		return TimePicker{}, cog.MakeBuildErrors("dashboard.timePicker", builder.errors)
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
