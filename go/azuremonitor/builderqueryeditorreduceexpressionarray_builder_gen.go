// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorReduceExpressionArray] = (*BuilderQueryEditorReduceExpressionArrayBuilder)(nil)

type BuilderQueryEditorReduceExpressionArrayBuilder struct {
	internal *BuilderQueryEditorReduceExpressionArray
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorReduceExpressionArrayBuilder() *BuilderQueryEditorReduceExpressionArrayBuilder {
	resource := NewBuilderQueryEditorReduceExpressionArray()
	builder := &BuilderQueryEditorReduceExpressionArrayBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorReduceExpressionArrayBuilder) Build() (BuilderQueryEditorReduceExpressionArray, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorReduceExpressionArray{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorReduceExpressionArrayBuilder) Expressions(expressions []cog.Builder[BuilderQueryEditorReduceExpression]) *BuilderQueryEditorReduceExpressionArrayBuilder {
	expressionsResources := make([]BuilderQueryEditorReduceExpression, 0, len(expressions))
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

func (builder *BuilderQueryEditorReduceExpressionArrayBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorReduceExpressionArrayBuilder {
	builder.internal.Type = typeArg

	return builder
}
