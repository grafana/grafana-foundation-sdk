// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HeatmapLegend] = (*HeatmapLegendBuilder)(nil)

// Controls legend options
type HeatmapLegendBuilder struct {
	internal *HeatmapLegend
	errors   cog.BuildErrors
}

func NewHeatmapLegendBuilder() *HeatmapLegendBuilder {
	resource := NewHeatmapLegend()
	builder := &HeatmapLegendBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *HeatmapLegendBuilder) Build() (HeatmapLegend, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapLegend{}, err
	}

	if len(builder.errors) > 0 {
		return HeatmapLegend{}, cog.MakeBuildErrors("heatmap.heatmapLegend", builder.errors)
	}

	return *builder.internal, nil
}

// Controls if the legend is shown
func (builder *HeatmapLegendBuilder) Show(show bool) *HeatmapLegendBuilder {
	builder.internal.Show = show

	return builder
}
