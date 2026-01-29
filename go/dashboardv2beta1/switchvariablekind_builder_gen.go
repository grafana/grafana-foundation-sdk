// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SwitchVariableKind] = (*SwitchVariableKindBuilder)(nil)

type SwitchVariableKindBuilder struct {
	internal *SwitchVariableKind
	errors   cog.BuildErrors
}

func NewSwitchVariableKindBuilder() *SwitchVariableKindBuilder {
	resource := NewSwitchVariableKind()
	builder := &SwitchVariableKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "SwitchVariable"

	return builder
}

func (builder *SwitchVariableKindBuilder) Build() (SwitchVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return SwitchVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return SwitchVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.switchVariableKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SwitchVariableKindBuilder) Spec(spec cog.Builder[SwitchVariableSpec]) *SwitchVariableKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}
