// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsReducer] = (*ExprTypeClassicConditionsConditionsReducerBuilder)(nil)

type ExprTypeClassicConditionsConditionsReducerBuilder struct {
	internal *ExprTypeClassicConditionsConditionsReducer
	errors   map[string]cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsReducerBuilder() *ExprTypeClassicConditionsConditionsReducerBuilder {
	resource := &ExprTypeClassicConditionsConditionsReducer{}
	builder := &ExprTypeClassicConditionsConditionsReducerBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsReducerBuilder) Build() (ExprTypeClassicConditionsConditionsReducer, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsReducer", err)...)
	}

	if len(errs) != 0 {
		return ExprTypeClassicConditionsConditionsReducer{}, errs
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsReducerBuilder) Type(typeArg string) *ExprTypeClassicConditionsConditionsReducerBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsReducerBuilder) applyDefaults() {
}
