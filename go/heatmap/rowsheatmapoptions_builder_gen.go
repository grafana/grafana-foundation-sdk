// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[RowsHeatmapOptions] = (*RowsHeatmapOptionsBuilder)(nil)

// Controls frame rows options
type RowsHeatmapOptionsBuilder struct {
	internal *RowsHeatmapOptions
	errors   map[string]cog.BuildErrors
}

func NewRowsHeatmapOptionsBuilder() *RowsHeatmapOptionsBuilder {
	resource := &RowsHeatmapOptions{}
	builder := &RowsHeatmapOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *RowsHeatmapOptionsBuilder) Build() (RowsHeatmapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowsHeatmapOptions{}, err
	}

	return *builder.internal, nil
}

// Sets the name of the cell when not calculating from data
func (builder *RowsHeatmapOptionsBuilder) Value(value string) *RowsHeatmapOptionsBuilder {
	builder.internal.Value = &value

	return builder
}

// Controls tick alignment when not calculating from data
func (builder *RowsHeatmapOptionsBuilder) Layout(layout common.HeatmapCellLayout) *RowsHeatmapOptionsBuilder {
	builder.internal.Layout = &layout

	return builder
}

func (builder *RowsHeatmapOptionsBuilder) applyDefaults() {
}
