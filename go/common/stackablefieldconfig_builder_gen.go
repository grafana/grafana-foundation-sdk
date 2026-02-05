// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StackableFieldConfig] = (*StackableFieldConfigBuilder)(nil)

// TODO docs
type StackableFieldConfigBuilder struct {
	internal *StackableFieldConfig
	errors   cog.BuildErrors
}

func NewStackableFieldConfigBuilder() *StackableFieldConfigBuilder {
	resource := NewStackableFieldConfig()
	builder := &StackableFieldConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *StackableFieldConfigBuilder) Build() (StackableFieldConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return StackableFieldConfig{}, err
	}

	if len(builder.errors) > 0 {
		return StackableFieldConfig{}, cog.MakeBuildErrors("common.stackableFieldConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *StackableFieldConfigBuilder) Stacking(stacking cog.Builder[StackingConfig]) *StackableFieldConfigBuilder {
	stackingResource, err := stacking.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Stacking = &stackingResource

	return builder
}
