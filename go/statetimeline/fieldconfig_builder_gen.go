// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statetimeline

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[FieldConfig] = (*FieldConfigBuilder)(nil)

type FieldConfigBuilder struct {
	internal *FieldConfig
	errors   cog.BuildErrors
}

func NewFieldConfigBuilder() *FieldConfigBuilder {
	resource := NewFieldConfig()
	builder := &FieldConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FieldConfigBuilder) Build() (FieldConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return FieldConfig{}, err
	}

	if len(builder.errors) > 0 {
		return FieldConfig{}, cog.MakeBuildErrors("statetimeline.fieldConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FieldConfigBuilder) LineWidth(lineWidth uint32) *FieldConfigBuilder {
	builder.internal.LineWidth = &lineWidth

	return builder
}

func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder {
	hideFromResource, err := hideFrom.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.HideFrom = &hideFromResource

	return builder
}

func (builder *FieldConfigBuilder) FillOpacity(fillOpacity uint32) *FieldConfigBuilder {
	builder.internal.FillOpacity = &fillOpacity

	return builder
}
