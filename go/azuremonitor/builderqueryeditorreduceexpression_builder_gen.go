// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorReduceExpression] = (*BuilderQueryEditorReduceExpressionBuilder)(nil)

type BuilderQueryEditorReduceExpressionBuilder struct {
	internal *BuilderQueryEditorReduceExpression
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorReduceExpressionBuilder() *BuilderQueryEditorReduceExpressionBuilder {
	resource := NewBuilderQueryEditorReduceExpression()
	builder := &BuilderQueryEditorReduceExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorReduceExpressionBuilder) Build() (BuilderQueryEditorReduceExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorReduceExpression{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorReduceExpressionBuilder) Property(property cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorReduceExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors["property"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Property = &propertyResource

	return builder
}

func (builder *BuilderQueryEditorReduceExpressionBuilder) Reduce(reduce cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorReduceExpressionBuilder {
	reduceResource, err := reduce.Build()
	if err != nil {
		builder.errors["reduce"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Reduce = &reduceResource

	return builder
}

func (builder *BuilderQueryEditorReduceExpressionBuilder) Parameters(parameters []cog.Builder[BuilderQueryEditorFunctionParameterExpression]) *BuilderQueryEditorReduceExpressionBuilder {
	parametersResources := make([]BuilderQueryEditorFunctionParameterExpression, 0, len(parameters))
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

func (builder *BuilderQueryEditorReduceExpressionBuilder) Focus(focus bool) *BuilderQueryEditorReduceExpressionBuilder {
	builder.internal.Focus = &focus

	return builder
}
