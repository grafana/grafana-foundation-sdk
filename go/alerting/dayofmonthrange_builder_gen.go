// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DayOfMonthRange] = (*DayOfMonthRangeBuilder)(nil)

type DayOfMonthRangeBuilder struct {
	internal *DayOfMonthRange
	errors   cog.BuildErrors
}

func NewDayOfMonthRangeBuilder() *DayOfMonthRangeBuilder {
	resource := NewDayOfMonthRange()
	builder := &DayOfMonthRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DayOfMonthRangeBuilder) Build() (DayOfMonthRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return DayOfMonthRange{}, err
	}

	if len(builder.errors) > 0 {
		return DayOfMonthRange{}, cog.MakeBuildErrors("alerting.dayOfMonthRange", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DayOfMonthRangeBuilder) Begin(begin int32) *DayOfMonthRangeBuilder {
	builder.internal.Begin = &begin

	return builder
}

func (builder *DayOfMonthRangeBuilder) End(end int32) *DayOfMonthRangeBuilder {
	builder.internal.End = &end

	return builder
}
