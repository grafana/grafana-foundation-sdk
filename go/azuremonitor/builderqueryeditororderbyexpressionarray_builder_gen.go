// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorOrderByExpressionArray] = (*BuilderQueryEditorOrderByExpressionArrayBuilder)(nil)

type BuilderQueryEditorOrderByExpressionArrayBuilder struct {
	internal *BuilderQueryEditorOrderByExpressionArray
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorOrderByExpressionArrayBuilder() *BuilderQueryEditorOrderByExpressionArrayBuilder {
	resource := NewBuilderQueryEditorOrderByExpressionArray()
	builder := &BuilderQueryEditorOrderByExpressionArrayBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorOrderByExpressionArrayBuilder) Build() (BuilderQueryEditorOrderByExpressionArray, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorOrderByExpressionArray{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorOrderByExpressionArrayBuilder) Expressions(expressions []cog.Builder[BuilderQueryEditorOrderByExpression]) *BuilderQueryEditorOrderByExpressionArrayBuilder {
	expressionsResources := make([]BuilderQueryEditorOrderByExpression, 0, len(expressions))
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

func (builder *BuilderQueryEditorOrderByExpressionArrayBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorOrderByExpressionArrayBuilder {
	builder.internal.Type = typeArg

	return builder
}
