// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableValueSingle] = (*VariableValueSingleBuilder)(nil)

type VariableValueSingleBuilder struct {
	internal *VariableValueSingle
	errors   cog.BuildErrors
}

func NewVariableValueSingleBuilder() *VariableValueSingleBuilder {
	resource := NewVariableValueSingle()
	builder := &VariableValueSingleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *VariableValueSingleBuilder) Build() (VariableValueSingle, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableValueSingle{}, err
	}

	if len(builder.errors) > 0 {
		return VariableValueSingle{}, cog.MakeBuildErrors("dashboardv2beta1.variableValueSingle", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *VariableValueSingleBuilder) String(stringArg string) *VariableValueSingleBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *VariableValueSingleBuilder) Bool(boolArg bool) *VariableValueSingleBuilder {
	builder.internal.Bool = &boolArg

	return builder
}

func (builder *VariableValueSingleBuilder) Float64(float64Arg float64) *VariableValueSingleBuilder {
	builder.internal.Float64 = &float64Arg

	return builder
}

func (builder *VariableValueSingleBuilder) CustomVariableValue(customVariableValue CustomVariableValue) *VariableValueSingleBuilder {
	builder.internal.CustomVariableValue = &customVariableValue

	return builder
}
