// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TextDimensionConfig] = (*TextDimensionConfigBuilder)(nil)

type TextDimensionConfigBuilder struct {
	internal *TextDimensionConfig
	errors   cog.BuildErrors
}

func NewTextDimensionConfigBuilder() *TextDimensionConfigBuilder {
	resource := NewTextDimensionConfig()
	builder := &TextDimensionConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TextDimensionConfigBuilder) Build() (TextDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return TextDimensionConfig{}, err
	}

	if len(builder.errors) > 0 {
		return TextDimensionConfig{}, cog.MakeBuildErrors("common.textDimensionConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TextDimensionConfigBuilder) RecordError(path string, err error) *TextDimensionConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *TextDimensionConfigBuilder) Mode(mode TextDimensionMode) *TextDimensionConfigBuilder {
	builder.internal.Mode = mode

	return builder
}

// fixed: T -- will be added by each element
func (builder *TextDimensionConfigBuilder) Field(field string) *TextDimensionConfigBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *TextDimensionConfigBuilder) Fixed(fixed string) *TextDimensionConfigBuilder {
	builder.internal.Fixed = &fixed

	return builder
}
