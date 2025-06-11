// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardGraphPanelLegend] = (*DashboardGraphPanelLegendBuilder)(nil)

type DashboardGraphPanelLegendBuilder struct {
	internal *DashboardGraphPanelLegend
	errors   cog.BuildErrors
}

func NewDashboardGraphPanelLegendBuilder() *DashboardGraphPanelLegendBuilder {
	resource := NewDashboardGraphPanelLegend()
	builder := &DashboardGraphPanelLegendBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DashboardGraphPanelLegendBuilder) Build() (DashboardGraphPanelLegend, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardGraphPanelLegend{}, err
	}

	if len(builder.errors) > 0 {
		return DashboardGraphPanelLegend{}, cog.MakeBuildErrors("dashboard.dashboardGraphPanelLegend", builder.errors)
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
