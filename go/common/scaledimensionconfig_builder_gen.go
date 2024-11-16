// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ScaleDimensionConfig] = (*ScaleDimensionConfigBuilder)(nil)

type ScaleDimensionConfigBuilder struct {
	internal *ScaleDimensionConfig
	errors   map[string]cog.BuildErrors
}

func NewScaleDimensionConfigBuilder() *ScaleDimensionConfigBuilder {
	resource := NewScaleDimensionConfig()
	builder := &ScaleDimensionConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ScaleDimensionConfigBuilder) Build() (ScaleDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ScaleDimensionConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *ScaleDimensionConfigBuilder) Min(min float64) *ScaleDimensionConfigBuilder {
	builder.internal.Min = min

	return builder
}

func (builder *ScaleDimensionConfigBuilder) Max(max float64) *ScaleDimensionConfigBuilder {
	builder.internal.Max = max

	return builder
}

func (builder *ScaleDimensionConfigBuilder) Fixed(fixed float64) *ScaleDimensionConfigBuilder {
	builder.internal.Fixed = &fixed

	return builder
}

// fixed: T -- will be added by each element
func (builder *ScaleDimensionConfigBuilder) Field(field string) *ScaleDimensionConfigBuilder {
	builder.internal.Field = &field

	return builder
}

// | *"linear"
func (builder *ScaleDimensionConfigBuilder) Mode(mode ScaleDimensionMode) *ScaleDimensionConfigBuilder {
	builder.internal.Mode = &mode

	return builder
}
