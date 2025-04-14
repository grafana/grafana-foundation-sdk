// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorGroupByExpression] = (*BuilderQueryEditorGroupByExpressionBuilder)(nil)

type BuilderQueryEditorGroupByExpressionBuilder struct {
	internal *BuilderQueryEditorGroupByExpression
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorGroupByExpressionBuilder() *BuilderQueryEditorGroupByExpressionBuilder {
	resource := NewBuilderQueryEditorGroupByExpression()
	builder := &BuilderQueryEditorGroupByExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorGroupByExpressionBuilder) Build() (BuilderQueryEditorGroupByExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorGroupByExpression{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorGroupByExpressionBuilder) Property(property cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorGroupByExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors["property"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Property = &propertyResource

	return builder
}

func (builder *BuilderQueryEditorGroupByExpressionBuilder) Interval(interval cog.Builder[BuilderQueryEditorProperty]) *BuilderQueryEditorGroupByExpressionBuilder {
	intervalResource, err := interval.Build()
	if err != nil {
		builder.errors["interval"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Interval = &intervalResource

	return builder
}

func (builder *BuilderQueryEditorGroupByExpressionBuilder) Focus(focus bool) *BuilderQueryEditorGroupByExpressionBuilder {
	builder.internal.Focus = &focus

	return builder
}

func (builder *BuilderQueryEditorGroupByExpressionBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorGroupByExpressionBuilder {
	builder.internal.Type = &typeArg

	return builder
}
