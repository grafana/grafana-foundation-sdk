// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeReduceTimeRange] = (*ExprTypeReduceTimeRangeBuilder)(nil)

type ExprTypeReduceTimeRangeBuilder struct {
	internal *ExprTypeReduceTimeRange
	errors   cog.BuildErrors
}

func NewExprTypeReduceTimeRangeBuilder() *ExprTypeReduceTimeRangeBuilder {
	resource := NewExprTypeReduceTimeRange()
	builder := &ExprTypeReduceTimeRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeReduceTimeRangeBuilder) Build() (ExprTypeReduceTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeReduceTimeRange{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeReduceTimeRange{}, cog.MakeBuildErrors("expr.exprTypeReduceTimeRange", builder.errors)
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeReduceTimeRangeBuilder) From(from string) *ExprTypeReduceTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeReduceTimeRangeBuilder) To(to string) *ExprTypeReduceTimeRangeBuilder {
	builder.internal.To = to

	return builder
}
