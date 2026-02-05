// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableValue] = (*VariableValueBuilder)(nil)

// Variable types
type VariableValueBuilder struct {
	internal *VariableValue
	errors   cog.BuildErrors
}

func NewVariableValueBuilder() *VariableValueBuilder {
	resource := NewVariableValue()
	builder := &VariableValueBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *VariableValueBuilder) Build() (VariableValue, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableValue{}, err
	}

	if len(builder.errors) > 0 {
		return VariableValue{}, cog.MakeBuildErrors("dashboardv2beta1.variableValue", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *VariableValueBuilder) String(stringArg string) *VariableValueBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *VariableValueBuilder) Bool(boolArg bool) *VariableValueBuilder {
	builder.internal.Bool = &boolArg

	return builder
}

func (builder *VariableValueBuilder) Float64(float64Arg float64) *VariableValueBuilder {
	builder.internal.Float64 = &float64Arg

	return builder
}

func (builder *VariableValueBuilder) CustomVariableValue(customVariableValue CustomVariableValue) *VariableValueBuilder {
	builder.internal.CustomVariableValue = &customVariableValue

	return builder
}

func (builder *VariableValueBuilder) ArrayOfVariableValueSingle(arrayOfVariableValueSingle []cog.Builder[VariableValueSingle]) *VariableValueBuilder {
	arrayOfVariableValueSingleResources := make([]VariableValueSingle, 0, len(arrayOfVariableValueSingle))
	for _, r1 := range arrayOfVariableValueSingle {
		arrayOfVariableValueSingleDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		arrayOfVariableValueSingleResources = append(arrayOfVariableValueSingleResources, arrayOfVariableValueSingleDepth1)
	}
	builder.internal.ArrayOfVariableValueSingle = arrayOfVariableValueSingleResources

	return builder
}
