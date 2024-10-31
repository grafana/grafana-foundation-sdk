// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeIntervalTimeRange] = (*TimeIntervalTimeRangeBuilder)(nil)

type TimeIntervalTimeRangeBuilder struct {
	internal *TimeIntervalTimeRange
	errors   map[string]cog.BuildErrors
}

func NewTimeIntervalTimeRangeBuilder() *TimeIntervalTimeRangeBuilder {
	resource := &TimeIntervalTimeRange{}
	builder := &TimeIntervalTimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TimeIntervalTimeRangeBuilder) Build() (TimeIntervalTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeIntervalTimeRange{}, err
	}

	return *builder.internal, nil
}

func (builder *TimeIntervalTimeRangeBuilder) EndTime(endTime string) *TimeIntervalTimeRangeBuilder {
	builder.internal.EndTime = &endTime

	return builder
}

func (builder *TimeIntervalTimeRangeBuilder) StartTime(startTime string) *TimeIntervalTimeRangeBuilder {
	builder.internal.StartTime = &startTime

	return builder
}

func (builder *TimeIntervalTimeRangeBuilder) applyDefaults() {
}
