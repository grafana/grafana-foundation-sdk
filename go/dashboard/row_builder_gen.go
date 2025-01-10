// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RowPanel] = (*RowBuilder)(nil)

// Row panel
type RowBuilder struct {
	internal *RowPanel
	errors   map[string]cog.BuildErrors
}

func NewRowBuilder(title string) *RowBuilder {
	resource := NewRowPanel()
	builder := &RowBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	builder.internal.Type = "row"
	builder.internal.Title = &title

	return builder
}

func (builder *RowBuilder) Build() (RowPanel, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowPanel{}, err
	}

	return *builder.internal, nil
}

// Whether this row should be collapsed or not.
// Note: panels added directly to a row will be stripped by Grafana unless the row is collapsed
func (builder *RowBuilder) Collapsed(collapsed bool) *RowBuilder {
	builder.internal.Collapsed = collapsed

	return builder
}

// Row title
func (builder *RowBuilder) Title(title string) *RowBuilder {
	builder.internal.Title = &title

	return builder
}

// Name of default datasource for the row
func (builder *RowBuilder) Datasource(datasource DataSourceRef) *RowBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Row grid position
func (builder *RowBuilder) GridPos(gridPos GridPos) *RowBuilder {
	builder.internal.GridPos = &gridPos

	return builder
}

// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
func (builder *RowBuilder) Id(id uint32) *RowBuilder {
	builder.internal.Id = id

	return builder
}

// List of panels in the row
// Note: since panels added directly to a row will be stripped by Grafana unless the row is collapsed,
// this option will set the current row as collapsed.
func (builder *RowBuilder) WithPanel(panel cog.Builder[Panel]) *RowBuilder {
	panelResource, err := panel.Build()
	if err != nil {
		builder.errors["panels"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Panels = append(builder.internal.Panels, panelResource)
	builder.internal.Collapsed = true

	return builder
}

// Name of template variable to repeat for.
func (builder *RowBuilder) Repeat(repeat string) *RowBuilder {
	builder.internal.Repeat = &repeat

	return builder
}
