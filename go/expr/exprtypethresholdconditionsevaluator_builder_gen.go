// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeThresholdConditionsEvaluator] = (*ExprTypeThresholdConditionsEvaluatorBuilder)(nil)

type ExprTypeThresholdConditionsEvaluatorBuilder struct {
	internal *ExprTypeThresholdConditionsEvaluator
	errors   map[string]cog.BuildErrors
}

func NewExprTypeThresholdConditionsEvaluatorBuilder() *ExprTypeThresholdConditionsEvaluatorBuilder {
	resource := &ExprTypeThresholdConditionsEvaluator{}
	builder := &ExprTypeThresholdConditionsEvaluatorBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Build() (ExprTypeThresholdConditionsEvaluator, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeThresholdConditionsEvaluator", err)...)
	}

	if len(errs) != 0 {
		return ExprTypeThresholdConditionsEvaluator{}, errs
	}

	return *builder.internal, nil
}

func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Params(params []float64) *ExprTypeThresholdConditionsEvaluatorBuilder {
	builder.internal.Params = params

	return builder
}

// e.g. "gt"
func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Type(typeArg TypeThresholdType) *ExprTypeThresholdConditionsEvaluatorBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) applyDefaults() {
}
