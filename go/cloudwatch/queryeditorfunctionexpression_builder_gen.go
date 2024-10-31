// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorFunctionExpression] = (*QueryEditorFunctionExpressionBuilder)(nil)

type QueryEditorFunctionExpressionBuilder struct {
	internal *QueryEditorFunctionExpression
	errors   map[string]cog.BuildErrors
}

func NewQueryEditorFunctionExpressionBuilder() *QueryEditorFunctionExpressionBuilder {
	resource := &QueryEditorFunctionExpression{}
	builder := &QueryEditorFunctionExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "function"

	return builder
}

func (builder *QueryEditorFunctionExpressionBuilder) Build() (QueryEditorFunctionExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorFunctionExpression{}, err
	}

	return *builder.internal, nil
}

func (builder *QueryEditorFunctionExpressionBuilder) Name(name string) *QueryEditorFunctionExpressionBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *QueryEditorFunctionExpressionBuilder) Parameters(parameters []cog.Builder[QueryEditorFunctionParameterExpression]) *QueryEditorFunctionExpressionBuilder {
	parametersResources := make([]QueryEditorFunctionParameterExpression, 0, len(parameters))
	for _, r1 := range parameters {
		parametersDepth1, err := r1.Build()
		if err != nil {
			builder.errors["parameters"] = err.(cog.BuildErrors)
			return builder
		}
		parametersResources = append(parametersResources, parametersDepth1)
	}
	builder.internal.Parameters = parametersResources

	return builder
}

func (builder *QueryEditorFunctionExpressionBuilder) applyDefaults() {
}
