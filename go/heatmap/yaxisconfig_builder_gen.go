// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[YAxisConfig] = (*YAxisConfigBuilder)(nil)

// Configuration options for the yAxis
type YAxisConfigBuilder struct {
	internal *YAxisConfig
	errors   map[string]cog.BuildErrors
}

func NewYAxisConfigBuilder() *YAxisConfigBuilder {
	resource := &YAxisConfig{}
	builder := &YAxisConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *YAxisConfigBuilder) Build() (YAxisConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return YAxisConfig{}, err
	}

	return *builder.internal, nil
}

// Sets the yAxis unit
func (builder *YAxisConfigBuilder) Unit(unit string) *YAxisConfigBuilder {
	builder.internal.Unit = &unit

	return builder
}

// Reverses the yAxis
func (builder *YAxisConfigBuilder) Reverse(reverse bool) *YAxisConfigBuilder {
	builder.internal.Reverse = &reverse

	return builder
}

// Controls the number of decimals for yAxis values
func (builder *YAxisConfigBuilder) Decimals(decimals float32) *YAxisConfigBuilder {
	builder.internal.Decimals = &decimals

	return builder
}

// Sets the minimum value for the yAxis
func (builder *YAxisConfigBuilder) Min(min float32) *YAxisConfigBuilder {
	builder.internal.Min = &min

	return builder
}

func (builder *YAxisConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *YAxisConfigBuilder {
	builder.internal.AxisPlacement = &axisPlacement

	return builder
}

func (builder *YAxisConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *YAxisConfigBuilder {
	builder.internal.AxisColorMode = &axisColorMode

	return builder
}

func (builder *YAxisConfigBuilder) AxisLabel(axisLabel string) *YAxisConfigBuilder {
	builder.internal.AxisLabel = &axisLabel

	return builder
}

func (builder *YAxisConfigBuilder) AxisWidth(axisWidth float64) *YAxisConfigBuilder {
	builder.internal.AxisWidth = &axisWidth

	return builder
}

func (builder *YAxisConfigBuilder) AxisSoftMin(axisSoftMin float64) *YAxisConfigBuilder {
	builder.internal.AxisSoftMin = &axisSoftMin

	return builder
}

func (builder *YAxisConfigBuilder) AxisSoftMax(axisSoftMax float64) *YAxisConfigBuilder {
	builder.internal.AxisSoftMax = &axisSoftMax

	return builder
}

func (builder *YAxisConfigBuilder) AxisGridShow(axisGridShow bool) *YAxisConfigBuilder {
	builder.internal.AxisGridShow = &axisGridShow

	return builder
}

func (builder *YAxisConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *YAxisConfigBuilder {
	scaleDistributionResource, err := scaleDistribution.Build()
	if err != nil {
		builder.errors["scaleDistribution"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.ScaleDistribution = &scaleDistributionResource

	return builder
}

func (builder *YAxisConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *YAxisConfigBuilder {
	builder.internal.AxisCenteredZero = &axisCenteredZero

	return builder
}

// Sets the maximum value for the yAxis
func (builder *YAxisConfigBuilder) Max(max float32) *YAxisConfigBuilder {
	builder.internal.Max = &max

	return builder
}

func (builder *YAxisConfigBuilder) AxisBorderShow(axisBorderShow bool) *YAxisConfigBuilder {
	builder.internal.AxisBorderShow = &axisBorderShow

	return builder
}

func (builder *YAxisConfigBuilder) applyDefaults() {
}
