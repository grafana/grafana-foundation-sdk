// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[PanelModel] = (*PanelModelBuilder)(nil)

// Dashboard panels are the basic visualization building blocks.
type PanelModelBuilder struct {
	internal *PanelModel
	errors   cog.BuildErrors
}

func NewPanelModelBuilder() *PanelModelBuilder {
	resource := NewPanelModel()
	builder := &PanelModelBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PanelModelBuilder) Build() (PanelModel, error) {
	if err := builder.internal.Validate(); err != nil {
		return PanelModel{}, err
	}

	if len(builder.errors) > 0 {
		return PanelModel{}, cog.MakeBuildErrors("librarypanel.panelModel", builder.errors)
	}

	return *builder.internal, nil
}

// The panel plugin type id. This is used to find the plugin to display the panel.
func (builder *PanelModelBuilder) Type(typeArg string) *PanelModelBuilder {
	builder.internal.Type = typeArg

	return builder
}

// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
func (builder *PanelModelBuilder) PluginVersion(pluginVersion string) *PanelModelBuilder {
	builder.internal.PluginVersion = &pluginVersion

	return builder
}

// Depends on the panel plugin. See the plugin documentation for details.
func (builder *PanelModelBuilder) Targets(targets []cog.Builder[variants.Dataquery]) *PanelModelBuilder {
	targetsResources := make([]variants.Dataquery, 0, len(targets))
	for _, r1 := range targets {
		targetsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		targetsResources = append(targetsResources, targetsDepth1)
	}
	builder.internal.Targets = targetsResources

	return builder
}

// Panel title.
func (builder *PanelModelBuilder) Title(title string) *PanelModelBuilder {
	builder.internal.Title = &title

	return builder
}

// Panel description.
func (builder *PanelModelBuilder) Description(description string) *PanelModelBuilder {
	builder.internal.Description = &description

	return builder
}

// Whether to display the panel without a background.
func (builder *PanelModelBuilder) Transparent(transparent bool) *PanelModelBuilder {
	builder.internal.Transparent = &transparent

	return builder
}

// The datasource used in all targets.
func (builder *PanelModelBuilder) Datasource(datasource common.DataSourceRef) *PanelModelBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Panel links.
func (builder *PanelModelBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *PanelModelBuilder {
	linksResources := make([]dashboard.DashboardLink, 0, len(links))
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

// Name of template variable to repeat for.
func (builder *PanelModelBuilder) Repeat(repeat string) *PanelModelBuilder {
	builder.internal.Repeat = &repeat

	return builder
}

// Direction to repeat in if 'repeat' is set.
// `h` for horizontal, `v` for vertical.
func (builder *PanelModelBuilder) RepeatDirection(repeatDirection PanelModelRepeatDirection) *PanelModelBuilder {
	builder.internal.RepeatDirection = &repeatDirection

	return builder
}

// Option for repeated panels that controls max items per row
// Only relevant for horizontally repeated panels
func (builder *PanelModelBuilder) MaxPerRow(maxPerRow float64) *PanelModelBuilder {
	builder.internal.MaxPerRow = &maxPerRow

	return builder
}

// The maximum number of data points that the panel queries are retrieving.
func (builder *PanelModelBuilder) MaxDataPoints(maxDataPoints float64) *PanelModelBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// List of transformations that are applied to the panel data before rendering.
// When there are multiple transformations, Grafana applies them in the order they are listed.
// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
func (builder *PanelModelBuilder) Transformations(transformations []dashboard.DataTransformerConfig) *PanelModelBuilder {
	builder.internal.Transformations = transformations

	return builder
}

// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
// This value must be formatted as a number followed by a valid time
// identifier like: "40s", "3d", etc.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelModelBuilder) Interval(interval string) *PanelModelBuilder {
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
func (builder *PanelModelBuilder) TimeFrom(timeFrom string) *PanelModelBuilder {
	builder.internal.TimeFrom = &timeFrom

	return builder
}

// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelModelBuilder) TimeShift(timeShift string) *PanelModelBuilder {
	builder.internal.TimeShift = &timeShift

	return builder
}

// Controls if the timeFrom or timeShift overrides are shown in the panel header
func (builder *PanelModelBuilder) HideTimeOverride(hideTimeOverride bool) *PanelModelBuilder {
	builder.internal.HideTimeOverride = &hideTimeOverride

	return builder
}

// Sets panel queries cache timeout.
func (builder *PanelModelBuilder) CacheTimeout(cacheTimeout string) *PanelModelBuilder {
	builder.internal.CacheTimeout = &cacheTimeout

	return builder
}

// Overrides the data source configured time-to-live for a query cache item in milliseconds
func (builder *PanelModelBuilder) QueryCachingTTL(queryCachingTTL float64) *PanelModelBuilder {
	builder.internal.QueryCachingTTL = &queryCachingTTL

	return builder
}

// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
func (builder *PanelModelBuilder) Options(options any) *PanelModelBuilder {
	builder.internal.Options = &options

	return builder
}

// Field options allow you to change how the data is displayed in your visualizations.
func (builder *PanelModelBuilder) FieldConfig(fieldConfig dashboard.FieldConfigSource) *PanelModelBuilder {
	builder.internal.FieldConfig = &fieldConfig

	return builder
}
