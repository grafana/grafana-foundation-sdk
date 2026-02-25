// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingVariableKind] = (*ConditionalRenderingVariableBuilder)(nil)

type ConditionalRenderingVariableBuilder struct {
	internal *ConditionalRenderingVariableKind
	errors   cog.BuildErrors
}

func NewConditionalRenderingVariableBuilder() *ConditionalRenderingVariableBuilder {
	resource := NewConditionalRenderingVariableKind()
	builder := &ConditionalRenderingVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "ConditionalRenderingVariable"

	return builder
}

func (builder *ConditionalRenderingVariableBuilder) Build() (ConditionalRenderingVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingVariableBuilder) RecordError(path string, err error) *ConditionalRenderingVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConditionalRenderingVariableBuilder) Variable(variable string) *ConditionalRenderingVariableBuilder {
	builder.internal.Spec.Variable = variable

	return builder
}

func (builder *ConditionalRenderingVariableBuilder) Operator(operator ConditionalRenderingVariableSpecOperator) *ConditionalRenderingVariableBuilder {
	builder.internal.Spec.Operator = operator

	return builder
}

func (builder *ConditionalRenderingVariableBuilder) Value(value string) *ConditionalRenderingVariableBuilder {
	builder.internal.Spec.Value = value

	return builder
}
