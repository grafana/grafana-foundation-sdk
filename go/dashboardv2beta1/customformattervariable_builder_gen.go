// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CustomFormatterVariable] = (*CustomFormatterVariableBuilder)(nil)

// Custom formatter variable
type CustomFormatterVariableBuilder struct {
	internal *CustomFormatterVariable
	errors   cog.BuildErrors
}

func NewCustomFormatterVariableBuilder() *CustomFormatterVariableBuilder {
	resource := NewCustomFormatterVariable()
	builder := &CustomFormatterVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CustomFormatterVariableBuilder) Build() (CustomFormatterVariable, error) {
	if err := builder.internal.Validate(); err != nil {
		return CustomFormatterVariable{}, err
	}

	if len(builder.errors) > 0 {
		return CustomFormatterVariable{}, cog.MakeBuildErrors("dashboardv2beta1.customFormatterVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CustomFormatterVariableBuilder) Name(name string) *CustomFormatterVariableBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *CustomFormatterVariableBuilder) Type(typeArg VariableType) *CustomFormatterVariableBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *CustomFormatterVariableBuilder) Multi(multi bool) *CustomFormatterVariableBuilder {
	builder.internal.Multi = multi

	return builder
}

func (builder *CustomFormatterVariableBuilder) IncludeAll(includeAll bool) *CustomFormatterVariableBuilder {
	builder.internal.IncludeAll = includeAll

	return builder
}
