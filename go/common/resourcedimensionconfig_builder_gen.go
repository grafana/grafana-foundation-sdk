// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResourceDimensionConfig] = (*ResourceDimensionConfigBuilder)(nil)

// Links to a resource (image/svg path)
type ResourceDimensionConfigBuilder struct {
	internal *ResourceDimensionConfig
	errors   map[string]cog.BuildErrors
}

func NewResourceDimensionConfigBuilder() *ResourceDimensionConfigBuilder {
	resource := &ResourceDimensionConfig{}
	builder := &ResourceDimensionConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ResourceDimensionConfigBuilder) Build() (ResourceDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ResourceDimensionConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *ResourceDimensionConfigBuilder) Mode(mode ResourceDimensionMode) *ResourceDimensionConfigBuilder {
	builder.internal.Mode = mode

	return builder
}

// fixed: T -- will be added by each element
func (builder *ResourceDimensionConfigBuilder) Field(field string) *ResourceDimensionConfigBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *ResourceDimensionConfigBuilder) Fixed(fixed string) *ResourceDimensionConfigBuilder {
	builder.internal.Fixed = &fixed

	return builder
}

func (builder *ResourceDimensionConfigBuilder) applyDefaults() {
}
