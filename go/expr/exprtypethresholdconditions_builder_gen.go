// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeThresholdConditions] = (*ExprTypeThresholdConditionsBuilder)(nil)

type ExprTypeThresholdConditionsBuilder struct {
	internal *ExprTypeThresholdConditions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeThresholdConditionsBuilder() *ExprTypeThresholdConditionsBuilder {
	resource := &ExprTypeThresholdConditions{}
	builder := &ExprTypeThresholdConditionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeThresholdConditionsBuilder) Build() (ExprTypeThresholdConditions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeThresholdConditions{}, err
	}

	return *builder.internal, nil
}

func (builder *ExprTypeThresholdConditionsBuilder) Evaluator(evaluator cog.Builder[ExprTypeThresholdConditionsEvaluator]) *ExprTypeThresholdConditionsBuilder {
	evaluatorResource, err := evaluator.Build()
	if err != nil {
		builder.errors["evaluator"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Evaluator = evaluatorResource

	return builder
}

func (builder *ExprTypeThresholdConditionsBuilder) LoadedDimensions(loadedDimensions any) *ExprTypeThresholdConditionsBuilder {
	builder.internal.LoadedDimensions = &loadedDimensions

	return builder
}

func (builder *ExprTypeThresholdConditionsBuilder) UnloadEvaluator(unloadEvaluator cog.Builder[ExprTypeThresholdConditionsUnloadEvaluator]) *ExprTypeThresholdConditionsBuilder {
	unloadEvaluatorResource, err := unloadEvaluator.Build()
	if err != nil {
		builder.errors["unloadEvaluator"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.UnloadEvaluator = &unloadEvaluatorResource

	return builder
}

func (builder *ExprTypeThresholdConditionsBuilder) applyDefaults() {
}
