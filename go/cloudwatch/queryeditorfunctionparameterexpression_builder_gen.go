// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorFunctionParameterExpression] = (*QueryEditorFunctionParameterExpressionBuilder)(nil)

type QueryEditorFunctionParameterExpressionBuilder struct {
	internal *QueryEditorFunctionParameterExpression
	errors   map[string]cog.BuildErrors
}

func NewQueryEditorFunctionParameterExpressionBuilder() *QueryEditorFunctionParameterExpressionBuilder {
	resource := &QueryEditorFunctionParameterExpression{}
	builder := &QueryEditorFunctionParameterExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "functionParameter"

	return builder
}

func (builder *QueryEditorFunctionParameterExpressionBuilder) Build() (QueryEditorFunctionParameterExpression, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionParameterExpression", err)...)
	}

	if len(errs) != 0 {
		return QueryEditorFunctionParameterExpression{}, errs
	}

	return *builder.internal, nil
}

func (builder *QueryEditorFunctionParameterExpressionBuilder) Name(name string) *QueryEditorFunctionParameterExpressionBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *QueryEditorFunctionParameterExpressionBuilder) applyDefaults() {
}
