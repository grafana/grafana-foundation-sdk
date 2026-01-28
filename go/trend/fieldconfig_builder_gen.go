// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package trend

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[FieldConfig] = (*FieldConfigBuilder)(nil)

type FieldConfigBuilder struct {
	internal *FieldConfig
	errors   cog.BuildErrors
}

func NewFieldConfigBuilder() *FieldConfigBuilder {
	resource := NewFieldConfig()
	builder := &FieldConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FieldConfigBuilder) Build() (FieldConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return FieldConfig{}, err
	}

	if len(builder.errors) > 0 {
		return FieldConfig{}, cog.MakeBuildErrors("trend.fieldConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FieldConfigBuilder) DrawStyle(drawStyle common.GraphDrawStyle) *FieldConfigBuilder {
	builder.internal.DrawStyle = &drawStyle

	return builder
}

func (builder *FieldConfigBuilder) GradientMode(gradientMode common.GraphGradientMode) *FieldConfigBuilder {
	builder.internal.GradientMode = &gradientMode

	return builder
}

func (builder *FieldConfigBuilder) ThresholdsStyle(thresholdsStyle cog.Builder[common.GraphThresholdsStyleConfig]) *FieldConfigBuilder {
	thresholdsStyleResource, err := thresholdsStyle.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ThresholdsStyle = &thresholdsStyleResource

	return builder
}

func (builder *FieldConfigBuilder) LineColor(lineColor string) *FieldConfigBuilder {
	builder.internal.LineColor = &lineColor

	return builder
}

func (builder *FieldConfigBuilder) LineWidth(lineWidth float64) *FieldConfigBuilder {
	builder.internal.LineWidth = &lineWidth

	return builder
}

func (builder *FieldConfigBuilder) LineInterpolation(lineInterpolation common.LineInterpolation) *FieldConfigBuilder {
	builder.internal.LineInterpolation = &lineInterpolation

	return builder
}

func (builder *FieldConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *FieldConfigBuilder {
	lineStyleResource, err := lineStyle.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.LineStyle = &lineStyleResource

	return builder
}

func (builder *FieldConfigBuilder) FillColor(fillColor string) *FieldConfigBuilder {
	builder.internal.FillColor = &fillColor

	return builder
}

func (builder *FieldConfigBuilder) FillOpacity(fillOpacity float64) *FieldConfigBuilder {
	builder.internal.FillOpacity = &fillOpacity

	return builder
}

func (builder *FieldConfigBuilder) ShowPoints(showPoints common.VisibilityMode) *FieldConfigBuilder {
	builder.internal.ShowPoints = &showPoints

	return builder
}

func (builder *FieldConfigBuilder) PointSize(pointSize float64) *FieldConfigBuilder {
	builder.internal.PointSize = &pointSize

	return builder
}

func (builder *FieldConfigBuilder) PointColor(pointColor string) *FieldConfigBuilder {
	builder.internal.PointColor = &pointColor

	return builder
}

func (builder *FieldConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *FieldConfigBuilder {
	builder.internal.AxisPlacement = &axisPlacement

	return builder
}

func (builder *FieldConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *FieldConfigBuilder {
	builder.internal.AxisColorMode = &axisColorMode

	return builder
}

func (builder *FieldConfigBuilder) AxisLabel(axisLabel string) *FieldConfigBuilder {
	builder.internal.AxisLabel = &axisLabel

	return builder
}

func (builder *FieldConfigBuilder) AxisWidth(axisWidth float64) *FieldConfigBuilder {
	builder.internal.AxisWidth = &axisWidth

	return builder
}

func (builder *FieldConfigBuilder) AxisSoftMin(axisSoftMin float64) *FieldConfigBuilder {
	builder.internal.AxisSoftMin = &axisSoftMin

	return builder
}

func (builder *FieldConfigBuilder) AxisSoftMax(axisSoftMax float64) *FieldConfigBuilder {
	builder.internal.AxisSoftMax = &axisSoftMax

	return builder
}

func (builder *FieldConfigBuilder) AxisGridShow(axisGridShow bool) *FieldConfigBuilder {
	builder.internal.AxisGridShow = &axisGridShow

	return builder
}

func (builder *FieldConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *FieldConfigBuilder {
	scaleDistributionResource, err := scaleDistribution.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ScaleDistribution = &scaleDistributionResource

	return builder
}

func (builder *FieldConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *FieldConfigBuilder {
	builder.internal.AxisCenteredZero = &axisCenteredZero

	return builder
}

func (builder *FieldConfigBuilder) BarAlignment(barAlignment common.BarAlignment) *FieldConfigBuilder {
	builder.internal.BarAlignment = &barAlignment

	return builder
}

func (builder *FieldConfigBuilder) BarWidthFactor(barWidthFactor float64) *FieldConfigBuilder {
	builder.internal.BarWidthFactor = &barWidthFactor

	return builder
}

func (builder *FieldConfigBuilder) Stacking(stacking cog.Builder[common.StackingConfig]) *FieldConfigBuilder {
	stackingResource, err := stacking.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Stacking = &stackingResource

	return builder
}

func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder {
	hideFromResource, err := hideFrom.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.HideFrom = &hideFromResource

	return builder
}

func (builder *FieldConfigBuilder) Transform(transform common.GraphTransform) *FieldConfigBuilder {
	builder.internal.Transform = &transform

	return builder
}

// Indicate if null values should be treated as gaps or connected.
// When the value is a number, it represents the maximum delta in the
// X axis that should be considered connected.  For timeseries, this is milliseconds
func (builder *FieldConfigBuilder) SpanNulls(spanNulls common.BoolOrFloat64) *FieldConfigBuilder {
	builder.internal.SpanNulls = &spanNulls

	return builder
}

func (builder *FieldConfigBuilder) FillBelowTo(fillBelowTo string) *FieldConfigBuilder {
	builder.internal.FillBelowTo = &fillBelowTo

	return builder
}

func (builder *FieldConfigBuilder) PointSymbol(pointSymbol string) *FieldConfigBuilder {
	builder.internal.PointSymbol = &pointSymbol

	return builder
}

func (builder *FieldConfigBuilder) AxisBorderShow(axisBorderShow bool) *FieldConfigBuilder {
	builder.internal.AxisBorderShow = &axisBorderShow

	return builder
}

func (builder *FieldConfigBuilder) BarMaxWidth(barMaxWidth float64) *FieldConfigBuilder {
	builder.internal.BarMaxWidth = &barMaxWidth

	return builder
}

func (builder *FieldConfigBuilder) InsertNulls(insertNulls common.BoolOrUint32) *FieldConfigBuilder {
	builder.internal.InsertNulls = &insertNulls

	return builder
}
