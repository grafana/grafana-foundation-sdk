// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardGraphPanelLegend] = (*DashboardGraphPanelLegendBuilder)(nil)

type DashboardGraphPanelLegendBuilder struct {
	internal *DashboardGraphPanelLegend
	errors   map[string]cog.BuildErrors
}

func NewDashboardGraphPanelLegendBuilder() *DashboardGraphPanelLegendBuilder {
	resource := &DashboardGraphPanelLegend{}
	builder := &DashboardGraphPanelLegendBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *DashboardGraphPanelLegendBuilder) Build() (DashboardGraphPanelLegend, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardGraphPanelLegend{}, err
	}

	return *builder.internal, nil
}

func (builder *DashboardGraphPanelLegendBuilder) Show(show bool) *DashboardGraphPanelLegendBuilder {
	builder.internal.Show = show

	return builder
}

func (builder *DashboardGraphPanelLegendBuilder) Sort(sort string) *DashboardGraphPanelLegendBuilder {
	builder.internal.Sort = &sort

	return builder
}

func (builder *DashboardGraphPanelLegendBuilder) SortDesc(sortDesc bool) *DashboardGraphPanelLegendBuilder {
	builder.internal.SortDesc = &sortDesc

	return builder
}

func (builder *DashboardGraphPanelLegendBuilder) applyDefaults() {
	builder.Show(true)
}
