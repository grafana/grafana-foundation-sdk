// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorColumnsExpression] = (*BuilderQueryEditorColumnsExpressionBuilder)(nil)

type BuilderQueryEditorColumnsExpressionBuilder struct {
	internal *BuilderQueryEditorColumnsExpression
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorColumnsExpressionBuilder() *BuilderQueryEditorColumnsExpressionBuilder {
	resource := NewBuilderQueryEditorColumnsExpression()
	builder := &BuilderQueryEditorColumnsExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorColumnsExpressionBuilder) Build() (BuilderQueryEditorColumnsExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorColumnsExpression{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorColumnsExpressionBuilder) Columns(columns []string) *BuilderQueryEditorColumnsExpressionBuilder {
	builder.internal.Columns = columns

	return builder
}

func (builder *BuilderQueryEditorColumnsExpressionBuilder) Type(typeArg BuilderQueryEditorExpressionType) *BuilderQueryEditorColumnsExpressionBuilder {
	builder.internal.Type = typeArg

	return builder
}
