// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeInterval] = (*TimeIntervalBuilder)(nil)

type TimeIntervalBuilder struct {
	internal *TimeInterval
	errors   map[string]cog.BuildErrors
}

func NewTimeIntervalBuilder() *TimeIntervalBuilder {
	resource := &TimeInterval{}
	builder := &TimeIntervalBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TimeIntervalBuilder) Build() (TimeInterval, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TimeInterval", err)...)
	}

	if len(errs) != 0 {
		return TimeInterval{}, errs
	}

	return *builder.internal, nil
}

func (builder *TimeIntervalBuilder) Name(name string) *TimeIntervalBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *TimeIntervalBuilder) TimeIntervals(timeIntervals []cog.Builder[TimeInterval]) *TimeIntervalBuilder {
	timeIntervalsResources := make([]TimeInterval, 0, len(timeIntervals))
	for _, r1 := range timeIntervals {
		timeIntervalsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["time_intervals"] = err.(cog.BuildErrors)
			return builder
		}
		timeIntervalsResources = append(timeIntervalsResources, timeIntervalsDepth1)
	}
	builder.internal.TimeIntervals = timeIntervalsResources

	return builder
}

func (builder *TimeIntervalBuilder) applyDefaults() {
}
