// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StringOrBoolOrFloat64OrCustomVariableValue] = (*StringOrBoolOrFloat64OrCustomVariableValueBuilder)(nil)

type StringOrBoolOrFloat64OrCustomVariableValueBuilder struct {
	internal *StringOrBoolOrFloat64OrCustomVariableValue
	errors   cog.BuildErrors
}

func NewStringOrBoolOrFloat64OrCustomVariableValueBuilder() *StringOrBoolOrFloat64OrCustomVariableValueBuilder {
	resource := NewStringOrBoolOrFloat64OrCustomVariableValue()
	builder := &StringOrBoolOrFloat64OrCustomVariableValueBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) Build() (StringOrBoolOrFloat64OrCustomVariableValue, error) {
	if err := builder.internal.Validate(); err != nil {
		return StringOrBoolOrFloat64OrCustomVariableValue{}, err
	}

	if len(builder.errors) > 0 {
		return StringOrBoolOrFloat64OrCustomVariableValue{}, cog.MakeBuildErrors("dashboardv2beta1.stringOrBoolOrFloat64OrCustomVariableValue", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) String(stringArg string) *StringOrBoolOrFloat64OrCustomVariableValueBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) Bool(boolArg bool) *StringOrBoolOrFloat64OrCustomVariableValueBuilder {
	builder.internal.Bool = &boolArg

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) Float64(float64Arg float64) *StringOrBoolOrFloat64OrCustomVariableValueBuilder {
	builder.internal.Float64 = &float64Arg

	return builder
}

func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) CustomVariableValue(customVariableValue CustomVariableValue) *StringOrBoolOrFloat64OrCustomVariableValueBuilder {
	builder.internal.CustomVariableValue = &customVariableValue

	return builder
}
