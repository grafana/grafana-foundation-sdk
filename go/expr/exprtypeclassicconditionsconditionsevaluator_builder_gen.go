// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsEvaluator] = (*ExprTypeClassicConditionsConditionsEvaluatorBuilder)(nil)

type ExprTypeClassicConditionsConditionsEvaluatorBuilder struct {
	internal *ExprTypeClassicConditionsConditionsEvaluator
	errors   map[string]cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsEvaluatorBuilder() *ExprTypeClassicConditionsConditionsEvaluatorBuilder {
	resource := &ExprTypeClassicConditionsConditionsEvaluator{}
	builder := &ExprTypeClassicConditionsConditionsEvaluatorBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsEvaluatorBuilder) Build() (ExprTypeClassicConditionsConditionsEvaluator, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsEvaluator", err)...)
	}

	if len(errs) != 0 {
		return ExprTypeClassicConditionsConditionsEvaluator{}, errs
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

func (builder *ExprTypeClassicConditionsConditionsEvaluatorBuilder) applyDefaults() {
}
