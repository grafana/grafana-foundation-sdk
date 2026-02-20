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
	errors   cog.BuildErrors
}

func NewHeatmapTooltipBuilder() *HeatmapTooltipBuilder {
	resource := NewHeatmapTooltip()
	builder := &HeatmapTooltipBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *HeatmapTooltipBuilder) Build() (HeatmapTooltip, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapTooltip{}, err
	}

	if len(builder.errors) > 0 {
		return HeatmapTooltip{}, cog.MakeBuildErrors("heatmap.heatmapTooltip", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *HeatmapTooltipBuilder) RecordError(path string, err error) *HeatmapTooltipBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
