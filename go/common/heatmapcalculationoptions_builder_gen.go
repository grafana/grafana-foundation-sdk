// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HeatmapCalculationOptions] = (*HeatmapCalculationOptionsBuilder)(nil)

type HeatmapCalculationOptionsBuilder struct {
	internal *HeatmapCalculationOptions
	errors   cog.BuildErrors
}

func NewHeatmapCalculationOptionsBuilder() *HeatmapCalculationOptionsBuilder {
	resource := NewHeatmapCalculationOptions()
	builder := &HeatmapCalculationOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *HeatmapCalculationOptionsBuilder) Build() (HeatmapCalculationOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapCalculationOptions{}, err
	}

	if len(builder.errors) > 0 {
		return HeatmapCalculationOptions{}, cog.MakeBuildErrors("common.heatmapCalculationOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *HeatmapCalculationOptionsBuilder) RecordError(path string, err error) *HeatmapCalculationOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The number of buckets to use for the xAxis in the heatmap
func (builder *HeatmapCalculationOptionsBuilder) XBuckets(xBuckets cog.Builder[HeatmapCalculationBucketConfig]) *HeatmapCalculationOptionsBuilder {
	xBucketsResource, err := xBuckets.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.XBuckets = &xBucketsResource

	return builder
}

// The number of buckets to use for the yAxis in the heatmap
func (builder *HeatmapCalculationOptionsBuilder) YBuckets(yBuckets cog.Builder[HeatmapCalculationBucketConfig]) *HeatmapCalculationOptionsBuilder {
	yBucketsResource, err := yBuckets.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.YBuckets = &yBucketsResource

	return builder
}
