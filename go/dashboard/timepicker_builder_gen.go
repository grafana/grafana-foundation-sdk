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
	builder.internal.Hidden = &hidden

	return builder
}

// Interval options available in the refresh picker dropdown.
func (builder *TimePickerBuilder) RefreshIntervals(refreshIntervals []string) *TimePickerBuilder {
	builder.internal.RefreshIntervals = refreshIntervals

	return builder
}

// Quick ranges for time picker.
func (builder *TimePickerBuilder) QuickRanges(quickRanges []cog.Builder[TimeOption]) *TimePickerBuilder {
	quickRangesResources := make([]TimeOption, 0, len(quickRanges))
	for _, r1 := range quickRanges {
		quickRangesDepth1, err := r1.Build()
		if err != nil {
			builder.errors["quick_ranges"] = err.(cog.BuildErrors)
			return builder
		}
		quickRangesResources = append(quickRangesResources, quickRangesDepth1)
	}
	builder.internal.QuickRanges = quickRangesResources

	return builder
}

// Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
func (builder *TimePickerBuilder) NowDelay(nowDelay string) *TimePickerBuilder {
	builder.internal.NowDelay = &nowDelay

	return builder
}
