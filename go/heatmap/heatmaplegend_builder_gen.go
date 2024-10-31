// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HeatmapLegend] = (*HeatmapLegendBuilder)(nil)

// Controls legend options
type HeatmapLegendBuilder struct {
	internal *HeatmapLegend
	errors   map[string]cog.BuildErrors
}

func NewHeatmapLegendBuilder() *HeatmapLegendBuilder {
	resource := &HeatmapLegend{}
	builder := &HeatmapLegendBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *HeatmapLegendBuilder) Build() (HeatmapLegend, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapLegend{}, err
	}

	return *builder.internal, nil
}

// Controls if the legend is shown
func (builder *HeatmapLegendBuilder) Show(show bool) *HeatmapLegendBuilder {
	builder.internal.Show = show

	return builder
}

func (builder *HeatmapLegendBuilder) applyDefaults() {
}
