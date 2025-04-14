// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorGroupByExpressionArray] = (*BuilderQueryEditorGroupByExpressionArrayBuilder)(nil)

type BuilderQueryEditorGroupByExpressionArrayBuilder struct {
	internal *BuilderQueryEditorGroupByExpressionArray
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorGroupByExpressionArrayBuilder() *BuilderQueryEditorGroupByExpressionArrayBuilder {
	resource := NewBuilderQueryEditorGroupByExpressionArray()
	builder := &BuilderQueryEditorGroupByExpressionArrayBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorGroupByExpressionArrayBuilder) Build() (BuilderQueryEditorGroupByExpressionArray, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorGroupByExpressionArray{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorGroupByExpressionArrayBuilder) Expressions(expressions []cog.Builder[BuilderQueryEditorGroupByExpression]) *BuilderQueryEditorGroupByExpressionArrayBuilder {
	expressionsResources := make([]BuilderQueryEditorGroupByExpression, 0, len(expressions))
	for _, r1 := range expressions {
		expressionsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["expressions"] = err.(cog.BuildErrors)
			return builder
		}
		expressionsResources = append(expressionsResources, expressionsDepth1)
	}
	builder.internal.Expressions = expressionsResources

	return builder
}

func (builder *BuilderQueryEditorGroupByExpressionArrayBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorGroupByExpressionArrayBuilder {
	builder.internal.Type = typeArg

	return builder
}
