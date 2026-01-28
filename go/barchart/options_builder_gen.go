// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package barchart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[Options] = (*OptionsBuilder)(nil)

type OptionsBuilder struct {
	internal *Options
	errors   cog.BuildErrors
}

func NewOptionsBuilder() *OptionsBuilder {
	resource := NewOptions()
	builder := &OptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsBuilder) Build() (Options, error) {
	if err := builder.internal.Validate(); err != nil {
		return Options{}, err
	}

	if len(builder.errors) > 0 {
		return Options{}, cog.MakeBuildErrors("barchart.options", builder.errors)
	}

	return *builder.internal, nil
}

// Manually select which field from the dataset to represent the x field.
func (builder *OptionsBuilder) XField(xField string) *OptionsBuilder {
	builder.internal.XField = &xField

	return builder
}

// Use the color value for a sibling field to color each bar value.
func (builder *OptionsBuilder) ColorByField(colorByField string) *OptionsBuilder {
	builder.internal.ColorByField = &colorByField

	return builder
}

// Controls the orientation of the bar chart, either vertical or horizontal.
func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder {
	builder.internal.Orientation = orientation

	return builder
}

// Controls the radius of each bar.
func (builder *OptionsBuilder) BarRadius(barRadius float64) *OptionsBuilder {
	builder.internal.BarRadius = &barRadius

	return builder
}

// Controls the rotation of the x axis labels.
func (builder *OptionsBuilder) XTickLabelRotation(xTickLabelRotation int32) *OptionsBuilder {
	builder.internal.XTickLabelRotation = xTickLabelRotation

	return builder
}

// Sets the max length that a label can have before it is truncated.
func (builder *OptionsBuilder) XTickLabelMaxLength(xTickLabelMaxLength int32) *OptionsBuilder {
	builder.internal.XTickLabelMaxLength = xTickLabelMaxLength

	return builder
}

// Controls the spacing between x axis labels.
// negative values indicate backwards skipping behavior
func (builder *OptionsBuilder) XTickLabelSpacing(xTickLabelSpacing int32) *OptionsBuilder {
	builder.internal.XTickLabelSpacing = &xTickLabelSpacing

	return builder
}

// Controls whether bars are stacked or not, either normally or in percent mode.
func (builder *OptionsBuilder) Stacking(stacking common.StackingMode) *OptionsBuilder {
	builder.internal.Stacking = stacking

	return builder
}

// This controls whether values are shown on top or to the left of bars.
func (builder *OptionsBuilder) ShowValue(showValue common.VisibilityMode) *OptionsBuilder {
	builder.internal.ShowValue = showValue

	return builder
}

// Controls the width of bars. 1 = Max width, 0 = Min width.
func (builder *OptionsBuilder) BarWidth(barWidth float64) *OptionsBuilder {
	builder.internal.BarWidth = barWidth

	return builder
}

// Controls the width of groups. 1 = max with, 0 = min width.
func (builder *OptionsBuilder) GroupWidth(groupWidth float64) *OptionsBuilder {
	builder.internal.GroupWidth = groupWidth

	return builder
}

func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Legend = legendResource

	return builder
}

func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder {
	tooltipResource, err := tooltip.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Tooltip = tooltipResource

	return builder
}

func (builder *OptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *OptionsBuilder {
	textResource, err := text.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Text = &textResource

	return builder
}

// Enables mode which highlights the entire bar area and shows tooltip when cursor
// hovers over highlighted area
func (builder *OptionsBuilder) FullHighlight(fullHighlight bool) *OptionsBuilder {
	builder.internal.FullHighlight = fullHighlight

	return builder
}
