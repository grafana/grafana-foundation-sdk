// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorColumnsExpression] = (*BuilderQueryEditorColumnsExpressionBuilder)(nil)

type BuilderQueryEditorColumnsExpressionBuilder struct {
	internal *BuilderQueryEditorColumnsExpression
	errors   cog.BuildErrors
}

func NewBuilderQueryEditorColumnsExpressionBuilder() *BuilderQueryEditorColumnsExpressionBuilder {
	resource := NewBuilderQueryEditorColumnsExpression()
	builder := &BuilderQueryEditorColumnsExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryEditorColumnsExpressionBuilder) Build() (BuilderQueryEditorColumnsExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorColumnsExpression{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryEditorColumnsExpression{}, cog.MakeBuildErrors("azuremonitor.builderQueryEditorColumnsExpression", builder.errors)
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
