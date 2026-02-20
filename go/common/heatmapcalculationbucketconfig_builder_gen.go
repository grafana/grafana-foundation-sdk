// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HeatmapCalculationBucketConfig] = (*HeatmapCalculationBucketConfigBuilder)(nil)

type HeatmapCalculationBucketConfigBuilder struct {
	internal *HeatmapCalculationBucketConfig
	errors   cog.BuildErrors
}

func NewHeatmapCalculationBucketConfigBuilder() *HeatmapCalculationBucketConfigBuilder {
	resource := NewHeatmapCalculationBucketConfig()
	builder := &HeatmapCalculationBucketConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *HeatmapCalculationBucketConfigBuilder) Build() (HeatmapCalculationBucketConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapCalculationBucketConfig{}, err
	}

	if len(builder.errors) > 0 {
		return HeatmapCalculationBucketConfig{}, cog.MakeBuildErrors("common.heatmapCalculationBucketConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *HeatmapCalculationBucketConfigBuilder) RecordError(path string, err error) *HeatmapCalculationBucketConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Scale = &scaleResource

	return builder
}
