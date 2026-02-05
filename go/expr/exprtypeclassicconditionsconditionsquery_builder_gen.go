// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsQuery] = (*ExprTypeClassicConditionsConditionsQueryBuilder)(nil)

type ExprTypeClassicConditionsConditionsQueryBuilder struct {
	internal *ExprTypeClassicConditionsConditionsQuery
	errors   cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsQueryBuilder() *ExprTypeClassicConditionsConditionsQueryBuilder {
	resource := NewExprTypeClassicConditionsConditionsQuery()
	builder := &ExprTypeClassicConditionsConditionsQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsQueryBuilder) Build() (ExprTypeClassicConditionsConditionsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsConditionsQuery{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeClassicConditionsConditionsQuery{}, cog.MakeBuildErrors("expr.exprTypeClassicConditionsConditionsQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsQueryBuilder) Params(params []string) *ExprTypeClassicConditionsConditionsQueryBuilder {
	builder.internal.Params = params

	return builder
}
