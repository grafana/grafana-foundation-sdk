// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BarConfig] = (*BarConfigBuilder)(nil)

// TODO docs
type BarConfigBuilder struct {
	internal *BarConfig
	errors   cog.BuildErrors
}

func NewBarConfigBuilder() *BarConfigBuilder {
	resource := NewBarConfig()
	builder := &BarConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BarConfigBuilder) Build() (BarConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return BarConfig{}, err
	}

	if len(builder.errors) > 0 {
		return BarConfig{}, cog.MakeBuildErrors("common.barConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BarConfigBuilder) RecordError(path string, err error) *BarConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *BarConfigBuilder) BarAlignment(barAlignment BarAlignment) *BarConfigBuilder {
	builder.internal.BarAlignment = &barAlignment

	return builder
}

func (builder *BarConfigBuilder) BarWidthFactor(barWidthFactor float64) *BarConfigBuilder {
	builder.internal.BarWidthFactor = &barWidthFactor

	return builder
}

func (builder *BarConfigBuilder) BarMaxWidth(barMaxWidth float64) *BarConfigBuilder {
	builder.internal.BarMaxWidth = &barMaxWidth

	return builder
}
