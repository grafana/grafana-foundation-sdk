// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorFunctionParameterExpression] = (*BuilderQueryEditorFunctionParameterExpressionBuilder)(nil)

type BuilderQueryEditorFunctionParameterExpressionBuilder struct {
	internal *BuilderQueryEditorFunctionParameterExpression
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorFunctionParameterExpressionBuilder() *BuilderQueryEditorFunctionParameterExpressionBuilder {
	resource := NewBuilderQueryEditorFunctionParameterExpression()
	builder := &BuilderQueryEditorFunctionParameterExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorFunctionParameterExpressionBuilder) Build() (BuilderQueryEditorFunctionParameterExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorFunctionParameterExpression{}, err
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
