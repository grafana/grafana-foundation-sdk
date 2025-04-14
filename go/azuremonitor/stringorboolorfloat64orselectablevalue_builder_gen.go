// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StringOrBoolOrFloat64OrSelectableValue] = (*StringOrBoolOrFloat64OrSelectableValueBuilder)(nil)

type StringOrBoolOrFloat64OrSelectableValueBuilder struct {
	internal *StringOrBoolOrFloat64OrSelectableValue
	errors   map[string]cog.BuildErrors
}

func NewStringOrBoolOrFloat64OrSelectableValueBuilder() *StringOrBoolOrFloat64OrSelectableValueBuilder {
	resource := NewStringOrBoolOrFloat64OrSelectableValue()
	builder := &StringOrBoolOrFloat64OrSelectableValueBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) Build() (StringOrBoolOrFloat64OrSelectableValue, error) {
	if err := builder.internal.Validate(); err != nil {
		return StringOrBoolOrFloat64OrSelectableValue{}, err
	}

	return *builder.internal, nil
}

func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) String(stringArg string) *StringOrBoolOrFloat64OrSelectableValueBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) Bool(boolArg bool) *StringOrBoolOrFloat64OrSelectableValueBuilder {
	builder.internal.Bool = &boolArg

	return builder
}

func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) Float64(float64Arg float64) *StringOrBoolOrFloat64OrSelectableValueBuilder {
	builder.internal.Float64 = &float64Arg

	return builder
}

func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) SelectableValue(selectableValue cog.Builder[SelectableValue]) *StringOrBoolOrFloat64OrSelectableValueBuilder {
	selectableValueResource, err := selectableValue.Build()
	if err != nil {
		builder.errors["SelectableValue"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.SelectableValue = &selectableValueResource

	return builder
}
