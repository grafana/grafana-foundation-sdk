// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConstantVariableKind] = (*ConstantVariableBuilder)(nil)

// Constant variable kind
type ConstantVariableBuilder struct {
	internal *ConstantVariableKind
	errors   cog.BuildErrors
}

func NewConstantVariableBuilder(name string) *ConstantVariableBuilder {
	resource := NewConstantVariableKind()
	builder := &ConstantVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "ConstantVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *ConstantVariableBuilder) Build() (ConstantVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConstantVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return ConstantVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.constantVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConstantVariableBuilder) RecordError(path string, err error) *ConstantVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConstantVariableBuilder) Spec(spec ConstantVariableSpec) *ConstantVariableBuilder {
	builder.internal.Spec = spec

	return builder
}

func (builder *ConstantVariableBuilder) Name(name string) *ConstantVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *ConstantVariableBuilder) Query(query string) *ConstantVariableBuilder {
	builder.internal.Spec.Query = query

	return builder
}

func (builder *ConstantVariableBuilder) Current(current VariableOption) *ConstantVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *ConstantVariableBuilder) Label(label string) *ConstantVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *ConstantVariableBuilder) Hide(hide VariableHide) *ConstantVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *ConstantVariableBuilder) SkipUrlSync(skipUrlSync bool) *ConstantVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *ConstantVariableBuilder) Description(description string) *ConstantVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}
