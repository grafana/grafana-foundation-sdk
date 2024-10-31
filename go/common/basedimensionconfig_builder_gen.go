// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseDimensionConfig] = (*BaseDimensionConfigBuilder)(nil)

type BaseDimensionConfigBuilder struct {
	internal *BaseDimensionConfig
	errors   map[string]cog.BuildErrors
}

func NewBaseDimensionConfigBuilder() *BaseDimensionConfigBuilder {
	resource := &BaseDimensionConfig{}
	builder := &BaseDimensionConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *BaseDimensionConfigBuilder) Build() (BaseDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseDimensionConfig{}, err
	}

	return *builder.internal, nil
}

// fixed: T -- will be added by each element
func (builder *BaseDimensionConfigBuilder) Field(field string) *BaseDimensionConfigBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *BaseDimensionConfigBuilder) applyDefaults() {
}
