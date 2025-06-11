// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorOrderByExpression] = (*BuilderQueryEditorOrderByExpressionBuilder)(nil)

type BuilderQueryEditorOrderByExpressionBuilder struct {
	internal *BuilderQueryEditorOrderByExpression
	errors   cog.BuildErrors
}

func NewBuilderQueryEditorOrderByExpressionBuilder() *BuilderQueryEditorOrderByExpressionBuilder {
	resource := NewBuilderQueryEditorOrderByExpression()
	builder := &BuilderQueryEditorOrderByExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryEditorOrderByExpressionBuilder) Build() (BuilderQueryEditorOrderByExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorOrderByExpression{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryEditorOrderByExpression{}, cog.MakeBuildErrors("azuremonitor.builderQueryEditorOrderByExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorOrderByExpressionBuilder) Property(property cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorOrderByExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

func (builder *BuilderQueryEditorOrderByExpressionBuilder) Order(order BuilderQueryEditorOrderByOptions) *BuilderQueryEditorOrderByExpressionBuilder {
	builder.internal.Order = order

	return builder
}

func (builder *BuilderQueryEditorOrderByExpressionBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorOrderByExpressionBuilder {
	builder.internal.Type = typeArg

	return builder
}
