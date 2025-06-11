// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GraphPanel] = (*GraphPanelBuilder)(nil)

// Support for legacy graph panel.
// @deprecated this a deprecated panel type
type GraphPanelBuilder struct {
	internal *GraphPanel
	errors   cog.BuildErrors
}

func NewGraphPanelBuilder() *GraphPanelBuilder {
	resource := NewGraphPanel()
	builder := &GraphPanelBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Type = "graph"

	return builder
}

func (builder *GraphPanelBuilder) Build() (GraphPanel, error) {
	if err := builder.internal.Validate(); err != nil {
		return GraphPanel{}, err
	}

	if len(builder.errors) > 0 {
		return GraphPanel{}, cog.MakeBuildErrors("dashboard.graphPanel", builder.errors)
	}

	return *builder.internal, nil
}

// @deprecated this is part of deprecated graph panel
func (builder *GraphPanelBuilder) Legend(legend cog.Builder[DashboardGraphPanelLegend]) *GraphPanelBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Legend = &legendResource

	return builder
}
