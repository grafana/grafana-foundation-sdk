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
	errors   map[string]cog.BuildErrors
}

func NewGraphPanelBuilder() *GraphPanelBuilder {
	resource := &GraphPanel{}
	builder := &GraphPanelBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "graph"

	return builder
}

func (builder *GraphPanelBuilder) Build() (GraphPanel, error) {
	if err := builder.internal.Validate(); err != nil {
		return GraphPanel{}, err
	}

	return *builder.internal, nil
}

// @deprecated this is part of deprecated graph panel
func (builder *GraphPanelBuilder) Legend(legend cog.Builder[DashboardGraphPanelLegend]) *GraphPanelBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors["legend"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Legend = &legendResource

	return builder
}

func (builder *GraphPanelBuilder) applyDefaults() {
}
