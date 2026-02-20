// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MonthRange] = (*MonthRangeBuilder)(nil)

type MonthRangeBuilder struct {
	internal *MonthRange
	errors   cog.BuildErrors
}

func NewMonthRangeBuilder() *MonthRangeBuilder {
	resource := NewMonthRange()
	builder := &MonthRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MonthRangeBuilder) Build() (MonthRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return MonthRange{}, err
	}

	if len(builder.errors) > 0 {
		return MonthRange{}, cog.MakeBuildErrors("alerting.monthRange", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MonthRangeBuilder) RecordError(path string, err error) *MonthRangeBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MonthRangeBuilder) Begin(begin int32) *MonthRangeBuilder {
	builder.internal.Begin = &begin

	return builder
}

func (builder *MonthRangeBuilder) End(end int32) *MonthRangeBuilder {
	builder.internal.End = &end

	return builder
}
