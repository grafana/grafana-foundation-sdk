// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[LineConfig] = (*LineConfigBuilder)(nil)

type LineConfigBuilder struct {
	internal *LineConfig
	errors   map[string]cog.BuildErrors
}

func NewLineConfigBuilder() *LineConfigBuilder {
	resource := &LineConfig{}
	builder := &LineConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LineConfigBuilder) Build() (LineConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return LineConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *LineConfigBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *LineConfigBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors["color"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *LineConfigBuilder) Width(width float64) *LineConfigBuilder {
	builder.internal.Width = &width

	return builder
}

func (builder *LineConfigBuilder) applyDefaults() {
}
