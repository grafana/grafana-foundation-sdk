// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorPropertyExpression] = (*BuilderQueryEditorPropertyExpressionBuilder)(nil)

type BuilderQueryEditorPropertyExpressionBuilder struct {
	internal *BuilderQueryEditorPropertyExpression
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorPropertyExpressionBuilder() *BuilderQueryEditorPropertyExpressionBuilder {
	resource := NewBuilderQueryEditorPropertyExpression()
	builder := &BuilderQueryEditorPropertyExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorPropertyExpressionBuilder) Build() (BuilderQueryEditorPropertyExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorPropertyExpression{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorPropertyExpressionBuilder) Property(property cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorPropertyExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors["property"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

func (builder *BuilderQueryEditorPropertyExpressionBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorPropertyExpressionBuilder {
	builder.internal.Type = typeArg

	return builder
}
