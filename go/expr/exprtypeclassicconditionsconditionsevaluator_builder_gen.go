// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsEvaluator] = (*ExprTypeClassicConditionsConditionsEvaluatorBuilder)(nil)

type ExprTypeClassicConditionsConditionsEvaluatorBuilder struct {
	internal *ExprTypeClassicConditionsConditionsEvaluator
	errors   cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsEvaluatorBuilder() *ExprTypeClassicConditionsConditionsEvaluatorBuilder {
	resource := NewExprTypeClassicConditionsConditionsEvaluator()
	builder := &ExprTypeClassicConditionsConditionsEvaluatorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsEvaluatorBuilder) Build() (ExprTypeClassicConditionsConditionsEvaluator, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsConditionsEvaluator{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeClassicConditionsConditionsEvaluator{}, cog.MakeBuildErrors("expr.exprTypeClassicConditionsConditionsEvaluator", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsEvaluatorBuilder) Params(params []float64) *ExprTypeClassicConditionsConditionsEvaluatorBuilder {
	builder.internal.Params = params

	return builder
}

// e.g. "gt"
func (builder *ExprTypeClassicConditionsConditionsEvaluatorBuilder) Type(typeArg string) *ExprTypeClassicConditionsConditionsEvaluatorBuilder {
	builder.internal.Type = typeArg

	return builder
}
