// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ColorDimensionConfig] = (*ColorDimensionConfigBuilder)(nil)

type ColorDimensionConfigBuilder struct {
	internal *ColorDimensionConfig
	errors   map[string]cog.BuildErrors
}

func NewColorDimensionConfigBuilder() *ColorDimensionConfigBuilder {
	resource := NewColorDimensionConfig()
	builder := &ColorDimensionConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ColorDimensionConfigBuilder) Build() (ColorDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ColorDimensionConfig{}, err
	}

	return *builder.internal, nil
}

// color value
func (builder *ColorDimensionConfigBuilder) Fixed(fixed string) *ColorDimensionConfigBuilder {
	builder.internal.Fixed = &fixed

	return builder
}

// fixed: T -- will be added by each element
func (builder *ColorDimensionConfigBuilder) Field(field string) *ColorDimensionConfigBuilder {
	builder.internal.Field = &field

	return builder
}
