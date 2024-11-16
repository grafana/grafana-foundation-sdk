// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HeatmapCalculationBucketConfig] = (*HeatmapCalculationBucketConfigBuilder)(nil)

type HeatmapCalculationBucketConfigBuilder struct {
	internal *HeatmapCalculationBucketConfig
	errors   map[string]cog.BuildErrors
}

func NewHeatmapCalculationBucketConfigBuilder() *HeatmapCalculationBucketConfigBuilder {
	resource := NewHeatmapCalculationBucketConfig()
	builder := &HeatmapCalculationBucketConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *HeatmapCalculationBucketConfigBuilder) Build() (HeatmapCalculationBucketConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapCalculationBucketConfig{}, err
	}

	return *builder.internal, nil
}

// Sets the bucket calculation mode
func (builder *HeatmapCalculationBucketConfigBuilder) Mode(mode HeatmapCalculationMode) *HeatmapCalculationBucketConfigBuilder {
	builder.internal.Mode = &mode

	return builder
}

// The number of buckets to use for the axis in the heatmap
func (builder *HeatmapCalculationBucketConfigBuilder) Value(value string) *HeatmapCalculationBucketConfigBuilder {
	builder.internal.Value = &value

	return builder
}

// Controls the scale of the buckets
func (builder *HeatmapCalculationBucketConfigBuilder) Scale(scale cog.Builder[ScaleDistributionConfig]) *HeatmapCalculationBucketConfigBuilder {
	scaleResource, err := scale.Build()
	if err != nil {
		builder.errors["scale"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Scale = &scaleResource

	return builder
}
