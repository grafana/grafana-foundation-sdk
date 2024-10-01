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
	resource := &ExprTypeClassicConditionsConditionsQuery{}
	builder := &ExprTypeClassicConditionsConditionsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsQueryBuilder) Build() (ExprTypeClassicConditionsConditionsQuery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsQuery", err)...)
	}

	if len(errs) != 0 {
		return ExprTypeClassicConditionsConditionsQuery{}, errs
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsQueryBuilder) Params(params []string) *ExprTypeClassicConditionsConditionsQueryBuilder {
	builder.internal.Params = params

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsQueryBuilder) applyDefaults() {
}
