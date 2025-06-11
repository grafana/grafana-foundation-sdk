// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsOperator] = (*ExprTypeClassicConditionsConditionsOperatorBuilder)(nil)

type ExprTypeClassicConditionsConditionsOperatorBuilder struct {
	internal *ExprTypeClassicConditionsConditionsOperator
	errors   cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsOperatorBuilder() *ExprTypeClassicConditionsConditionsOperatorBuilder {
	resource := NewExprTypeClassicConditionsConditionsOperator()
	builder := &ExprTypeClassicConditionsConditionsOperatorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsOperatorBuilder) Build() (ExprTypeClassicConditionsConditionsOperator, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsConditionsOperator{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeClassicConditionsConditionsOperator{}, cog.MakeBuildErrors("expr.exprTypeClassicConditionsConditionsOperator", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsOperatorBuilder) Type(typeArg ExprTypeClassicConditionsConditionsOperatorType) *ExprTypeClassicConditionsConditionsOperatorBuilder {
	builder.internal.Type = typeArg

	return builder
}
