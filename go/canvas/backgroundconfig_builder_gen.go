// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[BackgroundConfig] = (*BackgroundConfigBuilder)(nil)

type BackgroundConfigBuilder struct {
	internal *BackgroundConfig
	errors   map[string]cog.BuildErrors
}

func NewBackgroundConfigBuilder() *BackgroundConfigBuilder {
	resource := &BackgroundConfig{}
	builder := &BackgroundConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *BackgroundConfigBuilder) Build() (BackgroundConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return BackgroundConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *BackgroundConfigBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *BackgroundConfigBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors["color"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *BackgroundConfigBuilder) Image(image cog.Builder[common.ResourceDimensionConfig]) *BackgroundConfigBuilder {
	imageResource, err := image.Build()
	if err != nil {
		builder.errors["image"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Image = &imageResource

	return builder
}

func (builder *BackgroundConfigBuilder) Size(size BackgroundImageSize) *BackgroundConfigBuilder {
	builder.internal.Size = &size

	return builder
}

func (builder *BackgroundConfigBuilder) applyDefaults() {
}
