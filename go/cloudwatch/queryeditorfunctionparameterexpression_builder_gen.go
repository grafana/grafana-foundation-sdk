// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorFunctionParameterExpression] = (*QueryEditorFunctionParameterExpressionBuilder)(nil)

type QueryEditorFunctionParameterExpressionBuilder struct {
	internal *QueryEditorFunctionParameterExpression
	errors   cog.BuildErrors
}

func NewQueryEditorFunctionParameterExpressionBuilder() *QueryEditorFunctionParameterExpressionBuilder {
	resource := NewQueryEditorFunctionParameterExpression()
	builder := &QueryEditorFunctionParameterExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryEditorFunctionParameterExpressionBuilder) Build() (QueryEditorFunctionParameterExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorFunctionParameterExpression{}, err
	}

	if len(builder.errors) > 0 {
		return QueryEditorFunctionParameterExpression{}, cog.MakeBuildErrors("cloudwatch.queryEditorFunctionParameterExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryEditorFunctionParameterExpressionBuilder) RecordError(path string, err error) *QueryEditorFunctionParameterExpressionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryEditorFunctionParameterExpressionBuilder) Name(name string) *QueryEditorFunctionParameterExpressionBuilder {
	builder.internal.Name = &name

	return builder
}
