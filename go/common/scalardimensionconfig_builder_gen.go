// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ScalarDimensionConfig] = (*ScalarDimensionConfigBuilder)(nil)

type ScalarDimensionConfigBuilder struct {
	internal *ScalarDimensionConfig
	errors   cog.BuildErrors
}

func NewScalarDimensionConfigBuilder() *ScalarDimensionConfigBuilder {
	resource := NewScalarDimensionConfig()
	builder := &ScalarDimensionConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ScalarDimensionConfigBuilder) Build() (ScalarDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ScalarDimensionConfig{}, err
	}

	if len(builder.errors) > 0 {
		return ScalarDimensionConfig{}, cog.MakeBuildErrors("common.scalarDimensionConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ScalarDimensionConfigBuilder) Min(min float64) *ScalarDimensionConfigBuilder {
	builder.internal.Min = min

	return builder
}

func (builder *ScalarDimensionConfigBuilder) Max(max float64) *ScalarDimensionConfigBuilder {
	builder.internal.Max = max

	return builder
}

func (builder *ScalarDimensionConfigBuilder) Fixed(fixed float64) *ScalarDimensionConfigBuilder {
	builder.internal.Fixed = &fixed

	return builder
}

// fixed: T -- will be added by each element
func (builder *ScalarDimensionConfigBuilder) Field(field string) *ScalarDimensionConfigBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *ScalarDimensionConfigBuilder) Mode(mode ScalarDimensionMode) *ScalarDimensionConfigBuilder {
	builder.internal.Mode = &mode

	return builder
}
