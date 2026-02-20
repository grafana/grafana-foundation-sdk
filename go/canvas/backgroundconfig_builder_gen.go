// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[BackgroundConfig] = (*BackgroundConfigBuilder)(nil)

type BackgroundConfigBuilder struct {
	internal *BackgroundConfig
	errors   cog.BuildErrors
}

func NewBackgroundConfigBuilder() *BackgroundConfigBuilder {
	resource := NewBackgroundConfig()
	builder := &BackgroundConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BackgroundConfigBuilder) Build() (BackgroundConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return BackgroundConfig{}, err
	}

	if len(builder.errors) > 0 {
		return BackgroundConfig{}, cog.MakeBuildErrors("canvas.backgroundConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BackgroundConfigBuilder) RecordError(path string, err error) *BackgroundConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *BackgroundConfigBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *BackgroundConfigBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *BackgroundConfigBuilder) Image(image cog.Builder[common.ResourceDimensionConfig]) *BackgroundConfigBuilder {
	imageResource, err := image.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Image = &imageResource

	return builder
}

func (builder *BackgroundConfigBuilder) Size(size BackgroundImageSize) *BackgroundConfigBuilder {
	builder.internal.Size = &size

	return builder
}
