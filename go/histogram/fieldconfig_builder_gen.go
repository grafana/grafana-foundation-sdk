// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package histogram

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
		return FieldConfig{}, cog.MakeBuildErrors("histogram.fieldConfig", builder.errors)
	}

	return *builder.internal, nil
}

// Controls line width of the bars.
func (builder *FieldConfigBuilder) LineWidth(lineWidth uint32) *FieldConfigBuilder {
	builder.internal.LineWidth = &lineWidth

	return builder
}

// Controls the fill opacity of the bars.
func (builder *FieldConfigBuilder) FillOpacity(fillOpacity uint32) *FieldConfigBuilder {
	builder.internal.FillOpacity = &fillOpacity

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

func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder {
	hideFromResource, err := hideFrom.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.HideFrom = &hideFromResource

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

// Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
// Gradient appearance is influenced by the Fill opacity setting.
func (builder *FieldConfigBuilder) GradientMode(gradientMode common.GraphGradientMode) *FieldConfigBuilder {
	builder.internal.GradientMode = &gradientMode

	return builder
}

func (builder *FieldConfigBuilder) AxisBorderShow(axisBorderShow bool) *FieldConfigBuilder {
	builder.internal.AxisBorderShow = &axisBorderShow

	return builder
}
