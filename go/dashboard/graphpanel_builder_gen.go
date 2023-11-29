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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("GraphPanel", err)...)
	}

	if len(errs) != 0 {
		return GraphPanel{}, errs
	}

	return *builder.internal, nil
}

// @deprecated this is part of deprecated graph panel
func (builder *GraphPanelBuilder) Legend(legend struct {
	Show     bool    `json:"show"`
	Sort     *string `json:"sort,omitempty"`
	SortDesc *bool   `json:"sortDesc,omitempty"`
}) *GraphPanelBuilder {
	builder.internal.Legend = &legend

	return builder
}

func (builder *GraphPanelBuilder) applyDefaults() {
}
