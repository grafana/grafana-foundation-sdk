// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"encoding/json"
	"reflect"
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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
	Model LibrarypanelLibraryPanelModel `json:"model"`
	// Object storage metadata
	Meta *LibraryElementDTOMeta `json:"meta,omitempty"`
}

func (resource LibraryPanel) Equals(other LibraryPanel) bool {
	if resource.FolderUid == nil && other.FolderUid != nil || resource.FolderUid != nil && other.FolderUid == nil {
		return false
	}

	if resource.FolderUid != nil {
		if *resource.FolderUid != *other.FolderUid {
			return false
		}
	}
	if resource.Uid != other.Uid {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.Description == nil && other.Description != nil || resource.Description != nil && other.Description == nil {
		return false
	}

	if resource.Description != nil {
		if *resource.Description != *other.Description {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.SchemaVersion == nil && other.SchemaVersion != nil || resource.SchemaVersion != nil && other.SchemaVersion == nil {
		return false
	}

	if resource.SchemaVersion != nil {
		if *resource.SchemaVersion != *other.SchemaVersion {
			return false
		}
	}
	if resource.Version != other.Version {
		return false
	}
	if !resource.Model.Equals(other.Model) {
		return false
	}
	if resource.Meta == nil && other.Meta != nil || resource.Meta != nil && other.Meta == nil {
		return false
	}

	if resource.Meta != nil {
		if !resource.Meta.Equals(*other.Meta) {
			return false
		}
	}

	return true
}

type LibraryElementDTOMetaUser struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

func (resource LibraryElementDTOMetaUser) Equals(other LibraryElementDTOMetaUser) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.AvatarUrl != other.AvatarUrl {
		return false
	}

	return true
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

func (resource LibraryElementDTOMeta) Equals(other LibraryElementDTOMeta) bool {
	if resource.FolderName != other.FolderName {
		return false
	}
	if resource.FolderUid != other.FolderUid {
		return false
	}
	if resource.ConnectedDashboards != other.ConnectedDashboards {
		return false
	}
	if resource.Created != other.Created {
		return false
	}
	if resource.Updated != other.Updated {
		return false
	}
	if !resource.CreatedBy.Equals(other.CreatedBy) {
		return false
	}
	if !resource.UpdatedBy.Equals(other.UpdatedBy) {
		return false
	}

	return true
}

type LibraryPanelRepeatDirection string

const (
	LibraryPanelRepeatDirectionH LibraryPanelRepeatDirection = "h"
	LibraryPanelRepeatDirectionV LibraryPanelRepeatDirection = "v"
)

type LibrarypanelLibraryPanelModel struct {
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
}

func (resource *LibrarypanelLibraryPanelModel) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}
	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}

	if fields["type"] != nil {
		if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
			return err
		}
	}

	if fields["pluginVersion"] != nil {
		if err := json.Unmarshal(fields["pluginVersion"], &resource.PluginVersion); err != nil {
			return err
		}
	}

	if fields["title"] != nil {
		if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
			return err
		}
	}

	if fields["description"] != nil {
		if err := json.Unmarshal(fields["description"], &resource.Description); err != nil {
			return err
		}
	}

	if fields["transparent"] != nil {
		if err := json.Unmarshal(fields["transparent"], &resource.Transparent); err != nil {
			return err
		}
	}

	if fields["datasource"] != nil {
		if err := json.Unmarshal(fields["datasource"], &resource.Datasource); err != nil {
			return err
		}
	}

	if fields["links"] != nil {
		if err := json.Unmarshal(fields["links"], &resource.Links); err != nil {
			return err
		}
	}

	if fields["repeat"] != nil {
		if err := json.Unmarshal(fields["repeat"], &resource.Repeat); err != nil {
			return err
		}
	}

	if fields["repeatDirection"] != nil {
		if err := json.Unmarshal(fields["repeatDirection"], &resource.RepeatDirection); err != nil {
			return err
		}
	}

	if fields["maxPerRow"] != nil {
		if err := json.Unmarshal(fields["maxPerRow"], &resource.MaxPerRow); err != nil {
			return err
		}
	}

	if fields["maxDataPoints"] != nil {
		if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
			return err
		}
	}

	if fields["transformations"] != nil {
		if err := json.Unmarshal(fields["transformations"], &resource.Transformations); err != nil {
			return err
		}
	}

	if fields["interval"] != nil {
		if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
			return err
		}
	}

	if fields["timeFrom"] != nil {
		if err := json.Unmarshal(fields["timeFrom"], &resource.TimeFrom); err != nil {
			return err
		}
	}

	if fields["timeShift"] != nil {
		if err := json.Unmarshal(fields["timeShift"], &resource.TimeShift); err != nil {
			return err
		}
	}

	if fields["hideTimeOverride"] != nil {
		if err := json.Unmarshal(fields["hideTimeOverride"], &resource.HideTimeOverride); err != nil {
			return err
		}
	}

	if fields["cacheTimeout"] != nil {
		if err := json.Unmarshal(fields["cacheTimeout"], &resource.CacheTimeout); err != nil {
			return err
		}
	}

	if fields["queryCachingTTL"] != nil {
		if err := json.Unmarshal(fields["queryCachingTTL"], &resource.QueryCachingTTL); err != nil {
			return err
		}
	}

	if fields["options"] != nil {
		if err := json.Unmarshal(fields["options"], &resource.Options); err != nil {
			return err
		}
	}

	if fields["fieldConfig"] != nil {
		if err := json.Unmarshal(fields["fieldConfig"], &resource.FieldConfig); err != nil {
			return err
		}
	}

	dataqueryTypeHint := ""
	if resource.Datasource != nil && resource.Datasource.Type != nil {
		dataqueryTypeHint = *resource.Datasource.Type
	}

	if fields["targets"] != nil {
		targets, err := cog.UnmarshalDataqueryArray(fields["targets"], dataqueryTypeHint)
		if err != nil {
			return err
		}
		resource.Targets = targets
	}

	return nil
}

func (resource LibrarypanelLibraryPanelModel) Equals(other LibrarypanelLibraryPanelModel) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.PluginVersion == nil && other.PluginVersion != nil || resource.PluginVersion != nil && other.PluginVersion == nil {
		return false
	}

	if resource.PluginVersion != nil {
		if *resource.PluginVersion != *other.PluginVersion {
			return false
		}
	}

	if len(resource.Targets) != len(other.Targets) {
		return false
	}

	for i1 := range resource.Targets {
		if !resource.Targets[i1].Equals(other.Targets[i1]) {
			return false
		}
	}
	if resource.Title == nil && other.Title != nil || resource.Title != nil && other.Title == nil {
		return false
	}

	if resource.Title != nil {
		if *resource.Title != *other.Title {
			return false
		}
	}
	if resource.Description == nil && other.Description != nil || resource.Description != nil && other.Description == nil {
		return false
	}

	if resource.Description != nil {
		if *resource.Description != *other.Description {
			return false
		}
	}
	if resource.Transparent == nil && other.Transparent != nil || resource.Transparent != nil && other.Transparent == nil {
		return false
	}

	if resource.Transparent != nil {
		if *resource.Transparent != *other.Transparent {
			return false
		}
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}

	if len(resource.Links) != len(other.Links) {
		return false
	}

	for i1 := range resource.Links {
		if !resource.Links[i1].Equals(other.Links[i1]) {
			return false
		}
	}
	if resource.Repeat == nil && other.Repeat != nil || resource.Repeat != nil && other.Repeat == nil {
		return false
	}

	if resource.Repeat != nil {
		if *resource.Repeat != *other.Repeat {
			return false
		}
	}
	if resource.RepeatDirection == nil && other.RepeatDirection != nil || resource.RepeatDirection != nil && other.RepeatDirection == nil {
		return false
	}

	if resource.RepeatDirection != nil {
		if *resource.RepeatDirection != *other.RepeatDirection {
			return false
		}
	}
	if resource.MaxPerRow == nil && other.MaxPerRow != nil || resource.MaxPerRow != nil && other.MaxPerRow == nil {
		return false
	}

	if resource.MaxPerRow != nil {
		if *resource.MaxPerRow != *other.MaxPerRow {
			return false
		}
	}
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
			return false
		}
	}

	if len(resource.Transformations) != len(other.Transformations) {
		return false
	}

	for i1 := range resource.Transformations {
		if !resource.Transformations[i1].Equals(other.Transformations[i1]) {
			return false
		}
	}
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.TimeFrom == nil && other.TimeFrom != nil || resource.TimeFrom != nil && other.TimeFrom == nil {
		return false
	}

	if resource.TimeFrom != nil {
		if *resource.TimeFrom != *other.TimeFrom {
			return false
		}
	}
	if resource.TimeShift == nil && other.TimeShift != nil || resource.TimeShift != nil && other.TimeShift == nil {
		return false
	}

	if resource.TimeShift != nil {
		if *resource.TimeShift != *other.TimeShift {
			return false
		}
	}
	if resource.HideTimeOverride == nil && other.HideTimeOverride != nil || resource.HideTimeOverride != nil && other.HideTimeOverride == nil {
		return false
	}

	if resource.HideTimeOverride != nil {
		if *resource.HideTimeOverride != *other.HideTimeOverride {
			return false
		}
	}
	if resource.CacheTimeout == nil && other.CacheTimeout != nil || resource.CacheTimeout != nil && other.CacheTimeout == nil {
		return false
	}

	if resource.CacheTimeout != nil {
		if *resource.CacheTimeout != *other.CacheTimeout {
			return false
		}
	}
	if resource.QueryCachingTTL == nil && other.QueryCachingTTL != nil || resource.QueryCachingTTL != nil && other.QueryCachingTTL == nil {
		return false
	}

	if resource.QueryCachingTTL != nil {
		if *resource.QueryCachingTTL != *other.QueryCachingTTL {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Options, other.Options) {
		return false
	}
	if resource.FieldConfig == nil && other.FieldConfig != nil || resource.FieldConfig != nil && other.FieldConfig == nil {
		return false
	}

	if resource.FieldConfig != nil {
		if !resource.FieldConfig.Equals(*other.FieldConfig) {
			return false
		}
	}

	return true
}
