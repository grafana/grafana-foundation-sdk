// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"time"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type LibraryPanel struct {
	// Folder UID
	FolderUid *string `json:"folderUid,omitempty"`
	// Library element UID
	Uid string `json:"uid"`
	// Panel name (also saved in the model)
	Name string `json:"name"`
	// Panel description
	Description *string `json:"description,omitempty"`
	// The panel type (from inside the model)
	Type string `json:"type"`
	// Dashboard version when this was saved (zero if unknown)
	SchemaVersion *uint16 `json:"schemaVersion,omitempty"`
	// panel version, incremented each time the dashboard is updated.
	Version int64 `json:"version"`
	// TODO: should be the same panel schema defined in dashboard
	// Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
	Model struct {
		// The panel plugin type id. This is used to find the plugin to display the panel.
		Type string `json:"type"`
		// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
		PluginVersion *string `json:"pluginVersion,omitempty"`
		// Tags for the panel.
		Tags []string `json:"tags,omitempty"`
		// Depends on the panel plugin. See the plugin documentation for details.
		Targets []cogvariants.Dataquery `json:"targets,omitempty"`
		// Panel title.
		Title *string `json:"title,omitempty"`
		// Panel description.
		Description *string `json:"description,omitempty"`
		// Whether to display the panel without a background.
		Transparent bool `json:"transparent"`
		// The datasource used in all targets.
		Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
		// Panel links.
		Links []dashboard.DashboardLink `json:"links,omitempty"`
		// Name of template variable to repeat for.
		Repeat *string `json:"repeat,omitempty"`
		// Direction to repeat in if 'repeat' is set.
		// `h` for horizontal, `v` for vertical.
		RepeatDirection *LibraryPanelRepeatDirection `json:"repeatDirection,omitempty"`
		// Id of the repeating panel.
		RepeatPanelId *int64 `json:"repeatPanelId,omitempty"`
		// The maximum number of data points that the panel queries are retrieving.
		MaxDataPoints *float64 `json:"maxDataPoints,omitempty"`
		// List of transformations that are applied to the panel data before rendering.
		// When there are multiple transformations, Grafana applies them in the order they are listed.
		// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
		Transformations []dashboard.DataTransformerConfig `json:"transformations"`
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
		// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
		Options any `json:"options"`
		// Field options allow you to change how the data is displayed in your visualizations.
		FieldConfig dashboard.FieldConfigSource `json:"fieldConfig"`
	} `json:"model"`
	// Object storage metadata
	Meta *LibraryElementDTOMeta `json:"meta,omitempty"`
}

type LibraryElementDTOMetaUser struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type LibraryElementDTOMeta struct {
	FolderName          string                    `json:"folderName"`
	FolderUid           string                    `json:"folderUid"`
	ConnectedDashboards int64                     `json:"connectedDashboards"`
	Created             time.Time                 `json:"created"`
	Updated             time.Time                 `json:"updated"`
	CreatedBy           LibraryElementDTOMetaUser `json:"createdBy"`
	UpdatedBy           LibraryElementDTOMetaUser `json:"updatedBy"`
}

type LibraryPanelRepeatDirection string

const (
	LibraryPanelRepeatDirectionH LibraryPanelRepeatDirection = "h"
	LibraryPanelRepeatDirectionV LibraryPanelRepeatDirection = "v"
)
