// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ColorDimensionConfig] = (*ColorDimensionConfigBuilder)(nil)

type ColorDimensionConfigBuilder struct {
	internal *ColorDimensionConfig
	errors   cog.BuildErrors
}

func NewColorDimensionConfigBuilder() *ColorDimensionConfigBuilder {
	resource := NewColorDimensionConfig()
	builder := &ColorDimensionConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ColorDimensionConfigBuilder) Build() (ColorDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ColorDimensionConfig{}, err
	}

	if len(builder.errors) > 0 {
		return ColorDimensionConfig{}, cog.MakeBuildErrors("common.colorDimensionConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ColorDimensionConfigBuilder) RecordError(path string, err error) *ColorDimensionConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
