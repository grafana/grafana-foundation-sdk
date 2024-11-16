// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ScaleDistributionConfig] = (*ScaleDistributionConfigBuilder)(nil)

// TODO docs
type ScaleDistributionConfigBuilder struct {
	internal *ScaleDistributionConfig
	errors   map[string]cog.BuildErrors
}

func NewScaleDistributionConfigBuilder() *ScaleDistributionConfigBuilder {
	resource := NewScaleDistributionConfig()
	builder := &ScaleDistributionConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ScaleDistributionConfigBuilder) Build() (ScaleDistributionConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ScaleDistributionConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *ScaleDistributionConfigBuilder) Type(typeArg ScaleDistribution) *ScaleDistributionConfigBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *ScaleDistributionConfigBuilder) Log(log float64) *ScaleDistributionConfigBuilder {
	builder.internal.Log = &log

	return builder
}

func (builder *ScaleDistributionConfigBuilder) LinearThreshold(linearThreshold float64) *ScaleDistributionConfigBuilder {
	builder.internal.LinearThreshold = &linearThreshold

	return builder
}
