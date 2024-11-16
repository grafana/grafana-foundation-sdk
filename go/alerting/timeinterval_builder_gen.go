// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeInterval] = (*TimeIntervalBuilder)(nil)

// TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
// within the interval.
type TimeIntervalBuilder struct {
	internal *TimeInterval
	errors   map[string]cog.BuildErrors
}

func NewTimeIntervalBuilder() *TimeIntervalBuilder {
	resource := NewTimeInterval()
	builder := &TimeIntervalBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *TimeIntervalBuilder) Build() (TimeInterval, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeInterval{}, err
	}

	return *builder.internal, nil
}

func (builder *TimeIntervalBuilder) DaysOfMonth(daysOfMonth []string) *TimeIntervalBuilder {
	builder.internal.DaysOfMonth = daysOfMonth

	return builder
}

func (builder *TimeIntervalBuilder) Location(location string) *TimeIntervalBuilder {
	builder.internal.Location = &location

	return builder
}

func (builder *TimeIntervalBuilder) Months(months []string) *TimeIntervalBuilder {
	builder.internal.Months = months

	return builder
}

func (builder *TimeIntervalBuilder) Times(times []cog.Builder[TimeRange]) *TimeIntervalBuilder {
	timesResources := make([]TimeRange, 0, len(times))
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

func (builder *TimeIntervalBuilder) Weekdays(weekdays []string) *TimeIntervalBuilder {
	builder.internal.Weekdays = weekdays

	return builder
}

func (builder *TimeIntervalBuilder) Years(years []string) *TimeIntervalBuilder {
	builder.internal.Years = years

	return builder
}
