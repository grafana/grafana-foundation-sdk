// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[WeekdayRange] = (*WeekdayRangeBuilder)(nil)

type WeekdayRangeBuilder struct {
	internal *WeekdayRange
	errors   map[string]cog.BuildErrors
}

func NewWeekdayRangeBuilder() *WeekdayRangeBuilder {
	resource := NewWeekdayRange()
	builder := &WeekdayRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *WeekdayRangeBuilder) Build() (WeekdayRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return WeekdayRange{}, err
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
