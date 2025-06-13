// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorFunctionParameterExpression] = (*BuilderQueryEditorFunctionParameterExpressionBuilder)(nil)

type BuilderQueryEditorFunctionParameterExpressionBuilder struct {
	internal *BuilderQueryEditorFunctionParameterExpression
	errors   cog.BuildErrors
}

func NewBuilderQueryEditorFunctionParameterExpressionBuilder() *BuilderQueryEditorFunctionParameterExpressionBuilder {
	resource := NewBuilderQueryEditorFunctionParameterExpression()
	builder := &BuilderQueryEditorFunctionParameterExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryEditorFunctionParameterExpressionBuilder) Build() (BuilderQueryEditorFunctionParameterExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorFunctionParameterExpression{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryEditorFunctionParameterExpression{}, cog.MakeBuildErrors("azuremonitor.builderQueryEditorFunctionParameterExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorFunctionParameterExpressionBuilder) Value(value string) *BuilderQueryEditorFunctionParameterExpressionBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *BuilderQueryEditorFunctionParameterExpressionBuilder) FieldType(fieldType BuilderQueryEditorPropertyType) *BuilderQueryEditorFunctionParameterExpressionBuilder {
	builder.internal.FieldType = fieldType

	return builder
}

func (builder *BuilderQueryEditorFunctionParameterExpressionBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorFunctionParameterExpressionBuilder {
	builder.internal.Type = typeArg

	return builder
}
