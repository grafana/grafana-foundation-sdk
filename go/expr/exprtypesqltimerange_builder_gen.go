// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeSqlTimeRange] = (*ExprTypeSqlTimeRangeBuilder)(nil)

type ExprTypeSqlTimeRangeBuilder struct {
	internal *ExprTypeSqlTimeRange
	errors   cog.BuildErrors
}

func NewExprTypeSqlTimeRangeBuilder() *ExprTypeSqlTimeRangeBuilder {
	resource := NewExprTypeSqlTimeRange()
	builder := &ExprTypeSqlTimeRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeSqlTimeRangeBuilder) Build() (ExprTypeSqlTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeSqlTimeRange{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeSqlTimeRange{}, cog.MakeBuildErrors("expr.exprTypeSqlTimeRange", builder.errors)
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeSqlTimeRangeBuilder) From(from string) *ExprTypeSqlTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeSqlTimeRangeBuilder) To(to string) *ExprTypeSqlTimeRangeBuilder {
	builder.internal.To = to

	return builder
}
