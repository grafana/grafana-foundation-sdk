// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[YearRange] = (*YearRangeBuilder)(nil)

type YearRangeBuilder struct {
	internal *YearRange
	errors   cog.BuildErrors
}

func NewYearRangeBuilder() *YearRangeBuilder {
	resource := NewYearRange()
	builder := &YearRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *YearRangeBuilder) Build() (YearRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return YearRange{}, err
	}

	if len(builder.errors) > 0 {
		return YearRange{}, cog.MakeBuildErrors("alerting.yearRange", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *YearRangeBuilder) Begin(begin int32) *YearRangeBuilder {
	builder.internal.Begin = &begin

	return builder
}

func (builder *YearRangeBuilder) End(end int32) *YearRangeBuilder {
	builder.internal.End = &end

	return builder
}
