// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[ScatterSeriesConfig] = (*ScatterSeriesConfigBuilder)(nil)

type ScatterSeriesConfigBuilder struct {
	internal *ScatterSeriesConfig
	errors   map[string]cog.BuildErrors
}

func NewScatterSeriesConfigBuilder() *ScatterSeriesConfigBuilder {
	resource := &ScatterSeriesConfig{}
	builder := &ScatterSeriesConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ScatterSeriesConfigBuilder) Build() (ScatterSeriesConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return ScatterSeriesConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *ScatterSeriesConfigBuilder) X(x string) *ScatterSeriesConfigBuilder {
	builder.internal.X = &x

	return builder
}

func (builder *ScatterSeriesConfigBuilder) Y(y string) *ScatterSeriesConfigBuilder {
	builder.internal.Y = &y

	return builder
}

func (builder *ScatterSeriesConfigBuilder) Show(show ScatterShow) *ScatterSeriesConfigBuilder {
	builder.internal.Show = &show

	return builder
}

func (builder *ScatterSeriesConfigBuilder) PointSize(pointSize cog.Builder[common.ScaleDimensionConfig]) *ScatterSeriesConfigBuilder {
	pointSizeResource, err := pointSize.Build()
	if err != nil {
		builder.errors["pointSize"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.PointSize = &pointSizeResource

	return builder
}

func (builder *ScatterSeriesConfigBuilder) PointColor(pointColor cog.Builder[common.ColorDimensionConfig]) *ScatterSeriesConfigBuilder {
	pointColorResource, err := pointColor.Build()
	if err != nil {
		builder.errors["pointColor"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.PointColor = &pointColorResource

	return builder
}

func (builder *ScatterSeriesConfigBuilder) LineColor(lineColor cog.Builder[common.ColorDimensionConfig]) *ScatterSeriesConfigBuilder {
	lineColorResource, err := lineColor.Build()
	if err != nil {
		builder.errors["lineColor"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.LineColor = &lineColorResource

	return builder
}

func (builder *ScatterSeriesConfigBuilder) LineWidth(lineWidth int32) *ScatterSeriesConfigBuilder {
	builder.internal.LineWidth = &lineWidth

	return builder
}

func (builder *ScatterSeriesConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *ScatterSeriesConfigBuilder {
	lineStyleResource, err := lineStyle.Build()
	if err != nil {
		builder.errors["lineStyle"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.LineStyle = &lineStyleResource

	return builder
}

func (builder *ScatterSeriesConfigBuilder) Label(label common.VisibilityMode) *ScatterSeriesConfigBuilder {
	builder.internal.Label = &label

	return builder
}

func (builder *ScatterSeriesConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *ScatterSeriesConfigBuilder {
	hideFromResource, err := hideFrom.Build()
	if err != nil {
		builder.errors["hideFrom"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.HideFrom = &hideFromResource

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *ScatterSeriesConfigBuilder {
	builder.internal.AxisPlacement = &axisPlacement

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *ScatterSeriesConfigBuilder {
	builder.internal.AxisColorMode = &axisColorMode

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisLabel(axisLabel string) *ScatterSeriesConfigBuilder {
	builder.internal.AxisLabel = &axisLabel

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisWidth(axisWidth float64) *ScatterSeriesConfigBuilder {
	builder.internal.AxisWidth = &axisWidth

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisSoftMin(axisSoftMin float64) *ScatterSeriesConfigBuilder {
	builder.internal.AxisSoftMin = &axisSoftMin

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisSoftMax(axisSoftMax float64) *ScatterSeriesConfigBuilder {
	builder.internal.AxisSoftMax = &axisSoftMax

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisGridShow(axisGridShow bool) *ScatterSeriesConfigBuilder {
	builder.internal.AxisGridShow = &axisGridShow

	return builder
}

func (builder *ScatterSeriesConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *ScatterSeriesConfigBuilder {
	scaleDistributionResource, err := scaleDistribution.Build()
	if err != nil {
		builder.errors["scaleDistribution"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.ScaleDistribution = &scaleDistributionResource

	return builder
}

func (builder *ScatterSeriesConfigBuilder) Name(name string) *ScatterSeriesConfigBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *ScatterSeriesConfigBuilder) LabelValue(labelValue cog.Builder[common.TextDimensionConfig]) *ScatterSeriesConfigBuilder {
	labelValueResource, err := labelValue.Build()
	if err != nil {
		builder.errors["labelValue"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.LabelValue = &labelValueResource

	return builder
}

func (builder *ScatterSeriesConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *ScatterSeriesConfigBuilder {
	builder.internal.AxisCenteredZero = &axisCenteredZero

	return builder
}

func (builder *ScatterSeriesConfigBuilder) applyDefaults() {
	builder.Show("points")
	builder.Label("auto")
}
