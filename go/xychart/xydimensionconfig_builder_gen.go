// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XYDimensionConfig] = (*XYDimensionConfigBuilder)(nil)

type XYDimensionConfigBuilder struct {
	internal *XYDimensionConfig
	errors   map[string]cog.BuildErrors
}

func NewXYDimensionConfigBuilder() *XYDimensionConfigBuilder {
	resource := &XYDimensionConfig{}
	builder := &XYDimensionConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *XYDimensionConfigBuilder) Build() (XYDimensionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return XYDimensionConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *XYDimensionConfigBuilder) Frame(frame int32) *XYDimensionConfigBuilder {
	builder.internal.Frame = frame

	return builder
}

func (builder *XYDimensionConfigBuilder) X(x string) *XYDimensionConfigBuilder {
	builder.internal.X = &x

	return builder
}

func (builder *XYDimensionConfigBuilder) Exclude(exclude []string) *XYDimensionConfigBuilder {
	builder.internal.Exclude = exclude

	return builder
}

func (builder *XYDimensionConfigBuilder) applyDefaults() {
}
