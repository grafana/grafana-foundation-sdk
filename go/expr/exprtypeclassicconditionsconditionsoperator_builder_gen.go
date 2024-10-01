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
	resource := &ExprTypeClassicConditionsConditionsOperator{}
	builder := &ExprTypeClassicConditionsConditionsOperatorBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsOperatorBuilder) Build() (ExprTypeClassicConditionsConditionsOperator, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsOperator", err)...)
	}

	if len(errs) != 0 {
		return ExprTypeClassicConditionsConditionsOperator{}, errs
	}

	return *builder.internal, nil
}

func (builder *ExprTypeClassicConditionsConditionsOperatorBuilder) Type(typeArg TypeClassicConditionsType) *ExprTypeClassicConditionsConditionsOperatorBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *ExprTypeClassicConditionsConditionsOperatorBuilder) applyDefaults() {
}
