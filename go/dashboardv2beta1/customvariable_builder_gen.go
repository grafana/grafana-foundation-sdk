// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CustomVariableKind] = (*CustomVariableBuilder)(nil)

// Custom variable kind
type CustomVariableBuilder struct {
	internal *CustomVariableKind
	errors   cog.BuildErrors
}

func NewCustomVariableBuilder(name string) *CustomVariableBuilder {
	resource := NewCustomVariableKind()
	builder := &CustomVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "CustomVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *CustomVariableBuilder) Build() (CustomVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return CustomVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return CustomVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.customVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CustomVariableBuilder) RecordError(path string, err error) *CustomVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *CustomVariableBuilder) Name(name string) *CustomVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *CustomVariableBuilder) Query(query string) *CustomVariableBuilder {
	builder.internal.Spec.Query = query

	return builder
}

func (builder *CustomVariableBuilder) Current(current VariableOption) *CustomVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *CustomVariableBuilder) Options(options []VariableOption) *CustomVariableBuilder {
	builder.internal.Spec.Options = options

	return builder
}

func (builder *CustomVariableBuilder) Multi(multi bool) *CustomVariableBuilder {
	builder.internal.Spec.Multi = multi

	return builder
}

func (builder *CustomVariableBuilder) IncludeAll(includeAll bool) *CustomVariableBuilder {
	builder.internal.Spec.IncludeAll = includeAll

	return builder
}

func (builder *CustomVariableBuilder) AllValue(allValue string) *CustomVariableBuilder {
	builder.internal.Spec.AllValue = &allValue

	return builder
}

func (builder *CustomVariableBuilder) Label(label string) *CustomVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *CustomVariableBuilder) Hide(hide VariableHide) *CustomVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *CustomVariableBuilder) SkipUrlSync(skipUrlSync bool) *CustomVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *CustomVariableBuilder) Description(description string) *CustomVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}

func (builder *CustomVariableBuilder) AllowCustomValue(allowCustomValue bool) *CustomVariableBuilder {
	builder.internal.Spec.AllowCustomValue = allowCustomValue

	return builder
}
