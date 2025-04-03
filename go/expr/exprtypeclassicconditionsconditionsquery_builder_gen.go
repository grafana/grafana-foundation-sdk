// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsQuery] = (*ExprTypeClassicConditionsConditionsQueryBuilder)(nil)

type ExprTypeClassicConditionsConditionsQueryBuilder struct {
	internal *ExprTypeClassicConditionsConditionsQuery
	errors   map[string]cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsQueryBuilder() *ExprTypeClassicConditionsConditionsQueryBuilder {
	resource := NewExprTypeClassicConditionsConditionsQuery()
	builder := &ExprTypeClassicConditionsConditionsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsQueryBuilder) Build() (ExprTypeClassicConditionsConditionsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsConditionsQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsQueryBuilder) Params(params []string) *ExprTypeClassicConditionsConditionsQueryBuilder {
	builder.internal.Params = params

	return builder
}
