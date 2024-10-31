// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[LibrarypanelLibraryPanelModel] = (*LibrarypanelLibraryPanelModelBuilder)(nil)

type LibrarypanelLibraryPanelModelBuilder struct {
	internal *LibrarypanelLibraryPanelModel
	errors   map[string]cog.BuildErrors
}

func NewLibrarypanelLibraryPanelModelBuilder() *LibrarypanelLibraryPanelModelBuilder {
	resource := &LibrarypanelLibraryPanelModel{}
	builder := &LibrarypanelLibraryPanelModelBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LibrarypanelLibraryPanelModelBuilder) Build() (LibrarypanelLibraryPanelModel, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibrarypanelLibraryPanelModel{}, err
	}

	return *builder.internal, nil
}

// The panel plugin type id. This is used to find the plugin to display the panel.
func (builder *LibrarypanelLibraryPanelModelBuilder) Type(typeArg string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Type = typeArg

	return builder
}

// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
func (builder *LibrarypanelLibraryPanelModelBuilder) PluginVersion(pluginVersion string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.PluginVersion = &pluginVersion

	return builder
}

// Tags for the panel.
func (builder *LibrarypanelLibraryPanelModelBuilder) Tags(tags []string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Tags = tags

	return builder
}

// Depends on the panel plugin. See the plugin documentation for details.
func (builder *LibrarypanelLibraryPanelModelBuilder) Targets(targets []cog.Builder[variants.Dataquery]) *LibrarypanelLibraryPanelModelBuilder {
	targetsResources := make([]variants.Dataquery, 0, len(targets))
	for _, r1 := range targets {
		targetsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["targets"] = err.(cog.BuildErrors)
			return builder
		}
		targetsResources = append(targetsResources, targetsDepth1)
	}
	builder.internal.Targets = targetsResources

	return builder
}

// Panel title.
func (builder *LibrarypanelLibraryPanelModelBuilder) Title(title string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Title = &title

	return builder
}

// Panel description.
func (builder *LibrarypanelLibraryPanelModelBuilder) Description(description string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Description = &description

	return builder
}

// Whether to display the panel without a background.
func (builder *LibrarypanelLibraryPanelModelBuilder) Transparent(transparent bool) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Transparent = &transparent

	return builder
}

// The datasource used in all targets.
func (builder *LibrarypanelLibraryPanelModelBuilder) Datasource(datasource dashboard.DataSourceRef) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Panel links.
func (builder *LibrarypanelLibraryPanelModelBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *LibrarypanelLibraryPanelModelBuilder {
	linksResources := make([]dashboard.DashboardLink, 0, len(links))
	for _, r1 := range links {
		linksDepth1, err := r1.Build()
		if err != nil {
			builder.errors["links"] = err.(cog.BuildErrors)
			return builder
		}
		linksResources = append(linksResources, linksDepth1)
	}
	builder.internal.Links = linksResources

	return builder
}

// Name of template variable to repeat for.
func (builder *LibrarypanelLibraryPanelModelBuilder) Repeat(repeat string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Repeat = &repeat

	return builder
}

// Direction to repeat in if 'repeat' is set.
// `h` for horizontal, `v` for vertical.
func (builder *LibrarypanelLibraryPanelModelBuilder) RepeatDirection(repeatDirection LibraryPanelRepeatDirection) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.RepeatDirection = &repeatDirection

	return builder
}

// Option for repeated panels that controls max items per row
// Only relevant for horizontally repeated panels
func (builder *LibrarypanelLibraryPanelModelBuilder) MaxPerRow(maxPerRow float64) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.MaxPerRow = &maxPerRow

	return builder
}

// The maximum number of data points that the panel queries are retrieving.
func (builder *LibrarypanelLibraryPanelModelBuilder) MaxDataPoints(maxDataPoints float64) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// List of transformations that are applied to the panel data before rendering.
// When there are multiple transformations, Grafana applies them in the order they are listed.
// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
func (builder *LibrarypanelLibraryPanelModelBuilder) Transformations(transformations []dashboard.DataTransformerConfig) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Transformations = transformations

	return builder
}

// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
// This value must be formatted as a number followed by a valid time
// identifier like: "40s", "3d", etc.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *LibrarypanelLibraryPanelModelBuilder) Interval(interval string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Interval = &interval

	return builder
}

// Overrides the relative time range for individual panels,
// which causes them to be different than what is selected in
// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
// time periods or days on the same dashboard.
// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *LibrarypanelLibraryPanelModelBuilder) TimeFrom(timeFrom string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.TimeFrom = &timeFrom

	return builder
}

// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *LibrarypanelLibraryPanelModelBuilder) TimeShift(timeShift string) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.TimeShift = &timeShift

	return builder
}

// Controls if the timeFrom or timeShift overrides are shown in the panel header
func (builder *LibrarypanelLibraryPanelModelBuilder) HideTimeOverride(hideTimeOverride bool) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.HideTimeOverride = &hideTimeOverride

	return builder
}

// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
func (builder *LibrarypanelLibraryPanelModelBuilder) Options(options any) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.Options = &options

	return builder
}

// Field options allow you to change how the data is displayed in your visualizations.
func (builder *LibrarypanelLibraryPanelModelBuilder) FieldConfig(fieldConfig dashboard.FieldConfigSource) *LibrarypanelLibraryPanelModelBuilder {
	builder.internal.FieldConfig = &fieldConfig

	return builder
}

func (builder *LibrarypanelLibraryPanelModelBuilder) applyDefaults() {
	builder.Transparent(false)
}
