// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorWhereExpressionItems] = (*BuilderQueryEditorWhereExpressionItemsBuilder)(nil)

type BuilderQueryEditorWhereExpressionItemsBuilder struct {
	internal *BuilderQueryEditorWhereExpressionItems
	errors   cog.BuildErrors
}

func NewBuilderQueryEditorWhereExpressionItemsBuilder() *BuilderQueryEditorWhereExpressionItemsBuilder {
	resource := NewBuilderQueryEditorWhereExpressionItems()
	builder := &BuilderQueryEditorWhereExpressionItemsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Build() (BuilderQueryEditorWhereExpressionItems, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorWhereExpressionItems{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryEditorWhereExpressionItems{}, cog.MakeBuildErrors("azuremonitor.builderQueryEditorWhereExpressionItems", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Property(property cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorWhereExpressionItemsBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Operator(operator cog.Builder[BuilderQueryEditorOperator]) *BuilderQueryEditorWhereExpressionItemsBuilder {
	operatorResource, err := operator.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Operator = operatorResource

	return builder
}

func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorWhereExpressionItemsBuilder {
	builder.internal.Type = typeArg

	return builder
}
