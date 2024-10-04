// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[HeatmapTooltip] = (*HeatmapTooltipBuilder)(nil)

// Controls tooltip options
type HeatmapTooltipBuilder struct {
	internal *HeatmapTooltip
	errors   map[string]cog.BuildErrors
}

func NewHeatmapTooltipBuilder() *HeatmapTooltipBuilder {
	resource := &HeatmapTooltip{}
	builder := &HeatmapTooltipBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *HeatmapTooltipBuilder) Build() (HeatmapTooltip, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("HeatmapTooltip", err)...)
	}

	if len(errs) != 0 {
		return HeatmapTooltip{}, errs
	}

	return *builder.internal, nil
}

// Controls how the tooltip is shown
func (builder *HeatmapTooltipBuilder) Mode(mode common.TooltipDisplayMode) *HeatmapTooltipBuilder {
	builder.internal.Mode = mode

	return builder
}

func (builder *HeatmapTooltipBuilder) MaxHeight(maxHeight float64) *HeatmapTooltipBuilder {
	builder.internal.MaxHeight = &maxHeight

	return builder
}

func (builder *HeatmapTooltipBuilder) MaxWidth(maxWidth float64) *HeatmapTooltipBuilder {
	builder.internal.MaxWidth = &maxWidth

	return builder
}

// Controls if the tooltip shows a histogram of the y-axis values
func (builder *HeatmapTooltipBuilder) YHistogram(yHistogram bool) *HeatmapTooltipBuilder {
	builder.internal.YHistogram = &yHistogram

	return builder
}

// Controls if the tooltip shows a color scale in header
func (builder *HeatmapTooltipBuilder) ShowColorScale(showColorScale bool) *HeatmapTooltipBuilder {
	builder.internal.ShowColorScale = &showColorScale

	return builder
}

func (builder *HeatmapTooltipBuilder) applyDefaults() {
}
