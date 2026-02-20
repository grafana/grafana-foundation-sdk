// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeMathTimeRange] = (*ExprTypeMathTimeRangeBuilder)(nil)

type ExprTypeMathTimeRangeBuilder struct {
	internal *ExprTypeMathTimeRange
	errors   cog.BuildErrors
}

func NewExprTypeMathTimeRangeBuilder() *ExprTypeMathTimeRangeBuilder {
	resource := NewExprTypeMathTimeRange()
	builder := &ExprTypeMathTimeRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeMathTimeRangeBuilder) Build() (ExprTypeMathTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeMathTimeRange{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeMathTimeRange{}, cog.MakeBuildErrors("expr.exprTypeMathTimeRange", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeMathTimeRangeBuilder) RecordError(path string, err error) *ExprTypeMathTimeRangeBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// From is the start time of the query.
func (builder *ExprTypeMathTimeRangeBuilder) From(from string) *ExprTypeMathTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeMathTimeRangeBuilder) To(to string) *ExprTypeMathTimeRangeBuilder {
	builder.internal.To = to

	return builder
}
