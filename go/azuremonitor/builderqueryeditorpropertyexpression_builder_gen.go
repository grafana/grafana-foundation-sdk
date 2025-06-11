// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorPropertyExpression] = (*BuilderQueryEditorPropertyExpressionBuilder)(nil)

type BuilderQueryEditorPropertyExpressionBuilder struct {
	internal *BuilderQueryEditorPropertyExpression
	errors   cog.BuildErrors
}

func NewBuilderQueryEditorPropertyExpressionBuilder() *BuilderQueryEditorPropertyExpressionBuilder {
	resource := NewBuilderQueryEditorPropertyExpression()
	builder := &BuilderQueryEditorPropertyExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryEditorPropertyExpressionBuilder) Build() (BuilderQueryEditorPropertyExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorPropertyExpression{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryEditorPropertyExpression{}, cog.MakeBuildErrors("azuremonitor.builderQueryEditorPropertyExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorPropertyExpressionBuilder) Property(property cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorPropertyExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

func (builder *BuilderQueryEditorPropertyExpressionBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorPropertyExpressionBuilder {
	builder.internal.Type = typeArg

	return builder
}
