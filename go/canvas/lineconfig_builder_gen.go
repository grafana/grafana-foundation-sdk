// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[LineConfig] = (*LineConfigBuilder)(nil)

type LineConfigBuilder struct {
	internal *LineConfig
	errors   cog.BuildErrors
}

func NewLineConfigBuilder() *LineConfigBuilder {
	resource := NewLineConfig()
	builder := &LineConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LineConfigBuilder) Build() (LineConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return LineConfig{}, err
	}

	if len(builder.errors) > 0 {
		return LineConfig{}, cog.MakeBuildErrors("canvas.lineConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LineConfigBuilder) RecordError(path string, err error) *LineConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *LineConfigBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *LineConfigBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *LineConfigBuilder) Width(width float64) *LineConfigBuilder {
	builder.internal.Width = &width

	return builder
}

func (builder *LineConfigBuilder) Radius(radius float64) *LineConfigBuilder {
	builder.internal.Radius = &radius

	return builder
}
