// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorWhereExpressionArray] = (*BuilderQueryEditorWhereExpressionArrayBuilder)(nil)

type BuilderQueryEditorWhereExpressionArrayBuilder struct {
	internal *BuilderQueryEditorWhereExpressionArray
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorWhereExpressionArrayBuilder() *BuilderQueryEditorWhereExpressionArrayBuilder {
	resource := NewBuilderQueryEditorWhereExpressionArray()
	builder := &BuilderQueryEditorWhereExpressionArrayBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorWhereExpressionArrayBuilder) Build() (BuilderQueryEditorWhereExpressionArray, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorWhereExpressionArray{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorWhereExpressionArrayBuilder) Expressions(expressions []cog.Builder[BuilderQueryEditorWhereExpression]) *BuilderQueryEditorWhereExpressionArrayBuilder {
	expressionsResources := make([]BuilderQueryEditorWhereExpression, 0, len(expressions))
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

func (builder *BuilderQueryEditorWhereExpressionArrayBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorWhereExpressionArrayBuilder {
	builder.internal.Type = typeArg

	return builder
}
