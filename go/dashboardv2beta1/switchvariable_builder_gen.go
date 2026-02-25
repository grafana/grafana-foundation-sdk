// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SwitchVariableKind] = (*SwitchVariableBuilder)(nil)

type SwitchVariableBuilder struct {
	internal *SwitchVariableKind
	errors   cog.BuildErrors
}

func NewSwitchVariableBuilder(name string) *SwitchVariableBuilder {
	resource := NewSwitchVariableKind()
	builder := &SwitchVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "SwitchVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *SwitchVariableBuilder) Build() (SwitchVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return SwitchVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return SwitchVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.switchVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SwitchVariableBuilder) RecordError(path string, err error) *SwitchVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SwitchVariableBuilder) Name(name string) *SwitchVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *SwitchVariableBuilder) Current(current string) *SwitchVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *SwitchVariableBuilder) EnabledValue(enabledValue string) *SwitchVariableBuilder {
	builder.internal.Spec.EnabledValue = enabledValue

	return builder
}

func (builder *SwitchVariableBuilder) DisabledValue(disabledValue string) *SwitchVariableBuilder {
	builder.internal.Spec.DisabledValue = disabledValue

	return builder
}

func (builder *SwitchVariableBuilder) Label(label string) *SwitchVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *SwitchVariableBuilder) Hide(hide VariableHide) *SwitchVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *SwitchVariableBuilder) SkipUrlSync(skipUrlSync bool) *SwitchVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *SwitchVariableBuilder) Description(description string) *SwitchVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}
