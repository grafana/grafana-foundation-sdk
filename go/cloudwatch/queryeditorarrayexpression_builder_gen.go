// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorArrayExpression] = (*QueryEditorArrayExpressionBuilder)(nil)

type QueryEditorArrayExpressionBuilder struct {
	internal *QueryEditorArrayExpression
	errors   cog.BuildErrors
}

func NewQueryEditorArrayExpressionBuilder() *QueryEditorArrayExpressionBuilder {
	resource := NewQueryEditorArrayExpression()
	builder := &QueryEditorArrayExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryEditorArrayExpressionBuilder) Build() (QueryEditorArrayExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorArrayExpression{}, err
	}

	if len(builder.errors) > 0 {
		return QueryEditorArrayExpression{}, cog.MakeBuildErrors("cloudwatch.queryEditorArrayExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryEditorArrayExpressionBuilder) Type(typeArg QueryEditorArrayExpressionType) *QueryEditorArrayExpressionBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *QueryEditorArrayExpressionBuilder) Expressions(expressions []QueryEditorExpression) *QueryEditorArrayExpressionBuilder {
	builder.internal.Expressions = expressions

	return builder
}
