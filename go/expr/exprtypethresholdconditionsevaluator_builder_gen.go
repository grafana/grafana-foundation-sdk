// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeThresholdConditionsEvaluator] = (*ExprTypeThresholdConditionsEvaluatorBuilder)(nil)

type ExprTypeThresholdConditionsEvaluatorBuilder struct {
	internal *ExprTypeThresholdConditionsEvaluator
	errors   cog.BuildErrors
}

func NewExprTypeThresholdConditionsEvaluatorBuilder() *ExprTypeThresholdConditionsEvaluatorBuilder {
	resource := NewExprTypeThresholdConditionsEvaluator()
	builder := &ExprTypeThresholdConditionsEvaluatorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Build() (ExprTypeThresholdConditionsEvaluator, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeThresholdConditionsEvaluator{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeThresholdConditionsEvaluator{}, cog.MakeBuildErrors("expr.exprTypeThresholdConditionsEvaluator", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) RecordError(path string, err error) *ExprTypeThresholdConditionsEvaluatorBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Params(params []float64) *ExprTypeThresholdConditionsEvaluatorBuilder {
	builder.internal.Params = params

	return builder
}

// e.g. "gt"
func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Type(typeArg ExprTypeThresholdConditionsEvaluatorType) *ExprTypeThresholdConditionsEvaluatorBuilder {
	builder.internal.Type = typeArg

	return builder
}
