// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseDimensionConfig] = (*BaseDimensionConfigBuilder)(nil)

type BaseDimensionConfigBuilder struct {
	internal *BaseDimensionConfig
	errors   cog.BuildErrors
}

func NewBaseDimensionConfigBuilder() *BaseDimensionConfigBuilder {
	resource := NewBaseDimensionConfig()
	builder := &BaseDimensionConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BaseDimensionConfigBuilder) Build() (BaseDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseDimensionConfig{}, err
	}

	if len(builder.errors) > 0 {
		return BaseDimensionConfig{}, cog.MakeBuildErrors("common.baseDimensionConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BaseDimensionConfigBuilder) RecordError(path string, err error) *BaseDimensionConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// fixed: T -- will be added by each element
func (builder *BaseDimensionConfigBuilder) Field(field string) *BaseDimensionConfigBuilder {
	builder.internal.Field = &field

	return builder
}
