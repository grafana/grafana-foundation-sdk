// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle] = (*StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder)(nil)

type StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder struct {
	internal *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle
	errors   cog.BuildErrors
}

func NewStringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder() *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder {
	resource := NewStringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle()
	builder := &StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder) Build() (StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle, error) {
	if err := builder.internal.Validate(); err != nil {
		return StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle{}, err
	}

	if len(builder.errors) > 0 {
		return StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle{}, cog.MakeBuildErrors("dashboardv2beta1.stringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder) String(stringArg string) *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder) Bool(boolArg bool) *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder {
	builder.internal.Bool = &boolArg

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder) Float64(float64Arg float64) *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder {
	builder.internal.Float64 = &float64Arg

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder) CustomVariableValue(customVariableValue CustomVariableValue) *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder {
	builder.internal.CustomVariableValue = &customVariableValue

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder) ArrayOfVariableValueSingle(arrayOfVariableValueSingle []cog.Builder[VariableValueSingle]) *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder {
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
