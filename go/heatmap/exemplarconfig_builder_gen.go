// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExemplarConfig] = (*ExemplarConfigBuilder)(nil)

// Controls exemplar options
type ExemplarConfigBuilder struct {
	internal *ExemplarConfig
	errors   map[string]cog.BuildErrors
}

func NewExemplarConfigBuilder() *ExemplarConfigBuilder {
	resource := &ExemplarConfig{}
	builder := &ExemplarConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExemplarConfigBuilder) Build() (ExemplarConfig, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExemplarConfig", err)...)
	}

	if len(errs) != 0 {
		return ExemplarConfig{}, errs
	}

	return *builder.internal, nil
}

// Sets the color of the exemplar markers
func (builder *ExemplarConfigBuilder) Color(color string) *ExemplarConfigBuilder {
	builder.internal.Color = color

	return builder
}

func (builder *ExemplarConfigBuilder) applyDefaults() {
}
