// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HideableFieldConfig] = (*HideableFieldConfigBuilder)(nil)

// TODO docs
type HideableFieldConfigBuilder struct {
	internal *HideableFieldConfig
	errors   cog.BuildErrors
}

func NewHideableFieldConfigBuilder() *HideableFieldConfigBuilder {
	resource := NewHideableFieldConfig()
	builder := &HideableFieldConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *HideableFieldConfigBuilder) Build() (HideableFieldConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return HideableFieldConfig{}, err
	}

	if len(builder.errors) > 0 {
		return HideableFieldConfig{}, cog.MakeBuildErrors("common.hideableFieldConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *HideableFieldConfigBuilder) HideFrom(hideFrom cog.Builder[HideSeriesConfig]) *HideableFieldConfigBuilder {
	hideFromResource, err := hideFrom.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.HideFrom = &hideFromResource

	return builder
}
