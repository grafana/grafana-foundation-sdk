// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorArrayExpression] = (*QueryEditorArrayExpressionBuilder)(nil)

type QueryEditorArrayExpressionBuilder struct {
	internal *QueryEditorArrayExpression
	errors   map[string]cog.BuildErrors
}

func NewQueryEditorArrayExpressionBuilder() *QueryEditorArrayExpressionBuilder {
	resource := &QueryEditorArrayExpression{}
	builder := &QueryEditorArrayExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *QueryEditorArrayExpressionBuilder) Build() (QueryEditorArrayExpression, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorArrayExpression", err)...)
	}

	if len(errs) != 0 {
		return QueryEditorArrayExpression{}, errs
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

func (builder *QueryEditorArrayExpressionBuilder) applyDefaults() {
}
