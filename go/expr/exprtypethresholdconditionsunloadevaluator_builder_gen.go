// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeThresholdConditionsUnloadEvaluator] = (*ExprTypeThresholdConditionsUnloadEvaluatorBuilder)(nil)

type ExprTypeThresholdConditionsUnloadEvaluatorBuilder struct {
	internal *ExprTypeThresholdConditionsUnloadEvaluator
	errors   cog.BuildErrors
}

func NewExprTypeThresholdConditionsUnloadEvaluatorBuilder() *ExprTypeThresholdConditionsUnloadEvaluatorBuilder {
	resource := NewExprTypeThresholdConditionsUnloadEvaluator()
	builder := &ExprTypeThresholdConditionsUnloadEvaluatorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeThresholdConditionsUnloadEvaluatorBuilder) Build() (ExprTypeThresholdConditionsUnloadEvaluator, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeThresholdConditionsUnloadEvaluator{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeThresholdConditionsUnloadEvaluator{}, cog.MakeBuildErrors("expr.exprTypeThresholdConditionsUnloadEvaluator", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeThresholdConditionsUnloadEvaluatorBuilder) Params(params []float64) *ExprTypeThresholdConditionsUnloadEvaluatorBuilder {
	builder.internal.Params = params

	return builder
}

// e.g. "gt"
func (builder *ExprTypeThresholdConditionsUnloadEvaluatorBuilder) Type(typeArg ExprTypeThresholdConditionsUnloadEvaluatorType) *ExprTypeThresholdConditionsUnloadEvaluatorBuilder {
	builder.internal.Type = typeArg

	return builder
}
