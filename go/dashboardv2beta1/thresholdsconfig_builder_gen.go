// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ThresholdsConfig] = (*ThresholdsConfigBuilder)(nil)

type ThresholdsConfigBuilder struct {
	internal *ThresholdsConfig
	errors   cog.BuildErrors
}

func NewThresholdsConfigBuilder() *ThresholdsConfigBuilder {
	resource := NewThresholdsConfig()
	builder := &ThresholdsConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ThresholdsConfigBuilder) Build() (ThresholdsConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ThresholdsConfig{}, err
	}

	if len(builder.errors) > 0 {
		return ThresholdsConfig{}, cog.MakeBuildErrors("dashboardv2beta1.thresholdsConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ThresholdsConfigBuilder) RecordError(path string, err error) *ThresholdsConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ThresholdsConfigBuilder) Mode(mode ThresholdsMode) *ThresholdsConfigBuilder {
	builder.internal.Mode = mode

	return builder
}

func (builder *ThresholdsConfigBuilder) Steps(steps []Threshold) *ThresholdsConfigBuilder {
	builder.internal.Steps = steps

	return builder
}
