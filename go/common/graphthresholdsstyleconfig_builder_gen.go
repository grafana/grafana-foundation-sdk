// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GraphThresholdsStyleConfig] = (*GraphThresholdsStyleConfigBuilder)(nil)

// TODO docs
type GraphThresholdsStyleConfigBuilder struct {
	internal *GraphThresholdsStyleConfig
	errors   cog.BuildErrors
}

func NewGraphThresholdsStyleConfigBuilder() *GraphThresholdsStyleConfigBuilder {
	resource := NewGraphThresholdsStyleConfig()
	builder := &GraphThresholdsStyleConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *GraphThresholdsStyleConfigBuilder) Build() (GraphThresholdsStyleConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return GraphThresholdsStyleConfig{}, err
	}

	if len(builder.errors) > 0 {
		return GraphThresholdsStyleConfig{}, cog.MakeBuildErrors("common.graphThresholdsStyleConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *GraphThresholdsStyleConfigBuilder) Mode(mode GraphTresholdsStyleMode) *GraphThresholdsStyleConfigBuilder {
	builder.internal.Mode = mode

	return builder
}
