// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExemplarConfig] = (*ExemplarConfigBuilder)(nil)

// Controls exemplar options
type ExemplarConfigBuilder struct {
	internal *ExemplarConfig
	errors   cog.BuildErrors
}

func NewExemplarConfigBuilder() *ExemplarConfigBuilder {
	resource := NewExemplarConfig()
	builder := &ExemplarConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExemplarConfigBuilder) Build() (ExemplarConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExemplarConfig{}, err
	}

	if len(builder.errors) > 0 {
		return ExemplarConfig{}, cog.MakeBuildErrors("heatmap.exemplarConfig", builder.errors)
	}

	return *builder.internal, nil
}

// Sets the color of the exemplar markers
func (builder *ExemplarConfigBuilder) Color(color string) *ExemplarConfigBuilder {
	builder.internal.Color = color

	return builder
}
