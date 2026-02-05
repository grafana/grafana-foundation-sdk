// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SwitchVariableSpec] = (*SwitchVariableSpecBuilder)(nil)

type SwitchVariableSpecBuilder struct {
	internal *SwitchVariableSpec
	errors   cog.BuildErrors
}

func NewSwitchVariableSpecBuilder() *SwitchVariableSpecBuilder {
	resource := NewSwitchVariableSpec()
	builder := &SwitchVariableSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SwitchVariableSpecBuilder) Build() (SwitchVariableSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return SwitchVariableSpec{}, err
	}

	if len(builder.errors) > 0 {
		return SwitchVariableSpec{}, cog.MakeBuildErrors("dashboardv2beta1.switchVariableSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SwitchVariableSpecBuilder) Name(name string) *SwitchVariableSpecBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *SwitchVariableSpecBuilder) Current(current string) *SwitchVariableSpecBuilder {
	builder.internal.Current = current

	return builder
}

func (builder *SwitchVariableSpecBuilder) EnabledValue(enabledValue string) *SwitchVariableSpecBuilder {
	builder.internal.EnabledValue = enabledValue

	return builder
}

func (builder *SwitchVariableSpecBuilder) DisabledValue(disabledValue string) *SwitchVariableSpecBuilder {
	builder.internal.DisabledValue = disabledValue

	return builder
}

func (builder *SwitchVariableSpecBuilder) Label(label string) *SwitchVariableSpecBuilder {
	builder.internal.Label = &label

	return builder
}

func (builder *SwitchVariableSpecBuilder) Hide(hide VariableHide) *SwitchVariableSpecBuilder {
	builder.internal.Hide = hide

	return builder
}

func (builder *SwitchVariableSpecBuilder) SkipUrlSync(skipUrlSync bool) *SwitchVariableSpecBuilder {
	builder.internal.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *SwitchVariableSpecBuilder) Description(description string) *SwitchVariableSpecBuilder {
	builder.internal.Description = &description

	return builder
}
