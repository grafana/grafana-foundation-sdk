// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboard] = (*DashboardBuilder)(nil)

type DashboardBuilder struct {
	internal *Dashboard
	errors   cog.BuildErrors
}

func NewDashboardBuilder(title string) *DashboardBuilder {
	resource := NewDashboard()
	builder := &DashboardBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Title = title

	return builder
}

func (builder *DashboardBuilder) Build() (Dashboard, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboard{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboard{}, cog.MakeBuildErrors("dashboardv2beta1.dashboard", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DashboardBuilder) Annotations(annotations []cog.Builder[AnnotationQueryKind]) *DashboardBuilder {
	annotationsResources := make([]AnnotationQueryKind, 0, len(annotations))
	for _, r1 := range annotations {
		annotationsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		annotationsResources = append(annotationsResources, annotationsDepth1)
	}
	builder.internal.Annotations = annotationsResources

	return builder
}

// Configuration of dashboard cursor sync behavior.
// "Off" for no shared crosshair or tooltip (default).
// "Crosshair" for shared crosshair.
// "Tooltip" for shared crosshair AND shared tooltip.
func (builder *DashboardBuilder) CursorSync(cursorSync DashboardCursorSync) *DashboardBuilder {
	builder.internal.CursorSync = cursorSync

	return builder
}

// Description of dashboard.
func (builder *DashboardBuilder) Description(description string) *DashboardBuilder {
	builder.internal.Description = &description

	return builder
}

// Whether a dashboard is editable or not.
func (builder *DashboardBuilder) Editable(editable bool) *DashboardBuilder {
	builder.internal.Editable = &editable

	return builder
}

func (builder *DashboardBuilder) Elements(elements map[string]Element) *DashboardBuilder {
	builder.internal.Elements = elements

	return builder
}

func (builder *DashboardBuilder) Panel(key string, panelKind cog.Builder[PanelKind]) *DashboardBuilder {
	if builder.internal.Elements == nil {
		builder.internal.Elements = map[string]Element{}
	}
	panelKindResource, err := panelKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Elements[key] = Element{
		PanelKind: &panelKindResource,
	}

	return builder
}

func (builder *DashboardBuilder) LibraryPanel(key string, libraryPanelKind cog.Builder[LibraryPanelKind]) *DashboardBuilder {
	if builder.internal.Elements == nil {
		builder.internal.Elements = map[string]Element{}
	}
	libraryPanelKindResource, err := libraryPanelKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Elements[key] = Element{
		LibraryPanelKind: &libraryPanelKindResource,
	}

	return builder
}

func (builder *DashboardBuilder) GridLayout(gridLayoutKind cog.Builder[GridLayoutKind]) *DashboardBuilder {
	gridLayoutKindResource, err := gridLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		GridLayoutKind: &gridLayoutKindResource,
	}

	return builder
}

func (builder *DashboardBuilder) RowsLayout(rowsLayoutKind cog.Builder[RowsLayoutKind]) *DashboardBuilder {
	rowsLayoutKindResource, err := rowsLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		RowsLayoutKind: &rowsLayoutKindResource,
	}

	return builder
}

func (builder *DashboardBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[AutoGridLayoutKind]) *DashboardBuilder {
	autoGridLayoutKindResource, err := autoGridLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		AutoGridLayoutKind: &autoGridLayoutKindResource,
	}

	return builder
}

func (builder *DashboardBuilder) TabsLayout(tabsLayoutKind cog.Builder[TabsLayoutKind]) *DashboardBuilder {
	tabsLayoutKindResource, err := tabsLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		TabsLayoutKind: &tabsLayoutKindResource,
	}

	return builder
}

// Links with references to other dashboards or external websites.
func (builder *DashboardBuilder) Links(links []cog.Builder[DashboardLink]) *DashboardBuilder {
	linksResources := make([]DashboardLink, 0, len(links))
	for _, r1 := range links {
		linksDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		linksResources = append(linksResources, linksDepth1)
	}
	builder.internal.Links = linksResources

	return builder
}

// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
// This will keep data "moving left" regardless of the query refresh rate. This setting helps
// avoid dashboards presenting stale live data.
func (builder *DashboardBuilder) LiveNow(liveNow bool) *DashboardBuilder {
	builder.internal.LiveNow = &liveNow

	return builder
}

// When set to true, the dashboard will load all panels in the dashboard when it's loaded.
func (builder *DashboardBuilder) Preload(preload bool) *DashboardBuilder {
	builder.internal.Preload = preload

	return builder
}

// Plugins only. The version of the dashboard installed together with the plugin.
// This is used to determine if the dashboard should be updated when the plugin is updated.
func (builder *DashboardBuilder) Revision(revision uint16) *DashboardBuilder {
	builder.internal.Revision = &revision

	return builder
}

// Tags associated with dashboard.
func (builder *DashboardBuilder) Tags(tags []string) *DashboardBuilder {
	builder.internal.Tags = tags

	return builder
}

func (builder *DashboardBuilder) TimeSettings(timeSettings cog.Builder[TimeSettingsSpec]) *DashboardBuilder {
	timeSettingsResource, err := timeSettings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TimeSettings = timeSettingsResource

	return builder
}

// Title of dashboard.
func (builder *DashboardBuilder) Title(title string) *DashboardBuilder {
	builder.internal.Title = title

	return builder
}

// Configured template variables.
func (builder *DashboardBuilder) Variables(variables []VariableKind) *DashboardBuilder {
	builder.internal.Variables = variables

	return builder
}

func (builder *DashboardBuilder) QueryVariable(queryVariableKind cog.Builder[QueryVariableKind]) *DashboardBuilder {
	queryVariableKindResource, err := queryVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		QueryVariableKind: &queryVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) TextVariable(textVariableKind cog.Builder[TextVariableKind]) *DashboardBuilder {
	textVariableKindResource, err := textVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		TextVariableKind: &textVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) ConstantVariable(constantVariableKind cog.Builder[ConstantVariableKind]) *DashboardBuilder {
	constantVariableKindResource, err := constantVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		ConstantVariableKind: &constantVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) DatasourceVariable(datasourceVariableKind cog.Builder[DatasourceVariableKind]) *DashboardBuilder {
	datasourceVariableKindResource, err := datasourceVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		DatasourceVariableKind: &datasourceVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) IntervalVariable(intervalVariableKind cog.Builder[IntervalVariableKind]) *DashboardBuilder {
	intervalVariableKindResource, err := intervalVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		IntervalVariableKind: &intervalVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) CustomVariable(customVariableKind cog.Builder[CustomVariableKind]) *DashboardBuilder {
	customVariableKindResource, err := customVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		CustomVariableKind: &customVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) GroupByVariable(groupByVariableKind cog.Builder[GroupByVariableKind]) *DashboardBuilder {
	groupByVariableKindResource, err := groupByVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		GroupByVariableKind: &groupByVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) AdhocVariable(adhocVariableKind cog.Builder[AdhocVariableKind]) *DashboardBuilder {
	adhocVariableKindResource, err := adhocVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		AdhocVariableKind: &adhocVariableKindResource,
	})

	return builder
}

func (builder *DashboardBuilder) SwitchVariableKind(switchVariableKind cog.Builder[SwitchVariableKind]) *DashboardBuilder {
	switchVariableKindResource, err := switchVariableKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Variables = append(builder.internal.Variables, VariableKind{
		SwitchVariableKind: &switchVariableKindResource,
	})

	return builder
}
