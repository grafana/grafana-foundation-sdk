// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeRange] = (*TimeRangeBuilder)(nil)

type TimeRangeBuilder struct {
	internal *TimeRange
	errors   map[string]cog.BuildErrors
}

func NewTimeRangeBuilder() *TimeRangeBuilder {
	resource := NewTimeRange()
	builder := &TimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *TimeRangeBuilder) Build() (TimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeRange{}, err
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *TimeRangeBuilder) From(from string) *TimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *TimeRangeBuilder) To(to string) *TimeRangeBuilder {
	builder.internal.To = to

	return builder
}
