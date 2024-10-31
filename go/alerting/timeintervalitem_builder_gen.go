// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeIntervalItem] = (*TimeIntervalItemBuilder)(nil)

type TimeIntervalItemBuilder struct {
	internal *TimeIntervalItem
	errors   map[string]cog.BuildErrors
}

func NewTimeIntervalItemBuilder() *TimeIntervalItemBuilder {
	resource := &TimeIntervalItem{}
	builder := &TimeIntervalItemBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TimeIntervalItemBuilder) Build() (TimeIntervalItem, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeIntervalItem{}, err
	}

	return *builder.internal, nil
}

func (builder *TimeIntervalItemBuilder) DaysOfMonth(daysOfMonth []string) *TimeIntervalItemBuilder {
	builder.internal.DaysOfMonth = daysOfMonth

	return builder
}

func (builder *TimeIntervalItemBuilder) Location(location string) *TimeIntervalItemBuilder {
	builder.internal.Location = &location

	return builder
}

func (builder *TimeIntervalItemBuilder) Months(months []string) *TimeIntervalItemBuilder {
	builder.internal.Months = months

	return builder
}

func (builder *TimeIntervalItemBuilder) Times(times []cog.Builder[TimeIntervalTimeRange]) *TimeIntervalItemBuilder {
	timesResources := make([]TimeIntervalTimeRange, 0, len(times))
	for _, r1 := range times {
		timesDepth1, err := r1.Build()
		if err != nil {
			builder.errors["times"] = err.(cog.BuildErrors)
			return builder
		}
		timesResources = append(timesResources, timesDepth1)
	}
	builder.internal.Times = timesResources

	return builder
}

func (builder *TimeIntervalItemBuilder) Weekdays(weekdays []string) *TimeIntervalItemBuilder {
	builder.internal.Weekdays = weekdays

	return builder
}

func (builder *TimeIntervalItemBuilder) Years(years []string) *TimeIntervalItemBuilder {
	builder.internal.Years = years

	return builder
}

func (builder *TimeIntervalItemBuilder) applyDefaults() {
}
