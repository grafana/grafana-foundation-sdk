// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeSettingsSpec] = (*TimeSettingsBuilder)(nil)

// Time configuration
// It defines the default time config for the time picker, the refresh picker for the specific dashboard.
type TimeSettingsBuilder struct {
	internal *TimeSettingsSpec
	errors   cog.BuildErrors
}

func NewTimeSettingsBuilder() *TimeSettingsBuilder {
	resource := NewTimeSettingsSpec()
	builder := &TimeSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TimeSettingsBuilder) Build() (TimeSettingsSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeSettingsSpec{}, err
	}

	if len(builder.errors) > 0 {
		return TimeSettingsSpec{}, cog.MakeBuildErrors("dashboardv2beta1.timeSettings", builder.errors)
	}

	return *builder.internal, nil
}

// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
func (builder *TimeSettingsBuilder) Timezone(timezone string) *TimeSettingsBuilder {
	builder.internal.Timezone = &timezone

	return builder
}

// Start time range for dashboard.
// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
func (builder *TimeSettingsBuilder) From(from string) *TimeSettingsBuilder {
	builder.internal.From = from

	return builder
}

// End time range for dashboard.
// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
func (builder *TimeSettingsBuilder) To(to string) *TimeSettingsBuilder {
	builder.internal.To = to

	return builder
}

// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
// v1: refresh
func (builder *TimeSettingsBuilder) AutoRefresh(autoRefresh string) *TimeSettingsBuilder {
	builder.internal.AutoRefresh = autoRefresh

	return builder
}

// Interval options available in the refresh picker dropdown.
// v1: timepicker.refresh_intervals
func (builder *TimeSettingsBuilder) AutoRefreshIntervals(autoRefreshIntervals []string) *TimeSettingsBuilder {
	builder.internal.AutoRefreshIntervals = autoRefreshIntervals

	return builder
}

// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
// v1: timepicker.quick_ranges , not exposed in the UI
func (builder *TimeSettingsBuilder) QuickRanges(quickRanges []cog.Builder[TimeRangeOption]) *TimeSettingsBuilder {
	quickRangesResources := make([]TimeRangeOption, 0, len(quickRanges))
	for _, r1 := range quickRanges {
		quickRangesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		quickRangesResources = append(quickRangesResources, quickRangesDepth1)
	}
	builder.internal.QuickRanges = quickRangesResources

	return builder
}

// Whether timepicker is visible or not.
// v1: timepicker.hidden
func (builder *TimeSettingsBuilder) HideTimepicker(hideTimepicker bool) *TimeSettingsBuilder {
	builder.internal.HideTimepicker = hideTimepicker

	return builder
}

// Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
func (builder *TimeSettingsBuilder) WeekStart(weekStart TimeSettingsSpecWeekStart) *TimeSettingsBuilder {
	builder.internal.WeekStart = &weekStart

	return builder
}

// The month that the fiscal year starts on. 0 = January, 11 = December
func (builder *TimeSettingsBuilder) FiscalYearStartMonth(fiscalYearStartMonth int64) *TimeSettingsBuilder {
	builder.internal.FiscalYearStartMonth = fiscalYearStartMonth

	return builder
}

// Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
// v1: timepicker.nowDelay
func (builder *TimeSettingsBuilder) NowDelay(nowDelay string) *TimeSettingsBuilder {
	builder.internal.NowDelay = &nowDelay

	return builder
}
