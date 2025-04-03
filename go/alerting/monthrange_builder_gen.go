// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MonthRange] = (*MonthRangeBuilder)(nil)

type MonthRangeBuilder struct {
	internal *MonthRange
	errors   map[string]cog.BuildErrors
}

func NewMonthRangeBuilder() *MonthRangeBuilder {
	resource := NewMonthRange()
	builder := &MonthRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *MonthRangeBuilder) Build() (MonthRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return MonthRange{}, err
	}

	return *builder.internal, nil
}

func (builder *MonthRangeBuilder) Begin(begin int32) *MonthRangeBuilder {
	builder.internal.Begin = &begin

	return builder
}

func (builder *MonthRangeBuilder) End(end int32) *MonthRangeBuilder {
	builder.internal.End = &end

	return builder
}
