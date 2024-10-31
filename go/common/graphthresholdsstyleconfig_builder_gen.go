// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GraphThresholdsStyleConfig] = (*GraphThresholdsStyleConfigBuilder)(nil)

// TODO docs
type GraphThresholdsStyleConfigBuilder struct {
	internal *GraphThresholdsStyleConfig
	errors   map[string]cog.BuildErrors
}

func NewGraphThresholdsStyleConfigBuilder() *GraphThresholdsStyleConfigBuilder {
	resource := &GraphThresholdsStyleConfig{}
	builder := &GraphThresholdsStyleConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *GraphThresholdsStyleConfigBuilder) Build() (GraphThresholdsStyleConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return GraphThresholdsStyleConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *GraphThresholdsStyleConfigBuilder) Mode(mode GraphThresholdsStyleMode) *GraphThresholdsStyleConfigBuilder {
	builder.internal.Mode = mode

	return builder
}

func (builder *GraphThresholdsStyleConfigBuilder) applyDefaults() {
}
