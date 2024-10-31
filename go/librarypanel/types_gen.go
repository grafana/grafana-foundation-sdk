// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanel` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LibraryPanel) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "folderUid"
	if fields["folderUid"] != nil {
		if string(fields["folderUid"]) != "null" {
			if err := json.Unmarshal(fields["folderUid"], &resource.FolderUid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folderUid", err)...)
			}

		}
		delete(fields, "folderUid")

	}
	// Field "uid"
	if fields["uid"] != nil {
		if string(fields["uid"]) != "null" {
			if err := json.Unmarshal(fields["uid"], &resource.Uid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("uid", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("uid", errors.New("required field is null"))...)

		}
		delete(fields, "uid")
	} else {
		errs = append(errs, cog.MakeBuildErrors("uid", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}
	// Field "description"
	if fields["description"] != nil {
		if string(fields["description"]) != "null" {
			if err := json.Unmarshal(fields["description"], &resource.Description); err != nil {
				errs = append(errs, cog.MakeBuildErrors("description", err)...)
			}

		}
		delete(fields, "description")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "schemaVersion"
	if fields["schemaVersion"] != nil {
		if string(fields["schemaVersion"]) != "null" {
			if err := json.Unmarshal(fields["schemaVersion"], &resource.SchemaVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("schemaVersion", err)...)
			}

		}
		delete(fields, "schemaVersion")

	}
	// Field "version"
	if fields["version"] != nil {
		if string(fields["version"]) != "null" {
			if err := json.Unmarshal(fields["version"], &resource.Version); err != nil {
				errs = append(errs, cog.MakeBuildErrors("version", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("version", errors.New("required field is null"))...)

		}
		delete(fields, "version")
	} else {
		errs = append(errs, cog.MakeBuildErrors("version", errors.New("required field is missing from input"))...)
	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {

			resource.Model = LibrarypanelLibraryPanelModel{}
			if err := resource.Model.UnmarshalJSONStrict(fields["model"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "meta"
	if fields["meta"] != nil {
		if string(fields["meta"]) != "null" {

			resource.Meta = &LibraryElementDTOMeta{}
			if err := resource.Meta.UnmarshalJSONStrict(fields["meta"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("meta", err)...)
			}

		}
		delete(fields, "meta")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LibraryPanel", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LibraryPanel` objects.
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

// Validate checks all the validation constraints that may be defined on `LibraryPanel` fields for violations and returns them.
func (resource LibraryPanel) Validate() error {
	var errs cog.BuildErrors
	if !(len([]rune(resource.Name)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"name",
			errors.New("must be >= 1"),
		)...)
	}
	if !(len([]rune(resource.Type)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be >= 1"),
		)...)
	}
	if err := resource.Model.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("model", err)...)
	}
	if resource.Meta != nil {
		if err := resource.Meta.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("meta", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type LibraryElementDTOMetaUser struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryElementDTOMetaUser` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LibraryElementDTOMetaUser) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "id"
	if fields["id"] != nil {
		if string(fields["id"]) != "null" {
			if err := json.Unmarshal(fields["id"], &resource.Id); err != nil {
				errs = append(errs, cog.MakeBuildErrors("id", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is null"))...)

		}
		delete(fields, "id")
	} else {
		errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}
	// Field "avatarUrl"
	if fields["avatarUrl"] != nil {
		if string(fields["avatarUrl"]) != "null" {
			if err := json.Unmarshal(fields["avatarUrl"], &resource.AvatarUrl); err != nil {
				errs = append(errs, cog.MakeBuildErrors("avatarUrl", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("avatarUrl", errors.New("required field is null"))...)

		}
		delete(fields, "avatarUrl")
	} else {
		errs = append(errs, cog.MakeBuildErrors("avatarUrl", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LibraryElementDTOMetaUser", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LibraryElementDTOMetaUser` objects.
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

// Validate checks all the validation constraints that may be defined on `LibraryElementDTOMetaUser` fields for violations and returns them.
func (resource LibraryElementDTOMetaUser) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryElementDTOMeta` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LibraryElementDTOMeta) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "folderName"
	if fields["folderName"] != nil {
		if string(fields["folderName"]) != "null" {
			if err := json.Unmarshal(fields["folderName"], &resource.FolderName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folderName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("folderName", errors.New("required field is null"))...)

		}
		delete(fields, "folderName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("folderName", errors.New("required field is missing from input"))...)
	}
	// Field "folderUid"
	if fields["folderUid"] != nil {
		if string(fields["folderUid"]) != "null" {
			if err := json.Unmarshal(fields["folderUid"], &resource.FolderUid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folderUid", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("folderUid", errors.New("required field is null"))...)

		}
		delete(fields, "folderUid")
	} else {
		errs = append(errs, cog.MakeBuildErrors("folderUid", errors.New("required field is missing from input"))...)
	}
	// Field "connectedDashboards"
	if fields["connectedDashboards"] != nil {
		if string(fields["connectedDashboards"]) != "null" {
			if err := json.Unmarshal(fields["connectedDashboards"], &resource.ConnectedDashboards); err != nil {
				errs = append(errs, cog.MakeBuildErrors("connectedDashboards", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("connectedDashboards", errors.New("required field is null"))...)

		}
		delete(fields, "connectedDashboards")
	} else {
		errs = append(errs, cog.MakeBuildErrors("connectedDashboards", errors.New("required field is missing from input"))...)
	}
	// Field "created"
	if fields["created"] != nil {
		if string(fields["created"]) != "null" {
			if err := json.Unmarshal(fields["created"], &resource.Created); err != nil {
				errs = append(errs, cog.MakeBuildErrors("created", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("created", errors.New("required field is null"))...)

		}
		delete(fields, "created")
	} else {
		errs = append(errs, cog.MakeBuildErrors("created", errors.New("required field is missing from input"))...)
	}
	// Field "updated"
	if fields["updated"] != nil {
		if string(fields["updated"]) != "null" {
			if err := json.Unmarshal(fields["updated"], &resource.Updated); err != nil {
				errs = append(errs, cog.MakeBuildErrors("updated", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("updated", errors.New("required field is null"))...)

		}
		delete(fields, "updated")
	} else {
		errs = append(errs, cog.MakeBuildErrors("updated", errors.New("required field is missing from input"))...)
	}
	// Field "createdBy"
	if fields["createdBy"] != nil {
		if string(fields["createdBy"]) != "null" {

			resource.CreatedBy = LibraryElementDTOMetaUser{}
			if err := resource.CreatedBy.UnmarshalJSONStrict(fields["createdBy"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("createdBy", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("createdBy", errors.New("required field is null"))...)

		}
		delete(fields, "createdBy")
	} else {
		errs = append(errs, cog.MakeBuildErrors("createdBy", errors.New("required field is missing from input"))...)
	}
	// Field "updatedBy"
	if fields["updatedBy"] != nil {
		if string(fields["updatedBy"]) != "null" {

			resource.UpdatedBy = LibraryElementDTOMetaUser{}
			if err := resource.UpdatedBy.UnmarshalJSONStrict(fields["updatedBy"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("updatedBy", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("updatedBy", errors.New("required field is null"))...)

		}
		delete(fields, "updatedBy")
	} else {
		errs = append(errs, cog.MakeBuildErrors("updatedBy", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LibraryElementDTOMeta", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LibraryElementDTOMeta` objects.
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

// Validate checks all the validation constraints that may be defined on `LibraryElementDTOMeta` fields for violations and returns them.
func (resource LibraryElementDTOMeta) Validate() error {
	var errs cog.BuildErrors
	if err := resource.CreatedBy.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("createdBy", err)...)
	}
	if err := resource.UpdatedBy.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("updatedBy", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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
	// Tags for the panel.
	Tags []string `json:"tags,omitempty"`
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
	// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
	Options any `json:"options,omitempty"`
	// Field options allow you to change how the data is displayed in your visualizations.
	FieldConfig *dashboard.FieldConfigSource `json:"fieldConfig,omitempty"`
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode LibrarypanelLibraryPanelModel from JSON.
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
			return fmt.Errorf("error decoding field 'type': %w", err)
		}
	}

	if fields["pluginVersion"] != nil {
		if err := json.Unmarshal(fields["pluginVersion"], &resource.PluginVersion); err != nil {
			return fmt.Errorf("error decoding field 'pluginVersion': %w", err)
		}
	}

	if fields["tags"] != nil {
		if err := json.Unmarshal(fields["tags"], &resource.Tags); err != nil {
			return fmt.Errorf("error decoding field 'tags': %w", err)
		}
	}

	if fields["title"] != nil {
		if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
			return fmt.Errorf("error decoding field 'title': %w", err)
		}
	}

	if fields["description"] != nil {
		if err := json.Unmarshal(fields["description"], &resource.Description); err != nil {
			return fmt.Errorf("error decoding field 'description': %w", err)
		}
	}

	if fields["transparent"] != nil {
		if err := json.Unmarshal(fields["transparent"], &resource.Transparent); err != nil {
			return fmt.Errorf("error decoding field 'transparent': %w", err)
		}
	}

	if fields["datasource"] != nil {
		if err := json.Unmarshal(fields["datasource"], &resource.Datasource); err != nil {
			return fmt.Errorf("error decoding field 'datasource': %w", err)
		}
	}

	if fields["links"] != nil {
		if err := json.Unmarshal(fields["links"], &resource.Links); err != nil {
			return fmt.Errorf("error decoding field 'links': %w", err)
		}
	}

	if fields["repeat"] != nil {
		if err := json.Unmarshal(fields["repeat"], &resource.Repeat); err != nil {
			return fmt.Errorf("error decoding field 'repeat': %w", err)
		}
	}

	if fields["repeatDirection"] != nil {
		if err := json.Unmarshal(fields["repeatDirection"], &resource.RepeatDirection); err != nil {
			return fmt.Errorf("error decoding field 'repeatDirection': %w", err)
		}
	}

	if fields["maxPerRow"] != nil {
		if err := json.Unmarshal(fields["maxPerRow"], &resource.MaxPerRow); err != nil {
			return fmt.Errorf("error decoding field 'maxPerRow': %w", err)
		}
	}

	if fields["maxDataPoints"] != nil {
		if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
			return fmt.Errorf("error decoding field 'maxDataPoints': %w", err)
		}
	}

	if fields["transformations"] != nil {
		if err := json.Unmarshal(fields["transformations"], &resource.Transformations); err != nil {
			return fmt.Errorf("error decoding field 'transformations': %w", err)
		}
	}

	if fields["interval"] != nil {
		if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
			return fmt.Errorf("error decoding field 'interval': %w", err)
		}
	}

	if fields["timeFrom"] != nil {
		if err := json.Unmarshal(fields["timeFrom"], &resource.TimeFrom); err != nil {
			return fmt.Errorf("error decoding field 'timeFrom': %w", err)
		}
	}

	if fields["timeShift"] != nil {
		if err := json.Unmarshal(fields["timeShift"], &resource.TimeShift); err != nil {
			return fmt.Errorf("error decoding field 'timeShift': %w", err)
		}
	}

	if fields["hideTimeOverride"] != nil {
		if err := json.Unmarshal(fields["hideTimeOverride"], &resource.HideTimeOverride); err != nil {
			return fmt.Errorf("error decoding field 'hideTimeOverride': %w", err)
		}
	}

	if fields["options"] != nil {
		if err := json.Unmarshal(fields["options"], &resource.Options); err != nil {
			return fmt.Errorf("error decoding field 'options': %w", err)
		}
	}

	if fields["fieldConfig"] != nil {
		if err := json.Unmarshal(fields["fieldConfig"], &resource.FieldConfig); err != nil {
			return fmt.Errorf("error decoding field 'fieldConfig': %w", err)
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibrarypanelLibraryPanelModel` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LibrarypanelLibraryPanelModel) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "pluginVersion"
	if fields["pluginVersion"] != nil {
		if string(fields["pluginVersion"]) != "null" {
			if err := json.Unmarshal(fields["pluginVersion"], &resource.PluginVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pluginVersion", err)...)
			}

		}
		delete(fields, "pluginVersion")

	}
	// Field "tags"
	if fields["tags"] != nil {
		if string(fields["tags"]) != "null" {

			if err := json.Unmarshal(fields["tags"], &resource.Tags); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tags", err)...)
			}

		}
		delete(fields, "tags")

	}
	// Field "targets"
	if fields["targets"] != nil {
		if string(fields["targets"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["targets"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 variants.Dataquery

				dataquery, err := cog.StrictUnmarshalDataquery(partialArray[i1], "")
				if err != nil {
					errs = append(errs, cog.MakeBuildErrors("targets["+strconv.Itoa(i1)+"]", err)...)
				} else {
					result1 = dataquery
				}
				resource.Targets = append(resource.Targets, result1)
			}

		}
		delete(fields, "targets")

	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}

		}
		delete(fields, "title")

	}
	// Field "description"
	if fields["description"] != nil {
		if string(fields["description"]) != "null" {
			if err := json.Unmarshal(fields["description"], &resource.Description); err != nil {
				errs = append(errs, cog.MakeBuildErrors("description", err)...)
			}

		}
		delete(fields, "description")

	}
	// Field "transparent"
	if fields["transparent"] != nil {
		if string(fields["transparent"]) != "null" {
			if err := json.Unmarshal(fields["transparent"], &resource.Transparent); err != nil {
				errs = append(errs, cog.MakeBuildErrors("transparent", err)...)
			}

		}
		delete(fields, "transparent")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "links"
	if fields["links"] != nil {
		if string(fields["links"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["links"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 dashboard.DashboardLink

				result1 = dashboard.DashboardLink{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("links["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Links = append(resource.Links, result1)
			}

		}
		delete(fields, "links")

	}
	// Field "repeat"
	if fields["repeat"] != nil {
		if string(fields["repeat"]) != "null" {
			if err := json.Unmarshal(fields["repeat"], &resource.Repeat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
			}

		}
		delete(fields, "repeat")

	}
	// Field "repeatDirection"
	if fields["repeatDirection"] != nil {
		if string(fields["repeatDirection"]) != "null" {
			if err := json.Unmarshal(fields["repeatDirection"], &resource.RepeatDirection); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeatDirection", err)...)
			}

		}
		delete(fields, "repeatDirection")

	}
	// Field "maxPerRow"
	if fields["maxPerRow"] != nil {
		if string(fields["maxPerRow"]) != "null" {
			if err := json.Unmarshal(fields["maxPerRow"], &resource.MaxPerRow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxPerRow", err)...)
			}

		}
		delete(fields, "maxPerRow")

	}
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

	}
	// Field "transformations"
	if fields["transformations"] != nil {
		if string(fields["transformations"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["transformations"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 dashboard.DataTransformerConfig

				result1 = dashboard.DataTransformerConfig{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("transformations["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Transformations = append(resource.Transformations, result1)
			}

		}
		delete(fields, "transformations")

	}
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}

		}
		delete(fields, "interval")

	}
	// Field "timeFrom"
	if fields["timeFrom"] != nil {
		if string(fields["timeFrom"]) != "null" {
			if err := json.Unmarshal(fields["timeFrom"], &resource.TimeFrom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeFrom", err)...)
			}

		}
		delete(fields, "timeFrom")

	}
	// Field "timeShift"
	if fields["timeShift"] != nil {
		if string(fields["timeShift"]) != "null" {
			if err := json.Unmarshal(fields["timeShift"], &resource.TimeShift); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeShift", err)...)
			}

		}
		delete(fields, "timeShift")

	}
	// Field "hideTimeOverride"
	if fields["hideTimeOverride"] != nil {
		if string(fields["hideTimeOverride"]) != "null" {
			if err := json.Unmarshal(fields["hideTimeOverride"], &resource.HideTimeOverride); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideTimeOverride", err)...)
			}

		}
		delete(fields, "hideTimeOverride")

	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {
			if err := json.Unmarshal(fields["options"], &resource.Options); err != nil {
				errs = append(errs, cog.MakeBuildErrors("options", err)...)
			}

		}
		delete(fields, "options")

	}
	// Field "fieldConfig"
	if fields["fieldConfig"] != nil {
		if string(fields["fieldConfig"]) != "null" {

			resource.FieldConfig = &dashboard.FieldConfigSource{}
			if err := resource.FieldConfig.UnmarshalJSONStrict(fields["fieldConfig"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fieldConfig", err)...)
			}

		}
		delete(fields, "fieldConfig")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LibrarypanelLibraryPanelModel", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LibrarypanelLibraryPanelModel` objects.
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

	if len(resource.Tags) != len(other.Tags) {
		return false
	}

	for i1 := range resource.Tags {
		if resource.Tags[i1] != other.Tags[i1] {
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

// Validate checks all the validation constraints that may be defined on `LibrarypanelLibraryPanelModel` fields for violations and returns them.
func (resource LibrarypanelLibraryPanelModel) Validate() error {
	var errs cog.BuildErrors
	if !(len([]rune(resource.Type)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be >= 1"),
		)...)
	}

	for i1 := range resource.Targets {
		if err := resource.Targets[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("targets["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	for i1 := range resource.Links {
		if err := resource.Links[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("links["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.Transformations {
		if err := resource.Transformations[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("transformations["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.FieldConfig != nil {
		if err := resource.FieldConfig.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("fieldConfig", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
