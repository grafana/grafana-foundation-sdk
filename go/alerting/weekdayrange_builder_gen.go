// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[WeekdayRange] = (*WeekdayRangeBuilder)(nil)

type WeekdayRangeBuilder struct {
	internal *WeekdayRange
	errors   cog.BuildErrors
}

func NewWeekdayRangeBuilder() *WeekdayRangeBuilder {
	resource := NewWeekdayRange()
	builder := &WeekdayRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *WeekdayRangeBuilder) Build() (WeekdayRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return WeekdayRange{}, err
	}

	if len(builder.errors) > 0 {
		return WeekdayRange{}, cog.MakeBuildErrors("alerting.weekdayRange", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *WeekdayRangeBuilder) Begin(begin int32) *WeekdayRangeBuilder {
	builder.internal.Begin = &begin

	return builder
}

func (builder *WeekdayRangeBuilder) End(end int32) *WeekdayRangeBuilder {
	builder.internal.End = &end

	return builder
}
