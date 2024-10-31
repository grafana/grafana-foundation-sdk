// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HeatmapCalculationOptions] = (*HeatmapCalculationOptionsBuilder)(nil)

type HeatmapCalculationOptionsBuilder struct {
	internal *HeatmapCalculationOptions
	errors   map[string]cog.BuildErrors
}

func NewHeatmapCalculationOptionsBuilder() *HeatmapCalculationOptionsBuilder {
	resource := &HeatmapCalculationOptions{}
	builder := &HeatmapCalculationOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *HeatmapCalculationOptionsBuilder) Build() (HeatmapCalculationOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapCalculationOptions{}, err
	}

	return *builder.internal, nil
}

// The number of buckets to use for the xAxis in the heatmap
func (builder *HeatmapCalculationOptionsBuilder) XBuckets(xBuckets cog.Builder[HeatmapCalculationBucketConfig]) *HeatmapCalculationOptionsBuilder {
	xBucketsResource, err := xBuckets.Build()
	if err != nil {
		builder.errors["xBuckets"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.XBuckets = &xBucketsResource

	return builder
}

// The number of buckets to use for the yAxis in the heatmap
func (builder *HeatmapCalculationOptionsBuilder) YBuckets(yBuckets cog.Builder[HeatmapCalculationBucketConfig]) *HeatmapCalculationOptionsBuilder {
	yBucketsResource, err := yBuckets.Build()
	if err != nil {
		builder.errors["yBuckets"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.YBuckets = &yBucketsResource

	return builder
}

func (builder *HeatmapCalculationOptionsBuilder) applyDefaults() {
}
