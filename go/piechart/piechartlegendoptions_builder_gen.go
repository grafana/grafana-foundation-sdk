// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package piechart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[PieChartLegendOptions] = (*PieChartLegendOptionsBuilder)(nil)

type PieChartLegendOptionsBuilder struct {
	internal *PieChartLegendOptions
	errors   map[string]cog.BuildErrors
}

func NewPieChartLegendOptionsBuilder() *PieChartLegendOptionsBuilder {
	resource := &PieChartLegendOptions{}
	builder := &PieChartLegendOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PieChartLegendOptionsBuilder) Build() (PieChartLegendOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return PieChartLegendOptions{}, err
	}

	return *builder.internal, nil
}

func (builder *PieChartLegendOptionsBuilder) Values(values []PieChartLegendValues) *PieChartLegendOptionsBuilder {
	builder.internal.Values = values

	return builder
}

func (builder *PieChartLegendOptionsBuilder) DisplayMode(displayMode common.LegendDisplayMode) *PieChartLegendOptionsBuilder {
	builder.internal.DisplayMode = displayMode

	return builder
}

func (builder *PieChartLegendOptionsBuilder) Placement(placement common.LegendPlacement) *PieChartLegendOptionsBuilder {
	builder.internal.Placement = placement

	return builder
}

func (builder *PieChartLegendOptionsBuilder) ShowLegend(showLegend bool) *PieChartLegendOptionsBuilder {
	builder.internal.ShowLegend = showLegend

	return builder
}

func (builder *PieChartLegendOptionsBuilder) AsTable(asTable bool) *PieChartLegendOptionsBuilder {
	builder.internal.AsTable = &asTable

	return builder
}

func (builder *PieChartLegendOptionsBuilder) IsVisible(isVisible bool) *PieChartLegendOptionsBuilder {
	builder.internal.IsVisible = &isVisible

	return builder
}

func (builder *PieChartLegendOptionsBuilder) SortBy(sortBy string) *PieChartLegendOptionsBuilder {
	builder.internal.SortBy = &sortBy

	return builder
}

func (builder *PieChartLegendOptionsBuilder) SortDesc(sortDesc bool) *PieChartLegendOptionsBuilder {
	builder.internal.SortDesc = &sortDesc

	return builder
}

func (builder *PieChartLegendOptionsBuilder) Width(width float64) *PieChartLegendOptionsBuilder {
	builder.internal.Width = &width

	return builder
}

func (builder *PieChartLegendOptionsBuilder) Calcs(calcs []string) *PieChartLegendOptionsBuilder {
	builder.internal.Calcs = calcs

	return builder
}

func (builder *PieChartLegendOptionsBuilder) applyDefaults() {
}
