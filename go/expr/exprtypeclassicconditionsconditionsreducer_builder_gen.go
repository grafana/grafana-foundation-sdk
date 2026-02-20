// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsReducer] = (*ExprTypeClassicConditionsConditionsReducerBuilder)(nil)

type ExprTypeClassicConditionsConditionsReducerBuilder struct {
	internal *ExprTypeClassicConditionsConditionsReducer
	errors   cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsReducerBuilder() *ExprTypeClassicConditionsConditionsReducerBuilder {
	resource := NewExprTypeClassicConditionsConditionsReducer()
	builder := &ExprTypeClassicConditionsConditionsReducerBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsReducerBuilder) Build() (ExprTypeClassicConditionsConditionsReducer, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsConditionsReducer{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeClassicConditionsConditionsReducer{}, cog.MakeBuildErrors("expr.exprTypeClassicConditionsConditionsReducer", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsReducerBuilder) RecordError(path string, err error) *ExprTypeClassicConditionsConditionsReducerBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ExprTypeClassicConditionsConditionsReducerBuilder) Type(typeArg string) *ExprTypeClassicConditionsConditionsReducerBuilder {
	builder.internal.Type = typeArg

	return builder
}
