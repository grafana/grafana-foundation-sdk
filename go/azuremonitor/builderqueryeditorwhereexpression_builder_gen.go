// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorWhereExpression] = (*BuilderQueryEditorWhereExpressionBuilder)(nil)

type BuilderQueryEditorWhereExpressionBuilder struct {
	internal *BuilderQueryEditorWhereExpression
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorWhereExpressionBuilder() *BuilderQueryEditorWhereExpressionBuilder {
	resource := NewBuilderQueryEditorWhereExpression()
	builder := &BuilderQueryEditorWhereExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorWhereExpressionBuilder) Build() (BuilderQueryEditorWhereExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorWhereExpression{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorWhereExpressionBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorWhereExpressionBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *BuilderQueryEditorWhereExpressionBuilder) Expressions(expressions []cog.Builder[BuilderQueryEditorWhereExpressionItems]) *BuilderQueryEditorWhereExpressionBuilder {
	expressionsResources := make([]BuilderQueryEditorWhereExpressionItems, 0, len(expressions))
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
