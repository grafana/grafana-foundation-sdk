// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingVariableSpec] = (*ConditionalRenderingVariableSpecBuilder)(nil)

type ConditionalRenderingVariableSpecBuilder struct {
	internal *ConditionalRenderingVariableSpec
	errors   cog.BuildErrors
}

func NewConditionalRenderingVariableSpecBuilder() *ConditionalRenderingVariableSpecBuilder {
	resource := NewConditionalRenderingVariableSpec()
	builder := &ConditionalRenderingVariableSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ConditionalRenderingVariableSpecBuilder) Build() (ConditionalRenderingVariableSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingVariableSpec{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingVariableSpec{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingVariableSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingVariableSpecBuilder) Variable(variable string) *ConditionalRenderingVariableSpecBuilder {
	builder.internal.Variable = variable

	return builder
}

func (builder *ConditionalRenderingVariableSpecBuilder) Operator(operator ConditionalRenderingVariableSpecOperator) *ConditionalRenderingVariableSpecBuilder {
	builder.internal.Operator = operator

	return builder
}

func (builder *ConditionalRenderingVariableSpecBuilder) Value(value string) *ConditionalRenderingVariableSpecBuilder {
	builder.internal.Value = value

	return builder
}
