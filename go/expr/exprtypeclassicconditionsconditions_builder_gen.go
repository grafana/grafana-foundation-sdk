// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditions] = (*ExprTypeClassicConditionsConditionsBuilder)(nil)

type ExprTypeClassicConditionsConditionsBuilder struct {
	internal *ExprTypeClassicConditionsConditions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsBuilder() *ExprTypeClassicConditionsConditionsBuilder {
	resource := &ExprTypeClassicConditionsConditions{}
	builder := &ExprTypeClassicConditionsConditionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsBuilder) Build() (ExprTypeClassicConditionsConditions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsConditions{}, err
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsBuilder) Evaluator(evaluator cog.Builder[ExprTypeClassicConditionsConditionsEvaluator]) *ExprTypeClassicConditionsConditionsBuilder {
	evaluatorResource, err := evaluator.Build()
	if err != nil {
		builder.errors["evaluator"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Evaluator = evaluatorResource

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsBuilder) Operator(operator cog.Builder[ExprTypeClassicConditionsConditionsOperator]) *ExprTypeClassicConditionsConditionsBuilder {
	operatorResource, err := operator.Build()
	if err != nil {
		builder.errors["operator"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Operator = operatorResource

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsBuilder) Query(query cog.Builder[ExprTypeClassicConditionsConditionsQuery]) *ExprTypeClassicConditionsConditionsBuilder {
	queryResource, err := query.Build()
	if err != nil {
		builder.errors["query"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Query = queryResource

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsBuilder) Reducer(reducer cog.Builder[ExprTypeClassicConditionsConditionsReducer]) *ExprTypeClassicConditionsConditionsBuilder {
	reducerResource, err := reducer.Build()
	if err != nil {
		builder.errors["reducer"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Reducer = reducerResource

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsBuilder) applyDefaults() {
}
