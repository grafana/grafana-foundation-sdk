// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsConditionsOperator] = (*ExprTypeClassicConditionsConditionsOperatorBuilder)(nil)

type ExprTypeClassicConditionsConditionsOperatorBuilder struct {
	internal *ExprTypeClassicConditionsConditionsOperator
	errors   map[string]cog.BuildErrors
}

func NewExprTypeClassicConditionsConditionsOperatorBuilder() *ExprTypeClassicConditionsConditionsOperatorBuilder {
	resource := NewExprTypeClassicConditionsConditionsOperator()
	builder := &ExprTypeClassicConditionsConditionsOperatorBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsOperatorBuilder) Build() (ExprTypeClassicConditionsConditionsOperator, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsConditionsOperator{}, err
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsOperatorBuilder) Type(typeArg ExprTypeClassicConditionsConditionsOperatorType) *ExprTypeClassicConditionsConditionsOperatorBuilder {
	builder.internal.Type = typeArg

	return builder
}
