// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsTimeRange] = (*ExprTypeClassicConditionsTimeRangeBuilder)(nil)

type ExprTypeClassicConditionsTimeRangeBuilder struct {
	internal *ExprTypeClassicConditionsTimeRange
	errors   cog.BuildErrors
}

func NewExprTypeClassicConditionsTimeRangeBuilder() *ExprTypeClassicConditionsTimeRangeBuilder {
	resource := NewExprTypeClassicConditionsTimeRange()
	builder := &ExprTypeClassicConditionsTimeRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeClassicConditionsTimeRangeBuilder) Build() (ExprTypeClassicConditionsTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsTimeRange{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeClassicConditionsTimeRange{}, cog.MakeBuildErrors("expr.exprTypeClassicConditionsTimeRange", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsTimeRangeBuilder) RecordError(path string, err error) *ExprTypeClassicConditionsTimeRangeBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// From is the start time of the query.
func (builder *ExprTypeClassicConditionsTimeRangeBuilder) From(from string) *ExprTypeClassicConditionsTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeClassicConditionsTimeRangeBuilder) To(to string) *ExprTypeClassicConditionsTimeRangeBuilder {
	builder.internal.To = to

	return builder
}
