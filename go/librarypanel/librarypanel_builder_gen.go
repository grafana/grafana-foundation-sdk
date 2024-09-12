// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[LibraryPanel] = (*LibraryPanelBuilder)(nil)

type LibraryPanelBuilder struct {
	internal *LibraryPanel
	errors   map[string]cog.BuildErrors
}

func NewLibraryPanelBuilder() *LibraryPanelBuilder {
	resource := &LibraryPanel{}
	builder := &LibraryPanelBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LibraryPanelBuilder) Build() (LibraryPanel, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("LibraryPanel", err)...)
	}

	if len(errs) != 0 {
		return LibraryPanel{}, errs
	}

	return *builder.internal, nil
}

// Folder UID
func (builder *LibraryPanelBuilder) FolderUid(folderUid string) *LibraryPanelBuilder {
	builder.internal.FolderUid = &folderUid

	return builder
}

// Library element UID
func (builder *LibraryPanelBuilder) Uid(uid string) *LibraryPanelBuilder {
	builder.internal.Uid = uid

	return builder
}

// Panel name (also saved in the model)
func (builder *LibraryPanelBuilder) Name(name string) *LibraryPanelBuilder {
	if !(len([]rune(name)) >= 1) {
		builder.errors["name"] = cog.MakeBuildErrors("name", errors.New("len([]rune(name)) must be >= 1"))
		return builder
	}
	builder.internal.Name = name

	return builder
}

// Panel description
func (builder *LibraryPanelBuilder) Description(description string) *LibraryPanelBuilder {
	builder.internal.Description = &description

	return builder
}

// The panel type (from inside the model)
func (builder *LibraryPanelBuilder) Type(typeArg string) *LibraryPanelBuilder {
	if !(len([]rune(typeArg)) >= 1) {
		builder.errors["typeArg"] = cog.MakeBuildErrors("typeArg", errors.New("len([]rune(typeArg)) must be >= 1"))
		return builder
	}
	builder.internal.Type = typeArg

	return builder
}

// Dashboard version when this was saved (zero if unknown)
func (builder *LibraryPanelBuilder) SchemaVersion(schemaVersion uint16) *LibraryPanelBuilder {
	builder.internal.SchemaVersion = &schemaVersion

	return builder
}

// panel version, incremented each time the dashboard is updated.
func (builder *LibraryPanelBuilder) Version(version int64) *LibraryPanelBuilder {
	builder.internal.Version = version

	return builder
}

// TODO: should be the same panel schema defined in dashboard
// Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
func (builder *LibraryPanelBuilder) Model(model struct {
	// The panel plugin type id. This is used to find the plugin to display the panel.
	Type string `json:"type"`
	// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
	PluginVersion *string `json:"pluginVersion,omitempty"`
	// Depends on the panel plugin. See the plugin documentation for details.
	Targets []variants.Dataquery `json:"targets,omitempty"`
	// Panel title.
	Title *string `json:"title,omitempty"`
	// Panel description.
	Description *string `json:"description,omitempty"`
	// Whether to display the panel without a background.
	Transparent *bool `json:"transparent,omitempty"`
	// The datasource used in all targets.
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// Panel links.
	Links []dashboard.DashboardLink `json:"links,omitempty"`
	// Name of template variable to repeat for.
	Repeat *string `json:"repeat,omitempty"`
	// Direction to repeat in if 'repeat' is set.
	// `h` for horizontal, `v` for vertical.
	RepeatDirection *LibraryPanelRepeatDirection `json:"repeatDirection,omitempty"`
	// Option for repeated panels that controls max items per row
	// Only relevant for horizontally repeated panels
	MaxPerRow *float64 `json:"maxPerRow,omitempty"`
	// The maximum number of data points that the panel queries are retrieving.
	MaxDataPoints *float64 `json:"maxDataPoints,omitempty"`
	// List of transformations that are applied to the panel data before rendering.
	// When there are multiple transformations, Grafana applies them in the order they are listed.
	// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
	Transformations []dashboard.DataTransformerConfig `json:"transformations,omitempty"`
	// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
	// This value must be formatted as a number followed by a valid time
	// identifier like: "40s", "3d", etc.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	Interval *string `json:"interval,omitempty"`
	// Overrides the relative time range for individual panels,
	// which causes them to be different than what is selected in
	// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
	// time periods or days on the same dashboard.
	// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
	// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	TimeFrom *string `json:"timeFrom,omitempty"`
	// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
	// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	TimeShift *string `json:"timeShift,omitempty"`
	// Controls if the timeFrom or timeShift overrides are shown in the panel header
	HideTimeOverride *bool `json:"hideTimeOverride,omitempty"`
	// Sets panel queries cache timeout.
	CacheTimeout *string `json:"cacheTimeout,omitempty"`
	// Overrides the data source configured time-to-live for a query cache item in milliseconds
	QueryCachingTTL *float64 `json:"queryCachingTTL,omitempty"`
	// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
	Options any `json:"options,omitempty"`
	// Field options allow you to change how the data is displayed in your visualizations.
	FieldConfig *dashboard.FieldConfigSource `json:"fieldConfig,omitempty"`
}) *LibraryPanelBuilder {
	builder.internal.Model = model

	return builder
}

// Object storage metadata
func (builder *LibraryPanelBuilder) Meta(meta cog.Builder[LibraryElementDTOMeta]) *LibraryPanelBuilder {
	metaResource, err := meta.Build()
	if err != nil {
		builder.errors["meta"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Meta = &metaResource

	return builder
}

func (builder *LibraryPanelBuilder) applyDefaults() {
}
