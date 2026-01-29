// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type Dashboard struct {
	Annotations []AnnotationQueryKind `json:"annotations"`
	// Configuration of dashboard cursor sync behavior.
	// "Off" for no shared crosshair or tooltip (default).
	// "Crosshair" for shared crosshair.
	// "Tooltip" for shared crosshair AND shared tooltip.
	CursorSync DashboardCursorSync `json:"cursorSync"`
	// Description of dashboard.
	Description *string `json:"description,omitempty"`
	// Whether a dashboard is editable or not.
	Editable *bool                                                              `json:"editable,omitempty"`
	Elements map[string]Element                                                 `json:"elements"`
	Layout   GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind `json:"layout"`
	// Links with references to other dashboards or external websites.
	Links []DashboardLink `json:"links"`
	// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
	// This will keep data "moving left" regardless of the query refresh rate. This setting helps
	// avoid dashboards presenting stale live data.
	LiveNow *bool `json:"liveNow,omitempty"`
	// When set to true, the dashboard will load all panels in the dashboard when it's loaded.
	Preload bool `json:"preload"`
	// Plugins only. The version of the dashboard installed together with the plugin.
	// This is used to determine if the dashboard should be updated when the plugin is updated.
	Revision *uint16 `json:"revision,omitempty"`
	// Tags associated with dashboard.
	Tags         []string         `json:"tags"`
	TimeSettings TimeSettingsSpec `json:"timeSettings"`
	// Title of dashboard.
	Title string `json:"title"`
	// Configured template variables.
	Variables []VariableKind `json:"variables"`
}

// NewDashboard creates a new Dashboard object.
func NewDashboard() *Dashboard {
	return &Dashboard{
		Annotations:  []AnnotationQueryKind{},
		CursorSync:   DashboardCursorSyncOff,
		Editable:     (func(input bool) *bool { return &input })(true),
		Elements:     map[string]Element{},
		Layout:       *NewGridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind(),
		Links:        []DashboardLink{},
		Preload:      false,
		Tags:         []string{},
		TimeSettings: *NewTimeSettingsSpec(),
		Variables:    []VariableKind{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboard` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboard) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "annotations"
	if fields["annotations"] != nil {
		if string(fields["annotations"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["annotations"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AnnotationQueryKind

				result1 = AnnotationQueryKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("annotations["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Annotations = append(resource.Annotations, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("annotations", errors.New("required field is null"))...)

		}
		delete(fields, "annotations")
	} else {
		errs = append(errs, cog.MakeBuildErrors("annotations", errors.New("required field is missing from input"))...)
	}
	// Field "cursorSync"
	if fields["cursorSync"] != nil {
		if string(fields["cursorSync"]) != "null" {
			if err := json.Unmarshal(fields["cursorSync"], &resource.CursorSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cursorSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("cursorSync", errors.New("required field is null"))...)

		}
		delete(fields, "cursorSync")

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
	// Field "editable"
	if fields["editable"] != nil {
		if string(fields["editable"]) != "null" {
			if err := json.Unmarshal(fields["editable"], &resource.Editable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("editable", err)...)
			}

		}
		delete(fields, "editable")

	}
	// Field "elements"
	if fields["elements"] != nil {
		if string(fields["elements"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["elements"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]Element, len(partialMap))
			for key1 := range partialMap {
				var result1 Element

				result1 = Element{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("elements["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Elements = parsedMap1
		} else {
			errs = append(errs, cog.MakeBuildErrors("elements", errors.New("required field is null"))...)

		}
		delete(fields, "elements")
	} else {
		errs = append(errs, cog.MakeBuildErrors("elements", errors.New("required field is missing from input"))...)
	}
	// Field "layout"
	if fields["layout"] != nil {
		if string(fields["layout"]) != "null" {

			resource.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{}
			if err := resource.Layout.UnmarshalJSONStrict(fields["layout"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("layout", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("layout", errors.New("required field is null"))...)

		}
		delete(fields, "layout")
	} else {
		errs = append(errs, cog.MakeBuildErrors("layout", errors.New("required field is missing from input"))...)
	}
	// Field "links"
	if fields["links"] != nil {
		if string(fields["links"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["links"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 DashboardLink

				result1 = DashboardLink{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("links["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Links = append(resource.Links, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("links", errors.New("required field is null"))...)

		}
		delete(fields, "links")
	} else {
		errs = append(errs, cog.MakeBuildErrors("links", errors.New("required field is missing from input"))...)
	}
	// Field "liveNow"
	if fields["liveNow"] != nil {
		if string(fields["liveNow"]) != "null" {
			if err := json.Unmarshal(fields["liveNow"], &resource.LiveNow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("liveNow", err)...)
			}

		}
		delete(fields, "liveNow")

	}
	// Field "preload"
	if fields["preload"] != nil {
		if string(fields["preload"]) != "null" {
			if err := json.Unmarshal(fields["preload"], &resource.Preload); err != nil {
				errs = append(errs, cog.MakeBuildErrors("preload", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("preload", errors.New("required field is null"))...)

		}
		delete(fields, "preload")

	}
	// Field "revision"
	if fields["revision"] != nil {
		if string(fields["revision"]) != "null" {
			if err := json.Unmarshal(fields["revision"], &resource.Revision); err != nil {
				errs = append(errs, cog.MakeBuildErrors("revision", err)...)
			}

		}
		delete(fields, "revision")

	}
	// Field "tags"
	if fields["tags"] != nil {
		if string(fields["tags"]) != "null" {

			if err := json.Unmarshal(fields["tags"], &resource.Tags); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tags", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is null"))...)

		}
		delete(fields, "tags")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is missing from input"))...)
	}
	// Field "timeSettings"
	if fields["timeSettings"] != nil {
		if string(fields["timeSettings"]) != "null" {

			resource.TimeSettings = TimeSettingsSpec{}
			if err := resource.TimeSettings.UnmarshalJSONStrict(fields["timeSettings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeSettings", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("timeSettings", errors.New("required field is null"))...)

		}
		delete(fields, "timeSettings")
	} else {
		errs = append(errs, cog.MakeBuildErrors("timeSettings", errors.New("required field is missing from input"))...)
	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}
	// Field "variables"
	if fields["variables"] != nil {
		if string(fields["variables"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["variables"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 VariableKind

				result1 = VariableKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("variables["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Variables = append(resource.Variables, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("variables", errors.New("required field is null"))...)

		}
		delete(fields, "variables")
	} else {
		errs = append(errs, cog.MakeBuildErrors("variables", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboard", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboard` objects.
func (resource Dashboard) Equals(other Dashboard) bool {

	if len(resource.Annotations) != len(other.Annotations) {
		return false
	}

	for i1 := range resource.Annotations {
		if !resource.Annotations[i1].Equals(other.Annotations[i1]) {
			return false
		}
	}
	if resource.CursorSync != other.CursorSync {
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
	if resource.Editable == nil && other.Editable != nil || resource.Editable != nil && other.Editable == nil {
		return false
	}

	if resource.Editable != nil {
		if *resource.Editable != *other.Editable {
			return false
		}
	}

	if len(resource.Elements) != len(other.Elements) {
		return false
	}

	for key1 := range resource.Elements {
		if !resource.Elements[key1].Equals(other.Elements[key1]) {
			return false
		}
	}
	if !resource.Layout.Equals(other.Layout) {
		return false
	}

	if len(resource.Links) != len(other.Links) {
		return false
	}

	for i1 := range resource.Links {
		if !resource.Links[i1].Equals(other.Links[i1]) {
			return false
		}
	}
	if resource.LiveNow == nil && other.LiveNow != nil || resource.LiveNow != nil && other.LiveNow == nil {
		return false
	}

	if resource.LiveNow != nil {
		if *resource.LiveNow != *other.LiveNow {
			return false
		}
	}
	if resource.Preload != other.Preload {
		return false
	}
	if resource.Revision == nil && other.Revision != nil || resource.Revision != nil && other.Revision == nil {
		return false
	}

	if resource.Revision != nil {
		if *resource.Revision != *other.Revision {
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
	if !resource.TimeSettings.Equals(other.TimeSettings) {
		return false
	}
	if resource.Title != other.Title {
		return false
	}

	if len(resource.Variables) != len(other.Variables) {
		return false
	}

	for i1 := range resource.Variables {
		if !resource.Variables[i1].Equals(other.Variables[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboard` fields for violations and returns them.
func (resource Dashboard) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Annotations {
		if err := resource.Annotations[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("annotations["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for key1 := range resource.Elements {
		if err := resource.Elements[key1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("elements["+key1+"]", err)...)
		}
	}
	if err := resource.Layout.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("layout", err)...)
	}

	for i1 := range resource.Links {
		if err := resource.Links[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("links["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if err := resource.TimeSettings.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("timeSettings", err)...)
	}

	for i1 := range resource.Variables {
		if err := resource.Variables[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("variables["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AnnotationQueryKind struct {
	Kind string              `json:"kind"`
	Spec AnnotationQuerySpec `json:"spec"`
}

// NewAnnotationQueryKind creates a new AnnotationQueryKind object.
func NewAnnotationQueryKind() *AnnotationQueryKind {
	return &AnnotationQueryKind{
		Kind: "AnnotationQuery",
		Spec: *NewAnnotationQuerySpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationQueryKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AnnotationQueryKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = AnnotationQuerySpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AnnotationQueryKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AnnotationQueryKind` objects.
func (resource AnnotationQueryKind) Equals(other AnnotationQueryKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AnnotationQueryKind` fields for violations and returns them.
func (resource AnnotationQueryKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AnnotationQuerySpec struct {
	Query     DataQueryKind          `json:"query"`
	Enable    bool                   `json:"enable"`
	Hide      bool                   `json:"hide"`
	IconColor string                 `json:"iconColor"`
	Name      string                 `json:"name"`
	BuiltIn   *bool                  `json:"builtIn,omitempty"`
	Filter    *AnnotationPanelFilter `json:"filter,omitempty"`
	// Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
	Placement *string `json:"placement,omitempty"`
	// Mappings define how to convert data frame fields to annotation event fields.
	Mappings map[string]AnnotationEventFieldMapping `json:"mappings,omitempty"`
	// Catch-all field for datasource-specific properties. Should not be available in as code tooling.
	LegacyOptions map[string]any `json:"legacyOptions,omitempty"`
}

// NewAnnotationQuerySpec creates a new AnnotationQuerySpec object.
func NewAnnotationQuerySpec() *AnnotationQuerySpec {
	return &AnnotationQuerySpec{
		Query:     *NewDataQueryKind(),
		BuiltIn:   (func(input bool) *bool { return &input })(false),
		Placement: (func(input string) *string { return &input })(AnnotationQueryPlacement),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationQuerySpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AnnotationQuerySpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {

			resource.Query = DataQueryKind{}
			if err := resource.Query.UnmarshalJSONStrict(fields["query"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "enable"
	if fields["enable"] != nil {
		if string(fields["enable"]) != "null" {
			if err := json.Unmarshal(fields["enable"], &resource.Enable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("enable", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("enable", errors.New("required field is null"))...)

		}
		delete(fields, "enable")
	} else {
		errs = append(errs, cog.MakeBuildErrors("enable", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")
	} else {
		errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is missing from input"))...)
	}
	// Field "iconColor"
	if fields["iconColor"] != nil {
		if string(fields["iconColor"]) != "null" {
			if err := json.Unmarshal(fields["iconColor"], &resource.IconColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("iconColor", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("iconColor", errors.New("required field is null"))...)

		}
		delete(fields, "iconColor")
	} else {
		errs = append(errs, cog.MakeBuildErrors("iconColor", errors.New("required field is missing from input"))...)
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
	// Field "builtIn"
	if fields["builtIn"] != nil {
		if string(fields["builtIn"]) != "null" {
			if err := json.Unmarshal(fields["builtIn"], &resource.BuiltIn); err != nil {
				errs = append(errs, cog.MakeBuildErrors("builtIn", err)...)
			}

		}
		delete(fields, "builtIn")

	}
	// Field "filter"
	if fields["filter"] != nil {
		if string(fields["filter"]) != "null" {

			resource.Filter = &AnnotationPanelFilter{}
			if err := resource.Filter.UnmarshalJSONStrict(fields["filter"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filter", err)...)
			}

		}
		delete(fields, "filter")

	}
	// Field "placement"
	if fields["placement"] != nil {
		if string(fields["placement"]) != "null" {
			if err := json.Unmarshal(fields["placement"], &resource.Placement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("placement", err)...)
			}

		}
		delete(fields, "placement")

	}
	// Field "mappings"
	if fields["mappings"] != nil {
		if string(fields["mappings"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["mappings"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]AnnotationEventFieldMapping, len(partialMap))
			for key1 := range partialMap {
				var result1 AnnotationEventFieldMapping

				result1 = AnnotationEventFieldMapping{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("mappings["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Mappings = parsedMap1

		}
		delete(fields, "mappings")

	}
	// Field "legacyOptions"
	if fields["legacyOptions"] != nil {
		if string(fields["legacyOptions"]) != "null" {

			if err := json.Unmarshal(fields["legacyOptions"], &resource.LegacyOptions); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legacyOptions", err)...)
			}

		}
		delete(fields, "legacyOptions")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AnnotationQuerySpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AnnotationQuerySpec` objects.
func (resource AnnotationQuerySpec) Equals(other AnnotationQuerySpec) bool {
	if !resource.Query.Equals(other.Query) {
		return false
	}
	if resource.Enable != other.Enable {
		return false
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.IconColor != other.IconColor {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.BuiltIn == nil && other.BuiltIn != nil || resource.BuiltIn != nil && other.BuiltIn == nil {
		return false
	}

	if resource.BuiltIn != nil {
		if *resource.BuiltIn != *other.BuiltIn {
			return false
		}
	}
	if resource.Filter == nil && other.Filter != nil || resource.Filter != nil && other.Filter == nil {
		return false
	}

	if resource.Filter != nil {
		if !resource.Filter.Equals(*other.Filter) {
			return false
		}
	}
	if resource.Placement == nil && other.Placement != nil || resource.Placement != nil && other.Placement == nil {
		return false
	}

	if resource.Placement != nil {
		if *resource.Placement != *other.Placement {
			return false
		}
	}

	if len(resource.Mappings) != len(other.Mappings) {
		return false
	}

	for key1 := range resource.Mappings {
		if !resource.Mappings[key1].Equals(other.Mappings[key1]) {
			return false
		}
	}

	if len(resource.LegacyOptions) != len(other.LegacyOptions) {
		return false
	}

	for key1 := range resource.LegacyOptions {
		// is DeepEqual good enough here?
		if !reflect.DeepEqual(resource.LegacyOptions[key1], other.LegacyOptions[key1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AnnotationQuerySpec` fields for violations and returns them.
func (resource AnnotationQuerySpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Query.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("query", err)...)
	}
	if resource.Filter != nil {
		if err := resource.Filter.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filter", err)...)
		}
	}

	for key1 := range resource.Mappings {
		if err := resource.Mappings[key1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("mappings["+key1+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type DataQueryKind struct {
	Kind    string `json:"kind"`
	Group   string `json:"group"`
	Version string `json:"version"`
	// New type for datasource reference
	// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
	Datasource *Dashboardv2beta1DataQueryKindDatasource `json:"datasource,omitempty"`
	Spec       any                                      `json:"spec"`
}

// NewDataQueryKind creates a new DataQueryKind object.
func NewDataQueryKind() *DataQueryKind {
	return &DataQueryKind{
		Kind:    "DataQuery",
		Version: "v0",
	}
}
func (resource *DataQueryKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	if fields["kind"] != nil {
		if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
			return fmt.Errorf("error decoding field 'kind': %w", err)
		}
	}
	if fields["group"] != nil {
		if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
			return fmt.Errorf("error decoding field 'group': %w", err)
		}
	}
	if fields["version"] != nil {
		if err := json.Unmarshal(fields["version"], &resource.Version); err != nil {
			return fmt.Errorf("error decoding field 'version': %w", err)
		}
	}
	if fields["datasource"] != nil {
		if err := json.Unmarshal(fields["datasource"], &resource.Datasource); err != nil {
			return fmt.Errorf("error decoding field 'datasource': %w", err)
		}
	}

	if fields["spec"] != nil {
		dataquery, err := cog.UnmarshalDataquery(fields["spec"], resource.Group)
		if err != nil {
			return fmt.Errorf("error decoding field 'spec': %w", err)
		}
		resource.Spec = dataquery
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataQueryKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DataQueryKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is null"))...)

		}
		delete(fields, "group")
	} else {
		errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is missing from input"))...)
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

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &Dashboardv2beta1DataQueryKindDatasource{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {
			dataquery, err := cog.UnmarshalDataquery(fields["spec"], resource.Group)
			if err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			} else {
				resource.Spec = dataquery
			}

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DataQueryKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DataQueryKind` objects.
func (resource DataQueryKind) Equals(other DataQueryKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Group != other.Group {
		return false
	}
	if resource.Version != other.Version {
		return false
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Spec, other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DataQueryKind` fields for violations and returns them.
func (resource DataQueryKind) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AnnotationPanelFilter struct {
	// Should the specified panels be included or excluded
	Exclude *bool `json:"exclude,omitempty"`
	// Panel IDs that should be included or excluded
	Ids []uint32 `json:"ids"`
}

// NewAnnotationPanelFilter creates a new AnnotationPanelFilter object.
func NewAnnotationPanelFilter() *AnnotationPanelFilter {
	return &AnnotationPanelFilter{
		Exclude: (func(input bool) *bool { return &input })(false),
		Ids:     []uint32{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationPanelFilter` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AnnotationPanelFilter) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "exclude"
	if fields["exclude"] != nil {
		if string(fields["exclude"]) != "null" {
			if err := json.Unmarshal(fields["exclude"], &resource.Exclude); err != nil {
				errs = append(errs, cog.MakeBuildErrors("exclude", err)...)
			}

		}
		delete(fields, "exclude")

	}
	// Field "ids"
	if fields["ids"] != nil {
		if string(fields["ids"]) != "null" {

			if err := json.Unmarshal(fields["ids"], &resource.Ids); err != nil {
				errs = append(errs, cog.MakeBuildErrors("ids", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("ids", errors.New("required field is null"))...)

		}
		delete(fields, "ids")
	} else {
		errs = append(errs, cog.MakeBuildErrors("ids", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AnnotationPanelFilter", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AnnotationPanelFilter` objects.
func (resource AnnotationPanelFilter) Equals(other AnnotationPanelFilter) bool {
	if resource.Exclude == nil && other.Exclude != nil || resource.Exclude != nil && other.Exclude == nil {
		return false
	}

	if resource.Exclude != nil {
		if *resource.Exclude != *other.Exclude {
			return false
		}
	}

	if len(resource.Ids) != len(other.Ids) {
		return false
	}

	for i1 := range resource.Ids {
		if resource.Ids[i1] != other.Ids[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AnnotationPanelFilter` fields for violations and returns them.
func (resource AnnotationPanelFilter) Validate() error {
	return nil
}

// Annotation Query placement. Defines where the annotation query should be displayed.
// - "inControlsMenu" renders the annotation query in the dashboard controls dropdown menu
const AnnotationQueryPlacement = "inControlsMenu"

// Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
type AnnotationEventFieldMapping struct {
	// Source type for the field value
	Source *string `json:"source,omitempty"`
	// Constant value to use when source is "text"
	Value *string `json:"value,omitempty"`
	// Regular expression to apply to the field value
	Regex *string `json:"regex,omitempty"`
}

// NewAnnotationEventFieldMapping creates a new AnnotationEventFieldMapping object.
func NewAnnotationEventFieldMapping() *AnnotationEventFieldMapping {
	return &AnnotationEventFieldMapping{
		Source: (func(input string) *string { return &input })("field"),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationEventFieldMapping` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AnnotationEventFieldMapping) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "source"
	if fields["source"] != nil {
		if string(fields["source"]) != "null" {
			if err := json.Unmarshal(fields["source"], &resource.Source); err != nil {
				errs = append(errs, cog.MakeBuildErrors("source", err)...)
			}

		}
		delete(fields, "source")

	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}

		}
		delete(fields, "value")

	}
	// Field "regex"
	if fields["regex"] != nil {
		if string(fields["regex"]) != "null" {
			if err := json.Unmarshal(fields["regex"], &resource.Regex); err != nil {
				errs = append(errs, cog.MakeBuildErrors("regex", err)...)
			}

		}
		delete(fields, "regex")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AnnotationEventFieldMapping", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AnnotationEventFieldMapping` objects.
func (resource AnnotationEventFieldMapping) Equals(other AnnotationEventFieldMapping) bool {
	if resource.Source == nil && other.Source != nil || resource.Source != nil && other.Source == nil {
		return false
	}

	if resource.Source != nil {
		if *resource.Source != *other.Source {
			return false
		}
	}
	if resource.Value == nil && other.Value != nil || resource.Value != nil && other.Value == nil {
		return false
	}

	if resource.Value != nil {
		if *resource.Value != *other.Value {
			return false
		}
	}
	if resource.Regex == nil && other.Regex != nil || resource.Regex != nil && other.Regex == nil {
		return false
	}

	if resource.Regex != nil {
		if *resource.Regex != *other.Regex {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AnnotationEventFieldMapping` fields for violations and returns them.
func (resource AnnotationEventFieldMapping) Validate() error {
	return nil
}

// "Off" for no shared crosshair or tooltip (default).
// "Crosshair" for shared crosshair.
// "Tooltip" for shared crosshair AND shared tooltip.
type DashboardCursorSync string

const (
	DashboardCursorSyncCrosshair DashboardCursorSync = "Crosshair"
	DashboardCursorSyncTooltip   DashboardCursorSync = "Tooltip"
	DashboardCursorSyncOff       DashboardCursorSync = "Off"
)

// Supported dashboard elements
// |* more element types in the future
type Element = PanelKindOrLibraryPanelKind

// NewElement creates a new Element object.
func NewElement() *Element {
	return NewPanelKindOrLibraryPanelKind()
}

type PanelKind struct {
	Kind string    `json:"kind"`
	Spec PanelSpec `json:"spec"`
}

// NewPanelKind creates a new PanelKind object.
func NewPanelKind() *PanelKind {
	return &PanelKind{
		Kind: "Panel",
		Spec: *NewPanelSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PanelKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = PanelSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PanelKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PanelKind` objects.
func (resource PanelKind) Equals(other PanelKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PanelKind` fields for violations and returns them.
func (resource PanelKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type PanelSpec struct {
	Id          float64        `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Links       []DataLink     `json:"links"`
	Data        QueryGroupKind `json:"data"`
	VizConfig   VizConfigKind  `json:"vizConfig"`
	Transparent *bool          `json:"transparent,omitempty"`
}

// NewPanelSpec creates a new PanelSpec object.
func NewPanelSpec() *PanelSpec {
	return &PanelSpec{
		Links:     []DataLink{},
		Data:      *NewQueryGroupKind(),
		VizConfig: *NewVizConfigKind(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PanelSpec) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}
	// Field "description"
	if fields["description"] != nil {
		if string(fields["description"]) != "null" {
			if err := json.Unmarshal(fields["description"], &resource.Description); err != nil {
				errs = append(errs, cog.MakeBuildErrors("description", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("description", errors.New("required field is null"))...)

		}
		delete(fields, "description")
	} else {
		errs = append(errs, cog.MakeBuildErrors("description", errors.New("required field is missing from input"))...)
	}
	// Field "links"
	if fields["links"] != nil {
		if string(fields["links"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["links"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 DataLink

				result1 = DataLink{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("links["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Links = append(resource.Links, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("links", errors.New("required field is null"))...)

		}
		delete(fields, "links")
	} else {
		errs = append(errs, cog.MakeBuildErrors("links", errors.New("required field is missing from input"))...)
	}
	// Field "data"
	if fields["data"] != nil {
		if string(fields["data"]) != "null" {

			resource.Data = QueryGroupKind{}
			if err := resource.Data.UnmarshalJSONStrict(fields["data"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("data", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("data", errors.New("required field is null"))...)

		}
		delete(fields, "data")
	} else {
		errs = append(errs, cog.MakeBuildErrors("data", errors.New("required field is missing from input"))...)
	}
	// Field "vizConfig"
	if fields["vizConfig"] != nil {
		if string(fields["vizConfig"]) != "null" {

			resource.VizConfig = VizConfigKind{}
			if err := resource.VizConfig.UnmarshalJSONStrict(fields["vizConfig"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("vizConfig", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("vizConfig", errors.New("required field is null"))...)

		}
		delete(fields, "vizConfig")
	} else {
		errs = append(errs, cog.MakeBuildErrors("vizConfig", errors.New("required field is missing from input"))...)
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PanelSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PanelSpec` objects.
func (resource PanelSpec) Equals(other PanelSpec) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Title != other.Title {
		return false
	}
	if resource.Description != other.Description {
		return false
	}

	if len(resource.Links) != len(other.Links) {
		return false
	}

	for i1 := range resource.Links {
		if !resource.Links[i1].Equals(other.Links[i1]) {
			return false
		}
	}
	if !resource.Data.Equals(other.Data) {
		return false
	}
	if !resource.VizConfig.Equals(other.VizConfig) {
		return false
	}
	if resource.Transparent == nil && other.Transparent != nil || resource.Transparent != nil && other.Transparent == nil {
		return false
	}

	if resource.Transparent != nil {
		if *resource.Transparent != *other.Transparent {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PanelSpec` fields for violations and returns them.
func (resource PanelSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Links {
		if err := resource.Links[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("links["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if err := resource.Data.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("data", err)...)
	}
	if err := resource.VizConfig.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("vizConfig", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type DataLink struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	TargetBlank *bool  `json:"targetBlank,omitempty"`
}

// NewDataLink creates a new DataLink object.
func NewDataLink() *DataLink {
	return &DataLink{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataLink` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DataLink) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}
	// Field "url"
	if fields["url"] != nil {
		if string(fields["url"]) != "null" {
			if err := json.Unmarshal(fields["url"], &resource.Url); err != nil {
				errs = append(errs, cog.MakeBuildErrors("url", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("url", errors.New("required field is null"))...)

		}
		delete(fields, "url")
	} else {
		errs = append(errs, cog.MakeBuildErrors("url", errors.New("required field is missing from input"))...)
	}
	// Field "targetBlank"
	if fields["targetBlank"] != nil {
		if string(fields["targetBlank"]) != "null" {
			if err := json.Unmarshal(fields["targetBlank"], &resource.TargetBlank); err != nil {
				errs = append(errs, cog.MakeBuildErrors("targetBlank", err)...)
			}

		}
		delete(fields, "targetBlank")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DataLink", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DataLink` objects.
func (resource DataLink) Equals(other DataLink) bool {
	if resource.Title != other.Title {
		return false
	}
	if resource.Url != other.Url {
		return false
	}
	if resource.TargetBlank == nil && other.TargetBlank != nil || resource.TargetBlank != nil && other.TargetBlank == nil {
		return false
	}

	if resource.TargetBlank != nil {
		if *resource.TargetBlank != *other.TargetBlank {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DataLink` fields for violations and returns them.
func (resource DataLink) Validate() error {
	return nil
}

type QueryGroupKind struct {
	Kind string         `json:"kind"`
	Spec QueryGroupSpec `json:"spec"`
}

// NewQueryGroupKind creates a new QueryGroupKind object.
func NewQueryGroupKind() *QueryGroupKind {
	return &QueryGroupKind{
		Kind: "QueryGroup",
		Spec: *NewQueryGroupSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryGroupKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryGroupKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = QueryGroupSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryGroupKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryGroupKind` objects.
func (resource QueryGroupKind) Equals(other QueryGroupKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryGroupKind` fields for violations and returns them.
func (resource QueryGroupKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryGroupSpec struct {
	Queries         []PanelQueryKind     `json:"queries"`
	Transformations []TransformationKind `json:"transformations"`
	QueryOptions    QueryOptionsSpec     `json:"queryOptions"`
}

// NewQueryGroupSpec creates a new QueryGroupSpec object.
func NewQueryGroupSpec() *QueryGroupSpec {
	return &QueryGroupSpec{
		Queries:         []PanelQueryKind{},
		Transformations: []TransformationKind{},
		QueryOptions:    *NewQueryOptionsSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryGroupSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryGroupSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "queries"
	if fields["queries"] != nil {
		if string(fields["queries"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["queries"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 PanelQueryKind

				result1 = PanelQueryKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("queries["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Queries = append(resource.Queries, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("queries", errors.New("required field is null"))...)

		}
		delete(fields, "queries")
	} else {
		errs = append(errs, cog.MakeBuildErrors("queries", errors.New("required field is missing from input"))...)
	}
	// Field "transformations"
	if fields["transformations"] != nil {
		if string(fields["transformations"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["transformations"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TransformationKind

				result1 = TransformationKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("transformations["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Transformations = append(resource.Transformations, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("transformations", errors.New("required field is null"))...)

		}
		delete(fields, "transformations")
	} else {
		errs = append(errs, cog.MakeBuildErrors("transformations", errors.New("required field is missing from input"))...)
	}
	// Field "queryOptions"
	if fields["queryOptions"] != nil {
		if string(fields["queryOptions"]) != "null" {

			resource.QueryOptions = QueryOptionsSpec{}
			if err := resource.QueryOptions.UnmarshalJSONStrict(fields["queryOptions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryOptions", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("queryOptions", errors.New("required field is null"))...)

		}
		delete(fields, "queryOptions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("queryOptions", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryGroupSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryGroupSpec` objects.
func (resource QueryGroupSpec) Equals(other QueryGroupSpec) bool {

	if len(resource.Queries) != len(other.Queries) {
		return false
	}

	for i1 := range resource.Queries {
		if !resource.Queries[i1].Equals(other.Queries[i1]) {
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
	if !resource.QueryOptions.Equals(other.QueryOptions) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryGroupSpec` fields for violations and returns them.
func (resource QueryGroupSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Queries {
		if err := resource.Queries[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("queries["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.Transformations {
		if err := resource.Transformations[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("transformations["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if err := resource.QueryOptions.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("queryOptions", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type PanelQueryKind struct {
	Kind string         `json:"kind"`
	Spec PanelQuerySpec `json:"spec"`
}

// NewPanelQueryKind creates a new PanelQueryKind object.
func NewPanelQueryKind() *PanelQueryKind {
	return &PanelQueryKind{
		Kind: "PanelQuery",
		Spec: *NewPanelQuerySpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelQueryKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PanelQueryKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = PanelQuerySpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PanelQueryKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PanelQueryKind` objects.
func (resource PanelQueryKind) Equals(other PanelQueryKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PanelQueryKind` fields for violations and returns them.
func (resource PanelQueryKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type PanelQuerySpec struct {
	Query  DataQueryKind `json:"query"`
	RefId  string        `json:"refId"`
	Hidden bool          `json:"hidden"`
}

// NewPanelQuerySpec creates a new PanelQuerySpec object.
func NewPanelQuerySpec() *PanelQuerySpec {
	return &PanelQuerySpec{
		Query: *NewDataQueryKind(),
		RefId: "A",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelQuerySpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PanelQuerySpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {

			resource.Query = DataQueryKind{}
			if err := resource.Query.UnmarshalJSONStrict(fields["query"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")

	}
	// Field "hidden"
	if fields["hidden"] != nil {
		if string(fields["hidden"]) != "null" {
			if err := json.Unmarshal(fields["hidden"], &resource.Hidden); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hidden", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hidden", errors.New("required field is null"))...)

		}
		delete(fields, "hidden")
	} else {
		errs = append(errs, cog.MakeBuildErrors("hidden", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PanelQuerySpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PanelQuerySpec` objects.
func (resource PanelQuerySpec) Equals(other PanelQuerySpec) bool {
	if !resource.Query.Equals(other.Query) {
		return false
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hidden != other.Hidden {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PanelQuerySpec` fields for violations and returns them.
func (resource PanelQuerySpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Query.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("query", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TransformationKind struct {
	// The kind of a TransformationKind is the transformation ID
	Kind string                `json:"kind"`
	Spec DataTransformerConfig `json:"spec"`
}

// NewTransformationKind creates a new TransformationKind object.
func NewTransformationKind() *TransformationKind {
	return &TransformationKind{
		Spec: *NewDataTransformerConfig(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TransformationKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TransformationKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = DataTransformerConfig{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TransformationKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TransformationKind` objects.
func (resource TransformationKind) Equals(other TransformationKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TransformationKind` fields for violations and returns them.
func (resource TransformationKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Transformations allow to manipulate data returned by a query before the system applies a visualization.
// Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
// use the output of one transformation as the input to another transformation, etc.
type DataTransformerConfig struct {
	// Unique identifier of transformer
	Id string `json:"id"`
	// Disabled transformations are skipped
	Disabled *bool `json:"disabled,omitempty"`
	// Optional frame matcher. When missing it will be applied to all results
	Filter *MatcherConfig `json:"filter,omitempty"`
	// Where to pull DataFrames from as input to transformation
	Topic *DataTopic `json:"topic,omitempty"`
	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	Options any `json:"options"`
}

// NewDataTransformerConfig creates a new DataTransformerConfig object.
func NewDataTransformerConfig() *DataTransformerConfig {
	return &DataTransformerConfig{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataTransformerConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DataTransformerConfig) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "disabled"
	if fields["disabled"] != nil {
		if string(fields["disabled"]) != "null" {
			if err := json.Unmarshal(fields["disabled"], &resource.Disabled); err != nil {
				errs = append(errs, cog.MakeBuildErrors("disabled", err)...)
			}

		}
		delete(fields, "disabled")

	}
	// Field "filter"
	if fields["filter"] != nil {
		if string(fields["filter"]) != "null" {

			resource.Filter = &MatcherConfig{}
			if err := resource.Filter.UnmarshalJSONStrict(fields["filter"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filter", err)...)
			}

		}
		delete(fields, "filter")

	}
	// Field "topic"
	if fields["topic"] != nil {
		if string(fields["topic"]) != "null" {
			if err := json.Unmarshal(fields["topic"], &resource.Topic); err != nil {
				errs = append(errs, cog.MakeBuildErrors("topic", err)...)
			}

		}
		delete(fields, "topic")

	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {
			if err := json.Unmarshal(fields["options"], &resource.Options); err != nil {
				errs = append(errs, cog.MakeBuildErrors("options", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DataTransformerConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DataTransformerConfig` objects.
func (resource DataTransformerConfig) Equals(other DataTransformerConfig) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Disabled == nil && other.Disabled != nil || resource.Disabled != nil && other.Disabled == nil {
		return false
	}

	if resource.Disabled != nil {
		if *resource.Disabled != *other.Disabled {
			return false
		}
	}
	if resource.Filter == nil && other.Filter != nil || resource.Filter != nil && other.Filter == nil {
		return false
	}

	if resource.Filter != nil {
		if !resource.Filter.Equals(*other.Filter) {
			return false
		}
	}
	if resource.Topic == nil && other.Topic != nil || resource.Topic != nil && other.Topic == nil {
		return false
	}

	if resource.Topic != nil {
		if *resource.Topic != *other.Topic {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Options, other.Options) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DataTransformerConfig` fields for violations and returns them.
func (resource DataTransformerConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.Filter != nil {
		if err := resource.Filter.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filter", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
type MatcherConfig struct {
	// The matcher id. This is used to find the matcher implementation from registry.
	Id string `json:"id"`
	// The matcher options. This is specific to the matcher implementation.
	Options any `json:"options,omitempty"`
}

// NewMatcherConfig creates a new MatcherConfig object.
func NewMatcherConfig() *MatcherConfig {
	return &MatcherConfig{
		Id: "",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MatcherConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MatcherConfig) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MatcherConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MatcherConfig` objects.
func (resource MatcherConfig) Equals(other MatcherConfig) bool {
	if resource.Id != other.Id {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Options, other.Options) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MatcherConfig` fields for violations and returns them.
func (resource MatcherConfig) Validate() error {
	return nil
}

// A topic is attached to DataFrame metadata in query results.
// This specifies where the data should be used.
type DataTopic string

const (
	DataTopicSeries      DataTopic = "series"
	DataTopicAnnotations DataTopic = "annotations"
	DataTopicAlertStates DataTopic = "alertStates"
)

type QueryOptionsSpec struct {
	TimeFrom         *string `json:"timeFrom,omitempty"`
	MaxDataPoints    *int64  `json:"maxDataPoints,omitempty"`
	TimeShift        *string `json:"timeShift,omitempty"`
	QueryCachingTTL  *int64  `json:"queryCachingTTL,omitempty"`
	Interval         *string `json:"interval,omitempty"`
	CacheTimeout     *string `json:"cacheTimeout,omitempty"`
	HideTimeOverride *bool   `json:"hideTimeOverride,omitempty"`
	TimeCompare      *string `json:"timeCompare,omitempty"`
}

// NewQueryOptionsSpec creates a new QueryOptionsSpec object.
func NewQueryOptionsSpec() *QueryOptionsSpec {
	return &QueryOptionsSpec{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryOptionsSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryOptionsSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

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
	// Field "queryCachingTTL"
	if fields["queryCachingTTL"] != nil {
		if string(fields["queryCachingTTL"]) != "null" {
			if err := json.Unmarshal(fields["queryCachingTTL"], &resource.QueryCachingTTL); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryCachingTTL", err)...)
			}

		}
		delete(fields, "queryCachingTTL")

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
	// Field "cacheTimeout"
	if fields["cacheTimeout"] != nil {
		if string(fields["cacheTimeout"]) != "null" {
			if err := json.Unmarshal(fields["cacheTimeout"], &resource.CacheTimeout); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cacheTimeout", err)...)
			}

		}
		delete(fields, "cacheTimeout")

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
	// Field "timeCompare"
	if fields["timeCompare"] != nil {
		if string(fields["timeCompare"]) != "null" {
			if err := json.Unmarshal(fields["timeCompare"], &resource.TimeCompare); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeCompare", err)...)
			}

		}
		delete(fields, "timeCompare")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryOptionsSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryOptionsSpec` objects.
func (resource QueryOptionsSpec) Equals(other QueryOptionsSpec) bool {
	if resource.TimeFrom == nil && other.TimeFrom != nil || resource.TimeFrom != nil && other.TimeFrom == nil {
		return false
	}

	if resource.TimeFrom != nil {
		if *resource.TimeFrom != *other.TimeFrom {
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
	if resource.TimeShift == nil && other.TimeShift != nil || resource.TimeShift != nil && other.TimeShift == nil {
		return false
	}

	if resource.TimeShift != nil {
		if *resource.TimeShift != *other.TimeShift {
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
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
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
	if resource.HideTimeOverride == nil && other.HideTimeOverride != nil || resource.HideTimeOverride != nil && other.HideTimeOverride == nil {
		return false
	}

	if resource.HideTimeOverride != nil {
		if *resource.HideTimeOverride != *other.HideTimeOverride {
			return false
		}
	}
	if resource.TimeCompare == nil && other.TimeCompare != nil || resource.TimeCompare != nil && other.TimeCompare == nil {
		return false
	}

	if resource.TimeCompare != nil {
		if *resource.TimeCompare != *other.TimeCompare {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryOptionsSpec` fields for violations and returns them.
func (resource QueryOptionsSpec) Validate() error {
	return nil
}

type VizConfigKind struct {
	Kind string `json:"kind"`
	// The group is the plugin ID
	Group   string        `json:"group"`
	Version string        `json:"version"`
	Spec    VizConfigSpec `json:"spec"`
}

// NewVizConfigKind creates a new VizConfigKind object.
func NewVizConfigKind() *VizConfigKind {
	return &VizConfigKind{
		Kind: "VizConfig",
		Spec: *NewVizConfigSpec(),
	}
}
func (resource *VizConfigKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	if fields["kind"] != nil {
		if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
			return fmt.Errorf("error decoding field 'kind': %w", err)
		}
	}
	if fields["group"] != nil {
		if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
			return fmt.Errorf("error decoding field 'group': %w", err)
		}
	}
	if fields["version"] != nil {
		if err := json.Unmarshal(fields["version"], &resource.Version); err != nil {
			return fmt.Errorf("error decoding field 'version': %w", err)
		}
	}
	if fields["spec"] != nil {
		spec := VizConfigSpec{}
		specFields := make(map[string]json.RawMessage)
		if err := json.Unmarshal(fields["spec"], &specFields); err != nil {
			return err
		}

		if specFields["options"] != nil {
			variantCfg, found := cog.ConfigForPanelcfgVariant(resource.Group)
			if found && variantCfg.OptionsUnmarshaler != nil {
				options, err := variantCfg.OptionsUnmarshaler(specFields["options"])
				if err != nil {
					return err
				}
				spec.Options = options
			} else {
				if err := json.Unmarshal(specFields["options"], &spec.Options); err != nil {
					return err
				}
			}
		}

		if specFields["fieldConfig"] != nil {
			if err := json.Unmarshal(specFields["fieldConfig"], &spec.FieldConfig); err != nil {
				return fmt.Errorf("error decoding field 'fieldConfig': %w", err)
			}

			variantCfg, found := cog.ConfigForPanelcfgVariant(resource.Group)
			if found && variantCfg.FieldConfigUnmarshaler != nil {
				fakeFieldConfigSource := struct {
					Defaults struct {
						Custom json.RawMessage `json:"custom"`
					} `json:"defaults"`
				}{}
				if err := json.Unmarshal(specFields["fieldConfig"], &fakeFieldConfigSource); err != nil {
					return err
				}

				if fakeFieldConfigSource.Defaults.Custom != nil {
					customFieldConfig, err := variantCfg.FieldConfigUnmarshaler(fakeFieldConfigSource.Defaults.Custom)
					if err != nil {
						return err
					}

					spec.FieldConfig.Defaults.Custom = customFieldConfig
				}
			}
		}

		resource.Spec = spec
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizConfigKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *VizConfigKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is null"))...)

		}
		delete(fields, "group")
	} else {
		errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is missing from input"))...)
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
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {
			spec := VizConfigSpec{}
			specFields := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["spec"], &specFields); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			} else {
				if specFields["options"] != nil {
					variantCfg, found := cog.ConfigForPanelcfgVariant(resource.Group)
					if found && variantCfg.StrictOptionsUnmarshaler != nil {
						options, err := variantCfg.StrictOptionsUnmarshaler(specFields["options"])
						if err != nil {
							errs = append(errs, cog.MakeBuildErrors("spec.options", err)...)
						} else {
							spec.Options = options
						}
					} else {
						if err := json.Unmarshal(specFields["options"], &spec.Options); err != nil {
							errs = append(errs, cog.MakeBuildErrors("spec.options", err)...)
						}
					}
				}

				if specFields["fieldConfig"] != nil {
					if err := json.Unmarshal(specFields["fieldConfig"], &spec.FieldConfig); err != nil {
						return fmt.Errorf("error decoding field 'fieldConfig': %w", err)
					}

					variantCfg, found := cog.ConfigForPanelcfgVariant(resource.Group)
					if found && variantCfg.FieldConfigUnmarshaler != nil {
						fakeFieldConfigSource := struct {
							Defaults struct {
								Custom json.RawMessage `json:"custom"`
							} `json:"defaults"`
						}{}
						if err := json.Unmarshal(specFields["fieldConfig"], &fakeFieldConfigSource); err != nil {
							return err
						}

						if fakeFieldConfigSource.Defaults.Custom != nil {
							customFieldConfig, err := variantCfg.FieldConfigUnmarshaler(fakeFieldConfigSource.Defaults.Custom)
							if err != nil {
								return err
							}

							spec.FieldConfig.Defaults.Custom = customFieldConfig
						}
					}
				}

			}

			resource.Spec = spec

		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("VizConfigKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `VizConfigKind` objects.
func (resource VizConfigKind) Equals(other VizConfigKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Group != other.Group {
		return false
	}
	if resource.Version != other.Version {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `VizConfigKind` fields for violations and returns them.
func (resource VizConfigKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// --- Kinds ---
type VizConfigSpec struct {
	Options     any               `json:"options"`
	FieldConfig FieldConfigSource `json:"fieldConfig"`
}

// NewVizConfigSpec creates a new VizConfigSpec object.
func NewVizConfigSpec() *VizConfigSpec {
	return &VizConfigSpec{
		FieldConfig: *NewFieldConfigSource(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizConfigSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *VizConfigSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {
			if err := json.Unmarshal(fields["options"], &resource.Options); err != nil {
				errs = append(errs, cog.MakeBuildErrors("options", err)...)
			}

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}
	// Field "fieldConfig"
	if fields["fieldConfig"] != nil {
		if string(fields["fieldConfig"]) != "null" {

			resource.FieldConfig = FieldConfigSource{}
			if err := resource.FieldConfig.UnmarshalJSONStrict(fields["fieldConfig"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fieldConfig", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("fieldConfig", errors.New("required field is null"))...)

		}
		delete(fields, "fieldConfig")
	} else {
		errs = append(errs, cog.MakeBuildErrors("fieldConfig", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("VizConfigSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `VizConfigSpec` objects.
func (resource VizConfigSpec) Equals(other VizConfigSpec) bool {
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Options, other.Options) {
		return false
	}
	if !resource.FieldConfig.Equals(other.FieldConfig) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `VizConfigSpec` fields for violations and returns them.
func (resource VizConfigSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.FieldConfig.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("fieldConfig", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
type FieldConfigSource struct {
	// Defaults are the options applied to all fields.
	Defaults FieldConfig `json:"defaults"`
	// Overrides are the options applied to specific fields overriding the defaults.
	Overrides []Dashboardv2beta1FieldConfigSourceOverrides `json:"overrides"`
}

// NewFieldConfigSource creates a new FieldConfigSource object.
func NewFieldConfigSource() *FieldConfigSource {
	return &FieldConfigSource{
		Defaults:  *NewFieldConfig(),
		Overrides: []Dashboardv2beta1FieldConfigSourceOverrides{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfigSource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FieldConfigSource) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "defaults"
	if fields["defaults"] != nil {
		if string(fields["defaults"]) != "null" {

			resource.Defaults = FieldConfig{}
			if err := resource.Defaults.UnmarshalJSONStrict(fields["defaults"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("defaults", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("defaults", errors.New("required field is null"))...)

		}
		delete(fields, "defaults")
	} else {
		errs = append(errs, cog.MakeBuildErrors("defaults", errors.New("required field is missing from input"))...)
	}
	// Field "overrides"
	if fields["overrides"] != nil {
		if string(fields["overrides"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["overrides"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Dashboardv2beta1FieldConfigSourceOverrides

				result1 = Dashboardv2beta1FieldConfigSourceOverrides{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("overrides["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Overrides = append(resource.Overrides, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("overrides", errors.New("required field is null"))...)

		}
		delete(fields, "overrides")
	} else {
		errs = append(errs, cog.MakeBuildErrors("overrides", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FieldConfigSource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FieldConfigSource` objects.
func (resource FieldConfigSource) Equals(other FieldConfigSource) bool {
	if !resource.Defaults.Equals(other.Defaults) {
		return false
	}

	if len(resource.Overrides) != len(other.Overrides) {
		return false
	}

	for i1 := range resource.Overrides {
		if !resource.Overrides[i1].Equals(other.Overrides[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FieldConfigSource` fields for violations and returns them.
func (resource FieldConfigSource) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Defaults.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("defaults", err)...)
	}

	for i1 := range resource.Overrides {
		if err := resource.Overrides[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("overrides["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
type FieldConfig struct {
	// The display value for this field.  This supports template variables blank is auto
	DisplayName *string `json:"displayName,omitempty"`
	// This can be used by data sources that return and explicit naming structure for values and labels
	// When this property is configured, this value is used rather than the default naming strategy.
	DisplayNameFromDS *string `json:"displayNameFromDS,omitempty"`
	// Human readable field metadata
	Description *string `json:"description,omitempty"`
	// An explicit path to the field in the datasource.  When the frame meta includes a path,
	// This will default to `${frame.meta.path}/${field.name}
	//
	// When defined, this value can be used as an identifier within the datasource scope, and
	// may be used to update the results
	Path *string `json:"path,omitempty"`
	// True if data source can write a value to the path. Auth/authz are supported separately
	Writeable *bool `json:"writeable,omitempty"`
	// True if data source field supports ad-hoc filters
	Filterable *bool `json:"filterable,omitempty"`
	// Unit a field should use. The unit you select is applied to all fields except time.
	// You can use the units ID availables in Grafana or a custom unit.
	// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
	// As custom unit, you can use the following formats:
	// `suffix:<suffix>` for custom unit that should go after value.
	// `prefix:<prefix>` for custom unit that should go before value.
	// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
	// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
	// `count:<unit>` for a custom count unit.
	// `currency:<unit>` for custom a currency unit.
	Unit *string `json:"unit,omitempty"`
	// Specify the number of decimals Grafana includes in the rendered value.
	// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
	// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
	// To display all decimals, set the unit to `String`.
	Decimals *float64 `json:"decimals,omitempty"`
	// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	Min *float64 `json:"min,omitempty"`
	// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	Max *float64 `json:"max,omitempty"`
	// Convert input values into a display string
	Mappings []ValueMapping `json:"mappings,omitempty"`
	// Map numeric values to states
	Thresholds *ThresholdsConfig `json:"thresholds,omitempty"`
	// Panel color configuration
	Color *FieldColor `json:"color,omitempty"`
	// The behavior when clicking on a result
	Links []any `json:"links,omitempty"`
	// Define interactive HTTP requests that can be triggered from data visualizations.
	Actions []Action `json:"actions,omitempty"`
	// Alternative to empty string
	NoValue *string `json:"noValue,omitempty"`
	// custom is specified by the FieldConfig field
	// in panel plugin schemas.
	Custom any `json:"custom,omitempty"`
	// Calculate min max per field
	FieldMinMax *bool `json:"fieldMinMax,omitempty"`
	// How null values should be handled when calculating field stats
	// "null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero
	NullValueMode *NullValueMode `json:"nullValueMode,omitempty"`
}

// NewFieldConfig creates a new FieldConfig object.
func NewFieldConfig() *FieldConfig {
	return &FieldConfig{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FieldConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "displayName"
	if fields["displayName"] != nil {
		if string(fields["displayName"]) != "null" {
			if err := json.Unmarshal(fields["displayName"], &resource.DisplayName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayName", err)...)
			}

		}
		delete(fields, "displayName")

	}
	// Field "displayNameFromDS"
	if fields["displayNameFromDS"] != nil {
		if string(fields["displayNameFromDS"]) != "null" {
			if err := json.Unmarshal(fields["displayNameFromDS"], &resource.DisplayNameFromDS); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayNameFromDS", err)...)
			}

		}
		delete(fields, "displayNameFromDS")

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
	// Field "path"
	if fields["path"] != nil {
		if string(fields["path"]) != "null" {
			if err := json.Unmarshal(fields["path"], &resource.Path); err != nil {
				errs = append(errs, cog.MakeBuildErrors("path", err)...)
			}

		}
		delete(fields, "path")

	}
	// Field "writeable"
	if fields["writeable"] != nil {
		if string(fields["writeable"]) != "null" {
			if err := json.Unmarshal(fields["writeable"], &resource.Writeable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("writeable", err)...)
			}

		}
		delete(fields, "writeable")

	}
	// Field "filterable"
	if fields["filterable"] != nil {
		if string(fields["filterable"]) != "null" {
			if err := json.Unmarshal(fields["filterable"], &resource.Filterable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filterable", err)...)
			}

		}
		delete(fields, "filterable")

	}
	// Field "unit"
	if fields["unit"] != nil {
		if string(fields["unit"]) != "null" {
			if err := json.Unmarshal(fields["unit"], &resource.Unit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("unit", err)...)
			}

		}
		delete(fields, "unit")

	}
	// Field "decimals"
	if fields["decimals"] != nil {
		if string(fields["decimals"]) != "null" {
			if err := json.Unmarshal(fields["decimals"], &resource.Decimals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("decimals", err)...)
			}

		}
		delete(fields, "decimals")

	}
	// Field "min"
	if fields["min"] != nil {
		if string(fields["min"]) != "null" {
			if err := json.Unmarshal(fields["min"], &resource.Min); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min", err)...)
			}

		}
		delete(fields, "min")

	}
	// Field "max"
	if fields["max"] != nil {
		if string(fields["max"]) != "null" {
			if err := json.Unmarshal(fields["max"], &resource.Max); err != nil {
				errs = append(errs, cog.MakeBuildErrors("max", err)...)
			}

		}
		delete(fields, "max")

	}
	// Field "mappings"
	if fields["mappings"] != nil {
		if string(fields["mappings"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["mappings"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 ValueMapping

				result1 = ValueMapping{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("mappings["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Mappings = append(resource.Mappings, result1)
			}

		}
		delete(fields, "mappings")

	}
	// Field "thresholds"
	if fields["thresholds"] != nil {
		if string(fields["thresholds"]) != "null" {

			resource.Thresholds = &ThresholdsConfig{}
			if err := resource.Thresholds.UnmarshalJSONStrict(fields["thresholds"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("thresholds", err)...)
			}

		}
		delete(fields, "thresholds")

	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {

			resource.Color = &FieldColor{}
			if err := resource.Color.UnmarshalJSONStrict(fields["color"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}

		}
		delete(fields, "color")

	}
	// Field "links"
	if fields["links"] != nil {
		if string(fields["links"]) != "null" {

			if err := json.Unmarshal(fields["links"], &resource.Links); err != nil {
				errs = append(errs, cog.MakeBuildErrors("links", err)...)
			}

		}
		delete(fields, "links")

	}
	// Field "actions"
	if fields["actions"] != nil {
		if string(fields["actions"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["actions"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Action

				result1 = Action{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("actions["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Actions = append(resource.Actions, result1)
			}

		}
		delete(fields, "actions")

	}
	// Field "noValue"
	if fields["noValue"] != nil {
		if string(fields["noValue"]) != "null" {
			if err := json.Unmarshal(fields["noValue"], &resource.NoValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("noValue", err)...)
			}

		}
		delete(fields, "noValue")

	}
	// Field "custom"
	if fields["custom"] != nil {
		if string(fields["custom"]) != "null" {
			if err := json.Unmarshal(fields["custom"], &resource.Custom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("custom", err)...)
			}

		}
		delete(fields, "custom")

	}
	// Field "fieldMinMax"
	if fields["fieldMinMax"] != nil {
		if string(fields["fieldMinMax"]) != "null" {
			if err := json.Unmarshal(fields["fieldMinMax"], &resource.FieldMinMax); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fieldMinMax", err)...)
			}

		}
		delete(fields, "fieldMinMax")

	}
	// Field "nullValueMode"
	if fields["nullValueMode"] != nil {
		if string(fields["nullValueMode"]) != "null" {
			if err := json.Unmarshal(fields["nullValueMode"], &resource.NullValueMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("nullValueMode", err)...)
			}

		}
		delete(fields, "nullValueMode")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FieldConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FieldConfig` objects.
func (resource FieldConfig) Equals(other FieldConfig) bool {
	if resource.DisplayName == nil && other.DisplayName != nil || resource.DisplayName != nil && other.DisplayName == nil {
		return false
	}

	if resource.DisplayName != nil {
		if *resource.DisplayName != *other.DisplayName {
			return false
		}
	}
	if resource.DisplayNameFromDS == nil && other.DisplayNameFromDS != nil || resource.DisplayNameFromDS != nil && other.DisplayNameFromDS == nil {
		return false
	}

	if resource.DisplayNameFromDS != nil {
		if *resource.DisplayNameFromDS != *other.DisplayNameFromDS {
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
	if resource.Path == nil && other.Path != nil || resource.Path != nil && other.Path == nil {
		return false
	}

	if resource.Path != nil {
		if *resource.Path != *other.Path {
			return false
		}
	}
	if resource.Writeable == nil && other.Writeable != nil || resource.Writeable != nil && other.Writeable == nil {
		return false
	}

	if resource.Writeable != nil {
		if *resource.Writeable != *other.Writeable {
			return false
		}
	}
	if resource.Filterable == nil && other.Filterable != nil || resource.Filterable != nil && other.Filterable == nil {
		return false
	}

	if resource.Filterable != nil {
		if *resource.Filterable != *other.Filterable {
			return false
		}
	}
	if resource.Unit == nil && other.Unit != nil || resource.Unit != nil && other.Unit == nil {
		return false
	}

	if resource.Unit != nil {
		if *resource.Unit != *other.Unit {
			return false
		}
	}
	if resource.Decimals == nil && other.Decimals != nil || resource.Decimals != nil && other.Decimals == nil {
		return false
	}

	if resource.Decimals != nil {
		if *resource.Decimals != *other.Decimals {
			return false
		}
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if *resource.Min != *other.Min {
			return false
		}
	}
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if *resource.Max != *other.Max {
			return false
		}
	}

	if len(resource.Mappings) != len(other.Mappings) {
		return false
	}

	for i1 := range resource.Mappings {
		if !resource.Mappings[i1].Equals(other.Mappings[i1]) {
			return false
		}
	}
	if resource.Thresholds == nil && other.Thresholds != nil || resource.Thresholds != nil && other.Thresholds == nil {
		return false
	}

	if resource.Thresholds != nil {
		if !resource.Thresholds.Equals(*other.Thresholds) {
			return false
		}
	}
	if resource.Color == nil && other.Color != nil || resource.Color != nil && other.Color == nil {
		return false
	}

	if resource.Color != nil {
		if !resource.Color.Equals(*other.Color) {
			return false
		}
	}

	if len(resource.Links) != len(other.Links) {
		return false
	}

	for i1 := range resource.Links {
		// is DeepEqual good enough here?
		if !reflect.DeepEqual(resource.Links[i1], other.Links[i1]) {
			return false
		}
	}

	if len(resource.Actions) != len(other.Actions) {
		return false
	}

	for i1 := range resource.Actions {
		if !resource.Actions[i1].Equals(other.Actions[i1]) {
			return false
		}
	}
	if resource.NoValue == nil && other.NoValue != nil || resource.NoValue != nil && other.NoValue == nil {
		return false
	}

	if resource.NoValue != nil {
		if *resource.NoValue != *other.NoValue {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Custom, other.Custom) {
		return false
	}
	if resource.FieldMinMax == nil && other.FieldMinMax != nil || resource.FieldMinMax != nil && other.FieldMinMax == nil {
		return false
	}

	if resource.FieldMinMax != nil {
		if *resource.FieldMinMax != *other.FieldMinMax {
			return false
		}
	}
	if resource.NullValueMode == nil && other.NullValueMode != nil || resource.NullValueMode != nil && other.NullValueMode == nil {
		return false
	}

	if resource.NullValueMode != nil {
		if *resource.NullValueMode != *other.NullValueMode {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.
func (resource FieldConfig) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Mappings {
		if err := resource.Mappings[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("mappings["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Thresholds != nil {
		if err := resource.Thresholds.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("thresholds", err)...)
		}
	}
	if resource.Color != nil {
		if err := resource.Color.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("color", err)...)
		}
	}

	for i1 := range resource.Actions {
		if err := resource.Actions[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("actions["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ValueMapping = ValueMapOrRangeMapOrRegexMapOrSpecialValueMap

// NewValueMapping creates a new ValueMapping object.
func NewValueMapping() *ValueMapping {
	return NewValueMapOrRangeMapOrRegexMapOrSpecialValueMap()
}

// Maps text values to a color or different display text and color.
// For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
type ValueMap struct {
	Type MappingType `json:"type"`
	// Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
	Options map[string]ValueMappingResult `json:"options"`
}

// NewValueMap creates a new ValueMap object.
func NewValueMap() *ValueMap {
	return &ValueMap{
		Type:    MappingTypeValue,
		Options: map[string]ValueMappingResult{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ValueMap` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ValueMap) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["options"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]ValueMappingResult, len(partialMap))
			for key1 := range partialMap {
				var result1 ValueMappingResult

				result1 = ValueMappingResult{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("options["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Options = parsedMap1
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ValueMap", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ValueMap` objects.
func (resource ValueMap) Equals(other ValueMap) bool {
	if resource.Type != other.Type {
		return false
	}

	if len(resource.Options) != len(other.Options) {
		return false
	}

	for key1 := range resource.Options {
		if !resource.Options[key1].Equals(other.Options[key1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ValueMap` fields for violations and returns them.
func (resource ValueMap) Validate() error {
	var errs cog.BuildErrors
	if resource.Type != "value" {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("must be value"))...)
	}

	for key1 := range resource.Options {
		if err := resource.Options[key1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("options["+key1+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Supported value mapping types
// `value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
// `range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
// `regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
// `special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.
type MappingType string

const (
	MappingTypeValue   MappingType = "value"
	MappingTypeRange   MappingType = "range"
	MappingTypeRegex   MappingType = "regex"
	MappingTypeSpecial MappingType = "special"
)

// Result used as replacement with text and color when the value matches
type ValueMappingResult struct {
	// Text to display when the value matches
	Text *string `json:"text,omitempty"`
	// Text to use when the value matches
	Color *string `json:"color,omitempty"`
	// Icon to display when the value matches. Only specific visualizations.
	Icon *string `json:"icon,omitempty"`
	// Position in the mapping array. Only used internally.
	Index *int32 `json:"index,omitempty"`
}

// NewValueMappingResult creates a new ValueMappingResult object.
func NewValueMappingResult() *ValueMappingResult {
	return &ValueMappingResult{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ValueMappingResult` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ValueMappingResult) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {
			if err := json.Unmarshal(fields["text"], &resource.Text); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}

		}
		delete(fields, "text")

	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {
			if err := json.Unmarshal(fields["color"], &resource.Color); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}

		}
		delete(fields, "color")

	}
	// Field "icon"
	if fields["icon"] != nil {
		if string(fields["icon"]) != "null" {
			if err := json.Unmarshal(fields["icon"], &resource.Icon); err != nil {
				errs = append(errs, cog.MakeBuildErrors("icon", err)...)
			}

		}
		delete(fields, "icon")

	}
	// Field "index"
	if fields["index"] != nil {
		if string(fields["index"]) != "null" {
			if err := json.Unmarshal(fields["index"], &resource.Index); err != nil {
				errs = append(errs, cog.MakeBuildErrors("index", err)...)
			}

		}
		delete(fields, "index")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ValueMappingResult", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ValueMappingResult` objects.
func (resource ValueMappingResult) Equals(other ValueMappingResult) bool {
	if resource.Text == nil && other.Text != nil || resource.Text != nil && other.Text == nil {
		return false
	}

	if resource.Text != nil {
		if *resource.Text != *other.Text {
			return false
		}
	}
	if resource.Color == nil && other.Color != nil || resource.Color != nil && other.Color == nil {
		return false
	}

	if resource.Color != nil {
		if *resource.Color != *other.Color {
			return false
		}
	}
	if resource.Icon == nil && other.Icon != nil || resource.Icon != nil && other.Icon == nil {
		return false
	}

	if resource.Icon != nil {
		if *resource.Icon != *other.Icon {
			return false
		}
	}
	if resource.Index == nil && other.Index != nil || resource.Index != nil && other.Index == nil {
		return false
	}

	if resource.Index != nil {
		if *resource.Index != *other.Index {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ValueMappingResult` fields for violations and returns them.
func (resource ValueMappingResult) Validate() error {
	return nil
}

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
type RangeMap struct {
	Type MappingType `json:"type"`
	// Range to match against and the result to apply when the value is within the range
	Options Dashboardv2beta1RangeMapOptions `json:"options"`
}

// NewRangeMap creates a new RangeMap object.
func NewRangeMap() *RangeMap {
	return &RangeMap{
		Type:    MappingTypeRange,
		Options: *NewDashboardv2beta1RangeMapOptions(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RangeMap` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RangeMap) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			resource.Options = Dashboardv2beta1RangeMapOptions{}
			if err := resource.Options.UnmarshalJSONStrict(fields["options"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("options", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RangeMap", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RangeMap` objects.
func (resource RangeMap) Equals(other RangeMap) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Options.Equals(other.Options) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RangeMap` fields for violations and returns them.
func (resource RangeMap) Validate() error {
	var errs cog.BuildErrors
	if resource.Type != "range" {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("must be range"))...)
	}
	if err := resource.Options.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("options", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
type RegexMap struct {
	Type MappingType `json:"type"`
	// Regular expression to match against and the result to apply when the value matches the regex
	Options Dashboardv2beta1RegexMapOptions `json:"options"`
}

// NewRegexMap creates a new RegexMap object.
func NewRegexMap() *RegexMap {
	return &RegexMap{
		Type:    MappingTypeRegex,
		Options: *NewDashboardv2beta1RegexMapOptions(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RegexMap` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RegexMap) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			resource.Options = Dashboardv2beta1RegexMapOptions{}
			if err := resource.Options.UnmarshalJSONStrict(fields["options"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("options", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RegexMap", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RegexMap` objects.
func (resource RegexMap) Equals(other RegexMap) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Options.Equals(other.Options) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RegexMap` fields for violations and returns them.
func (resource RegexMap) Validate() error {
	var errs cog.BuildErrors
	if resource.Type != "regex" {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("must be regex"))...)
	}
	if err := resource.Options.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("options", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
// See SpecialValueMatch to see the list of special values.
// For example, you can configure a special value mapping so that null values appear as N/A.
type SpecialValueMap struct {
	Type    MappingType                            `json:"type"`
	Options Dashboardv2beta1SpecialValueMapOptions `json:"options"`
}

// NewSpecialValueMap creates a new SpecialValueMap object.
func NewSpecialValueMap() *SpecialValueMap {
	return &SpecialValueMap{
		Type:    MappingTypeSpecial,
		Options: *NewDashboardv2beta1SpecialValueMapOptions(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SpecialValueMap` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SpecialValueMap) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			resource.Options = Dashboardv2beta1SpecialValueMapOptions{}
			if err := resource.Options.UnmarshalJSONStrict(fields["options"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("options", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SpecialValueMap", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SpecialValueMap` objects.
func (resource SpecialValueMap) Equals(other SpecialValueMap) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Options.Equals(other.Options) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SpecialValueMap` fields for violations and returns them.
func (resource SpecialValueMap) Validate() error {
	var errs cog.BuildErrors
	if resource.Type != "special" {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("must be special"))...)
	}
	if err := resource.Options.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("options", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Special value types supported by the `SpecialValueMap`
type SpecialValueMatch string

const (
	SpecialValueMatchTrue       SpecialValueMatch = "true"
	SpecialValueMatchFalse      SpecialValueMatch = "false"
	SpecialValueMatchNull       SpecialValueMatch = "null"
	SpecialValueMatchNaN        SpecialValueMatch = "nan"
	SpecialValueMatchNullAndNaN SpecialValueMatch = "null+nan"
	SpecialValueMatchEmpty      SpecialValueMatch = "empty"
)

type ThresholdsConfig struct {
	Mode  ThresholdsMode `json:"mode"`
	Steps []Threshold    `json:"steps"`
}

// NewThresholdsConfig creates a new ThresholdsConfig object.
func NewThresholdsConfig() *ThresholdsConfig {
	return &ThresholdsConfig{
		Steps: []Threshold{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ThresholdsConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ThresholdsConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "steps"
	if fields["steps"] != nil {
		if string(fields["steps"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["steps"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Threshold

				result1 = Threshold{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("steps["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Steps = append(resource.Steps, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("steps", errors.New("required field is null"))...)

		}
		delete(fields, "steps")
	} else {
		errs = append(errs, cog.MakeBuildErrors("steps", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ThresholdsConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ThresholdsConfig` objects.
func (resource ThresholdsConfig) Equals(other ThresholdsConfig) bool {
	if resource.Mode != other.Mode {
		return false
	}

	if len(resource.Steps) != len(other.Steps) {
		return false
	}

	for i1 := range resource.Steps {
		if !resource.Steps[i1].Equals(other.Steps[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ThresholdsConfig` fields for violations and returns them.
func (resource ThresholdsConfig) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Steps {
		if err := resource.Steps[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("steps["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ThresholdsMode string

const (
	ThresholdsModeAbsolute   ThresholdsMode = "absolute"
	ThresholdsModePercentage ThresholdsMode = "percentage"
)

type Threshold struct {
	// Value null means -Infinity
	Value *float64 `json:"value"`
	Color string   `json:"color"`
}

// NewThreshold creates a new Threshold object.
func NewThreshold() *Threshold {
	return &Threshold{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Threshold` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Threshold) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {
			if err := json.Unmarshal(fields["color"], &resource.Color); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("color", errors.New("required field is null"))...)

		}
		delete(fields, "color")
	} else {
		errs = append(errs, cog.MakeBuildErrors("color", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Threshold", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Threshold` objects.
func (resource Threshold) Equals(other Threshold) bool {
	if resource.Value == nil && other.Value != nil || resource.Value != nil && other.Value == nil {
		return false
	}

	if resource.Value != nil {
		if *resource.Value != *other.Value {
			return false
		}
	}
	if resource.Color != other.Color {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Threshold` fields for violations and returns them.
func (resource Threshold) Validate() error {
	return nil
}

// Map a field to a color.
type FieldColor struct {
	// The main color scheme mode.
	Mode FieldColorModeId `json:"mode"`
	// The fixed color value for fixed or shades color modes.
	FixedColor *string `json:"fixedColor,omitempty"`
	// Some visualizations need to know how to assign a series color from by value color schemes.
	SeriesBy *FieldColorSeriesByMode `json:"seriesBy,omitempty"`
}

// NewFieldColor creates a new FieldColor object.
func NewFieldColor() *FieldColor {
	return &FieldColor{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldColor` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FieldColor) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "fixedColor"
	if fields["fixedColor"] != nil {
		if string(fields["fixedColor"]) != "null" {
			if err := json.Unmarshal(fields["fixedColor"], &resource.FixedColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixedColor", err)...)
			}

		}
		delete(fields, "fixedColor")

	}
	// Field "seriesBy"
	if fields["seriesBy"] != nil {
		if string(fields["seriesBy"]) != "null" {
			if err := json.Unmarshal(fields["seriesBy"], &resource.SeriesBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("seriesBy", err)...)
			}

		}
		delete(fields, "seriesBy")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FieldColor", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FieldColor` objects.
func (resource FieldColor) Equals(other FieldColor) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.FixedColor == nil && other.FixedColor != nil || resource.FixedColor != nil && other.FixedColor == nil {
		return false
	}

	if resource.FixedColor != nil {
		if *resource.FixedColor != *other.FixedColor {
			return false
		}
	}
	if resource.SeriesBy == nil && other.SeriesBy != nil || resource.SeriesBy != nil && other.SeriesBy == nil {
		return false
	}

	if resource.SeriesBy != nil {
		if *resource.SeriesBy != *other.SeriesBy {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FieldColor` fields for violations and returns them.
func (resource FieldColor) Validate() error {
	return nil
}

// Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.
// Continuous color interpolates a color using the percentage of a value relative to min and max.
// Accepted values are:
// `thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold
// `palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations
// `palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations
// `continuous-viridis`: Continuous Viridis palette mode
// `continuous-magma`: Continuous Magma palette mode
// `continuous-plasma`: Continuous Plasma palette mode
// `continuous-inferno`: Continuous Inferno palette mode
// `continuous-cividis`: Continuous Cividis palette mode
// `continuous-GrYlRd`: Continuous Green-Yellow-Red palette mode
// `continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode
// `continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode
// `continuous-YlRd`: Continuous Yellow-Red palette mode
// `continuous-BlPu`: Continuous Blue-Purple palette mode
// `continuous-YlBl`: Continuous Yellow-Blue palette mode
// `continuous-blues`: Continuous Blue palette mode
// `continuous-reds`: Continuous Red palette mode
// `continuous-greens`: Continuous Green palette mode
// `continuous-purples`: Continuous Purple palette mode
// `shades`: Shades of a single color. Specify a single color, useful in an override rule.
// `fixed`: Fixed color mode. Specify a single color, useful in an override rule.
type FieldColorModeId string

const (
	FieldColorModeIdThresholds           FieldColorModeId = "thresholds"
	FieldColorModeIdPaletteClassic       FieldColorModeId = "palette-classic"
	FieldColorModeIdPaletteClassicByName FieldColorModeId = "palette-classic-by-name"
	FieldColorModeIdContinuousViridis    FieldColorModeId = "continuous-viridis"
	FieldColorModeIdContinuousMagma      FieldColorModeId = "continuous-magma"
	FieldColorModeIdContinuousPlasma     FieldColorModeId = "continuous-plasma"
	FieldColorModeIdContinuousInferno    FieldColorModeId = "continuous-inferno"
	FieldColorModeIdContinuousCividis    FieldColorModeId = "continuous-cividis"
	FieldColorModeIdContinuousGrYlRd     FieldColorModeId = "continuous-GrYlRd"
	FieldColorModeIdContinuousRdYlGr     FieldColorModeId = "continuous-RdYlGr"
	FieldColorModeIdContinuousBlYlRd     FieldColorModeId = "continuous-BlYlRd"
	FieldColorModeIdContinuousYlRd       FieldColorModeId = "continuous-YlRd"
	FieldColorModeIdContinuousBlPu       FieldColorModeId = "continuous-BlPu"
	FieldColorModeIdContinuousYlBl       FieldColorModeId = "continuous-YlBl"
	FieldColorModeIdContinuousBlues      FieldColorModeId = "continuous-blues"
	FieldColorModeIdContinuousReds       FieldColorModeId = "continuous-reds"
	FieldColorModeIdContinuousGreens     FieldColorModeId = "continuous-greens"
	FieldColorModeIdContinuousPurples    FieldColorModeId = "continuous-purples"
	FieldColorModeIdFixed                FieldColorModeId = "fixed"
	FieldColorModeIdShades               FieldColorModeId = "shades"
)

// Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
type FieldColorSeriesByMode string

const (
	FieldColorSeriesByModeMin  FieldColorSeriesByMode = "min"
	FieldColorSeriesByModeMax  FieldColorSeriesByMode = "max"
	FieldColorSeriesByModeLast FieldColorSeriesByMode = "last"
)

type Action struct {
	Type         ActionType                   `json:"type"`
	Title        string                       `json:"title"`
	Fetch        *FetchOptions                `json:"fetch,omitempty"`
	Infinity     *InfinityOptions             `json:"infinity,omitempty"`
	Confirmation *string                      `json:"confirmation,omitempty"`
	OneClick     *bool                        `json:"oneClick,omitempty"`
	Variables    []ActionVariable             `json:"variables,omitempty"`
	Style        *Dashboardv2beta1ActionStyle `json:"style,omitempty"`
}

// NewAction creates a new Action object.
func NewAction() *Action {
	return &Action{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Action` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Action) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}
	// Field "fetch"
	if fields["fetch"] != nil {
		if string(fields["fetch"]) != "null" {

			resource.Fetch = &FetchOptions{}
			if err := resource.Fetch.UnmarshalJSONStrict(fields["fetch"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fetch", err)...)
			}

		}
		delete(fields, "fetch")

	}
	// Field "infinity"
	if fields["infinity"] != nil {
		if string(fields["infinity"]) != "null" {

			resource.Infinity = &InfinityOptions{}
			if err := resource.Infinity.UnmarshalJSONStrict(fields["infinity"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("infinity", err)...)
			}

		}
		delete(fields, "infinity")

	}
	// Field "confirmation"
	if fields["confirmation"] != nil {
		if string(fields["confirmation"]) != "null" {
			if err := json.Unmarshal(fields["confirmation"], &resource.Confirmation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("confirmation", err)...)
			}

		}
		delete(fields, "confirmation")

	}
	// Field "oneClick"
	if fields["oneClick"] != nil {
		if string(fields["oneClick"]) != "null" {
			if err := json.Unmarshal(fields["oneClick"], &resource.OneClick); err != nil {
				errs = append(errs, cog.MakeBuildErrors("oneClick", err)...)
			}

		}
		delete(fields, "oneClick")

	}
	// Field "variables"
	if fields["variables"] != nil {
		if string(fields["variables"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["variables"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 ActionVariable

				result1 = ActionVariable{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("variables["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Variables = append(resource.Variables, result1)
			}

		}
		delete(fields, "variables")

	}
	// Field "style"
	if fields["style"] != nil {
		if string(fields["style"]) != "null" {

			resource.Style = &Dashboardv2beta1ActionStyle{}
			if err := resource.Style.UnmarshalJSONStrict(fields["style"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("style", err)...)
			}

		}
		delete(fields, "style")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Action", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Action` objects.
func (resource Action) Equals(other Action) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Title != other.Title {
		return false
	}
	if resource.Fetch == nil && other.Fetch != nil || resource.Fetch != nil && other.Fetch == nil {
		return false
	}

	if resource.Fetch != nil {
		if !resource.Fetch.Equals(*other.Fetch) {
			return false
		}
	}
	if resource.Infinity == nil && other.Infinity != nil || resource.Infinity != nil && other.Infinity == nil {
		return false
	}

	if resource.Infinity != nil {
		if !resource.Infinity.Equals(*other.Infinity) {
			return false
		}
	}
	if resource.Confirmation == nil && other.Confirmation != nil || resource.Confirmation != nil && other.Confirmation == nil {
		return false
	}

	if resource.Confirmation != nil {
		if *resource.Confirmation != *other.Confirmation {
			return false
		}
	}
	if resource.OneClick == nil && other.OneClick != nil || resource.OneClick != nil && other.OneClick == nil {
		return false
	}

	if resource.OneClick != nil {
		if *resource.OneClick != *other.OneClick {
			return false
		}
	}

	if len(resource.Variables) != len(other.Variables) {
		return false
	}

	for i1 := range resource.Variables {
		if !resource.Variables[i1].Equals(other.Variables[i1]) {
			return false
		}
	}
	if resource.Style == nil && other.Style != nil || resource.Style != nil && other.Style == nil {
		return false
	}

	if resource.Style != nil {
		if !resource.Style.Equals(*other.Style) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Action` fields for violations and returns them.
func (resource Action) Validate() error {
	var errs cog.BuildErrors
	if resource.Fetch != nil {
		if err := resource.Fetch.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("fetch", err)...)
		}
	}
	if resource.Infinity != nil {
		if err := resource.Infinity.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("infinity", err)...)
		}
	}

	for i1 := range resource.Variables {
		if err := resource.Variables[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("variables["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Style != nil {
		if err := resource.Style.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("style", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ActionType string

const (
	ActionTypeFetch    ActionType = "fetch"
	ActionTypeInfinity ActionType = "infinity"
)

type FetchOptions struct {
	Method HttpRequestMethod `json:"method"`
	Url    string            `json:"url"`
	Body   *string           `json:"body,omitempty"`
	// These are 2D arrays of strings, each representing a key-value pair
	// We are defining them this way because we can't generate a go struct that
	// that would have exactly two strings in each sub-array
	QueryParams [][]string `json:"queryParams,omitempty"`
	Headers     [][]string `json:"headers,omitempty"`
}

// NewFetchOptions creates a new FetchOptions object.
func NewFetchOptions() *FetchOptions {
	return &FetchOptions{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FetchOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FetchOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "method"
	if fields["method"] != nil {
		if string(fields["method"]) != "null" {
			if err := json.Unmarshal(fields["method"], &resource.Method); err != nil {
				errs = append(errs, cog.MakeBuildErrors("method", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("method", errors.New("required field is null"))...)

		}
		delete(fields, "method")
	} else {
		errs = append(errs, cog.MakeBuildErrors("method", errors.New("required field is missing from input"))...)
	}
	// Field "url"
	if fields["url"] != nil {
		if string(fields["url"]) != "null" {
			if err := json.Unmarshal(fields["url"], &resource.Url); err != nil {
				errs = append(errs, cog.MakeBuildErrors("url", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("url", errors.New("required field is null"))...)

		}
		delete(fields, "url")
	} else {
		errs = append(errs, cog.MakeBuildErrors("url", errors.New("required field is missing from input"))...)
	}
	// Field "body"
	if fields["body"] != nil {
		if string(fields["body"]) != "null" {
			if err := json.Unmarshal(fields["body"], &resource.Body); err != nil {
				errs = append(errs, cog.MakeBuildErrors("body", err)...)
			}

		}
		delete(fields, "body")

	}
	// Field "queryParams"
	if fields["queryParams"] != nil {
		if string(fields["queryParams"]) != "null" {

			if err := json.Unmarshal(fields["queryParams"], &resource.QueryParams); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryParams", err)...)
			}

		}
		delete(fields, "queryParams")

	}
	// Field "headers"
	if fields["headers"] != nil {
		if string(fields["headers"]) != "null" {

			if err := json.Unmarshal(fields["headers"], &resource.Headers); err != nil {
				errs = append(errs, cog.MakeBuildErrors("headers", err)...)
			}

		}
		delete(fields, "headers")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FetchOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FetchOptions` objects.
func (resource FetchOptions) Equals(other FetchOptions) bool {
	if resource.Method != other.Method {
		return false
	}
	if resource.Url != other.Url {
		return false
	}
	if resource.Body == nil && other.Body != nil || resource.Body != nil && other.Body == nil {
		return false
	}

	if resource.Body != nil {
		if *resource.Body != *other.Body {
			return false
		}
	}

	if len(resource.QueryParams) != len(other.QueryParams) {
		return false
	}

	for i1 := range resource.QueryParams {

		if len(resource.QueryParams[i1]) != len(other.QueryParams[i1]) {
			return false
		}

		for i2 := range resource.QueryParams[i1] {
			if resource.QueryParams[i1][i2] != other.QueryParams[i1][i2] {
				return false
			}
		}
	}

	if len(resource.Headers) != len(other.Headers) {
		return false
	}

	for i1 := range resource.Headers {

		if len(resource.Headers[i1]) != len(other.Headers[i1]) {
			return false
		}

		for i2 := range resource.Headers[i1] {
			if resource.Headers[i1][i2] != other.Headers[i1][i2] {
				return false
			}
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FetchOptions` fields for violations and returns them.
func (resource FetchOptions) Validate() error {
	return nil
}

type HttpRequestMethod string

const (
	HttpRequestMethodGET    HttpRequestMethod = "GET"
	HttpRequestMethodPUT    HttpRequestMethod = "PUT"
	HttpRequestMethodPOST   HttpRequestMethod = "POST"
	HttpRequestMethodDELETE HttpRequestMethod = "DELETE"
	HttpRequestMethodPATCH  HttpRequestMethod = "PATCH"
)

type InfinityOptions struct {
	Method HttpRequestMethod `json:"method"`
	Url    string            `json:"url"`
	Body   *string           `json:"body,omitempty"`
	// These are 2D arrays of strings, each representing a key-value pair
	// We are defining them this way because we can't generate a go struct that
	// that would have exactly two strings in each sub-array
	QueryParams   [][]string `json:"queryParams,omitempty"`
	DatasourceUid string     `json:"datasourceUid"`
	Headers       [][]string `json:"headers,omitempty"`
}

// NewInfinityOptions creates a new InfinityOptions object.
func NewInfinityOptions() *InfinityOptions {
	return &InfinityOptions{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `InfinityOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *InfinityOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "method"
	if fields["method"] != nil {
		if string(fields["method"]) != "null" {
			if err := json.Unmarshal(fields["method"], &resource.Method); err != nil {
				errs = append(errs, cog.MakeBuildErrors("method", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("method", errors.New("required field is null"))...)

		}
		delete(fields, "method")
	} else {
		errs = append(errs, cog.MakeBuildErrors("method", errors.New("required field is missing from input"))...)
	}
	// Field "url"
	if fields["url"] != nil {
		if string(fields["url"]) != "null" {
			if err := json.Unmarshal(fields["url"], &resource.Url); err != nil {
				errs = append(errs, cog.MakeBuildErrors("url", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("url", errors.New("required field is null"))...)

		}
		delete(fields, "url")
	} else {
		errs = append(errs, cog.MakeBuildErrors("url", errors.New("required field is missing from input"))...)
	}
	// Field "body"
	if fields["body"] != nil {
		if string(fields["body"]) != "null" {
			if err := json.Unmarshal(fields["body"], &resource.Body); err != nil {
				errs = append(errs, cog.MakeBuildErrors("body", err)...)
			}

		}
		delete(fields, "body")

	}
	// Field "queryParams"
	if fields["queryParams"] != nil {
		if string(fields["queryParams"]) != "null" {

			if err := json.Unmarshal(fields["queryParams"], &resource.QueryParams); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryParams", err)...)
			}

		}
		delete(fields, "queryParams")

	}
	// Field "datasourceUid"
	if fields["datasourceUid"] != nil {
		if string(fields["datasourceUid"]) != "null" {
			if err := json.Unmarshal(fields["datasourceUid"], &resource.DatasourceUid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasourceUid", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("datasourceUid", errors.New("required field is null"))...)

		}
		delete(fields, "datasourceUid")
	} else {
		errs = append(errs, cog.MakeBuildErrors("datasourceUid", errors.New("required field is missing from input"))...)
	}
	// Field "headers"
	if fields["headers"] != nil {
		if string(fields["headers"]) != "null" {

			if err := json.Unmarshal(fields["headers"], &resource.Headers); err != nil {
				errs = append(errs, cog.MakeBuildErrors("headers", err)...)
			}

		}
		delete(fields, "headers")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("InfinityOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `InfinityOptions` objects.
func (resource InfinityOptions) Equals(other InfinityOptions) bool {
	if resource.Method != other.Method {
		return false
	}
	if resource.Url != other.Url {
		return false
	}
	if resource.Body == nil && other.Body != nil || resource.Body != nil && other.Body == nil {
		return false
	}

	if resource.Body != nil {
		if *resource.Body != *other.Body {
			return false
		}
	}

	if len(resource.QueryParams) != len(other.QueryParams) {
		return false
	}

	for i1 := range resource.QueryParams {

		if len(resource.QueryParams[i1]) != len(other.QueryParams[i1]) {
			return false
		}

		for i2 := range resource.QueryParams[i1] {
			if resource.QueryParams[i1][i2] != other.QueryParams[i1][i2] {
				return false
			}
		}
	}
	if resource.DatasourceUid != other.DatasourceUid {
		return false
	}

	if len(resource.Headers) != len(other.Headers) {
		return false
	}

	for i1 := range resource.Headers {

		if len(resource.Headers[i1]) != len(other.Headers[i1]) {
			return false
		}

		for i2 := range resource.Headers[i1] {
			if resource.Headers[i1][i2] != other.Headers[i1][i2] {
				return false
			}
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `InfinityOptions` fields for violations and returns them.
func (resource InfinityOptions) Validate() error {
	return nil
}

type ActionVariable struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewActionVariable creates a new ActionVariable object.
func NewActionVariable() *ActionVariable {
	return &ActionVariable{
		Type: ActionVariableType,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ActionVariable` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ActionVariable) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "key"
	if fields["key"] != nil {
		if string(fields["key"]) != "null" {
			if err := json.Unmarshal(fields["key"], &resource.Key); err != nil {
				errs = append(errs, cog.MakeBuildErrors("key", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is null"))...)

		}
		delete(fields, "key")
	} else {
		errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is missing from input"))...)
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ActionVariable", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ActionVariable` objects.
func (resource ActionVariable) Equals(other ActionVariable) bool {
	if resource.Key != other.Key {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ActionVariable` fields for violations and returns them.
func (resource ActionVariable) Validate() error {
	return nil
}

// Action variable type
const ActionVariableType = "string"

// How null values should be handled
type NullValueMode string

const (
	NullValueModeNull       NullValueMode = "null"
	NullValueModeConnected  NullValueMode = "connected"
	NullValueModeNullAsZero NullValueMode = "null as zero"
)

type DynamicConfigValue struct {
	Id    string `json:"id"`
	Value any    `json:"value,omitempty"`
}

// NewDynamicConfigValue creates a new DynamicConfigValue object.
func NewDynamicConfigValue() *DynamicConfigValue {
	return &DynamicConfigValue{
		Id: "",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DynamicConfigValue` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DynamicConfigValue) UnmarshalJSONStrict(raw []byte) error {
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

	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}

		}
		delete(fields, "value")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DynamicConfigValue", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DynamicConfigValue` objects.
func (resource DynamicConfigValue) Equals(other DynamicConfigValue) bool {
	if resource.Id != other.Id {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Value, other.Value) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DynamicConfigValue` fields for violations and returns them.
func (resource DynamicConfigValue) Validate() error {
	return nil
}

type LibraryPanelKind struct {
	Kind string               `json:"kind"`
	Spec LibraryPanelKindSpec `json:"spec"`
}

// NewLibraryPanelKind creates a new LibraryPanelKind object.
func NewLibraryPanelKind() *LibraryPanelKind {
	return &LibraryPanelKind{
		Kind: "LibraryPanel",
		Spec: *NewLibraryPanelKindSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanelKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LibraryPanelKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = LibraryPanelKindSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LibraryPanelKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LibraryPanelKind` objects.
func (resource LibraryPanelKind) Equals(other LibraryPanelKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `LibraryPanelKind` fields for violations and returns them.
func (resource LibraryPanelKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type LibraryPanelKindSpec struct {
	// Panel ID for the library panel in the dashboard
	Id float64 `json:"id"`
	// Title for the library panel in the dashboard
	Title        string          `json:"title"`
	LibraryPanel LibraryPanelRef `json:"libraryPanel"`
}

// NewLibraryPanelKindSpec creates a new LibraryPanelKindSpec object.
func NewLibraryPanelKindSpec() *LibraryPanelKindSpec {
	return &LibraryPanelKindSpec{
		LibraryPanel: *NewLibraryPanelRef(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanelKindSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LibraryPanelKindSpec) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}
	// Field "libraryPanel"
	if fields["libraryPanel"] != nil {
		if string(fields["libraryPanel"]) != "null" {

			resource.LibraryPanel = LibraryPanelRef{}
			if err := resource.LibraryPanel.UnmarshalJSONStrict(fields["libraryPanel"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("libraryPanel", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("libraryPanel", errors.New("required field is null"))...)

		}
		delete(fields, "libraryPanel")
	} else {
		errs = append(errs, cog.MakeBuildErrors("libraryPanel", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LibraryPanelKindSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LibraryPanelKindSpec` objects.
func (resource LibraryPanelKindSpec) Equals(other LibraryPanelKindSpec) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Title != other.Title {
		return false
	}
	if !resource.LibraryPanel.Equals(other.LibraryPanel) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `LibraryPanelKindSpec` fields for violations and returns them.
func (resource LibraryPanelKindSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.LibraryPanel.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("libraryPanel", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
type LibraryPanelRef struct {
	// Library panel name
	Name string `json:"name"`
	// Library panel uid
	Uid string `json:"uid"`
}

// NewLibraryPanelRef creates a new LibraryPanelRef object.
func NewLibraryPanelRef() *LibraryPanelRef {
	return &LibraryPanelRef{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanelRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LibraryPanelRef) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LibraryPanelRef", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LibraryPanelRef` objects.
func (resource LibraryPanelRef) Equals(other LibraryPanelRef) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Uid != other.Uid {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `LibraryPanelRef` fields for violations and returns them.
func (resource LibraryPanelRef) Validate() error {
	return nil
}

type GridLayoutKind struct {
	Kind string         `json:"kind"`
	Spec GridLayoutSpec `json:"spec"`
}

// NewGridLayoutKind creates a new GridLayoutKind object.
func NewGridLayoutKind() *GridLayoutKind {
	return &GridLayoutKind{
		Kind: "GridLayout",
		Spec: *NewGridLayoutSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GridLayoutKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = GridLayoutSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GridLayoutKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GridLayoutKind` objects.
func (resource GridLayoutKind) Equals(other GridLayoutKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GridLayoutKind` fields for violations and returns them.
func (resource GridLayoutKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type GridLayoutSpec struct {
	Items []GridLayoutItemKind `json:"items"`
}

// NewGridLayoutSpec creates a new GridLayoutSpec object.
func NewGridLayoutSpec() *GridLayoutSpec {
	return &GridLayoutSpec{
		Items: []GridLayoutItemKind{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GridLayoutSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "items"
	if fields["items"] != nil {
		if string(fields["items"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["items"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 GridLayoutItemKind

				result1 = GridLayoutItemKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Items = append(resource.Items, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is null"))...)

		}
		delete(fields, "items")
	} else {
		errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GridLayoutSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GridLayoutSpec` objects.
func (resource GridLayoutSpec) Equals(other GridLayoutSpec) bool {

	if len(resource.Items) != len(other.Items) {
		return false
	}

	for i1 := range resource.Items {
		if !resource.Items[i1].Equals(other.Items[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GridLayoutSpec` fields for violations and returns them.
func (resource GridLayoutSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Items {
		if err := resource.Items[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type GridLayoutItemKind struct {
	Kind string             `json:"kind"`
	Spec GridLayoutItemSpec `json:"spec"`
}

// NewGridLayoutItemKind creates a new GridLayoutItemKind object.
func NewGridLayoutItemKind() *GridLayoutItemKind {
	return &GridLayoutItemKind{
		Kind: "GridLayoutItem",
		Spec: *NewGridLayoutItemSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutItemKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GridLayoutItemKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = GridLayoutItemSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GridLayoutItemKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GridLayoutItemKind` objects.
func (resource GridLayoutItemKind) Equals(other GridLayoutItemKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GridLayoutItemKind` fields for violations and returns them.
func (resource GridLayoutItemKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type GridLayoutItemSpec struct {
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	// reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
	Element ElementReference `json:"element"`
	Repeat  *RepeatOptions   `json:"repeat,omitempty"`
}

// NewGridLayoutItemSpec creates a new GridLayoutItemSpec object.
func NewGridLayoutItemSpec() *GridLayoutItemSpec {
	return &GridLayoutItemSpec{
		Element: *NewElementReference(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutItemSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GridLayoutItemSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "x"
	if fields["x"] != nil {
		if string(fields["x"]) != "null" {
			if err := json.Unmarshal(fields["x"], &resource.X); err != nil {
				errs = append(errs, cog.MakeBuildErrors("x", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("x", errors.New("required field is null"))...)

		}
		delete(fields, "x")
	} else {
		errs = append(errs, cog.MakeBuildErrors("x", errors.New("required field is missing from input"))...)
	}
	// Field "y"
	if fields["y"] != nil {
		if string(fields["y"]) != "null" {
			if err := json.Unmarshal(fields["y"], &resource.Y); err != nil {
				errs = append(errs, cog.MakeBuildErrors("y", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("y", errors.New("required field is null"))...)

		}
		delete(fields, "y")
	} else {
		errs = append(errs, cog.MakeBuildErrors("y", errors.New("required field is missing from input"))...)
	}
	// Field "width"
	if fields["width"] != nil {
		if string(fields["width"]) != "null" {
			if err := json.Unmarshal(fields["width"], &resource.Width); err != nil {
				errs = append(errs, cog.MakeBuildErrors("width", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("width", errors.New("required field is null"))...)

		}
		delete(fields, "width")
	} else {
		errs = append(errs, cog.MakeBuildErrors("width", errors.New("required field is missing from input"))...)
	}
	// Field "height"
	if fields["height"] != nil {
		if string(fields["height"]) != "null" {
			if err := json.Unmarshal(fields["height"], &resource.Height); err != nil {
				errs = append(errs, cog.MakeBuildErrors("height", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("height", errors.New("required field is null"))...)

		}
		delete(fields, "height")
	} else {
		errs = append(errs, cog.MakeBuildErrors("height", errors.New("required field is missing from input"))...)
	}
	// Field "element"
	if fields["element"] != nil {
		if string(fields["element"]) != "null" {

			resource.Element = ElementReference{}
			if err := resource.Element.UnmarshalJSONStrict(fields["element"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("element", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("element", errors.New("required field is null"))...)

		}
		delete(fields, "element")
	} else {
		errs = append(errs, cog.MakeBuildErrors("element", errors.New("required field is missing from input"))...)
	}
	// Field "repeat"
	if fields["repeat"] != nil {
		if string(fields["repeat"]) != "null" {

			resource.Repeat = &RepeatOptions{}
			if err := resource.Repeat.UnmarshalJSONStrict(fields["repeat"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
			}

		}
		delete(fields, "repeat")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GridLayoutItemSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GridLayoutItemSpec` objects.
func (resource GridLayoutItemSpec) Equals(other GridLayoutItemSpec) bool {
	if resource.X != other.X {
		return false
	}
	if resource.Y != other.Y {
		return false
	}
	if resource.Width != other.Width {
		return false
	}
	if resource.Height != other.Height {
		return false
	}
	if !resource.Element.Equals(other.Element) {
		return false
	}
	if resource.Repeat == nil && other.Repeat != nil || resource.Repeat != nil && other.Repeat == nil {
		return false
	}

	if resource.Repeat != nil {
		if !resource.Repeat.Equals(*other.Repeat) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GridLayoutItemSpec` fields for violations and returns them.
func (resource GridLayoutItemSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Element.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("element", err)...)
	}
	if resource.Repeat != nil {
		if err := resource.Repeat.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElementReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

// NewElementReference creates a new ElementReference object.
func NewElementReference() *ElementReference {
	return &ElementReference{
		Kind: "ElementReference",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElementReference` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElementReference) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElementReference", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElementReference` objects.
func (resource ElementReference) Equals(other ElementReference) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElementReference` fields for violations and returns them.
func (resource ElementReference) Validate() error {
	return nil
}

type RepeatOptions struct {
	Mode      RepeatMode              `json:"mode"`
	Value     string                  `json:"value"`
	Direction *RepeatOptionsDirection `json:"direction,omitempty"`
	MaxPerRow *int64                  `json:"maxPerRow,omitempty"`
}

// NewRepeatOptions creates a new RepeatOptions object.
func NewRepeatOptions() *RepeatOptions {
	return &RepeatOptions{
		Mode: RepeatModeVariable,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RepeatOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RepeatOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}
	// Field "direction"
	if fields["direction"] != nil {
		if string(fields["direction"]) != "null" {
			if err := json.Unmarshal(fields["direction"], &resource.Direction); err != nil {
				errs = append(errs, cog.MakeBuildErrors("direction", err)...)
			}

		}
		delete(fields, "direction")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RepeatOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RepeatOptions` objects.
func (resource RepeatOptions) Equals(other RepeatOptions) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Value != other.Value {
		return false
	}
	if resource.Direction == nil && other.Direction != nil || resource.Direction != nil && other.Direction == nil {
		return false
	}

	if resource.Direction != nil {
		if *resource.Direction != *other.Direction {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `RepeatOptions` fields for violations and returns them.
func (resource RepeatOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.Mode != "variable" {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("must be variable"))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// other repeat modes will be added in the future: label, frame
type RepeatMode string

const (
	RepeatModeVariable RepeatMode = "variable"
)

type RowsLayoutKind struct {
	Kind string         `json:"kind"`
	Spec RowsLayoutSpec `json:"spec"`
}

// NewRowsLayoutKind creates a new RowsLayoutKind object.
func NewRowsLayoutKind() *RowsLayoutKind {
	return &RowsLayoutKind{
		Kind: "RowsLayout",
		Spec: *NewRowsLayoutSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsLayoutKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RowsLayoutKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = RowsLayoutSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RowsLayoutKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RowsLayoutKind` objects.
func (resource RowsLayoutKind) Equals(other RowsLayoutKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RowsLayoutKind` fields for violations and returns them.
func (resource RowsLayoutKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RowsLayoutSpec struct {
	Rows []RowsLayoutRowKind `json:"rows"`
}

// NewRowsLayoutSpec creates a new RowsLayoutSpec object.
func NewRowsLayoutSpec() *RowsLayoutSpec {
	return &RowsLayoutSpec{
		Rows: []RowsLayoutRowKind{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsLayoutSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RowsLayoutSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rows"
	if fields["rows"] != nil {
		if string(fields["rows"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["rows"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 RowsLayoutRowKind

				result1 = RowsLayoutRowKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("rows["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Rows = append(resource.Rows, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("rows", errors.New("required field is null"))...)

		}
		delete(fields, "rows")
	} else {
		errs = append(errs, cog.MakeBuildErrors("rows", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RowsLayoutSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RowsLayoutSpec` objects.
func (resource RowsLayoutSpec) Equals(other RowsLayoutSpec) bool {

	if len(resource.Rows) != len(other.Rows) {
		return false
	}

	for i1 := range resource.Rows {
		if !resource.Rows[i1].Equals(other.Rows[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RowsLayoutSpec` fields for violations and returns them.
func (resource RowsLayoutSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Rows {
		if err := resource.Rows[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("rows["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RowsLayoutRowKind struct {
	Kind string            `json:"kind"`
	Spec RowsLayoutRowSpec `json:"spec"`
}

// NewRowsLayoutRowKind creates a new RowsLayoutRowKind object.
func NewRowsLayoutRowKind() *RowsLayoutRowKind {
	return &RowsLayoutRowKind{
		Kind: "RowsLayoutRow",
		Spec: *NewRowsLayoutRowSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsLayoutRowKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RowsLayoutRowKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = RowsLayoutRowSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RowsLayoutRowKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RowsLayoutRowKind` objects.
func (resource RowsLayoutRowKind) Equals(other RowsLayoutRowKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RowsLayoutRowKind` fields for violations and returns them.
func (resource RowsLayoutRowKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RowsLayoutRowSpec struct {
	Title                *string                                                            `json:"title,omitempty"`
	Collapse             *bool                                                              `json:"collapse,omitempty"`
	HideHeader           *bool                                                              `json:"hideHeader,omitempty"`
	FillScreen           *bool                                                              `json:"fillScreen,omitempty"`
	ConditionalRendering *ConditionalRenderingGroupKind                                     `json:"conditionalRendering,omitempty"`
	Repeat               *RowRepeatOptions                                                  `json:"repeat,omitempty"`
	Layout               GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind `json:"layout"`
}

// NewRowsLayoutRowSpec creates a new RowsLayoutRowSpec object.
func NewRowsLayoutRowSpec() *RowsLayoutRowSpec {
	return &RowsLayoutRowSpec{
		Layout: *NewGridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsLayoutRowSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RowsLayoutRowSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "collapse"
	if fields["collapse"] != nil {
		if string(fields["collapse"]) != "null" {
			if err := json.Unmarshal(fields["collapse"], &resource.Collapse); err != nil {
				errs = append(errs, cog.MakeBuildErrors("collapse", err)...)
			}

		}
		delete(fields, "collapse")

	}
	// Field "hideHeader"
	if fields["hideHeader"] != nil {
		if string(fields["hideHeader"]) != "null" {
			if err := json.Unmarshal(fields["hideHeader"], &resource.HideHeader); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideHeader", err)...)
			}

		}
		delete(fields, "hideHeader")

	}
	// Field "fillScreen"
	if fields["fillScreen"] != nil {
		if string(fields["fillScreen"]) != "null" {
			if err := json.Unmarshal(fields["fillScreen"], &resource.FillScreen); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillScreen", err)...)
			}

		}
		delete(fields, "fillScreen")

	}
	// Field "conditionalRendering"
	if fields["conditionalRendering"] != nil {
		if string(fields["conditionalRendering"]) != "null" {

			resource.ConditionalRendering = &ConditionalRenderingGroupKind{}
			if err := resource.ConditionalRendering.UnmarshalJSONStrict(fields["conditionalRendering"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("conditionalRendering", err)...)
			}

		}
		delete(fields, "conditionalRendering")

	}
	// Field "repeat"
	if fields["repeat"] != nil {
		if string(fields["repeat"]) != "null" {

			resource.Repeat = &RowRepeatOptions{}
			if err := resource.Repeat.UnmarshalJSONStrict(fields["repeat"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
			}

		}
		delete(fields, "repeat")

	}
	// Field "layout"
	if fields["layout"] != nil {
		if string(fields["layout"]) != "null" {

			resource.Layout = GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind{}
			if err := resource.Layout.UnmarshalJSONStrict(fields["layout"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("layout", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("layout", errors.New("required field is null"))...)

		}
		delete(fields, "layout")
	} else {
		errs = append(errs, cog.MakeBuildErrors("layout", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RowsLayoutRowSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RowsLayoutRowSpec` objects.
func (resource RowsLayoutRowSpec) Equals(other RowsLayoutRowSpec) bool {
	if resource.Title == nil && other.Title != nil || resource.Title != nil && other.Title == nil {
		return false
	}

	if resource.Title != nil {
		if *resource.Title != *other.Title {
			return false
		}
	}
	if resource.Collapse == nil && other.Collapse != nil || resource.Collapse != nil && other.Collapse == nil {
		return false
	}

	if resource.Collapse != nil {
		if *resource.Collapse != *other.Collapse {
			return false
		}
	}
	if resource.HideHeader == nil && other.HideHeader != nil || resource.HideHeader != nil && other.HideHeader == nil {
		return false
	}

	if resource.HideHeader != nil {
		if *resource.HideHeader != *other.HideHeader {
			return false
		}
	}
	if resource.FillScreen == nil && other.FillScreen != nil || resource.FillScreen != nil && other.FillScreen == nil {
		return false
	}

	if resource.FillScreen != nil {
		if *resource.FillScreen != *other.FillScreen {
			return false
		}
	}
	if resource.ConditionalRendering == nil && other.ConditionalRendering != nil || resource.ConditionalRendering != nil && other.ConditionalRendering == nil {
		return false
	}

	if resource.ConditionalRendering != nil {
		if !resource.ConditionalRendering.Equals(*other.ConditionalRendering) {
			return false
		}
	}
	if resource.Repeat == nil && other.Repeat != nil || resource.Repeat != nil && other.Repeat == nil {
		return false
	}

	if resource.Repeat != nil {
		if !resource.Repeat.Equals(*other.Repeat) {
			return false
		}
	}
	if !resource.Layout.Equals(other.Layout) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RowsLayoutRowSpec` fields for violations and returns them.
func (resource RowsLayoutRowSpec) Validate() error {
	var errs cog.BuildErrors
	if resource.ConditionalRendering != nil {
		if err := resource.ConditionalRendering.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("conditionalRendering", err)...)
		}
	}
	if resource.Repeat != nil {
		if err := resource.Repeat.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
		}
	}
	if err := resource.Layout.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("layout", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ConditionalRenderingGroupKind struct {
	Kind string                        `json:"kind"`
	Spec ConditionalRenderingGroupSpec `json:"spec"`
}

// NewConditionalRenderingGroupKind creates a new ConditionalRenderingGroupKind object.
func NewConditionalRenderingGroupKind() *ConditionalRenderingGroupKind {
	return &ConditionalRenderingGroupKind{
		Kind: "ConditionalRenderingGroup",
		Spec: *NewConditionalRenderingGroupSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingGroupKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingGroupKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = ConditionalRenderingGroupSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingGroupKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingGroupKind` objects.
func (resource ConditionalRenderingGroupKind) Equals(other ConditionalRenderingGroupKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingGroupKind` fields for violations and returns them.
func (resource ConditionalRenderingGroupKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ConditionalRenderingGroupSpec struct {
	Visibility ConditionalRenderingGroupSpecVisibility                                                                 `json:"visibility"`
	Condition  ConditionalRenderingGroupSpecCondition                                                                  `json:"condition"`
	Items      []ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind `json:"items"`
}

// NewConditionalRenderingGroupSpec creates a new ConditionalRenderingGroupSpec object.
func NewConditionalRenderingGroupSpec() *ConditionalRenderingGroupSpec {
	return &ConditionalRenderingGroupSpec{
		Items: []ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingGroupSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingGroupSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "visibility"
	if fields["visibility"] != nil {
		if string(fields["visibility"]) != "null" {
			if err := json.Unmarshal(fields["visibility"], &resource.Visibility); err != nil {
				errs = append(errs, cog.MakeBuildErrors("visibility", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("visibility", errors.New("required field is null"))...)

		}
		delete(fields, "visibility")
	} else {
		errs = append(errs, cog.MakeBuildErrors("visibility", errors.New("required field is missing from input"))...)
	}
	// Field "condition"
	if fields["condition"] != nil {
		if string(fields["condition"]) != "null" {
			if err := json.Unmarshal(fields["condition"], &resource.Condition); err != nil {
				errs = append(errs, cog.MakeBuildErrors("condition", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("condition", errors.New("required field is null"))...)

		}
		delete(fields, "condition")
	} else {
		errs = append(errs, cog.MakeBuildErrors("condition", errors.New("required field is missing from input"))...)
	}
	// Field "items"
	if fields["items"] != nil {
		if string(fields["items"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["items"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind

				result1 = ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Items = append(resource.Items, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is null"))...)

		}
		delete(fields, "items")
	} else {
		errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingGroupSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingGroupSpec` objects.
func (resource ConditionalRenderingGroupSpec) Equals(other ConditionalRenderingGroupSpec) bool {
	if resource.Visibility != other.Visibility {
		return false
	}
	if resource.Condition != other.Condition {
		return false
	}

	if len(resource.Items) != len(other.Items) {
		return false
	}

	for i1 := range resource.Items {
		if !resource.Items[i1].Equals(other.Items[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingGroupSpec` fields for violations and returns them.
func (resource ConditionalRenderingGroupSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Items {
		if err := resource.Items[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ConditionalRenderingVariableKind struct {
	Kind string                           `json:"kind"`
	Spec ConditionalRenderingVariableSpec `json:"spec"`
}

// NewConditionalRenderingVariableKind creates a new ConditionalRenderingVariableKind object.
func NewConditionalRenderingVariableKind() *ConditionalRenderingVariableKind {
	return &ConditionalRenderingVariableKind{
		Kind: "ConditionalRenderingVariable",
		Spec: *NewConditionalRenderingVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = ConditionalRenderingVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingVariableKind` objects.
func (resource ConditionalRenderingVariableKind) Equals(other ConditionalRenderingVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingVariableKind` fields for violations and returns them.
func (resource ConditionalRenderingVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ConditionalRenderingVariableSpec struct {
	Variable string                                   `json:"variable"`
	Operator ConditionalRenderingVariableSpecOperator `json:"operator"`
	Value    string                                   `json:"value"`
}

// NewConditionalRenderingVariableSpec creates a new ConditionalRenderingVariableSpec object.
func NewConditionalRenderingVariableSpec() *ConditionalRenderingVariableSpec {
	return &ConditionalRenderingVariableSpec{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "variable"
	if fields["variable"] != nil {
		if string(fields["variable"]) != "null" {
			if err := json.Unmarshal(fields["variable"], &resource.Variable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("variable", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("variable", errors.New("required field is null"))...)

		}
		delete(fields, "variable")
	} else {
		errs = append(errs, cog.MakeBuildErrors("variable", errors.New("required field is missing from input"))...)
	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {
			if err := json.Unmarshal(fields["operator"], &resource.Operator); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is null"))...)

		}
		delete(fields, "operator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingVariableSpec` objects.
func (resource ConditionalRenderingVariableSpec) Equals(other ConditionalRenderingVariableSpec) bool {
	if resource.Variable != other.Variable {
		return false
	}
	if resource.Operator != other.Operator {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingVariableSpec` fields for violations and returns them.
func (resource ConditionalRenderingVariableSpec) Validate() error {
	return nil
}

type ConditionalRenderingDataKind struct {
	Kind string                       `json:"kind"`
	Spec ConditionalRenderingDataSpec `json:"spec"`
}

// NewConditionalRenderingDataKind creates a new ConditionalRenderingDataKind object.
func NewConditionalRenderingDataKind() *ConditionalRenderingDataKind {
	return &ConditionalRenderingDataKind{
		Kind: "ConditionalRenderingData",
		Spec: *NewConditionalRenderingDataSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingDataKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingDataKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = ConditionalRenderingDataSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingDataKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingDataKind` objects.
func (resource ConditionalRenderingDataKind) Equals(other ConditionalRenderingDataKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingDataKind` fields for violations and returns them.
func (resource ConditionalRenderingDataKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ConditionalRenderingDataSpec struct {
	Value bool `json:"value"`
}

// NewConditionalRenderingDataSpec creates a new ConditionalRenderingDataSpec object.
func NewConditionalRenderingDataSpec() *ConditionalRenderingDataSpec {
	return &ConditionalRenderingDataSpec{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingDataSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingDataSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingDataSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingDataSpec` objects.
func (resource ConditionalRenderingDataSpec) Equals(other ConditionalRenderingDataSpec) bool {
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingDataSpec` fields for violations and returns them.
func (resource ConditionalRenderingDataSpec) Validate() error {
	return nil
}

type ConditionalRenderingTimeRangeSizeKind struct {
	Kind string                                `json:"kind"`
	Spec ConditionalRenderingTimeRangeSizeSpec `json:"spec"`
}

// NewConditionalRenderingTimeRangeSizeKind creates a new ConditionalRenderingTimeRangeSizeKind object.
func NewConditionalRenderingTimeRangeSizeKind() *ConditionalRenderingTimeRangeSizeKind {
	return &ConditionalRenderingTimeRangeSizeKind{
		Kind: "ConditionalRenderingTimeRangeSize",
		Spec: *NewConditionalRenderingTimeRangeSizeSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingTimeRangeSizeKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingTimeRangeSizeKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = ConditionalRenderingTimeRangeSizeSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingTimeRangeSizeKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingTimeRangeSizeKind` objects.
func (resource ConditionalRenderingTimeRangeSizeKind) Equals(other ConditionalRenderingTimeRangeSizeKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingTimeRangeSizeKind` fields for violations and returns them.
func (resource ConditionalRenderingTimeRangeSizeKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ConditionalRenderingTimeRangeSizeSpec struct {
	Value string `json:"value"`
}

// NewConditionalRenderingTimeRangeSizeSpec creates a new ConditionalRenderingTimeRangeSizeSpec object.
func NewConditionalRenderingTimeRangeSizeSpec() *ConditionalRenderingTimeRangeSizeSpec {
	return &ConditionalRenderingTimeRangeSizeSpec{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingTimeRangeSizeSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingTimeRangeSizeSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingTimeRangeSizeSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConditionalRenderingTimeRangeSizeSpec` objects.
func (resource ConditionalRenderingTimeRangeSizeSpec) Equals(other ConditionalRenderingTimeRangeSizeSpec) bool {
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingTimeRangeSizeSpec` fields for violations and returns them.
func (resource ConditionalRenderingTimeRangeSizeSpec) Validate() error {
	return nil
}

type RowRepeatOptions struct {
	Mode  RepeatMode `json:"mode"`
	Value string     `json:"value"`
}

// NewRowRepeatOptions creates a new RowRepeatOptions object.
func NewRowRepeatOptions() *RowRepeatOptions {
	return &RowRepeatOptions{
		Mode: RepeatModeVariable,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowRepeatOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RowRepeatOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RowRepeatOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RowRepeatOptions` objects.
func (resource RowRepeatOptions) Equals(other RowRepeatOptions) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RowRepeatOptions` fields for violations and returns them.
func (resource RowRepeatOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.Mode != "variable" {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("must be variable"))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AutoGridLayoutKind struct {
	Kind string             `json:"kind"`
	Spec AutoGridLayoutSpec `json:"spec"`
}

// NewAutoGridLayoutKind creates a new AutoGridLayoutKind object.
func NewAutoGridLayoutKind() *AutoGridLayoutKind {
	return &AutoGridLayoutKind{
		Kind: "AutoGridLayout",
		Spec: *NewAutoGridLayoutSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AutoGridLayoutKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = AutoGridLayoutSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AutoGridLayoutKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AutoGridLayoutKind` objects.
func (resource AutoGridLayoutKind) Equals(other AutoGridLayoutKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AutoGridLayoutKind` fields for violations and returns them.
func (resource AutoGridLayoutKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AutoGridLayoutSpec struct {
	MaxColumnCount  *float64                          `json:"maxColumnCount,omitempty"`
	ColumnWidthMode AutoGridLayoutSpecColumnWidthMode `json:"columnWidthMode"`
	ColumnWidth     *float64                          `json:"columnWidth,omitempty"`
	RowHeightMode   AutoGridLayoutSpecRowHeightMode   `json:"rowHeightMode"`
	RowHeight       *float64                          `json:"rowHeight,omitempty"`
	FillScreen      *bool                             `json:"fillScreen,omitempty"`
	Items           []AutoGridLayoutItemKind          `json:"items"`
}

// NewAutoGridLayoutSpec creates a new AutoGridLayoutSpec object.
func NewAutoGridLayoutSpec() *AutoGridLayoutSpec {
	return &AutoGridLayoutSpec{
		MaxColumnCount:  (func(input float64) *float64 { return &input })(3),
		ColumnWidthMode: AutoGridLayoutSpecColumnWidthModeStandard,
		RowHeightMode:   AutoGridLayoutSpecRowHeightModeStandard,
		FillScreen:      (func(input bool) *bool { return &input })(false),
		Items:           []AutoGridLayoutItemKind{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AutoGridLayoutSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxColumnCount"
	if fields["maxColumnCount"] != nil {
		if string(fields["maxColumnCount"]) != "null" {
			if err := json.Unmarshal(fields["maxColumnCount"], &resource.MaxColumnCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxColumnCount", err)...)
			}

		}
		delete(fields, "maxColumnCount")

	}
	// Field "columnWidthMode"
	if fields["columnWidthMode"] != nil {
		if string(fields["columnWidthMode"]) != "null" {
			if err := json.Unmarshal(fields["columnWidthMode"], &resource.ColumnWidthMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("columnWidthMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("columnWidthMode", errors.New("required field is null"))...)

		}
		delete(fields, "columnWidthMode")

	}
	// Field "columnWidth"
	if fields["columnWidth"] != nil {
		if string(fields["columnWidth"]) != "null" {
			if err := json.Unmarshal(fields["columnWidth"], &resource.ColumnWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("columnWidth", err)...)
			}

		}
		delete(fields, "columnWidth")

	}
	// Field "rowHeightMode"
	if fields["rowHeightMode"] != nil {
		if string(fields["rowHeightMode"]) != "null" {
			if err := json.Unmarshal(fields["rowHeightMode"], &resource.RowHeightMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rowHeightMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("rowHeightMode", errors.New("required field is null"))...)

		}
		delete(fields, "rowHeightMode")

	}
	// Field "rowHeight"
	if fields["rowHeight"] != nil {
		if string(fields["rowHeight"]) != "null" {
			if err := json.Unmarshal(fields["rowHeight"], &resource.RowHeight); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rowHeight", err)...)
			}

		}
		delete(fields, "rowHeight")

	}
	// Field "fillScreen"
	if fields["fillScreen"] != nil {
		if string(fields["fillScreen"]) != "null" {
			if err := json.Unmarshal(fields["fillScreen"], &resource.FillScreen); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillScreen", err)...)
			}

		}
		delete(fields, "fillScreen")

	}
	// Field "items"
	if fields["items"] != nil {
		if string(fields["items"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["items"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AutoGridLayoutItemKind

				result1 = AutoGridLayoutItemKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Items = append(resource.Items, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is null"))...)

		}
		delete(fields, "items")
	} else {
		errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AutoGridLayoutSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AutoGridLayoutSpec` objects.
func (resource AutoGridLayoutSpec) Equals(other AutoGridLayoutSpec) bool {
	if resource.MaxColumnCount == nil && other.MaxColumnCount != nil || resource.MaxColumnCount != nil && other.MaxColumnCount == nil {
		return false
	}

	if resource.MaxColumnCount != nil {
		if *resource.MaxColumnCount != *other.MaxColumnCount {
			return false
		}
	}
	if resource.ColumnWidthMode != other.ColumnWidthMode {
		return false
	}
	if resource.ColumnWidth == nil && other.ColumnWidth != nil || resource.ColumnWidth != nil && other.ColumnWidth == nil {
		return false
	}

	if resource.ColumnWidth != nil {
		if *resource.ColumnWidth != *other.ColumnWidth {
			return false
		}
	}
	if resource.RowHeightMode != other.RowHeightMode {
		return false
	}
	if resource.RowHeight == nil && other.RowHeight != nil || resource.RowHeight != nil && other.RowHeight == nil {
		return false
	}

	if resource.RowHeight != nil {
		if *resource.RowHeight != *other.RowHeight {
			return false
		}
	}
	if resource.FillScreen == nil && other.FillScreen != nil || resource.FillScreen != nil && other.FillScreen == nil {
		return false
	}

	if resource.FillScreen != nil {
		if *resource.FillScreen != *other.FillScreen {
			return false
		}
	}

	if len(resource.Items) != len(other.Items) {
		return false
	}

	for i1 := range resource.Items {
		if !resource.Items[i1].Equals(other.Items[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AutoGridLayoutSpec` fields for violations and returns them.
func (resource AutoGridLayoutSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Items {
		if err := resource.Items[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AutoGridLayoutItemKind struct {
	Kind string                 `json:"kind"`
	Spec AutoGridLayoutItemSpec `json:"spec"`
}

// NewAutoGridLayoutItemKind creates a new AutoGridLayoutItemKind object.
func NewAutoGridLayoutItemKind() *AutoGridLayoutItemKind {
	return &AutoGridLayoutItemKind{
		Kind: "AutoGridLayoutItem",
		Spec: *NewAutoGridLayoutItemSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutItemKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AutoGridLayoutItemKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = AutoGridLayoutItemSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AutoGridLayoutItemKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AutoGridLayoutItemKind` objects.
func (resource AutoGridLayoutItemKind) Equals(other AutoGridLayoutItemKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AutoGridLayoutItemKind` fields for violations and returns them.
func (resource AutoGridLayoutItemKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AutoGridLayoutItemSpec struct {
	Element              ElementReference               `json:"element"`
	Repeat               *AutoGridRepeatOptions         `json:"repeat,omitempty"`
	ConditionalRendering *ConditionalRenderingGroupKind `json:"conditionalRendering,omitempty"`
}

// NewAutoGridLayoutItemSpec creates a new AutoGridLayoutItemSpec object.
func NewAutoGridLayoutItemSpec() *AutoGridLayoutItemSpec {
	return &AutoGridLayoutItemSpec{
		Element: *NewElementReference(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutItemSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AutoGridLayoutItemSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "element"
	if fields["element"] != nil {
		if string(fields["element"]) != "null" {

			resource.Element = ElementReference{}
			if err := resource.Element.UnmarshalJSONStrict(fields["element"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("element", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("element", errors.New("required field is null"))...)

		}
		delete(fields, "element")
	} else {
		errs = append(errs, cog.MakeBuildErrors("element", errors.New("required field is missing from input"))...)
	}
	// Field "repeat"
	if fields["repeat"] != nil {
		if string(fields["repeat"]) != "null" {

			resource.Repeat = &AutoGridRepeatOptions{}
			if err := resource.Repeat.UnmarshalJSONStrict(fields["repeat"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
			}

		}
		delete(fields, "repeat")

	}
	// Field "conditionalRendering"
	if fields["conditionalRendering"] != nil {
		if string(fields["conditionalRendering"]) != "null" {

			resource.ConditionalRendering = &ConditionalRenderingGroupKind{}
			if err := resource.ConditionalRendering.UnmarshalJSONStrict(fields["conditionalRendering"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("conditionalRendering", err)...)
			}

		}
		delete(fields, "conditionalRendering")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AutoGridLayoutItemSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AutoGridLayoutItemSpec` objects.
func (resource AutoGridLayoutItemSpec) Equals(other AutoGridLayoutItemSpec) bool {
	if !resource.Element.Equals(other.Element) {
		return false
	}
	if resource.Repeat == nil && other.Repeat != nil || resource.Repeat != nil && other.Repeat == nil {
		return false
	}

	if resource.Repeat != nil {
		if !resource.Repeat.Equals(*other.Repeat) {
			return false
		}
	}
	if resource.ConditionalRendering == nil && other.ConditionalRendering != nil || resource.ConditionalRendering != nil && other.ConditionalRendering == nil {
		return false
	}

	if resource.ConditionalRendering != nil {
		if !resource.ConditionalRendering.Equals(*other.ConditionalRendering) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AutoGridLayoutItemSpec` fields for violations and returns them.
func (resource AutoGridLayoutItemSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Element.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("element", err)...)
	}
	if resource.Repeat != nil {
		if err := resource.Repeat.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
		}
	}
	if resource.ConditionalRendering != nil {
		if err := resource.ConditionalRendering.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("conditionalRendering", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AutoGridRepeatOptions struct {
	Mode  RepeatMode `json:"mode"`
	Value string     `json:"value"`
}

// NewAutoGridRepeatOptions creates a new AutoGridRepeatOptions object.
func NewAutoGridRepeatOptions() *AutoGridRepeatOptions {
	return &AutoGridRepeatOptions{
		Mode: RepeatModeVariable,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridRepeatOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AutoGridRepeatOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AutoGridRepeatOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AutoGridRepeatOptions` objects.
func (resource AutoGridRepeatOptions) Equals(other AutoGridRepeatOptions) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AutoGridRepeatOptions` fields for violations and returns them.
func (resource AutoGridRepeatOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.Mode != "variable" {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("must be variable"))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TabsLayoutKind struct {
	Kind string         `json:"kind"`
	Spec TabsLayoutSpec `json:"spec"`
}

// NewTabsLayoutKind creates a new TabsLayoutKind object.
func NewTabsLayoutKind() *TabsLayoutKind {
	return &TabsLayoutKind{
		Kind: "TabsLayout",
		Spec: *NewTabsLayoutSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TabsLayoutKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = TabsLayoutSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TabsLayoutKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TabsLayoutKind` objects.
func (resource TabsLayoutKind) Equals(other TabsLayoutKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TabsLayoutKind` fields for violations and returns them.
func (resource TabsLayoutKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TabsLayoutSpec struct {
	Tabs []TabsLayoutTabKind `json:"tabs"`
}

// NewTabsLayoutSpec creates a new TabsLayoutSpec object.
func NewTabsLayoutSpec() *TabsLayoutSpec {
	return &TabsLayoutSpec{
		Tabs: []TabsLayoutTabKind{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TabsLayoutSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "tabs"
	if fields["tabs"] != nil {
		if string(fields["tabs"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["tabs"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TabsLayoutTabKind

				result1 = TabsLayoutTabKind{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("tabs["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Tabs = append(resource.Tabs, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tabs", errors.New("required field is null"))...)

		}
		delete(fields, "tabs")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tabs", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TabsLayoutSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TabsLayoutSpec` objects.
func (resource TabsLayoutSpec) Equals(other TabsLayoutSpec) bool {

	if len(resource.Tabs) != len(other.Tabs) {
		return false
	}

	for i1 := range resource.Tabs {
		if !resource.Tabs[i1].Equals(other.Tabs[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TabsLayoutSpec` fields for violations and returns them.
func (resource TabsLayoutSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Tabs {
		if err := resource.Tabs[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("tabs["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TabsLayoutTabKind struct {
	Kind string            `json:"kind"`
	Spec TabsLayoutTabSpec `json:"spec"`
}

// NewTabsLayoutTabKind creates a new TabsLayoutTabKind object.
func NewTabsLayoutTabKind() *TabsLayoutTabKind {
	return &TabsLayoutTabKind{
		Kind: "TabsLayoutTab",
		Spec: *NewTabsLayoutTabSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutTabKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TabsLayoutTabKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = TabsLayoutTabSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TabsLayoutTabKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TabsLayoutTabKind` objects.
func (resource TabsLayoutTabKind) Equals(other TabsLayoutTabKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TabsLayoutTabKind` fields for violations and returns them.
func (resource TabsLayoutTabKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TabsLayoutTabSpec struct {
	Title                *string                                                            `json:"title,omitempty"`
	Layout               GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind `json:"layout"`
	ConditionalRendering *ConditionalRenderingGroupKind                                     `json:"conditionalRendering,omitempty"`
	Repeat               *TabRepeatOptions                                                  `json:"repeat,omitempty"`
}

// NewTabsLayoutTabSpec creates a new TabsLayoutTabSpec object.
func NewTabsLayoutTabSpec() *TabsLayoutTabSpec {
	return &TabsLayoutTabSpec{
		Layout: *NewGridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutTabSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TabsLayoutTabSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "layout"
	if fields["layout"] != nil {
		if string(fields["layout"]) != "null" {

			resource.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{}
			if err := resource.Layout.UnmarshalJSONStrict(fields["layout"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("layout", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("layout", errors.New("required field is null"))...)

		}
		delete(fields, "layout")
	} else {
		errs = append(errs, cog.MakeBuildErrors("layout", errors.New("required field is missing from input"))...)
	}
	// Field "conditionalRendering"
	if fields["conditionalRendering"] != nil {
		if string(fields["conditionalRendering"]) != "null" {

			resource.ConditionalRendering = &ConditionalRenderingGroupKind{}
			if err := resource.ConditionalRendering.UnmarshalJSONStrict(fields["conditionalRendering"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("conditionalRendering", err)...)
			}

		}
		delete(fields, "conditionalRendering")

	}
	// Field "repeat"
	if fields["repeat"] != nil {
		if string(fields["repeat"]) != "null" {

			resource.Repeat = &TabRepeatOptions{}
			if err := resource.Repeat.UnmarshalJSONStrict(fields["repeat"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
			}

		}
		delete(fields, "repeat")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TabsLayoutTabSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TabsLayoutTabSpec` objects.
func (resource TabsLayoutTabSpec) Equals(other TabsLayoutTabSpec) bool {
	if resource.Title == nil && other.Title != nil || resource.Title != nil && other.Title == nil {
		return false
	}

	if resource.Title != nil {
		if *resource.Title != *other.Title {
			return false
		}
	}
	if !resource.Layout.Equals(other.Layout) {
		return false
	}
	if resource.ConditionalRendering == nil && other.ConditionalRendering != nil || resource.ConditionalRendering != nil && other.ConditionalRendering == nil {
		return false
	}

	if resource.ConditionalRendering != nil {
		if !resource.ConditionalRendering.Equals(*other.ConditionalRendering) {
			return false
		}
	}
	if resource.Repeat == nil && other.Repeat != nil || resource.Repeat != nil && other.Repeat == nil {
		return false
	}

	if resource.Repeat != nil {
		if !resource.Repeat.Equals(*other.Repeat) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TabsLayoutTabSpec` fields for violations and returns them.
func (resource TabsLayoutTabSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Layout.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("layout", err)...)
	}
	if resource.ConditionalRendering != nil {
		if err := resource.ConditionalRendering.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("conditionalRendering", err)...)
		}
	}
	if resource.Repeat != nil {
		if err := resource.Repeat.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("repeat", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TabRepeatOptions struct {
	Mode  RepeatMode `json:"mode"`
	Value string     `json:"value"`
}

// NewTabRepeatOptions creates a new TabRepeatOptions object.
func NewTabRepeatOptions() *TabRepeatOptions {
	return &TabRepeatOptions{
		Mode: RepeatModeVariable,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabRepeatOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TabRepeatOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TabRepeatOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TabRepeatOptions` objects.
func (resource TabRepeatOptions) Equals(other TabRepeatOptions) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TabRepeatOptions` fields for violations and returns them.
func (resource TabRepeatOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.Mode != "variable" {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("must be variable"))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Links with references to other dashboards or external resources
type DashboardLink struct {
	// Title to display with the link
	Title string `json:"title"`
	// Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
	// FIXME: The type is generated as `type: DashboardLinkType | dashboardLinkType.Link;` but it should be `type: DashboardLinkType`
	Type DashboardLinkType `json:"type"`
	// Icon name to be displayed with the link
	Icon string `json:"icon"`
	// Tooltip to display when the user hovers their mouse over it
	Tooltip string `json:"tooltip"`
	// Link URL. Only required/valid if the type is link
	Url *string `json:"url,omitempty"`
	// List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
	Tags []string `json:"tags"`
	// If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
	AsDropdown bool `json:"asDropdown"`
	// If true, the link will be opened in a new tab
	TargetBlank bool `json:"targetBlank"`
	// If true, includes current template variables values in the link as query params
	IncludeVars bool `json:"includeVars"`
	// If true, includes current time range in the link as query params
	KeepTime bool `json:"keepTime"`
	// Placement can be used to display the link somewhere else on the dashboard other than above the visualisations.
	Placement *string `json:"placement,omitempty"`
}

// NewDashboardLink creates a new DashboardLink object.
func NewDashboardLink() *DashboardLink {
	return &DashboardLink{
		Tags:        []string{},
		AsDropdown:  false,
		TargetBlank: false,
		IncludeVars: false,
		KeepTime:    false,
		Placement:   (func(input string) *string { return &input })(DashboardLinkPlacement),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardLink` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DashboardLink) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
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
	// Field "icon"
	if fields["icon"] != nil {
		if string(fields["icon"]) != "null" {
			if err := json.Unmarshal(fields["icon"], &resource.Icon); err != nil {
				errs = append(errs, cog.MakeBuildErrors("icon", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("icon", errors.New("required field is null"))...)

		}
		delete(fields, "icon")
	} else {
		errs = append(errs, cog.MakeBuildErrors("icon", errors.New("required field is missing from input"))...)
	}
	// Field "tooltip"
	if fields["tooltip"] != nil {
		if string(fields["tooltip"]) != "null" {
			if err := json.Unmarshal(fields["tooltip"], &resource.Tooltip); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is null"))...)

		}
		delete(fields, "tooltip")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is missing from input"))...)
	}
	// Field "url"
	if fields["url"] != nil {
		if string(fields["url"]) != "null" {
			if err := json.Unmarshal(fields["url"], &resource.Url); err != nil {
				errs = append(errs, cog.MakeBuildErrors("url", err)...)
			}

		}
		delete(fields, "url")

	}
	// Field "tags"
	if fields["tags"] != nil {
		if string(fields["tags"]) != "null" {

			if err := json.Unmarshal(fields["tags"], &resource.Tags); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tags", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is null"))...)

		}
		delete(fields, "tags")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is missing from input"))...)
	}
	// Field "asDropdown"
	if fields["asDropdown"] != nil {
		if string(fields["asDropdown"]) != "null" {
			if err := json.Unmarshal(fields["asDropdown"], &resource.AsDropdown); err != nil {
				errs = append(errs, cog.MakeBuildErrors("asDropdown", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("asDropdown", errors.New("required field is null"))...)

		}
		delete(fields, "asDropdown")

	}
	// Field "targetBlank"
	if fields["targetBlank"] != nil {
		if string(fields["targetBlank"]) != "null" {
			if err := json.Unmarshal(fields["targetBlank"], &resource.TargetBlank); err != nil {
				errs = append(errs, cog.MakeBuildErrors("targetBlank", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("targetBlank", errors.New("required field is null"))...)

		}
		delete(fields, "targetBlank")

	}
	// Field "includeVars"
	if fields["includeVars"] != nil {
		if string(fields["includeVars"]) != "null" {
			if err := json.Unmarshal(fields["includeVars"], &resource.IncludeVars); err != nil {
				errs = append(errs, cog.MakeBuildErrors("includeVars", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("includeVars", errors.New("required field is null"))...)

		}
		delete(fields, "includeVars")

	}
	// Field "keepTime"
	if fields["keepTime"] != nil {
		if string(fields["keepTime"]) != "null" {
			if err := json.Unmarshal(fields["keepTime"], &resource.KeepTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("keepTime", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("keepTime", errors.New("required field is null"))...)

		}
		delete(fields, "keepTime")

	}
	// Field "placement"
	if fields["placement"] != nil {
		if string(fields["placement"]) != "null" {
			if err := json.Unmarshal(fields["placement"], &resource.Placement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("placement", err)...)
			}

		}
		delete(fields, "placement")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DashboardLink", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DashboardLink` objects.
func (resource DashboardLink) Equals(other DashboardLink) bool {
	if resource.Title != other.Title {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Icon != other.Icon {
		return false
	}
	if resource.Tooltip != other.Tooltip {
		return false
	}
	if resource.Url == nil && other.Url != nil || resource.Url != nil && other.Url == nil {
		return false
	}

	if resource.Url != nil {
		if *resource.Url != *other.Url {
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
	if resource.AsDropdown != other.AsDropdown {
		return false
	}
	if resource.TargetBlank != other.TargetBlank {
		return false
	}
	if resource.IncludeVars != other.IncludeVars {
		return false
	}
	if resource.KeepTime != other.KeepTime {
		return false
	}
	if resource.Placement == nil && other.Placement != nil || resource.Placement != nil && other.Placement == nil {
		return false
	}

	if resource.Placement != nil {
		if *resource.Placement != *other.Placement {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DashboardLink` fields for violations and returns them.
func (resource DashboardLink) Validate() error {
	return nil
}

// Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
type DashboardLinkType string

const (
	DashboardLinkTypeLink       DashboardLinkType = "link"
	DashboardLinkTypeDashboards DashboardLinkType = "dashboards"
)

// Dashboard Link placement. Defines where the link should be displayed.
// - "inControlsMenu" renders the link in bottom part of the dashboard controls dropdown menu
const DashboardLinkPlacement = "inControlsMenu"

// Time configuration
// It defines the default time config for the time picker, the refresh picker for the specific dashboard.
type TimeSettingsSpec struct {
	// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
	Timezone *string `json:"timezone,omitempty"`
	// Start time range for dashboard.
	// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
	From string `json:"from"`
	// End time range for dashboard.
	// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
	To string `json:"to"`
	// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
	// v1: refresh
	AutoRefresh string `json:"autoRefresh"`
	// Interval options available in the refresh picker dropdown.
	// v1: timepicker.refresh_intervals
	AutoRefreshIntervals []string `json:"autoRefreshIntervals"`
	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	// v1: timepicker.quick_ranges , not exposed in the UI
	QuickRanges []TimeRangeOption `json:"quickRanges,omitempty"`
	// Whether timepicker is visible or not.
	// v1: timepicker.hidden
	HideTimepicker bool `json:"hideTimepicker"`
	// Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
	WeekStart *TimeSettingsSpecWeekStart `json:"weekStart,omitempty"`
	// The month that the fiscal year starts on. 0 = January, 11 = December
	FiscalYearStartMonth int64 `json:"fiscalYearStartMonth"`
	// Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
	// v1: timepicker.nowDelay
	NowDelay *string `json:"nowDelay,omitempty"`
}

// NewTimeSettingsSpec creates a new TimeSettingsSpec object.
func NewTimeSettingsSpec() *TimeSettingsSpec {
	return &TimeSettingsSpec{
		Timezone:             (func(input string) *string { return &input })("browser"),
		From:                 "now-6h",
		To:                   "now",
		AutoRefresh:          "",
		AutoRefreshIntervals: []string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"},
		HideTimepicker:       false,
		FiscalYearStartMonth: 0,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeSettingsSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeSettingsSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "timezone"
	if fields["timezone"] != nil {
		if string(fields["timezone"]) != "null" {
			if err := json.Unmarshal(fields["timezone"], &resource.Timezone); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timezone", err)...)
			}

		}
		delete(fields, "timezone")

	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")

	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")

	}
	// Field "autoRefresh"
	if fields["autoRefresh"] != nil {
		if string(fields["autoRefresh"]) != "null" {
			if err := json.Unmarshal(fields["autoRefresh"], &resource.AutoRefresh); err != nil {
				errs = append(errs, cog.MakeBuildErrors("autoRefresh", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("autoRefresh", errors.New("required field is null"))...)

		}
		delete(fields, "autoRefresh")

	}
	// Field "autoRefreshIntervals"
	if fields["autoRefreshIntervals"] != nil {
		if string(fields["autoRefreshIntervals"]) != "null" {

			if err := json.Unmarshal(fields["autoRefreshIntervals"], &resource.AutoRefreshIntervals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("autoRefreshIntervals", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("autoRefreshIntervals", errors.New("required field is null"))...)

		}
		delete(fields, "autoRefreshIntervals")

	}
	// Field "quickRanges"
	if fields["quickRanges"] != nil {
		if string(fields["quickRanges"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["quickRanges"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TimeRangeOption

				result1 = TimeRangeOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("quickRanges["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.QuickRanges = append(resource.QuickRanges, result1)
			}

		}
		delete(fields, "quickRanges")

	}
	// Field "hideTimepicker"
	if fields["hideTimepicker"] != nil {
		if string(fields["hideTimepicker"]) != "null" {
			if err := json.Unmarshal(fields["hideTimepicker"], &resource.HideTimepicker); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideTimepicker", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hideTimepicker", errors.New("required field is null"))...)

		}
		delete(fields, "hideTimepicker")

	}
	// Field "weekStart"
	if fields["weekStart"] != nil {
		if string(fields["weekStart"]) != "null" {
			if err := json.Unmarshal(fields["weekStart"], &resource.WeekStart); err != nil {
				errs = append(errs, cog.MakeBuildErrors("weekStart", err)...)
			}

		}
		delete(fields, "weekStart")

	}
	// Field "fiscalYearStartMonth"
	if fields["fiscalYearStartMonth"] != nil {
		if string(fields["fiscalYearStartMonth"]) != "null" {
			if err := json.Unmarshal(fields["fiscalYearStartMonth"], &resource.FiscalYearStartMonth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fiscalYearStartMonth", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("fiscalYearStartMonth", errors.New("required field is null"))...)

		}
		delete(fields, "fiscalYearStartMonth")

	}
	// Field "nowDelay"
	if fields["nowDelay"] != nil {
		if string(fields["nowDelay"]) != "null" {
			if err := json.Unmarshal(fields["nowDelay"], &resource.NowDelay); err != nil {
				errs = append(errs, cog.MakeBuildErrors("nowDelay", err)...)
			}

		}
		delete(fields, "nowDelay")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeSettingsSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeSettingsSpec` objects.
func (resource TimeSettingsSpec) Equals(other TimeSettingsSpec) bool {
	if resource.Timezone == nil && other.Timezone != nil || resource.Timezone != nil && other.Timezone == nil {
		return false
	}

	if resource.Timezone != nil {
		if *resource.Timezone != *other.Timezone {
			return false
		}
	}
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}
	if resource.AutoRefresh != other.AutoRefresh {
		return false
	}

	if len(resource.AutoRefreshIntervals) != len(other.AutoRefreshIntervals) {
		return false
	}

	for i1 := range resource.AutoRefreshIntervals {
		if resource.AutoRefreshIntervals[i1] != other.AutoRefreshIntervals[i1] {
			return false
		}
	}

	if len(resource.QuickRanges) != len(other.QuickRanges) {
		return false
	}

	for i1 := range resource.QuickRanges {
		if !resource.QuickRanges[i1].Equals(other.QuickRanges[i1]) {
			return false
		}
	}
	if resource.HideTimepicker != other.HideTimepicker {
		return false
	}
	if resource.WeekStart == nil && other.WeekStart != nil || resource.WeekStart != nil && other.WeekStart == nil {
		return false
	}

	if resource.WeekStart != nil {
		if *resource.WeekStart != *other.WeekStart {
			return false
		}
	}
	if resource.FiscalYearStartMonth != other.FiscalYearStartMonth {
		return false
	}
	if resource.NowDelay == nil && other.NowDelay != nil || resource.NowDelay != nil && other.NowDelay == nil {
		return false
	}

	if resource.NowDelay != nil {
		if *resource.NowDelay != *other.NowDelay {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TimeSettingsSpec` fields for violations and returns them.
func (resource TimeSettingsSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.QuickRanges {
		if err := resource.QuickRanges[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("quickRanges["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TimeRangeOption struct {
	Display string `json:"display"`
	From    string `json:"from"`
	To      string `json:"to"`
}

// NewTimeRangeOption creates a new TimeRangeOption object.
func NewTimeRangeOption() *TimeRangeOption {
	return &TimeRangeOption{
		Display: "Last 6 hours",
		From:    "now-6h",
		To:      "now",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeRangeOption` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeRangeOption) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "display"
	if fields["display"] != nil {
		if string(fields["display"]) != "null" {
			if err := json.Unmarshal(fields["display"], &resource.Display); err != nil {
				errs = append(errs, cog.MakeBuildErrors("display", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("display", errors.New("required field is null"))...)

		}
		delete(fields, "display")

	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")

	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeRangeOption", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeRangeOption` objects.
func (resource TimeRangeOption) Equals(other TimeRangeOption) bool {
	if resource.Display != other.Display {
		return false
	}
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TimeRangeOption` fields for violations and returns them.
func (resource TimeRangeOption) Validate() error {
	return nil
}

type VariableKind = QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind

// NewVariableKind creates a new VariableKind object.
func NewVariableKind() *VariableKind {
	return NewQueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind()
}

// Query variable kind
type QueryVariableKind struct {
	Kind string            `json:"kind"`
	Spec QueryVariableSpec `json:"spec"`
}

// NewQueryVariableKind creates a new QueryVariableKind object.
func NewQueryVariableKind() *QueryVariableKind {
	return &QueryVariableKind{
		Kind: "QueryVariable",
		Spec: *NewQueryVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = QueryVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryVariableKind` objects.
func (resource QueryVariableKind) Equals(other QueryVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryVariableKind` fields for violations and returns them.
func (resource QueryVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Query variable specification
type QueryVariableSpec struct {
	Name               string                               `json:"name"`
	Current            VariableOption                       `json:"current"`
	Label              *string                              `json:"label,omitempty"`
	Hide               VariableHide                         `json:"hide"`
	Refresh            VariableRefresh                      `json:"refresh"`
	SkipUrlSync        bool                                 `json:"skipUrlSync"`
	Description        *string                              `json:"description,omitempty"`
	Query              DataQueryKind                        `json:"query"`
	Regex              string                               `json:"regex"`
	RegexApplyTo       *VariableRegexApplyTo                `json:"regexApplyTo,omitempty"`
	Sort               VariableSort                         `json:"sort"`
	Definition         *string                              `json:"definition,omitempty"`
	Options            []VariableOption                     `json:"options"`
	Multi              bool                                 `json:"multi"`
	IncludeAll         bool                                 `json:"includeAll"`
	AllValue           *string                              `json:"allValue,omitempty"`
	Placeholder        *string                              `json:"placeholder,omitempty"`
	AllowCustomValue   bool                                 `json:"allowCustomValue"`
	StaticOptions      []VariableOption                     `json:"staticOptions,omitempty"`
	StaticOptionsOrder *QueryVariableSpecStaticOptionsOrder `json:"staticOptionsOrder,omitempty"`
}

// NewQueryVariableSpec creates a new QueryVariableSpec object.
func NewQueryVariableSpec() *QueryVariableSpec {
	return &QueryVariableSpec{
		Name: "",
		Current: VariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Hide:             VariableHideDontHide,
		Refresh:          VariableRefreshNever,
		SkipUrlSync:      false,
		Query:            *NewDataQueryKind(),
		Regex:            "",
		RegexApplyTo:     (func(input VariableRegexApplyTo) *VariableRegexApplyTo { return &input })(VariableRegexApplyToValue),
		Options:          []VariableOption{},
		Multi:            false,
		IncludeAll:       false,
		AllowCustomValue: true,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {

			resource.Current = VariableOption{}
			if err := resource.Current.UnmarshalJSONStrict(fields["current"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "refresh"
	if fields["refresh"] != nil {
		if string(fields["refresh"]) != "null" {
			if err := json.Unmarshal(fields["refresh"], &resource.Refresh); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refresh", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refresh", errors.New("required field is null"))...)

		}
		delete(fields, "refresh")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {

			resource.Query = DataQueryKind{}
			if err := resource.Query.UnmarshalJSONStrict(fields["query"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "regex"
	if fields["regex"] != nil {
		if string(fields["regex"]) != "null" {
			if err := json.Unmarshal(fields["regex"], &resource.Regex); err != nil {
				errs = append(errs, cog.MakeBuildErrors("regex", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("regex", errors.New("required field is null"))...)

		}
		delete(fields, "regex")

	}
	// Field "regexApplyTo"
	if fields["regexApplyTo"] != nil {
		if string(fields["regexApplyTo"]) != "null" {
			if err := json.Unmarshal(fields["regexApplyTo"], &resource.RegexApplyTo); err != nil {
				errs = append(errs, cog.MakeBuildErrors("regexApplyTo", err)...)
			}

		}
		delete(fields, "regexApplyTo")

	}
	// Field "sort"
	if fields["sort"] != nil {
		if string(fields["sort"]) != "null" {
			if err := json.Unmarshal(fields["sort"], &resource.Sort); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sort", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("sort", errors.New("required field is null"))...)

		}
		delete(fields, "sort")
	} else {
		errs = append(errs, cog.MakeBuildErrors("sort", errors.New("required field is missing from input"))...)
	}
	// Field "definition"
	if fields["definition"] != nil {
		if string(fields["definition"]) != "null" {
			if err := json.Unmarshal(fields["definition"], &resource.Definition); err != nil {
				errs = append(errs, cog.MakeBuildErrors("definition", err)...)
			}

		}
		delete(fields, "definition")

	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["options"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 VariableOption

				result1 = VariableOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Options = append(resource.Options, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}
	// Field "multi"
	if fields["multi"] != nil {
		if string(fields["multi"]) != "null" {
			if err := json.Unmarshal(fields["multi"], &resource.Multi); err != nil {
				errs = append(errs, cog.MakeBuildErrors("multi", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("multi", errors.New("required field is null"))...)

		}
		delete(fields, "multi")

	}
	// Field "includeAll"
	if fields["includeAll"] != nil {
		if string(fields["includeAll"]) != "null" {
			if err := json.Unmarshal(fields["includeAll"], &resource.IncludeAll); err != nil {
				errs = append(errs, cog.MakeBuildErrors("includeAll", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("includeAll", errors.New("required field is null"))...)

		}
		delete(fields, "includeAll")

	}
	// Field "allValue"
	if fields["allValue"] != nil {
		if string(fields["allValue"]) != "null" {
			if err := json.Unmarshal(fields["allValue"], &resource.AllValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allValue", err)...)
			}

		}
		delete(fields, "allValue")

	}
	// Field "placeholder"
	if fields["placeholder"] != nil {
		if string(fields["placeholder"]) != "null" {
			if err := json.Unmarshal(fields["placeholder"], &resource.Placeholder); err != nil {
				errs = append(errs, cog.MakeBuildErrors("placeholder", err)...)
			}

		}
		delete(fields, "placeholder")

	}
	// Field "allowCustomValue"
	if fields["allowCustomValue"] != nil {
		if string(fields["allowCustomValue"]) != "null" {
			if err := json.Unmarshal(fields["allowCustomValue"], &resource.AllowCustomValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allowCustomValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("allowCustomValue", errors.New("required field is null"))...)

		}
		delete(fields, "allowCustomValue")

	}
	// Field "staticOptions"
	if fields["staticOptions"] != nil {
		if string(fields["staticOptions"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["staticOptions"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 VariableOption

				result1 = VariableOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("staticOptions["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.StaticOptions = append(resource.StaticOptions, result1)
			}

		}
		delete(fields, "staticOptions")

	}
	// Field "staticOptionsOrder"
	if fields["staticOptionsOrder"] != nil {
		if string(fields["staticOptionsOrder"]) != "null" {
			if err := json.Unmarshal(fields["staticOptionsOrder"], &resource.StaticOptionsOrder); err != nil {
				errs = append(errs, cog.MakeBuildErrors("staticOptionsOrder", err)...)
			}

		}
		delete(fields, "staticOptionsOrder")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryVariableSpec` objects.
func (resource QueryVariableSpec) Equals(other QueryVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if !resource.Current.Equals(other.Current) {
		return false
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.Refresh != other.Refresh {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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
	if !resource.Query.Equals(other.Query) {
		return false
	}
	if resource.Regex != other.Regex {
		return false
	}
	if resource.RegexApplyTo == nil && other.RegexApplyTo != nil || resource.RegexApplyTo != nil && other.RegexApplyTo == nil {
		return false
	}

	if resource.RegexApplyTo != nil {
		if *resource.RegexApplyTo != *other.RegexApplyTo {
			return false
		}
	}
	if resource.Sort != other.Sort {
		return false
	}
	if resource.Definition == nil && other.Definition != nil || resource.Definition != nil && other.Definition == nil {
		return false
	}

	if resource.Definition != nil {
		if *resource.Definition != *other.Definition {
			return false
		}
	}

	if len(resource.Options) != len(other.Options) {
		return false
	}

	for i1 := range resource.Options {
		if !resource.Options[i1].Equals(other.Options[i1]) {
			return false
		}
	}
	if resource.Multi != other.Multi {
		return false
	}
	if resource.IncludeAll != other.IncludeAll {
		return false
	}
	if resource.AllValue == nil && other.AllValue != nil || resource.AllValue != nil && other.AllValue == nil {
		return false
	}

	if resource.AllValue != nil {
		if *resource.AllValue != *other.AllValue {
			return false
		}
	}
	if resource.Placeholder == nil && other.Placeholder != nil || resource.Placeholder != nil && other.Placeholder == nil {
		return false
	}

	if resource.Placeholder != nil {
		if *resource.Placeholder != *other.Placeholder {
			return false
		}
	}
	if resource.AllowCustomValue != other.AllowCustomValue {
		return false
	}

	if len(resource.StaticOptions) != len(other.StaticOptions) {
		return false
	}

	for i1 := range resource.StaticOptions {
		if !resource.StaticOptions[i1].Equals(other.StaticOptions[i1]) {
			return false
		}
	}
	if resource.StaticOptionsOrder == nil && other.StaticOptionsOrder != nil || resource.StaticOptionsOrder != nil && other.StaticOptionsOrder == nil {
		return false
	}

	if resource.StaticOptionsOrder != nil {
		if *resource.StaticOptionsOrder != *other.StaticOptionsOrder {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryVariableSpec` fields for violations and returns them.
func (resource QueryVariableSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Current.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("current", err)...)
	}
	if err := resource.Query.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("query", err)...)
	}

	for i1 := range resource.Options {
		if err := resource.Options[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.StaticOptions {
		if err := resource.StaticOptions[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("staticOptions["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Variable option specification
type VariableOption struct {
	// Whether the option is selected or not
	Selected *bool `json:"selected,omitempty"`
	// Text to be displayed for the option
	Text StringOrArrayOfString `json:"text"`
	// Value of the option
	Value StringOrArrayOfString `json:"value"`
}

// NewVariableOption creates a new VariableOption object.
func NewVariableOption() *VariableOption {
	return &VariableOption{
		Text:  *NewStringOrArrayOfString(),
		Value: *NewStringOrArrayOfString(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VariableOption` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *VariableOption) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "selected"
	if fields["selected"] != nil {
		if string(fields["selected"]) != "null" {
			if err := json.Unmarshal(fields["selected"], &resource.Selected); err != nil {
				errs = append(errs, cog.MakeBuildErrors("selected", err)...)
			}

		}
		delete(fields, "selected")

	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {

			resource.Text = StringOrArrayOfString{}
			if err := resource.Text.UnmarshalJSONStrict(fields["text"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("text", errors.New("required field is null"))...)

		}
		delete(fields, "text")
	} else {
		errs = append(errs, cog.MakeBuildErrors("text", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {

			resource.Value = StringOrArrayOfString{}
			if err := resource.Value.UnmarshalJSONStrict(fields["value"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("VariableOption", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `VariableOption` objects.
func (resource VariableOption) Equals(other VariableOption) bool {
	if resource.Selected == nil && other.Selected != nil || resource.Selected != nil && other.Selected == nil {
		return false
	}

	if resource.Selected != nil {
		if *resource.Selected != *other.Selected {
			return false
		}
	}
	if !resource.Text.Equals(other.Text) {
		return false
	}
	if !resource.Value.Equals(other.Value) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `VariableOption` fields for violations and returns them.
func (resource VariableOption) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Text.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("text", err)...)
	}
	if err := resource.Value.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("value", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Determine if the variable shows on dashboard
// Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).
type VariableHide string

const (
	VariableHideDontHide       VariableHide = "dontHide"
	VariableHideHideLabel      VariableHide = "hideLabel"
	VariableHideHideVariable   VariableHide = "hideVariable"
	VariableHideInControlsMenu VariableHide = "inControlsMenu"
)

// Options to config when to refresh a variable
// `never`: Never refresh the variable
// `onDashboardLoad`: Queries the data source every time the dashboard loads.
// `onTimeRangeChanged`: Queries the data source when the dashboard time range changes.
type VariableRefresh string

const (
	VariableRefreshNever              VariableRefresh = "never"
	VariableRefreshOnDashboardLoad    VariableRefresh = "onDashboardLoad"
	VariableRefreshOnTimeRangeChanged VariableRefresh = "onTimeRangeChanged"
)

// Determine whether regex applies to variable value or display text
// Accepted values are `value` (apply to value used in queries) or `text` (apply to display text shown to users)
type VariableRegexApplyTo string

const (
	VariableRegexApplyToValue VariableRegexApplyTo = "value"
	VariableRegexApplyToText  VariableRegexApplyTo = "text"
)

// Sort variable options
// Accepted values are:
// `disabled`: No sorting
// `alphabeticalAsc`: Alphabetical ASC
// `alphabeticalDesc`: Alphabetical DESC
// `numericalAsc`: Numerical ASC
// `numericalDesc`: Numerical DESC
// `alphabeticalCaseInsensitiveAsc`: Alphabetical Case Insensitive ASC
// `alphabeticalCaseInsensitiveDesc`: Alphabetical Case Insensitive DESC
// `naturalAsc`: Natural ASC
// `naturalDesc`: Natural DESC
// VariableSort enum with default value
type VariableSort string

const (
	VariableSortDisabled                        VariableSort = "disabled"
	VariableSortAlphabeticalAsc                 VariableSort = "alphabeticalAsc"
	VariableSortAlphabeticalDesc                VariableSort = "alphabeticalDesc"
	VariableSortNumericalAsc                    VariableSort = "numericalAsc"
	VariableSortNumericalDesc                   VariableSort = "numericalDesc"
	VariableSortAlphabeticalCaseInsensitiveAsc  VariableSort = "alphabeticalCaseInsensitiveAsc"
	VariableSortAlphabeticalCaseInsensitiveDesc VariableSort = "alphabeticalCaseInsensitiveDesc"
	VariableSortNaturalAsc                      VariableSort = "naturalAsc"
	VariableSortNaturalDesc                     VariableSort = "naturalDesc"
)

// Text variable kind
type TextVariableKind struct {
	Kind string           `json:"kind"`
	Spec TextVariableSpec `json:"spec"`
}

// NewTextVariableKind creates a new TextVariableKind object.
func NewTextVariableKind() *TextVariableKind {
	return &TextVariableKind{
		Kind: "TextVariable",
		Spec: *NewTextVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TextVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TextVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = TextVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TextVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TextVariableKind` objects.
func (resource TextVariableKind) Equals(other TextVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TextVariableKind` fields for violations and returns them.
func (resource TextVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Text variable specification
type TextVariableSpec struct {
	Name        string         `json:"name"`
	Current     VariableOption `json:"current"`
	Query       string         `json:"query"`
	Label       *string        `json:"label,omitempty"`
	Hide        VariableHide   `json:"hide"`
	SkipUrlSync bool           `json:"skipUrlSync"`
	Description *string        `json:"description,omitempty"`
}

// NewTextVariableSpec creates a new TextVariableSpec object.
func NewTextVariableSpec() *TextVariableSpec {
	return &TextVariableSpec{
		Name: "",
		Current: VariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Query:       "",
		Hide:        VariableHideDontHide,
		SkipUrlSync: false,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TextVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TextVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {

			resource.Current = VariableOption{}
			if err := resource.Current.UnmarshalJSONStrict(fields["current"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")

	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TextVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TextVariableSpec` objects.
func (resource TextVariableSpec) Equals(other TextVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if !resource.Current.Equals(other.Current) {
		return false
	}
	if resource.Query != other.Query {
		return false
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `TextVariableSpec` fields for violations and returns them.
func (resource TextVariableSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Current.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("current", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Constant variable kind
type ConstantVariableKind struct {
	Kind string               `json:"kind"`
	Spec ConstantVariableSpec `json:"spec"`
}

// NewConstantVariableKind creates a new ConstantVariableKind object.
func NewConstantVariableKind() *ConstantVariableKind {
	return &ConstantVariableKind{
		Kind: "ConstantVariable",
		Spec: *NewConstantVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConstantVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConstantVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = ConstantVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConstantVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConstantVariableKind` objects.
func (resource ConstantVariableKind) Equals(other ConstantVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConstantVariableKind` fields for violations and returns them.
func (resource ConstantVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Constant variable specification
type ConstantVariableSpec struct {
	Name        string         `json:"name"`
	Query       string         `json:"query"`
	Current     VariableOption `json:"current"`
	Label       *string        `json:"label,omitempty"`
	Hide        VariableHide   `json:"hide"`
	SkipUrlSync bool           `json:"skipUrlSync"`
	Description *string        `json:"description,omitempty"`
}

// NewConstantVariableSpec creates a new ConstantVariableSpec object.
func NewConstantVariableSpec() *ConstantVariableSpec {
	return &ConstantVariableSpec{
		Name:  "",
		Query: "",
		Current: VariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Hide:        VariableHideDontHide,
		SkipUrlSync: false,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConstantVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConstantVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {

			resource.Current = VariableOption{}
			if err := resource.Current.UnmarshalJSONStrict(fields["current"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConstantVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConstantVariableSpec` objects.
func (resource ConstantVariableSpec) Equals(other ConstantVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Query != other.Query {
		return false
	}
	if !resource.Current.Equals(other.Current) {
		return false
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `ConstantVariableSpec` fields for violations and returns them.
func (resource ConstantVariableSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Current.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("current", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Datasource variable kind
type DatasourceVariableKind struct {
	Kind string                 `json:"kind"`
	Spec DatasourceVariableSpec `json:"spec"`
}

// NewDatasourceVariableKind creates a new DatasourceVariableKind object.
func NewDatasourceVariableKind() *DatasourceVariableKind {
	return &DatasourceVariableKind{
		Kind: "DatasourceVariable",
		Spec: *NewDatasourceVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DatasourceVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DatasourceVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = DatasourceVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DatasourceVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DatasourceVariableKind` objects.
func (resource DatasourceVariableKind) Equals(other DatasourceVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DatasourceVariableKind` fields for violations and returns them.
func (resource DatasourceVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Datasource variable specification
type DatasourceVariableSpec struct {
	Name             string           `json:"name"`
	PluginId         string           `json:"pluginId"`
	Refresh          VariableRefresh  `json:"refresh"`
	Regex            string           `json:"regex"`
	Current          VariableOption   `json:"current"`
	Options          []VariableOption `json:"options"`
	Multi            bool             `json:"multi"`
	IncludeAll       bool             `json:"includeAll"`
	AllValue         *string          `json:"allValue,omitempty"`
	Label            *string          `json:"label,omitempty"`
	Hide             VariableHide     `json:"hide"`
	SkipUrlSync      bool             `json:"skipUrlSync"`
	Description      *string          `json:"description,omitempty"`
	AllowCustomValue bool             `json:"allowCustomValue"`
}

// NewDatasourceVariableSpec creates a new DatasourceVariableSpec object.
func NewDatasourceVariableSpec() *DatasourceVariableSpec {
	return &DatasourceVariableSpec{
		Name:     "",
		PluginId: "",
		Refresh:  VariableRefreshNever,
		Regex:    "",
		Current: VariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Options:          []VariableOption{},
		Multi:            false,
		IncludeAll:       false,
		Hide:             VariableHideDontHide,
		SkipUrlSync:      false,
		AllowCustomValue: true,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DatasourceVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DatasourceVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "pluginId"
	if fields["pluginId"] != nil {
		if string(fields["pluginId"]) != "null" {
			if err := json.Unmarshal(fields["pluginId"], &resource.PluginId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pluginId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("pluginId", errors.New("required field is null"))...)

		}
		delete(fields, "pluginId")

	}
	// Field "refresh"
	if fields["refresh"] != nil {
		if string(fields["refresh"]) != "null" {
			if err := json.Unmarshal(fields["refresh"], &resource.Refresh); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refresh", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refresh", errors.New("required field is null"))...)

		}
		delete(fields, "refresh")

	}
	// Field "regex"
	if fields["regex"] != nil {
		if string(fields["regex"]) != "null" {
			if err := json.Unmarshal(fields["regex"], &resource.Regex); err != nil {
				errs = append(errs, cog.MakeBuildErrors("regex", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("regex", errors.New("required field is null"))...)

		}
		delete(fields, "regex")

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {

			resource.Current = VariableOption{}
			if err := resource.Current.UnmarshalJSONStrict(fields["current"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")

	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["options"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 VariableOption

				result1 = VariableOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Options = append(resource.Options, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}
	// Field "multi"
	if fields["multi"] != nil {
		if string(fields["multi"]) != "null" {
			if err := json.Unmarshal(fields["multi"], &resource.Multi); err != nil {
				errs = append(errs, cog.MakeBuildErrors("multi", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("multi", errors.New("required field is null"))...)

		}
		delete(fields, "multi")

	}
	// Field "includeAll"
	if fields["includeAll"] != nil {
		if string(fields["includeAll"]) != "null" {
			if err := json.Unmarshal(fields["includeAll"], &resource.IncludeAll); err != nil {
				errs = append(errs, cog.MakeBuildErrors("includeAll", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("includeAll", errors.New("required field is null"))...)

		}
		delete(fields, "includeAll")

	}
	// Field "allValue"
	if fields["allValue"] != nil {
		if string(fields["allValue"]) != "null" {
			if err := json.Unmarshal(fields["allValue"], &resource.AllValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allValue", err)...)
			}

		}
		delete(fields, "allValue")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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
	// Field "allowCustomValue"
	if fields["allowCustomValue"] != nil {
		if string(fields["allowCustomValue"]) != "null" {
			if err := json.Unmarshal(fields["allowCustomValue"], &resource.AllowCustomValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allowCustomValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("allowCustomValue", errors.New("required field is null"))...)

		}
		delete(fields, "allowCustomValue")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DatasourceVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DatasourceVariableSpec` objects.
func (resource DatasourceVariableSpec) Equals(other DatasourceVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.PluginId != other.PluginId {
		return false
	}
	if resource.Refresh != other.Refresh {
		return false
	}
	if resource.Regex != other.Regex {
		return false
	}
	if !resource.Current.Equals(other.Current) {
		return false
	}

	if len(resource.Options) != len(other.Options) {
		return false
	}

	for i1 := range resource.Options {
		if !resource.Options[i1].Equals(other.Options[i1]) {
			return false
		}
	}
	if resource.Multi != other.Multi {
		return false
	}
	if resource.IncludeAll != other.IncludeAll {
		return false
	}
	if resource.AllValue == nil && other.AllValue != nil || resource.AllValue != nil && other.AllValue == nil {
		return false
	}

	if resource.AllValue != nil {
		if *resource.AllValue != *other.AllValue {
			return false
		}
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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
	if resource.AllowCustomValue != other.AllowCustomValue {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DatasourceVariableSpec` fields for violations and returns them.
func (resource DatasourceVariableSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Current.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("current", err)...)
	}

	for i1 := range resource.Options {
		if err := resource.Options[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Interval variable kind
type IntervalVariableKind struct {
	Kind string               `json:"kind"`
	Spec IntervalVariableSpec `json:"spec"`
}

// NewIntervalVariableKind creates a new IntervalVariableKind object.
func NewIntervalVariableKind() *IntervalVariableKind {
	return &IntervalVariableKind{
		Kind: "IntervalVariable",
		Spec: *NewIntervalVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `IntervalVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *IntervalVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = IntervalVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("IntervalVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `IntervalVariableKind` objects.
func (resource IntervalVariableKind) Equals(other IntervalVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `IntervalVariableKind` fields for violations and returns them.
func (resource IntervalVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Interval variable specification
type IntervalVariableSpec struct {
	Name        string           `json:"name"`
	Query       string           `json:"query"`
	Current     VariableOption   `json:"current"`
	Options     []VariableOption `json:"options"`
	Auto        bool             `json:"auto"`
	AutoMin     string           `json:"auto_min"`
	AutoCount   int64            `json:"auto_count"`
	Refresh     string           `json:"refresh"`
	Label       *string          `json:"label,omitempty"`
	Hide        VariableHide     `json:"hide"`
	SkipUrlSync bool             `json:"skipUrlSync"`
	Description *string          `json:"description,omitempty"`
}

// NewIntervalVariableSpec creates a new IntervalVariableSpec object.
func NewIntervalVariableSpec() *IntervalVariableSpec {
	return &IntervalVariableSpec{
		Name:  "",
		Query: "",
		Current: VariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Options:     []VariableOption{},
		Auto:        false,
		AutoMin:     "",
		AutoCount:   0,
		Refresh:     "onTimeRangeChanged",
		Hide:        VariableHideDontHide,
		SkipUrlSync: false,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `IntervalVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *IntervalVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {

			resource.Current = VariableOption{}
			if err := resource.Current.UnmarshalJSONStrict(fields["current"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")

	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["options"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 VariableOption

				result1 = VariableOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Options = append(resource.Options, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}
	// Field "auto"
	if fields["auto"] != nil {
		if string(fields["auto"]) != "null" {
			if err := json.Unmarshal(fields["auto"], &resource.Auto); err != nil {
				errs = append(errs, cog.MakeBuildErrors("auto", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("auto", errors.New("required field is null"))...)

		}
		delete(fields, "auto")

	}
	// Field "auto_min"
	if fields["auto_min"] != nil {
		if string(fields["auto_min"]) != "null" {
			if err := json.Unmarshal(fields["auto_min"], &resource.AutoMin); err != nil {
				errs = append(errs, cog.MakeBuildErrors("auto_min", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("auto_min", errors.New("required field is null"))...)

		}
		delete(fields, "auto_min")

	}
	// Field "auto_count"
	if fields["auto_count"] != nil {
		if string(fields["auto_count"]) != "null" {
			if err := json.Unmarshal(fields["auto_count"], &resource.AutoCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("auto_count", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("auto_count", errors.New("required field is null"))...)

		}
		delete(fields, "auto_count")

	}
	// Field "refresh"
	if fields["refresh"] != nil {
		if string(fields["refresh"]) != "null" {
			if err := json.Unmarshal(fields["refresh"], &resource.Refresh); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refresh", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refresh", errors.New("required field is null"))...)

		}
		delete(fields, "refresh")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refresh", errors.New("required field is missing from input"))...)
	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("IntervalVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `IntervalVariableSpec` objects.
func (resource IntervalVariableSpec) Equals(other IntervalVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Query != other.Query {
		return false
	}
	if !resource.Current.Equals(other.Current) {
		return false
	}

	if len(resource.Options) != len(other.Options) {
		return false
	}

	for i1 := range resource.Options {
		if !resource.Options[i1].Equals(other.Options[i1]) {
			return false
		}
	}
	if resource.Auto != other.Auto {
		return false
	}
	if resource.AutoMin != other.AutoMin {
		return false
	}
	if resource.AutoCount != other.AutoCount {
		return false
	}
	if resource.Refresh != other.Refresh {
		return false
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `IntervalVariableSpec` fields for violations and returns them.
func (resource IntervalVariableSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Current.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("current", err)...)
	}

	for i1 := range resource.Options {
		if err := resource.Options[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Custom variable kind
type CustomVariableKind struct {
	Kind string             `json:"kind"`
	Spec CustomVariableSpec `json:"spec"`
}

// NewCustomVariableKind creates a new CustomVariableKind object.
func NewCustomVariableKind() *CustomVariableKind {
	return &CustomVariableKind{
		Kind: "CustomVariable",
		Spec: *NewCustomVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CustomVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = CustomVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CustomVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CustomVariableKind` objects.
func (resource CustomVariableKind) Equals(other CustomVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CustomVariableKind` fields for violations and returns them.
func (resource CustomVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Custom variable specification
type CustomVariableSpec struct {
	Name             string                          `json:"name"`
	Query            string                          `json:"query"`
	Current          VariableOption                  `json:"current"`
	Options          []VariableOption                `json:"options"`
	Multi            bool                            `json:"multi"`
	IncludeAll       bool                            `json:"includeAll"`
	AllValue         *string                         `json:"allValue,omitempty"`
	Label            *string                         `json:"label,omitempty"`
	Hide             VariableHide                    `json:"hide"`
	SkipUrlSync      bool                            `json:"skipUrlSync"`
	Description      *string                         `json:"description,omitempty"`
	AllowCustomValue bool                            `json:"allowCustomValue"`
	ValuesFormat     *CustomVariableSpecValuesFormat `json:"valuesFormat,omitempty"`
}

// NewCustomVariableSpec creates a new CustomVariableSpec object.
func NewCustomVariableSpec() *CustomVariableSpec {
	return &CustomVariableSpec{
		Name:             "",
		Query:            "",
		Current:          *NewVariableOption(),
		Options:          []VariableOption{},
		Multi:            false,
		IncludeAll:       false,
		Hide:             VariableHideDontHide,
		SkipUrlSync:      false,
		AllowCustomValue: true,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CustomVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {

			resource.Current = VariableOption{}
			if err := resource.Current.UnmarshalJSONStrict(fields["current"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")
	} else {
		errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is missing from input"))...)
	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["options"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 VariableOption

				result1 = VariableOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Options = append(resource.Options, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}
	// Field "multi"
	if fields["multi"] != nil {
		if string(fields["multi"]) != "null" {
			if err := json.Unmarshal(fields["multi"], &resource.Multi); err != nil {
				errs = append(errs, cog.MakeBuildErrors("multi", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("multi", errors.New("required field is null"))...)

		}
		delete(fields, "multi")

	}
	// Field "includeAll"
	if fields["includeAll"] != nil {
		if string(fields["includeAll"]) != "null" {
			if err := json.Unmarshal(fields["includeAll"], &resource.IncludeAll); err != nil {
				errs = append(errs, cog.MakeBuildErrors("includeAll", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("includeAll", errors.New("required field is null"))...)

		}
		delete(fields, "includeAll")

	}
	// Field "allValue"
	if fields["allValue"] != nil {
		if string(fields["allValue"]) != "null" {
			if err := json.Unmarshal(fields["allValue"], &resource.AllValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allValue", err)...)
			}

		}
		delete(fields, "allValue")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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
	// Field "allowCustomValue"
	if fields["allowCustomValue"] != nil {
		if string(fields["allowCustomValue"]) != "null" {
			if err := json.Unmarshal(fields["allowCustomValue"], &resource.AllowCustomValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allowCustomValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("allowCustomValue", errors.New("required field is null"))...)

		}
		delete(fields, "allowCustomValue")

	}
	// Field "valuesFormat"
	if fields["valuesFormat"] != nil {
		if string(fields["valuesFormat"]) != "null" {
			if err := json.Unmarshal(fields["valuesFormat"], &resource.ValuesFormat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valuesFormat", err)...)
			}

		}
		delete(fields, "valuesFormat")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CustomVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CustomVariableSpec` objects.
func (resource CustomVariableSpec) Equals(other CustomVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Query != other.Query {
		return false
	}
	if !resource.Current.Equals(other.Current) {
		return false
	}

	if len(resource.Options) != len(other.Options) {
		return false
	}

	for i1 := range resource.Options {
		if !resource.Options[i1].Equals(other.Options[i1]) {
			return false
		}
	}
	if resource.Multi != other.Multi {
		return false
	}
	if resource.IncludeAll != other.IncludeAll {
		return false
	}
	if resource.AllValue == nil && other.AllValue != nil || resource.AllValue != nil && other.AllValue == nil {
		return false
	}

	if resource.AllValue != nil {
		if *resource.AllValue != *other.AllValue {
			return false
		}
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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
	if resource.AllowCustomValue != other.AllowCustomValue {
		return false
	}
	if resource.ValuesFormat == nil && other.ValuesFormat != nil || resource.ValuesFormat != nil && other.ValuesFormat == nil {
		return false
	}

	if resource.ValuesFormat != nil {
		if *resource.ValuesFormat != *other.ValuesFormat {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CustomVariableSpec` fields for violations and returns them.
func (resource CustomVariableSpec) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Current.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("current", err)...)
	}

	for i1 := range resource.Options {
		if err := resource.Options[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Group variable kind
type GroupByVariableKind struct {
	Kind       string                                         `json:"kind"`
	Group      string                                         `json:"group"`
	Datasource *Dashboardv2beta1GroupByVariableKindDatasource `json:"datasource,omitempty"`
	Spec       GroupByVariableSpec                            `json:"spec"`
}

// NewGroupByVariableKind creates a new GroupByVariableKind object.
func NewGroupByVariableKind() *GroupByVariableKind {
	return &GroupByVariableKind{
		Kind: "GroupByVariable",
		Spec: *NewGroupByVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GroupByVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GroupByVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is null"))...)

		}
		delete(fields, "group")
	} else {
		errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is missing from input"))...)
	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &Dashboardv2beta1GroupByVariableKindDatasource{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = GroupByVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GroupByVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GroupByVariableKind` objects.
func (resource GroupByVariableKind) Equals(other GroupByVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Group != other.Group {
		return false
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GroupByVariableKind` fields for violations and returns them.
func (resource GroupByVariableKind) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// GroupBy variable specification
type GroupByVariableSpec struct {
	Name         string           `json:"name"`
	DefaultValue *VariableOption  `json:"defaultValue,omitempty"`
	Current      VariableOption   `json:"current"`
	Options      []VariableOption `json:"options"`
	Multi        bool             `json:"multi"`
	Label        *string          `json:"label,omitempty"`
	Hide         VariableHide     `json:"hide"`
	SkipUrlSync  bool             `json:"skipUrlSync"`
	Description  *string          `json:"description,omitempty"`
}

// NewGroupByVariableSpec creates a new GroupByVariableSpec object.
func NewGroupByVariableSpec() *GroupByVariableSpec {
	return &GroupByVariableSpec{
		Name: "",
		Current: VariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Options:     []VariableOption{},
		Multi:       false,
		Hide:        VariableHideDontHide,
		SkipUrlSync: false,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GroupByVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GroupByVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "defaultValue"
	if fields["defaultValue"] != nil {
		if string(fields["defaultValue"]) != "null" {

			resource.DefaultValue = &VariableOption{}
			if err := resource.DefaultValue.UnmarshalJSONStrict(fields["defaultValue"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("defaultValue", err)...)
			}

		}
		delete(fields, "defaultValue")

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {

			resource.Current = VariableOption{}
			if err := resource.Current.UnmarshalJSONStrict(fields["current"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")

	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["options"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 VariableOption

				result1 = VariableOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Options = append(resource.Options, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is null"))...)

		}
		delete(fields, "options")
	} else {
		errs = append(errs, cog.MakeBuildErrors("options", errors.New("required field is missing from input"))...)
	}
	// Field "multi"
	if fields["multi"] != nil {
		if string(fields["multi"]) != "null" {
			if err := json.Unmarshal(fields["multi"], &resource.Multi); err != nil {
				errs = append(errs, cog.MakeBuildErrors("multi", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("multi", errors.New("required field is null"))...)

		}
		delete(fields, "multi")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GroupByVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GroupByVariableSpec` objects.
func (resource GroupByVariableSpec) Equals(other GroupByVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.DefaultValue == nil && other.DefaultValue != nil || resource.DefaultValue != nil && other.DefaultValue == nil {
		return false
	}

	if resource.DefaultValue != nil {
		if !resource.DefaultValue.Equals(*other.DefaultValue) {
			return false
		}
	}
	if !resource.Current.Equals(other.Current) {
		return false
	}

	if len(resource.Options) != len(other.Options) {
		return false
	}

	for i1 := range resource.Options {
		if !resource.Options[i1].Equals(other.Options[i1]) {
			return false
		}
	}
	if resource.Multi != other.Multi {
		return false
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `GroupByVariableSpec` fields for violations and returns them.
func (resource GroupByVariableSpec) Validate() error {
	var errs cog.BuildErrors
	if resource.DefaultValue != nil {
		if err := resource.DefaultValue.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("defaultValue", err)...)
		}
	}
	if err := resource.Current.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("current", err)...)
	}

	for i1 := range resource.Options {
		if err := resource.Options[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("options["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Adhoc variable kind
type AdhocVariableKind struct {
	Kind       string                                       `json:"kind"`
	Group      string                                       `json:"group"`
	Datasource *Dashboardv2beta1AdhocVariableKindDatasource `json:"datasource,omitempty"`
	Spec       AdhocVariableSpec                            `json:"spec"`
}

// NewAdhocVariableKind creates a new AdhocVariableKind object.
func NewAdhocVariableKind() *AdhocVariableKind {
	return &AdhocVariableKind{
		Kind: "AdhocVariable",
		Spec: *NewAdhocVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdhocVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AdhocVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is null"))...)

		}
		delete(fields, "group")
	} else {
		errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is missing from input"))...)
	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &Dashboardv2beta1AdhocVariableKindDatasource{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = AdhocVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AdhocVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AdhocVariableKind` objects.
func (resource AdhocVariableKind) Equals(other AdhocVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Group != other.Group {
		return false
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AdhocVariableKind` fields for violations and returns them.
func (resource AdhocVariableKind) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Adhoc variable specification
type AdhocVariableSpec struct {
	Name             string                  `json:"name"`
	BaseFilters      []AdHocFilterWithLabels `json:"baseFilters"`
	Filters          []AdHocFilterWithLabels `json:"filters"`
	DefaultKeys      []MetricFindValue       `json:"defaultKeys"`
	Label            *string                 `json:"label,omitempty"`
	Hide             VariableHide            `json:"hide"`
	SkipUrlSync      bool                    `json:"skipUrlSync"`
	Description      *string                 `json:"description,omitempty"`
	AllowCustomValue bool                    `json:"allowCustomValue"`
}

// NewAdhocVariableSpec creates a new AdhocVariableSpec object.
func NewAdhocVariableSpec() *AdhocVariableSpec {
	return &AdhocVariableSpec{
		Name:             "",
		BaseFilters:      []AdHocFilterWithLabels{},
		Filters:          []AdHocFilterWithLabels{},
		DefaultKeys:      []MetricFindValue{},
		Hide:             VariableHideDontHide,
		SkipUrlSync:      false,
		AllowCustomValue: true,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdhocVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AdhocVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "baseFilters"
	if fields["baseFilters"] != nil {
		if string(fields["baseFilters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["baseFilters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AdHocFilterWithLabels

				result1 = AdHocFilterWithLabels{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("baseFilters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.BaseFilters = append(resource.BaseFilters, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("baseFilters", errors.New("required field is null"))...)

		}
		delete(fields, "baseFilters")
	} else {
		errs = append(errs, cog.MakeBuildErrors("baseFilters", errors.New("required field is missing from input"))...)
	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["filters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AdHocFilterWithLabels

				result1 = AdHocFilterWithLabels{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Filters = append(resource.Filters, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is null"))...)

		}
		delete(fields, "filters")
	} else {
		errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is missing from input"))...)
	}
	// Field "defaultKeys"
	if fields["defaultKeys"] != nil {
		if string(fields["defaultKeys"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["defaultKeys"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 MetricFindValue

				result1 = MetricFindValue{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("defaultKeys["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.DefaultKeys = append(resource.DefaultKeys, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("defaultKeys", errors.New("required field is null"))...)

		}
		delete(fields, "defaultKeys")
	} else {
		errs = append(errs, cog.MakeBuildErrors("defaultKeys", errors.New("required field is missing from input"))...)
	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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
	// Field "allowCustomValue"
	if fields["allowCustomValue"] != nil {
		if string(fields["allowCustomValue"]) != "null" {
			if err := json.Unmarshal(fields["allowCustomValue"], &resource.AllowCustomValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allowCustomValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("allowCustomValue", errors.New("required field is null"))...)

		}
		delete(fields, "allowCustomValue")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AdhocVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AdhocVariableSpec` objects.
func (resource AdhocVariableSpec) Equals(other AdhocVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}

	if len(resource.BaseFilters) != len(other.BaseFilters) {
		return false
	}

	for i1 := range resource.BaseFilters {
		if !resource.BaseFilters[i1].Equals(other.BaseFilters[i1]) {
			return false
		}
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if !resource.Filters[i1].Equals(other.Filters[i1]) {
			return false
		}
	}

	if len(resource.DefaultKeys) != len(other.DefaultKeys) {
		return false
	}

	for i1 := range resource.DefaultKeys {
		if !resource.DefaultKeys[i1].Equals(other.DefaultKeys[i1]) {
			return false
		}
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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
	if resource.AllowCustomValue != other.AllowCustomValue {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AdhocVariableSpec` fields for violations and returns them.
func (resource AdhocVariableSpec) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.BaseFilters {
		if err := resource.BaseFilters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("baseFilters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.Filters {
		if err := resource.Filters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.DefaultKeys {
		if err := resource.DefaultKeys[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("defaultKeys["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Define the AdHocFilterWithLabels type
type AdHocFilterWithLabels struct {
	Key         string   `json:"key"`
	Operator    string   `json:"operator"`
	Value       string   `json:"value"`
	Values      []string `json:"values,omitempty"`
	KeyLabel    *string  `json:"keyLabel,omitempty"`
	ValueLabels []string `json:"valueLabels,omitempty"`
	ForceEdit   *bool    `json:"forceEdit,omitempty"`
	Origin      *string  `json:"origin,omitempty"`
	// @deprecated
	Condition *string `json:"condition,omitempty"`
}

// NewAdHocFilterWithLabels creates a new AdHocFilterWithLabels object.
func NewAdHocFilterWithLabels() *AdHocFilterWithLabels {
	return &AdHocFilterWithLabels{
		Origin: (func(input string) *string { return &input })(FilterOrigin),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdHocFilterWithLabels` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AdHocFilterWithLabels) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "key"
	if fields["key"] != nil {
		if string(fields["key"]) != "null" {
			if err := json.Unmarshal(fields["key"], &resource.Key); err != nil {
				errs = append(errs, cog.MakeBuildErrors("key", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is null"))...)

		}
		delete(fields, "key")
	} else {
		errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is missing from input"))...)
	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {
			if err := json.Unmarshal(fields["operator"], &resource.Operator); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is null"))...)

		}
		delete(fields, "operator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}
	// Field "values"
	if fields["values"] != nil {
		if string(fields["values"]) != "null" {

			if err := json.Unmarshal(fields["values"], &resource.Values); err != nil {
				errs = append(errs, cog.MakeBuildErrors("values", err)...)
			}

		}
		delete(fields, "values")

	}
	// Field "keyLabel"
	if fields["keyLabel"] != nil {
		if string(fields["keyLabel"]) != "null" {
			if err := json.Unmarshal(fields["keyLabel"], &resource.KeyLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("keyLabel", err)...)
			}

		}
		delete(fields, "keyLabel")

	}
	// Field "valueLabels"
	if fields["valueLabels"] != nil {
		if string(fields["valueLabels"]) != "null" {

			if err := json.Unmarshal(fields["valueLabels"], &resource.ValueLabels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valueLabels", err)...)
			}

		}
		delete(fields, "valueLabels")

	}
	// Field "forceEdit"
	if fields["forceEdit"] != nil {
		if string(fields["forceEdit"]) != "null" {
			if err := json.Unmarshal(fields["forceEdit"], &resource.ForceEdit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("forceEdit", err)...)
			}

		}
		delete(fields, "forceEdit")

	}
	// Field "origin"
	if fields["origin"] != nil {
		if string(fields["origin"]) != "null" {
			if err := json.Unmarshal(fields["origin"], &resource.Origin); err != nil {
				errs = append(errs, cog.MakeBuildErrors("origin", err)...)
			}

		}
		delete(fields, "origin")

	}
	// Field "condition"
	if fields["condition"] != nil {
		if string(fields["condition"]) != "null" {
			if err := json.Unmarshal(fields["condition"], &resource.Condition); err != nil {
				errs = append(errs, cog.MakeBuildErrors("condition", err)...)
			}

		}
		delete(fields, "condition")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AdHocFilterWithLabels", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AdHocFilterWithLabels` objects.
func (resource AdHocFilterWithLabels) Equals(other AdHocFilterWithLabels) bool {
	if resource.Key != other.Key {
		return false
	}
	if resource.Operator != other.Operator {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	if len(resource.Values) != len(other.Values) {
		return false
	}

	for i1 := range resource.Values {
		if resource.Values[i1] != other.Values[i1] {
			return false
		}
	}
	if resource.KeyLabel == nil && other.KeyLabel != nil || resource.KeyLabel != nil && other.KeyLabel == nil {
		return false
	}

	if resource.KeyLabel != nil {
		if *resource.KeyLabel != *other.KeyLabel {
			return false
		}
	}

	if len(resource.ValueLabels) != len(other.ValueLabels) {
		return false
	}

	for i1 := range resource.ValueLabels {
		if resource.ValueLabels[i1] != other.ValueLabels[i1] {
			return false
		}
	}
	if resource.ForceEdit == nil && other.ForceEdit != nil || resource.ForceEdit != nil && other.ForceEdit == nil {
		return false
	}

	if resource.ForceEdit != nil {
		if *resource.ForceEdit != *other.ForceEdit {
			return false
		}
	}
	if resource.Origin == nil && other.Origin != nil || resource.Origin != nil && other.Origin == nil {
		return false
	}

	if resource.Origin != nil {
		if *resource.Origin != *other.Origin {
			return false
		}
	}
	if resource.Condition == nil && other.Condition != nil || resource.Condition != nil && other.Condition == nil {
		return false
	}

	if resource.Condition != nil {
		if *resource.Condition != *other.Condition {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AdHocFilterWithLabels` fields for violations and returns them.
func (resource AdHocFilterWithLabels) Validate() error {
	return nil
}

// Determine the origin of the adhoc variable filter
const FilterOrigin = "dashboard"

// Define the MetricFindValue type
type MetricFindValue struct {
	Text       string           `json:"text"`
	Value      *StringOrFloat64 `json:"value,omitempty"`
	Group      *string          `json:"group,omitempty"`
	Expandable *bool            `json:"expandable,omitempty"`
}

// NewMetricFindValue creates a new MetricFindValue object.
func NewMetricFindValue() *MetricFindValue {
	return &MetricFindValue{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricFindValue` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricFindValue) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {
			if err := json.Unmarshal(fields["text"], &resource.Text); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("text", errors.New("required field is null"))...)

		}
		delete(fields, "text")
	} else {
		errs = append(errs, cog.MakeBuildErrors("text", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {

			resource.Value = &StringOrFloat64{}
			if err := resource.Value.UnmarshalJSONStrict(fields["value"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}

		}
		delete(fields, "value")

	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}

		}
		delete(fields, "group")

	}
	// Field "expandable"
	if fields["expandable"] != nil {
		if string(fields["expandable"]) != "null" {
			if err := json.Unmarshal(fields["expandable"], &resource.Expandable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expandable", err)...)
			}

		}
		delete(fields, "expandable")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricFindValue", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricFindValue` objects.
func (resource MetricFindValue) Equals(other MetricFindValue) bool {
	if resource.Text != other.Text {
		return false
	}
	if resource.Value == nil && other.Value != nil || resource.Value != nil && other.Value == nil {
		return false
	}

	if resource.Value != nil {
		if !resource.Value.Equals(*other.Value) {
			return false
		}
	}
	if resource.Group == nil && other.Group != nil || resource.Group != nil && other.Group == nil {
		return false
	}

	if resource.Group != nil {
		if *resource.Group != *other.Group {
			return false
		}
	}
	if resource.Expandable == nil && other.Expandable != nil || resource.Expandable != nil && other.Expandable == nil {
		return false
	}

	if resource.Expandable != nil {
		if *resource.Expandable != *other.Expandable {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MetricFindValue` fields for violations and returns them.
func (resource MetricFindValue) Validate() error {
	var errs cog.BuildErrors
	if resource.Value != nil {
		if err := resource.Value.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("value", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type SwitchVariableKind struct {
	Kind string             `json:"kind"`
	Spec SwitchVariableSpec `json:"spec"`
}

// NewSwitchVariableKind creates a new SwitchVariableKind object.
func NewSwitchVariableKind() *SwitchVariableKind {
	return &SwitchVariableKind{
		Kind: "SwitchVariable",
		Spec: *NewSwitchVariableSpec(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SwitchVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SwitchVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {

			resource.Spec = SwitchVariableSpec{}
			if err := resource.Spec.UnmarshalJSONStrict(fields["spec"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SwitchVariableKind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SwitchVariableKind` objects.
func (resource SwitchVariableKind) Equals(other SwitchVariableKind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Spec.Equals(other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SwitchVariableKind` fields for violations and returns them.
func (resource SwitchVariableKind) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Spec.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("spec", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type SwitchVariableSpec struct {
	Name          string       `json:"name"`
	Current       string       `json:"current"`
	EnabledValue  string       `json:"enabledValue"`
	DisabledValue string       `json:"disabledValue"`
	Label         *string      `json:"label,omitempty"`
	Hide          VariableHide `json:"hide"`
	SkipUrlSync   bool         `json:"skipUrlSync"`
	Description   *string      `json:"description,omitempty"`
}

// NewSwitchVariableSpec creates a new SwitchVariableSpec object.
func NewSwitchVariableSpec() *SwitchVariableSpec {
	return &SwitchVariableSpec{
		Name:          "",
		Current:       "false",
		EnabledValue:  "true",
		DisabledValue: "false",
		Hide:          VariableHideDontHide,
		SkipUrlSync:   false,
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SwitchVariableSpec` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SwitchVariableSpec) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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

	}
	// Field "current"
	if fields["current"] != nil {
		if string(fields["current"]) != "null" {
			if err := json.Unmarshal(fields["current"], &resource.Current); err != nil {
				errs = append(errs, cog.MakeBuildErrors("current", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("current", errors.New("required field is null"))...)

		}
		delete(fields, "current")

	}
	// Field "enabledValue"
	if fields["enabledValue"] != nil {
		if string(fields["enabledValue"]) != "null" {
			if err := json.Unmarshal(fields["enabledValue"], &resource.EnabledValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("enabledValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("enabledValue", errors.New("required field is null"))...)

		}
		delete(fields, "enabledValue")

	}
	// Field "disabledValue"
	if fields["disabledValue"] != nil {
		if string(fields["disabledValue"]) != "null" {
			if err := json.Unmarshal(fields["disabledValue"], &resource.DisabledValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("disabledValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("disabledValue", errors.New("required field is null"))...)

		}
		delete(fields, "disabledValue")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hide", errors.New("required field is null"))...)

		}
		delete(fields, "hide")

	}
	// Field "skipUrlSync"
	if fields["skipUrlSync"] != nil {
		if string(fields["skipUrlSync"]) != "null" {
			if err := json.Unmarshal(fields["skipUrlSync"], &resource.SkipUrlSync); err != nil {
				errs = append(errs, cog.MakeBuildErrors("skipUrlSync", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("skipUrlSync", errors.New("required field is null"))...)

		}
		delete(fields, "skipUrlSync")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SwitchVariableSpec", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SwitchVariableSpec` objects.
func (resource SwitchVariableSpec) Equals(other SwitchVariableSpec) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Current != other.Current {
		return false
	}
	if resource.EnabledValue != other.EnabledValue {
		return false
	}
	if resource.DisabledValue != other.DisabledValue {
		return false
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Hide != other.Hide {
		return false
	}
	if resource.SkipUrlSync != other.SkipUrlSync {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `SwitchVariableSpec` fields for violations and returns them.
func (resource SwitchVariableSpec) Validate() error {
	return nil
}

// Annotation event field source. Defines how to obtain the value for an annotation event field.
// - "field": Find the value with a matching key (default)
// - "text": Write a constant string into the value
// - "skip": Do not include the field
type AnnotationEventFieldSource string

const (
	AnnotationEventFieldSourceField AnnotationEventFieldSource = "field"
	AnnotationEventFieldSourceText  AnnotationEventFieldSource = "text"
	AnnotationEventFieldSourceSkip  AnnotationEventFieldSource = "skip"
)

// --- Common types ---
type Kind struct {
	Kind     string `json:"kind"`
	Spec     any    `json:"spec"`
	Metadata any    `json:"metadata,omitempty"`
}

// NewKind creates a new Kind object.
func NewKind() *Kind {
	return &Kind{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Kind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Kind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {
			if err := json.Unmarshal(fields["spec"], &resource.Spec); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}
	// Field "metadata"
	if fields["metadata"] != nil {
		if string(fields["metadata"]) != "null" {
			if err := json.Unmarshal(fields["metadata"], &resource.Metadata); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metadata", err)...)
			}

		}
		delete(fields, "metadata")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Kind", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Kind` objects.
func (resource Kind) Equals(other Kind) bool {
	if resource.Kind != other.Kind {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Spec, other.Spec) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Metadata, other.Metadata) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Kind` fields for violations and returns them.
func (resource Kind) Validate() error {
	return nil
}

// Variable types
type VariableValue = StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle

// NewVariableValue creates a new VariableValue object.
func NewVariableValue() *VariableValue {
	return NewStringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle()
}

type VariableValueSingle = StringOrBoolOrFloat64OrCustomVariableValue

// NewVariableValueSingle creates a new VariableValueSingle object.
func NewVariableValueSingle() *VariableValueSingle {
	return NewStringOrBoolOrFloat64OrCustomVariableValue()
}

// Custom variable value
type CustomVariableValue struct {
	// The format name or function used in the expression
	Formatter *string `json:"formatter"`
}

// NewCustomVariableValue creates a new CustomVariableValue object.
func NewCustomVariableValue() *CustomVariableValue {
	return &CustomVariableValue{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomVariableValue` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CustomVariableValue) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "formatter"
	if fields["formatter"] != nil {
		if string(fields["formatter"]) != "null" {
			if err := json.Unmarshal(fields["formatter"], &resource.Formatter); err != nil {
				errs = append(errs, cog.MakeBuildErrors("formatter", err)...)
			}

		}
		delete(fields, "formatter")
	} else {
		errs = append(errs, cog.MakeBuildErrors("formatter", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CustomVariableValue", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CustomVariableValue` objects.
func (resource CustomVariableValue) Equals(other CustomVariableValue) bool {
	if resource.Formatter == nil && other.Formatter != nil || resource.Formatter != nil && other.Formatter == nil {
		return false
	}

	if resource.Formatter != nil {
		if *resource.Formatter != *other.Formatter {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CustomVariableValue` fields for violations and returns them.
func (resource CustomVariableValue) Validate() error {
	return nil
}

// Dashboard variable type
// `query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.
// `adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).
// `constant`: 	Define a hidden constant.
// `datasource`: Quickly change the data source for an entire dashboard.
// `interval`: Interval variables represent time spans.
// `textbox`: Display a free text input field with an optional default value.
// `custom`: Define the variable options manually using a comma-separated list.
// `system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables
type VariableType string

const (
	VariableTypeQuery      VariableType = "query"
	VariableTypeAdhoc      VariableType = "adhoc"
	VariableTypeGroupby    VariableType = "groupby"
	VariableTypeConstant   VariableType = "constant"
	VariableTypeDatasource VariableType = "datasource"
	VariableTypeInterval   VariableType = "interval"
	VariableTypeTextbox    VariableType = "textbox"
	VariableTypeCustom     VariableType = "custom"
	VariableTypeSystem     VariableType = "system"
	VariableTypeSnapshot   VariableType = "snapshot"
	VariableTypeSwitch     VariableType = "switch"
)

// Custom formatter variable
type CustomFormatterVariable struct {
	Name       string       `json:"name"`
	Type       VariableType `json:"type"`
	Multi      bool         `json:"multi"`
	IncludeAll bool         `json:"includeAll"`
}

// NewCustomFormatterVariable creates a new CustomFormatterVariable object.
func NewCustomFormatterVariable() *CustomFormatterVariable {
	return &CustomFormatterVariable{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomFormatterVariable` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CustomFormatterVariable) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "multi"
	if fields["multi"] != nil {
		if string(fields["multi"]) != "null" {
			if err := json.Unmarshal(fields["multi"], &resource.Multi); err != nil {
				errs = append(errs, cog.MakeBuildErrors("multi", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("multi", errors.New("required field is null"))...)

		}
		delete(fields, "multi")
	} else {
		errs = append(errs, cog.MakeBuildErrors("multi", errors.New("required field is missing from input"))...)
	}
	// Field "includeAll"
	if fields["includeAll"] != nil {
		if string(fields["includeAll"]) != "null" {
			if err := json.Unmarshal(fields["includeAll"], &resource.IncludeAll); err != nil {
				errs = append(errs, cog.MakeBuildErrors("includeAll", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("includeAll", errors.New("required field is null"))...)

		}
		delete(fields, "includeAll")
	} else {
		errs = append(errs, cog.MakeBuildErrors("includeAll", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CustomFormatterVariable", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CustomFormatterVariable` objects.
func (resource CustomFormatterVariable) Equals(other CustomFormatterVariable) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Multi != other.Multi {
		return false
	}
	if resource.IncludeAll != other.IncludeAll {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CustomFormatterVariable` fields for violations and returns them.
func (resource CustomFormatterVariable) Validate() error {
	return nil
}

// FIXME: should we introduce this? --- Variable value option
type VariableValueOption struct {
	Label string              `json:"label"`
	Value VariableValueSingle `json:"value"`
	Group *string             `json:"group,omitempty"`
}

// NewVariableValueOption creates a new VariableValueOption object.
func NewVariableValueOption() *VariableValueOption {
	return &VariableValueOption{
		Value: *NewVariableValueSingle(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VariableValueOption` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *VariableValueOption) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is null"))...)

		}
		delete(fields, "label")
	} else {
		errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {

			resource.Value = VariableValueSingle{}
			if err := resource.Value.UnmarshalJSONStrict(fields["value"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}

		}
		delete(fields, "group")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("VariableValueOption", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `VariableValueOption` objects.
func (resource VariableValueOption) Equals(other VariableValueOption) bool {
	if resource.Label != other.Label {
		return false
	}
	if !resource.Value.Equals(other.Value) {
		return false
	}
	if resource.Group == nil && other.Group != nil || resource.Group != nil && other.Group == nil {
		return false
	}

	if resource.Group != nil {
		if *resource.Group != *other.Group {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `VariableValueOption` fields for violations and returns them.
func (resource VariableValueOption) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Value.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("value", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Dashboardv2beta1DataQueryKindDatasource struct {
	Name *string `json:"name,omitempty"`
}

// NewDashboardv2beta1DataQueryKindDatasource creates a new Dashboardv2beta1DataQueryKindDatasource object.
func NewDashboardv2beta1DataQueryKindDatasource() *Dashboardv2beta1DataQueryKindDatasource {
	return &Dashboardv2beta1DataQueryKindDatasource{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1DataQueryKindDatasource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1DataQueryKindDatasource) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1DataQueryKindDatasource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1DataQueryKindDatasource` objects.
func (resource Dashboardv2beta1DataQueryKindDatasource) Equals(other Dashboardv2beta1DataQueryKindDatasource) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1DataQueryKindDatasource` fields for violations and returns them.
func (resource Dashboardv2beta1DataQueryKindDatasource) Validate() error {
	return nil
}

type Dashboardv2beta1FieldConfigSourceOverrides struct {
	// Describes config override rules created when interacting with Grafana.
	SystemRef  *string              `json:"__systemRef,omitempty"`
	Matcher    MatcherConfig        `json:"matcher"`
	Properties []DynamicConfigValue `json:"properties"`
}

// NewDashboardv2beta1FieldConfigSourceOverrides creates a new Dashboardv2beta1FieldConfigSourceOverrides object.
func NewDashboardv2beta1FieldConfigSourceOverrides() *Dashboardv2beta1FieldConfigSourceOverrides {
	return &Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher:    *NewMatcherConfig(),
		Properties: []DynamicConfigValue{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1FieldConfigSourceOverrides` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1FieldConfigSourceOverrides) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "__systemRef"
	if fields["__systemRef"] != nil {
		if string(fields["__systemRef"]) != "null" {
			if err := json.Unmarshal(fields["__systemRef"], &resource.SystemRef); err != nil {
				errs = append(errs, cog.MakeBuildErrors("__systemRef", err)...)
			}

		}
		delete(fields, "__systemRef")

	}
	// Field "matcher"
	if fields["matcher"] != nil {
		if string(fields["matcher"]) != "null" {

			resource.Matcher = MatcherConfig{}
			if err := resource.Matcher.UnmarshalJSONStrict(fields["matcher"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is null"))...)

		}
		delete(fields, "matcher")
	} else {
		errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is missing from input"))...)
	}
	// Field "properties"
	if fields["properties"] != nil {
		if string(fields["properties"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["properties"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 DynamicConfigValue

				result1 = DynamicConfigValue{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("properties["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Properties = append(resource.Properties, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("properties", errors.New("required field is null"))...)

		}
		delete(fields, "properties")
	} else {
		errs = append(errs, cog.MakeBuildErrors("properties", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1FieldConfigSourceOverrides", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1FieldConfigSourceOverrides` objects.
func (resource Dashboardv2beta1FieldConfigSourceOverrides) Equals(other Dashboardv2beta1FieldConfigSourceOverrides) bool {
	if resource.SystemRef == nil && other.SystemRef != nil || resource.SystemRef != nil && other.SystemRef == nil {
		return false
	}

	if resource.SystemRef != nil {
		if *resource.SystemRef != *other.SystemRef {
			return false
		}
	}
	if !resource.Matcher.Equals(other.Matcher) {
		return false
	}

	if len(resource.Properties) != len(other.Properties) {
		return false
	}

	for i1 := range resource.Properties {
		if !resource.Properties[i1].Equals(other.Properties[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1FieldConfigSourceOverrides` fields for violations and returns them.
func (resource Dashboardv2beta1FieldConfigSourceOverrides) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Matcher.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
	}

	for i1 := range resource.Properties {
		if err := resource.Properties[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("properties["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Dashboardv2beta1RangeMapOptions struct {
	// Min value of the range. It can be null which means -Infinity
	From *float64 `json:"from"`
	// Max value of the range. It can be null which means +Infinity
	To *float64 `json:"to"`
	// Config to apply when the value is within the range
	Result ValueMappingResult `json:"result"`
}

// NewDashboardv2beta1RangeMapOptions creates a new Dashboardv2beta1RangeMapOptions object.
func NewDashboardv2beta1RangeMapOptions() *Dashboardv2beta1RangeMapOptions {
	return &Dashboardv2beta1RangeMapOptions{
		Result: *NewValueMappingResult(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1RangeMapOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1RangeMapOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}
	// Field "result"
	if fields["result"] != nil {
		if string(fields["result"]) != "null" {

			resource.Result = ValueMappingResult{}
			if err := resource.Result.UnmarshalJSONStrict(fields["result"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("result", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("result", errors.New("required field is null"))...)

		}
		delete(fields, "result")
	} else {
		errs = append(errs, cog.MakeBuildErrors("result", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1RangeMapOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1RangeMapOptions` objects.
func (resource Dashboardv2beta1RangeMapOptions) Equals(other Dashboardv2beta1RangeMapOptions) bool {
	if resource.From == nil && other.From != nil || resource.From != nil && other.From == nil {
		return false
	}

	if resource.From != nil {
		if *resource.From != *other.From {
			return false
		}
	}
	if resource.To == nil && other.To != nil || resource.To != nil && other.To == nil {
		return false
	}

	if resource.To != nil {
		if *resource.To != *other.To {
			return false
		}
	}
	if !resource.Result.Equals(other.Result) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1RangeMapOptions` fields for violations and returns them.
func (resource Dashboardv2beta1RangeMapOptions) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Result.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("result", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Dashboardv2beta1RegexMapOptions struct {
	// Regular expression to match against
	Pattern string `json:"pattern"`
	// Config to apply when the value matches the regex
	Result ValueMappingResult `json:"result"`
}

// NewDashboardv2beta1RegexMapOptions creates a new Dashboardv2beta1RegexMapOptions object.
func NewDashboardv2beta1RegexMapOptions() *Dashboardv2beta1RegexMapOptions {
	return &Dashboardv2beta1RegexMapOptions{
		Result: *NewValueMappingResult(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1RegexMapOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1RegexMapOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pattern"
	if fields["pattern"] != nil {
		if string(fields["pattern"]) != "null" {
			if err := json.Unmarshal(fields["pattern"], &resource.Pattern); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pattern", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("pattern", errors.New("required field is null"))...)

		}
		delete(fields, "pattern")
	} else {
		errs = append(errs, cog.MakeBuildErrors("pattern", errors.New("required field is missing from input"))...)
	}
	// Field "result"
	if fields["result"] != nil {
		if string(fields["result"]) != "null" {

			resource.Result = ValueMappingResult{}
			if err := resource.Result.UnmarshalJSONStrict(fields["result"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("result", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("result", errors.New("required field is null"))...)

		}
		delete(fields, "result")
	} else {
		errs = append(errs, cog.MakeBuildErrors("result", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1RegexMapOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1RegexMapOptions` objects.
func (resource Dashboardv2beta1RegexMapOptions) Equals(other Dashboardv2beta1RegexMapOptions) bool {
	if resource.Pattern != other.Pattern {
		return false
	}
	if !resource.Result.Equals(other.Result) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1RegexMapOptions` fields for violations and returns them.
func (resource Dashboardv2beta1RegexMapOptions) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Result.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("result", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Dashboardv2beta1SpecialValueMapOptions struct {
	// Special value to match against
	Match SpecialValueMatch `json:"match"`
	// Config to apply when the value matches the special value
	Result ValueMappingResult `json:"result"`
}

// NewDashboardv2beta1SpecialValueMapOptions creates a new Dashboardv2beta1SpecialValueMapOptions object.
func NewDashboardv2beta1SpecialValueMapOptions() *Dashboardv2beta1SpecialValueMapOptions {
	return &Dashboardv2beta1SpecialValueMapOptions{
		Result: *NewValueMappingResult(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1SpecialValueMapOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1SpecialValueMapOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "match"
	if fields["match"] != nil {
		if string(fields["match"]) != "null" {
			if err := json.Unmarshal(fields["match"], &resource.Match); err != nil {
				errs = append(errs, cog.MakeBuildErrors("match", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("match", errors.New("required field is null"))...)

		}
		delete(fields, "match")
	} else {
		errs = append(errs, cog.MakeBuildErrors("match", errors.New("required field is missing from input"))...)
	}
	// Field "result"
	if fields["result"] != nil {
		if string(fields["result"]) != "null" {

			resource.Result = ValueMappingResult{}
			if err := resource.Result.UnmarshalJSONStrict(fields["result"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("result", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("result", errors.New("required field is null"))...)

		}
		delete(fields, "result")
	} else {
		errs = append(errs, cog.MakeBuildErrors("result", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1SpecialValueMapOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1SpecialValueMapOptions` objects.
func (resource Dashboardv2beta1SpecialValueMapOptions) Equals(other Dashboardv2beta1SpecialValueMapOptions) bool {
	if resource.Match != other.Match {
		return false
	}
	if !resource.Result.Equals(other.Result) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1SpecialValueMapOptions` fields for violations and returns them.
func (resource Dashboardv2beta1SpecialValueMapOptions) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Result.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("result", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Dashboardv2beta1ActionStyle struct {
	BackgroundColor *string `json:"backgroundColor,omitempty"`
}

// NewDashboardv2beta1ActionStyle creates a new Dashboardv2beta1ActionStyle object.
func NewDashboardv2beta1ActionStyle() *Dashboardv2beta1ActionStyle {
	return &Dashboardv2beta1ActionStyle{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1ActionStyle` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1ActionStyle) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "backgroundColor"
	if fields["backgroundColor"] != nil {
		if string(fields["backgroundColor"]) != "null" {
			if err := json.Unmarshal(fields["backgroundColor"], &resource.BackgroundColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("backgroundColor", err)...)
			}

		}
		delete(fields, "backgroundColor")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1ActionStyle", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1ActionStyle` objects.
func (resource Dashboardv2beta1ActionStyle) Equals(other Dashboardv2beta1ActionStyle) bool {
	if resource.BackgroundColor == nil && other.BackgroundColor != nil || resource.BackgroundColor != nil && other.BackgroundColor == nil {
		return false
	}

	if resource.BackgroundColor != nil {
		if *resource.BackgroundColor != *other.BackgroundColor {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1ActionStyle` fields for violations and returns them.
func (resource Dashboardv2beta1ActionStyle) Validate() error {
	return nil
}

type Dashboardv2beta1GroupByVariableKindDatasource struct {
	Name *string `json:"name,omitempty"`
}

// NewDashboardv2beta1GroupByVariableKindDatasource creates a new Dashboardv2beta1GroupByVariableKindDatasource object.
func NewDashboardv2beta1GroupByVariableKindDatasource() *Dashboardv2beta1GroupByVariableKindDatasource {
	return &Dashboardv2beta1GroupByVariableKindDatasource{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1GroupByVariableKindDatasource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1GroupByVariableKindDatasource) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1GroupByVariableKindDatasource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1GroupByVariableKindDatasource` objects.
func (resource Dashboardv2beta1GroupByVariableKindDatasource) Equals(other Dashboardv2beta1GroupByVariableKindDatasource) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1GroupByVariableKindDatasource` fields for violations and returns them.
func (resource Dashboardv2beta1GroupByVariableKindDatasource) Validate() error {
	return nil
}

type Dashboardv2beta1AdhocVariableKindDatasource struct {
	Name *string `json:"name,omitempty"`
}

// NewDashboardv2beta1AdhocVariableKindDatasource creates a new Dashboardv2beta1AdhocVariableKindDatasource object.
func NewDashboardv2beta1AdhocVariableKindDatasource() *Dashboardv2beta1AdhocVariableKindDatasource {
	return &Dashboardv2beta1AdhocVariableKindDatasource{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1AdhocVariableKindDatasource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dashboardv2beta1AdhocVariableKindDatasource) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dashboardv2beta1AdhocVariableKindDatasource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Dashboardv2beta1AdhocVariableKindDatasource` objects.
func (resource Dashboardv2beta1AdhocVariableKindDatasource) Equals(other Dashboardv2beta1AdhocVariableKindDatasource) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dashboardv2beta1AdhocVariableKindDatasource` fields for violations and returns them.
func (resource Dashboardv2beta1AdhocVariableKindDatasource) Validate() error {
	return nil
}

type RepeatOptionsDirection string

const (
	RepeatOptionsDirectionH RepeatOptionsDirection = "h"
	RepeatOptionsDirectionV RepeatOptionsDirection = "v"
)

type ConditionalRenderingGroupSpecVisibility string

const (
	ConditionalRenderingGroupSpecVisibilityShow ConditionalRenderingGroupSpecVisibility = "show"
	ConditionalRenderingGroupSpecVisibilityHide ConditionalRenderingGroupSpecVisibility = "hide"
)

type ConditionalRenderingGroupSpecCondition string

const (
	ConditionalRenderingGroupSpecConditionAnd ConditionalRenderingGroupSpecCondition = "and"
	ConditionalRenderingGroupSpecConditionOr  ConditionalRenderingGroupSpecCondition = "or"
)

type ConditionalRenderingVariableSpecOperator string

const (
	ConditionalRenderingVariableSpecOperatorEquals     ConditionalRenderingVariableSpecOperator = "equals"
	ConditionalRenderingVariableSpecOperatorNotEquals  ConditionalRenderingVariableSpecOperator = "notEquals"
	ConditionalRenderingVariableSpecOperatorMatches    ConditionalRenderingVariableSpecOperator = "matches"
	ConditionalRenderingVariableSpecOperatorNotMatches ConditionalRenderingVariableSpecOperator = "notMatches"
)

type AutoGridLayoutSpecColumnWidthMode string

const (
	AutoGridLayoutSpecColumnWidthModeNarrow   AutoGridLayoutSpecColumnWidthMode = "narrow"
	AutoGridLayoutSpecColumnWidthModeStandard AutoGridLayoutSpecColumnWidthMode = "standard"
	AutoGridLayoutSpecColumnWidthModeWide     AutoGridLayoutSpecColumnWidthMode = "wide"
	AutoGridLayoutSpecColumnWidthModeCustom   AutoGridLayoutSpecColumnWidthMode = "custom"
)

type AutoGridLayoutSpecRowHeightMode string

const (
	AutoGridLayoutSpecRowHeightModeShort    AutoGridLayoutSpecRowHeightMode = "short"
	AutoGridLayoutSpecRowHeightModeStandard AutoGridLayoutSpecRowHeightMode = "standard"
	AutoGridLayoutSpecRowHeightModeTall     AutoGridLayoutSpecRowHeightMode = "tall"
	AutoGridLayoutSpecRowHeightModeCustom   AutoGridLayoutSpecRowHeightMode = "custom"
)

type TimeSettingsSpecWeekStart string

const (
	TimeSettingsSpecWeekStartSaturday TimeSettingsSpecWeekStart = "saturday"
	TimeSettingsSpecWeekStartMonday   TimeSettingsSpecWeekStart = "monday"
	TimeSettingsSpecWeekStartSunday   TimeSettingsSpecWeekStart = "sunday"
)

type QueryVariableSpecStaticOptionsOrder string

const (
	QueryVariableSpecStaticOptionsOrderBefore QueryVariableSpecStaticOptionsOrder = "before"
	QueryVariableSpecStaticOptionsOrderAfter  QueryVariableSpecStaticOptionsOrder = "after"
	QueryVariableSpecStaticOptionsOrderSorted QueryVariableSpecStaticOptionsOrder = "sorted"
)

type CustomVariableSpecValuesFormat string

const (
	CustomVariableSpecValuesFormatCsv  CustomVariableSpecValuesFormat = "csv"
	CustomVariableSpecValuesFormatJson CustomVariableSpecValuesFormat = "json"
)

type GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind struct {
	GridLayoutKind     *GridLayoutKind     `json:"GridLayoutKind,omitempty"`
	RowsLayoutKind     *RowsLayoutKind     `json:"RowsLayoutKind,omitempty"`
	AutoGridLayoutKind *AutoGridLayoutKind `json:"AutoGridLayoutKind,omitempty"`
	TabsLayoutKind     *TabsLayoutKind     `json:"TabsLayoutKind,omitempty"`
}

// NewGridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind creates a new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind object.
func NewGridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind() *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind {
	return &GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` as JSON.
func (resource GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) MarshalJSON() ([]byte, error) {
	if resource.GridLayoutKind != nil {
		return json.Marshal(resource.GridLayoutKind)
	}
	if resource.RowsLayoutKind != nil {
		return json.Marshal(resource.RowsLayoutKind)
	}
	if resource.AutoGridLayoutKind != nil {
		return json.Marshal(resource.AutoGridLayoutKind)
	}
	if resource.TabsLayoutKind != nil {
		return json.Marshal(resource.TabsLayoutKind)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` from JSON.
func (resource *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return nil
	}

	switch discriminator {
	case "AutoGridLayout":
		var autoGridLayoutKind AutoGridLayoutKind
		if err := json.Unmarshal(raw, &autoGridLayoutKind); err != nil {
			return err
		}

		resource.AutoGridLayoutKind = &autoGridLayoutKind
		return nil
	case "GridLayout":
		var gridLayoutKind GridLayoutKind
		if err := json.Unmarshal(raw, &gridLayoutKind); err != nil {
			return err
		}

		resource.GridLayoutKind = &gridLayoutKind
		return nil
	case "RowsLayout":
		var rowsLayoutKind RowsLayoutKind
		if err := json.Unmarshal(raw, &rowsLayoutKind); err != nil {
			return err
		}

		resource.RowsLayoutKind = &rowsLayoutKind
		return nil
	case "TabsLayout":
		var tabsLayoutKind TabsLayoutKind
		if err := json.Unmarshal(raw, &tabsLayoutKind); err != nil {
			return err
		}

		resource.TabsLayoutKind = &tabsLayoutKind
		return nil
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return fmt.Errorf("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "AutoGridLayout":
		autoGridLayoutKind := &AutoGridLayoutKind{}
		if err := autoGridLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.AutoGridLayoutKind = autoGridLayoutKind
		return nil
	case "GridLayout":
		gridLayoutKind := &GridLayoutKind{}
		if err := gridLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.GridLayoutKind = gridLayoutKind
		return nil
	case "RowsLayout":
		rowsLayoutKind := &RowsLayoutKind{}
		if err := rowsLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RowsLayoutKind = rowsLayoutKind
		return nil
	case "TabsLayout":
		tabsLayoutKind := &TabsLayoutKind{}
		if err := tabsLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TabsLayoutKind = tabsLayoutKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// Equals tests the equality of two `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` objects.
func (resource GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) Equals(other GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) bool {
	if resource.GridLayoutKind == nil && other.GridLayoutKind != nil || resource.GridLayoutKind != nil && other.GridLayoutKind == nil {
		return false
	}

	if resource.GridLayoutKind != nil {
		if !resource.GridLayoutKind.Equals(*other.GridLayoutKind) {
			return false
		}
	}
	if resource.RowsLayoutKind == nil && other.RowsLayoutKind != nil || resource.RowsLayoutKind != nil && other.RowsLayoutKind == nil {
		return false
	}

	if resource.RowsLayoutKind != nil {
		if !resource.RowsLayoutKind.Equals(*other.RowsLayoutKind) {
			return false
		}
	}
	if resource.AutoGridLayoutKind == nil && other.AutoGridLayoutKind != nil || resource.AutoGridLayoutKind != nil && other.AutoGridLayoutKind == nil {
		return false
	}

	if resource.AutoGridLayoutKind != nil {
		if !resource.AutoGridLayoutKind.Equals(*other.AutoGridLayoutKind) {
			return false
		}
	}
	if resource.TabsLayoutKind == nil && other.TabsLayoutKind != nil || resource.TabsLayoutKind != nil && other.TabsLayoutKind == nil {
		return false
	}

	if resource.TabsLayoutKind != nil {
		if !resource.TabsLayoutKind.Equals(*other.TabsLayoutKind) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` fields for violations and returns them.
func (resource GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) Validate() error {
	var errs cog.BuildErrors
	if resource.GridLayoutKind != nil {
		if err := resource.GridLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("GridLayoutKind", err)...)
		}
	}
	if resource.RowsLayoutKind != nil {
		if err := resource.RowsLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RowsLayoutKind", err)...)
		}
	}
	if resource.AutoGridLayoutKind != nil {
		if err := resource.AutoGridLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("AutoGridLayoutKind", err)...)
		}
	}
	if resource.TabsLayoutKind != nil {
		if err := resource.TabsLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TabsLayoutKind", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type PanelKindOrLibraryPanelKind struct {
	PanelKind        *PanelKind        `json:"PanelKind,omitempty"`
	LibraryPanelKind *LibraryPanelKind `json:"LibraryPanelKind,omitempty"`
}

// NewPanelKindOrLibraryPanelKind creates a new PanelKindOrLibraryPanelKind object.
func NewPanelKindOrLibraryPanelKind() *PanelKindOrLibraryPanelKind {
	return &PanelKindOrLibraryPanelKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `PanelKindOrLibraryPanelKind` as JSON.
func (resource PanelKindOrLibraryPanelKind) MarshalJSON() ([]byte, error) {
	if resource.PanelKind != nil {
		return json.Marshal(resource.PanelKind)
	}
	if resource.LibraryPanelKind != nil {
		return json.Marshal(resource.LibraryPanelKind)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `PanelKindOrLibraryPanelKind` from JSON.
func (resource *PanelKindOrLibraryPanelKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return nil
	}

	switch discriminator {
	case "LibraryPanel":
		var libraryPanelKind LibraryPanelKind
		if err := json.Unmarshal(raw, &libraryPanelKind); err != nil {
			return err
		}

		resource.LibraryPanelKind = &libraryPanelKind
		return nil
	case "Panel":
		var panelKind PanelKind
		if err := json.Unmarshal(raw, &panelKind); err != nil {
			return err
		}

		resource.PanelKind = &panelKind
		return nil
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelKindOrLibraryPanelKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PanelKindOrLibraryPanelKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return fmt.Errorf("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "LibraryPanel":
		libraryPanelKind := &LibraryPanelKind{}
		if err := libraryPanelKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.LibraryPanelKind = libraryPanelKind
		return nil
	case "Panel":
		panelKind := &PanelKind{}
		if err := panelKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.PanelKind = panelKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// Equals tests the equality of two `PanelKindOrLibraryPanelKind` objects.
func (resource PanelKindOrLibraryPanelKind) Equals(other PanelKindOrLibraryPanelKind) bool {
	if resource.PanelKind == nil && other.PanelKind != nil || resource.PanelKind != nil && other.PanelKind == nil {
		return false
	}

	if resource.PanelKind != nil {
		if !resource.PanelKind.Equals(*other.PanelKind) {
			return false
		}
	}
	if resource.LibraryPanelKind == nil && other.LibraryPanelKind != nil || resource.LibraryPanelKind != nil && other.LibraryPanelKind == nil {
		return false
	}

	if resource.LibraryPanelKind != nil {
		if !resource.LibraryPanelKind.Equals(*other.LibraryPanelKind) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PanelKindOrLibraryPanelKind` fields for violations and returns them.
func (resource PanelKindOrLibraryPanelKind) Validate() error {
	var errs cog.BuildErrors
	if resource.PanelKind != nil {
		if err := resource.PanelKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("PanelKind", err)...)
		}
	}
	if resource.LibraryPanelKind != nil {
		if err := resource.LibraryPanelKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("LibraryPanelKind", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ValueMapOrRangeMapOrRegexMapOrSpecialValueMap struct {
	ValueMap        *ValueMap        `json:"ValueMap,omitempty"`
	RangeMap        *RangeMap        `json:"RangeMap,omitempty"`
	RegexMap        *RegexMap        `json:"RegexMap,omitempty"`
	SpecialValueMap *SpecialValueMap `json:"SpecialValueMap,omitempty"`
}

// NewValueMapOrRangeMapOrRegexMapOrSpecialValueMap creates a new ValueMapOrRangeMapOrRegexMapOrSpecialValueMap object.
func NewValueMapOrRangeMapOrRegexMapOrSpecialValueMap() *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap {
	return &ValueMapOrRangeMapOrRegexMapOrSpecialValueMap{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` as JSON.
func (resource ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) MarshalJSON() ([]byte, error) {
	if resource.ValueMap != nil {
		return json.Marshal(resource.ValueMap)
	}
	if resource.RangeMap != nil {
		return json.Marshal(resource.RangeMap)
	}
	if resource.RegexMap != nil {
		return json.Marshal(resource.RegexMap)
	}
	if resource.SpecialValueMap != nil {
		return json.Marshal(resource.SpecialValueMap)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` from JSON.
func (resource *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return nil
	}

	switch discriminator {
	case "range":
		var rangeMap RangeMap
		if err := json.Unmarshal(raw, &rangeMap); err != nil {
			return err
		}

		resource.RangeMap = &rangeMap
		return nil
	case "regex":
		var regexMap RegexMap
		if err := json.Unmarshal(raw, &regexMap); err != nil {
			return err
		}

		resource.RegexMap = &regexMap
		return nil
	case "special":
		var specialValueMap SpecialValueMap
		if err := json.Unmarshal(raw, &specialValueMap); err != nil {
			return err
		}

		resource.SpecialValueMap = &specialValueMap
		return nil
	case "value":
		var valueMap ValueMap
		if err := json.Unmarshal(raw, &valueMap); err != nil {
			return err
		}

		resource.ValueMap = &valueMap
		return nil
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "range":
		rangeMap := &RangeMap{}
		if err := rangeMap.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RangeMap = rangeMap
		return nil
	case "regex":
		regexMap := &RegexMap{}
		if err := regexMap.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RegexMap = regexMap
		return nil
	case "special":
		specialValueMap := &SpecialValueMap{}
		if err := specialValueMap.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.SpecialValueMap = specialValueMap
		return nil
	case "value":
		valueMap := &ValueMap{}
		if err := valueMap.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ValueMap = valueMap
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` objects.
func (resource ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) Equals(other ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) bool {
	if resource.ValueMap == nil && other.ValueMap != nil || resource.ValueMap != nil && other.ValueMap == nil {
		return false
	}

	if resource.ValueMap != nil {
		if !resource.ValueMap.Equals(*other.ValueMap) {
			return false
		}
	}
	if resource.RangeMap == nil && other.RangeMap != nil || resource.RangeMap != nil && other.RangeMap == nil {
		return false
	}

	if resource.RangeMap != nil {
		if !resource.RangeMap.Equals(*other.RangeMap) {
			return false
		}
	}
	if resource.RegexMap == nil && other.RegexMap != nil || resource.RegexMap != nil && other.RegexMap == nil {
		return false
	}

	if resource.RegexMap != nil {
		if !resource.RegexMap.Equals(*other.RegexMap) {
			return false
		}
	}
	if resource.SpecialValueMap == nil && other.SpecialValueMap != nil || resource.SpecialValueMap != nil && other.SpecialValueMap == nil {
		return false
	}

	if resource.SpecialValueMap != nil {
		if !resource.SpecialValueMap.Equals(*other.SpecialValueMap) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ValueMapOrRangeMapOrRegexMapOrSpecialValueMap` fields for violations and returns them.
func (resource ValueMapOrRangeMapOrRegexMapOrSpecialValueMap) Validate() error {
	var errs cog.BuildErrors
	if resource.ValueMap != nil {
		if err := resource.ValueMap.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ValueMap", err)...)
		}
	}
	if resource.RangeMap != nil {
		if err := resource.RangeMap.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RangeMap", err)...)
		}
	}
	if resource.RegexMap != nil {
		if err := resource.RegexMap.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RegexMap", err)...)
		}
	}
	if resource.SpecialValueMap != nil {
		if err := resource.SpecialValueMap.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("SpecialValueMap", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind struct {
	GridLayoutKind     *GridLayoutKind     `json:"GridLayoutKind,omitempty"`
	AutoGridLayoutKind *AutoGridLayoutKind `json:"AutoGridLayoutKind,omitempty"`
	TabsLayoutKind     *TabsLayoutKind     `json:"TabsLayoutKind,omitempty"`
	RowsLayoutKind     *RowsLayoutKind     `json:"RowsLayoutKind,omitempty"`
}

// NewGridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind creates a new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind object.
func NewGridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind() *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind {
	return &GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` as JSON.
func (resource GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) MarshalJSON() ([]byte, error) {
	if resource.GridLayoutKind != nil {
		return json.Marshal(resource.GridLayoutKind)
	}
	if resource.AutoGridLayoutKind != nil {
		return json.Marshal(resource.AutoGridLayoutKind)
	}
	if resource.TabsLayoutKind != nil {
		return json.Marshal(resource.TabsLayoutKind)
	}
	if resource.RowsLayoutKind != nil {
		return json.Marshal(resource.RowsLayoutKind)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` from JSON.
func (resource *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return nil
	}

	switch discriminator {
	case "AutoGridLayout":
		var autoGridLayoutKind AutoGridLayoutKind
		if err := json.Unmarshal(raw, &autoGridLayoutKind); err != nil {
			return err
		}

		resource.AutoGridLayoutKind = &autoGridLayoutKind
		return nil
	case "GridLayout":
		var gridLayoutKind GridLayoutKind
		if err := json.Unmarshal(raw, &gridLayoutKind); err != nil {
			return err
		}

		resource.GridLayoutKind = &gridLayoutKind
		return nil
	case "RowsLayout":
		var rowsLayoutKind RowsLayoutKind
		if err := json.Unmarshal(raw, &rowsLayoutKind); err != nil {
			return err
		}

		resource.RowsLayoutKind = &rowsLayoutKind
		return nil
	case "TabsLayout":
		var tabsLayoutKind TabsLayoutKind
		if err := json.Unmarshal(raw, &tabsLayoutKind); err != nil {
			return err
		}

		resource.TabsLayoutKind = &tabsLayoutKind
		return nil
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return fmt.Errorf("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "AutoGridLayout":
		autoGridLayoutKind := &AutoGridLayoutKind{}
		if err := autoGridLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.AutoGridLayoutKind = autoGridLayoutKind
		return nil
	case "GridLayout":
		gridLayoutKind := &GridLayoutKind{}
		if err := gridLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.GridLayoutKind = gridLayoutKind
		return nil
	case "RowsLayout":
		rowsLayoutKind := &RowsLayoutKind{}
		if err := rowsLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RowsLayoutKind = rowsLayoutKind
		return nil
	case "TabsLayout":
		tabsLayoutKind := &TabsLayoutKind{}
		if err := tabsLayoutKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TabsLayoutKind = tabsLayoutKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// Equals tests the equality of two `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` objects.
func (resource GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) Equals(other GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) bool {
	if resource.GridLayoutKind == nil && other.GridLayoutKind != nil || resource.GridLayoutKind != nil && other.GridLayoutKind == nil {
		return false
	}

	if resource.GridLayoutKind != nil {
		if !resource.GridLayoutKind.Equals(*other.GridLayoutKind) {
			return false
		}
	}
	if resource.AutoGridLayoutKind == nil && other.AutoGridLayoutKind != nil || resource.AutoGridLayoutKind != nil && other.AutoGridLayoutKind == nil {
		return false
	}

	if resource.AutoGridLayoutKind != nil {
		if !resource.AutoGridLayoutKind.Equals(*other.AutoGridLayoutKind) {
			return false
		}
	}
	if resource.TabsLayoutKind == nil && other.TabsLayoutKind != nil || resource.TabsLayoutKind != nil && other.TabsLayoutKind == nil {
		return false
	}

	if resource.TabsLayoutKind != nil {
		if !resource.TabsLayoutKind.Equals(*other.TabsLayoutKind) {
			return false
		}
	}
	if resource.RowsLayoutKind == nil && other.RowsLayoutKind != nil || resource.RowsLayoutKind != nil && other.RowsLayoutKind == nil {
		return false
	}

	if resource.RowsLayoutKind != nil {
		if !resource.RowsLayoutKind.Equals(*other.RowsLayoutKind) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` fields for violations and returns them.
func (resource GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) Validate() error {
	var errs cog.BuildErrors
	if resource.GridLayoutKind != nil {
		if err := resource.GridLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("GridLayoutKind", err)...)
		}
	}
	if resource.AutoGridLayoutKind != nil {
		if err := resource.AutoGridLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("AutoGridLayoutKind", err)...)
		}
	}
	if resource.TabsLayoutKind != nil {
		if err := resource.TabsLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TabsLayoutKind", err)...)
		}
	}
	if resource.RowsLayoutKind != nil {
		if err := resource.RowsLayoutKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RowsLayoutKind", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind struct {
	ConditionalRenderingVariableKind      *ConditionalRenderingVariableKind      `json:"ConditionalRenderingVariableKind,omitempty"`
	ConditionalRenderingDataKind          *ConditionalRenderingDataKind          `json:"ConditionalRenderingDataKind,omitempty"`
	ConditionalRenderingTimeRangeSizeKind *ConditionalRenderingTimeRangeSizeKind `json:"ConditionalRenderingTimeRangeSizeKind,omitempty"`
}

// NewConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind creates a new ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind object.
func NewConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind() *ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind {
	return &ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind` as JSON.
func (resource ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) MarshalJSON() ([]byte, error) {
	if resource.ConditionalRenderingVariableKind != nil {
		return json.Marshal(resource.ConditionalRenderingVariableKind)
	}
	if resource.ConditionalRenderingDataKind != nil {
		return json.Marshal(resource.ConditionalRenderingDataKind)
	}
	if resource.ConditionalRenderingTimeRangeSizeKind != nil {
		return json.Marshal(resource.ConditionalRenderingTimeRangeSizeKind)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind` from JSON.
func (resource *ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return nil
	}

	switch discriminator {
	case "ConditionalRenderingData":
		var conditionalRenderingDataKind ConditionalRenderingDataKind
		if err := json.Unmarshal(raw, &conditionalRenderingDataKind); err != nil {
			return err
		}

		resource.ConditionalRenderingDataKind = &conditionalRenderingDataKind
		return nil
	case "ConditionalRenderingTimeRangeSize":
		var conditionalRenderingTimeRangeSizeKind ConditionalRenderingTimeRangeSizeKind
		if err := json.Unmarshal(raw, &conditionalRenderingTimeRangeSizeKind); err != nil {
			return err
		}

		resource.ConditionalRenderingTimeRangeSizeKind = &conditionalRenderingTimeRangeSizeKind
		return nil
	case "ConditionalRenderingVariable":
		var conditionalRenderingVariableKind ConditionalRenderingVariableKind
		if err := json.Unmarshal(raw, &conditionalRenderingVariableKind); err != nil {
			return err
		}

		resource.ConditionalRenderingVariableKind = &conditionalRenderingVariableKind
		return nil
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return fmt.Errorf("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "ConditionalRenderingData":
		conditionalRenderingDataKind := &ConditionalRenderingDataKind{}
		if err := conditionalRenderingDataKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ConditionalRenderingDataKind = conditionalRenderingDataKind
		return nil
	case "ConditionalRenderingTimeRangeSize":
		conditionalRenderingTimeRangeSizeKind := &ConditionalRenderingTimeRangeSizeKind{}
		if err := conditionalRenderingTimeRangeSizeKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ConditionalRenderingTimeRangeSizeKind = conditionalRenderingTimeRangeSizeKind
		return nil
	case "ConditionalRenderingVariable":
		conditionalRenderingVariableKind := &ConditionalRenderingVariableKind{}
		if err := conditionalRenderingVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ConditionalRenderingVariableKind = conditionalRenderingVariableKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// Equals tests the equality of two `ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind` objects.
func (resource ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) Equals(other ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) bool {
	if resource.ConditionalRenderingVariableKind == nil && other.ConditionalRenderingVariableKind != nil || resource.ConditionalRenderingVariableKind != nil && other.ConditionalRenderingVariableKind == nil {
		return false
	}

	if resource.ConditionalRenderingVariableKind != nil {
		if !resource.ConditionalRenderingVariableKind.Equals(*other.ConditionalRenderingVariableKind) {
			return false
		}
	}
	if resource.ConditionalRenderingDataKind == nil && other.ConditionalRenderingDataKind != nil || resource.ConditionalRenderingDataKind != nil && other.ConditionalRenderingDataKind == nil {
		return false
	}

	if resource.ConditionalRenderingDataKind != nil {
		if !resource.ConditionalRenderingDataKind.Equals(*other.ConditionalRenderingDataKind) {
			return false
		}
	}
	if resource.ConditionalRenderingTimeRangeSizeKind == nil && other.ConditionalRenderingTimeRangeSizeKind != nil || resource.ConditionalRenderingTimeRangeSizeKind != nil && other.ConditionalRenderingTimeRangeSizeKind == nil {
		return false
	}

	if resource.ConditionalRenderingTimeRangeSizeKind != nil {
		if !resource.ConditionalRenderingTimeRangeSizeKind.Equals(*other.ConditionalRenderingTimeRangeSizeKind) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind` fields for violations and returns them.
func (resource ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) Validate() error {
	var errs cog.BuildErrors
	if resource.ConditionalRenderingVariableKind != nil {
		if err := resource.ConditionalRenderingVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingVariableKind", err)...)
		}
	}
	if resource.ConditionalRenderingDataKind != nil {
		if err := resource.ConditionalRenderingDataKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingDataKind", err)...)
		}
	}
	if resource.ConditionalRenderingTimeRangeSizeKind != nil {
		if err := resource.ConditionalRenderingTimeRangeSizeKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ConditionalRenderingTimeRangeSizeKind", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind struct {
	QueryVariableKind      *QueryVariableKind      `json:"QueryVariableKind,omitempty"`
	TextVariableKind       *TextVariableKind       `json:"TextVariableKind,omitempty"`
	ConstantVariableKind   *ConstantVariableKind   `json:"ConstantVariableKind,omitempty"`
	DatasourceVariableKind *DatasourceVariableKind `json:"DatasourceVariableKind,omitempty"`
	IntervalVariableKind   *IntervalVariableKind   `json:"IntervalVariableKind,omitempty"`
	CustomVariableKind     *CustomVariableKind     `json:"CustomVariableKind,omitempty"`
	GroupByVariableKind    *GroupByVariableKind    `json:"GroupByVariableKind,omitempty"`
	AdhocVariableKind      *AdhocVariableKind      `json:"AdhocVariableKind,omitempty"`
	SwitchVariableKind     *SwitchVariableKind     `json:"SwitchVariableKind,omitempty"`
}

// NewQueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind creates a new QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind object.
func NewQueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind() *QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind {
	return &QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind` as JSON.
func (resource QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind) MarshalJSON() ([]byte, error) {
	if resource.QueryVariableKind != nil {
		return json.Marshal(resource.QueryVariableKind)
	}
	if resource.TextVariableKind != nil {
		return json.Marshal(resource.TextVariableKind)
	}
	if resource.ConstantVariableKind != nil {
		return json.Marshal(resource.ConstantVariableKind)
	}
	if resource.DatasourceVariableKind != nil {
		return json.Marshal(resource.DatasourceVariableKind)
	}
	if resource.IntervalVariableKind != nil {
		return json.Marshal(resource.IntervalVariableKind)
	}
	if resource.CustomVariableKind != nil {
		return json.Marshal(resource.CustomVariableKind)
	}
	if resource.GroupByVariableKind != nil {
		return json.Marshal(resource.GroupByVariableKind)
	}
	if resource.AdhocVariableKind != nil {
		return json.Marshal(resource.AdhocVariableKind)
	}
	if resource.SwitchVariableKind != nil {
		return json.Marshal(resource.SwitchVariableKind)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind` from JSON.
func (resource *QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return nil
	}

	switch discriminator {
	case "AdhocVariable":
		var adhocVariableKind AdhocVariableKind
		if err := json.Unmarshal(raw, &adhocVariableKind); err != nil {
			return err
		}

		resource.AdhocVariableKind = &adhocVariableKind
		return nil
	case "ConstantVariable":
		var constantVariableKind ConstantVariableKind
		if err := json.Unmarshal(raw, &constantVariableKind); err != nil {
			return err
		}

		resource.ConstantVariableKind = &constantVariableKind
		return nil
	case "CustomVariable":
		var customVariableKind CustomVariableKind
		if err := json.Unmarshal(raw, &customVariableKind); err != nil {
			return err
		}

		resource.CustomVariableKind = &customVariableKind
		return nil
	case "DatasourceVariable":
		var datasourceVariableKind DatasourceVariableKind
		if err := json.Unmarshal(raw, &datasourceVariableKind); err != nil {
			return err
		}

		resource.DatasourceVariableKind = &datasourceVariableKind
		return nil
	case "GroupByVariable":
		var groupByVariableKind GroupByVariableKind
		if err := json.Unmarshal(raw, &groupByVariableKind); err != nil {
			return err
		}

		resource.GroupByVariableKind = &groupByVariableKind
		return nil
	case "IntervalVariable":
		var intervalVariableKind IntervalVariableKind
		if err := json.Unmarshal(raw, &intervalVariableKind); err != nil {
			return err
		}

		resource.IntervalVariableKind = &intervalVariableKind
		return nil
	case "QueryVariable":
		var queryVariableKind QueryVariableKind
		if err := json.Unmarshal(raw, &queryVariableKind); err != nil {
			return err
		}

		resource.QueryVariableKind = &queryVariableKind
		return nil
	case "SwitchVariable":
		var switchVariableKind SwitchVariableKind
		if err := json.Unmarshal(raw, &switchVariableKind); err != nil {
			return err
		}

		resource.SwitchVariableKind = &switchVariableKind
		return nil
	case "TextVariable":
		var textVariableKind TextVariableKind
		if err := json.Unmarshal(raw, &textVariableKind); err != nil {
			return err
		}

		resource.TextVariableKind = &textVariableKind
		return nil
	}

	return nil
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return fmt.Errorf("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "AdhocVariable":
		adhocVariableKind := &AdhocVariableKind{}
		if err := adhocVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.AdhocVariableKind = adhocVariableKind
		return nil
	case "ConstantVariable":
		constantVariableKind := &ConstantVariableKind{}
		if err := constantVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ConstantVariableKind = constantVariableKind
		return nil
	case "CustomVariable":
		customVariableKind := &CustomVariableKind{}
		if err := customVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CustomVariableKind = customVariableKind
		return nil
	case "DatasourceVariable":
		datasourceVariableKind := &DatasourceVariableKind{}
		if err := datasourceVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.DatasourceVariableKind = datasourceVariableKind
		return nil
	case "GroupByVariable":
		groupByVariableKind := &GroupByVariableKind{}
		if err := groupByVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.GroupByVariableKind = groupByVariableKind
		return nil
	case "IntervalVariable":
		intervalVariableKind := &IntervalVariableKind{}
		if err := intervalVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.IntervalVariableKind = intervalVariableKind
		return nil
	case "QueryVariable":
		queryVariableKind := &QueryVariableKind{}
		if err := queryVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryVariableKind = queryVariableKind
		return nil
	case "SwitchVariable":
		switchVariableKind := &SwitchVariableKind{}
		if err := switchVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.SwitchVariableKind = switchVariableKind
		return nil
	case "TextVariable":
		textVariableKind := &TextVariableKind{}
		if err := textVariableKind.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TextVariableKind = textVariableKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// Equals tests the equality of two `QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind` objects.
func (resource QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind) Equals(other QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind) bool {
	if resource.QueryVariableKind == nil && other.QueryVariableKind != nil || resource.QueryVariableKind != nil && other.QueryVariableKind == nil {
		return false
	}

	if resource.QueryVariableKind != nil {
		if !resource.QueryVariableKind.Equals(*other.QueryVariableKind) {
			return false
		}
	}
	if resource.TextVariableKind == nil && other.TextVariableKind != nil || resource.TextVariableKind != nil && other.TextVariableKind == nil {
		return false
	}

	if resource.TextVariableKind != nil {
		if !resource.TextVariableKind.Equals(*other.TextVariableKind) {
			return false
		}
	}
	if resource.ConstantVariableKind == nil && other.ConstantVariableKind != nil || resource.ConstantVariableKind != nil && other.ConstantVariableKind == nil {
		return false
	}

	if resource.ConstantVariableKind != nil {
		if !resource.ConstantVariableKind.Equals(*other.ConstantVariableKind) {
			return false
		}
	}
	if resource.DatasourceVariableKind == nil && other.DatasourceVariableKind != nil || resource.DatasourceVariableKind != nil && other.DatasourceVariableKind == nil {
		return false
	}

	if resource.DatasourceVariableKind != nil {
		if !resource.DatasourceVariableKind.Equals(*other.DatasourceVariableKind) {
			return false
		}
	}
	if resource.IntervalVariableKind == nil && other.IntervalVariableKind != nil || resource.IntervalVariableKind != nil && other.IntervalVariableKind == nil {
		return false
	}

	if resource.IntervalVariableKind != nil {
		if !resource.IntervalVariableKind.Equals(*other.IntervalVariableKind) {
			return false
		}
	}
	if resource.CustomVariableKind == nil && other.CustomVariableKind != nil || resource.CustomVariableKind != nil && other.CustomVariableKind == nil {
		return false
	}

	if resource.CustomVariableKind != nil {
		if !resource.CustomVariableKind.Equals(*other.CustomVariableKind) {
			return false
		}
	}
	if resource.GroupByVariableKind == nil && other.GroupByVariableKind != nil || resource.GroupByVariableKind != nil && other.GroupByVariableKind == nil {
		return false
	}

	if resource.GroupByVariableKind != nil {
		if !resource.GroupByVariableKind.Equals(*other.GroupByVariableKind) {
			return false
		}
	}
	if resource.AdhocVariableKind == nil && other.AdhocVariableKind != nil || resource.AdhocVariableKind != nil && other.AdhocVariableKind == nil {
		return false
	}

	if resource.AdhocVariableKind != nil {
		if !resource.AdhocVariableKind.Equals(*other.AdhocVariableKind) {
			return false
		}
	}
	if resource.SwitchVariableKind == nil && other.SwitchVariableKind != nil || resource.SwitchVariableKind != nil && other.SwitchVariableKind == nil {
		return false
	}

	if resource.SwitchVariableKind != nil {
		if !resource.SwitchVariableKind.Equals(*other.SwitchVariableKind) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind` fields for violations and returns them.
func (resource QueryVariableKindOrTextVariableKindOrConstantVariableKindOrDatasourceVariableKindOrIntervalVariableKindOrCustomVariableKindOrGroupByVariableKindOrAdhocVariableKindOrSwitchVariableKind) Validate() error {
	var errs cog.BuildErrors
	if resource.QueryVariableKind != nil {
		if err := resource.QueryVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryVariableKind", err)...)
		}
	}
	if resource.TextVariableKind != nil {
		if err := resource.TextVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TextVariableKind", err)...)
		}
	}
	if resource.ConstantVariableKind != nil {
		if err := resource.ConstantVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ConstantVariableKind", err)...)
		}
	}
	if resource.DatasourceVariableKind != nil {
		if err := resource.DatasourceVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("DatasourceVariableKind", err)...)
		}
	}
	if resource.IntervalVariableKind != nil {
		if err := resource.IntervalVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("IntervalVariableKind", err)...)
		}
	}
	if resource.CustomVariableKind != nil {
		if err := resource.CustomVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CustomVariableKind", err)...)
		}
	}
	if resource.GroupByVariableKind != nil {
		if err := resource.GroupByVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("GroupByVariableKind", err)...)
		}
	}
	if resource.AdhocVariableKind != nil {
		if err := resource.AdhocVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("AdhocVariableKind", err)...)
		}
	}
	if resource.SwitchVariableKind != nil {
		if err := resource.SwitchVariableKind.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("SwitchVariableKind", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type StringOrArrayOfString struct {
	String        *string  `json:"String,omitempty"`
	ArrayOfString []string `json:"ArrayOfString,omitempty"`
}

// NewStringOrArrayOfString creates a new StringOrArrayOfString object.
func NewStringOrArrayOfString() *StringOrArrayOfString {
	return &StringOrArrayOfString{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrArrayOfString` as JSON.
func (resource StringOrArrayOfString) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.ArrayOfString != nil {
		return json.Marshal(resource.ArrayOfString)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.
func (resource *StringOrArrayOfString) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// ArrayOfString
	var ArrayOfString []string
	if err := json.Unmarshal(raw, &ArrayOfString); err != nil {
		errList = append(errList, err)
		resource.ArrayOfString = nil
	} else {
		resource.ArrayOfString = ArrayOfString
		return nil
	}

	return errors.Join(errList...)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrArrayOfString) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string

	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// ArrayOfString
	var ArrayOfString []string

	if err := json.Unmarshal(raw, &ArrayOfString); err != nil {
		errList = append(errList, err)
	} else {
		resource.ArrayOfString = ArrayOfString
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrArrayOfString", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrArrayOfString` objects.
func (resource StringOrArrayOfString) Equals(other StringOrArrayOfString) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}

	if len(resource.ArrayOfString) != len(other.ArrayOfString) {
		return false
	}

	for i1 := range resource.ArrayOfString {
		if resource.ArrayOfString[i1] != other.ArrayOfString[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrArrayOfString` fields for violations and returns them.
func (resource StringOrArrayOfString) Validate() error {
	return nil
}

type StringOrFloat64 struct {
	String  *string  `json:"String,omitempty"`
	Float64 *float64 `json:"Float64,omitempty"`
}

// NewStringOrFloat64 creates a new StringOrFloat64 object.
func NewStringOrFloat64() *StringOrFloat64 {
	return &StringOrFloat64{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrFloat64` as JSON.
func (resource StringOrFloat64) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.Float64 != nil {
		return json.Marshal(resource.Float64)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrFloat64` from JSON.
func (resource *StringOrFloat64) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
		resource.Float64 = nil
	} else {
		resource.Float64 = &Float64
		return nil
	}

	return errors.Join(errList...)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrFloat64` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrFloat64) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string

	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// Float64
	var Float64 float64

	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
	} else {
		resource.Float64 = &Float64
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrFloat64", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrFloat64` objects.
func (resource StringOrFloat64) Equals(other StringOrFloat64) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.Float64 == nil && other.Float64 != nil || resource.Float64 != nil && other.Float64 == nil {
		return false
	}

	if resource.Float64 != nil {
		if *resource.Float64 != *other.Float64 {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrFloat64` fields for violations and returns them.
func (resource StringOrFloat64) Validate() error {
	return nil
}

type StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle struct {
	String                     *string               `json:"String,omitempty"`
	Bool                       *bool                 `json:"Bool,omitempty"`
	Float64                    *float64              `json:"Float64,omitempty"`
	CustomVariableValue        *CustomVariableValue  `json:"CustomVariableValue,omitempty"`
	ArrayOfVariableValueSingle []VariableValueSingle `json:"ArrayOfVariableValueSingle,omitempty"`
}

// NewStringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle creates a new StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle object.
func NewStringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle() *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle {
	return &StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` as JSON.
func (resource StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}
	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}
	if resource.Float64 != nil {
		return json.Marshal(resource.Float64)
	}
	if resource.CustomVariableValue != nil {
		return json.Marshal(resource.CustomVariableValue)
	}
	if resource.ArrayOfVariableValueSingle != nil {
		return json.Marshal(resource.ArrayOfVariableValueSingle)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` from JSON.
func (resource *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
		resource.Float64 = nil
	} else {
		resource.Float64 = &Float64
		return nil
	}

	// CustomVariableValue
	var CustomVariableValue CustomVariableValue
	customVariableValuedec := json.NewDecoder(bytes.NewReader(raw))
	customVariableValuedec.DisallowUnknownFields()
	if err := customVariableValuedec.Decode(&CustomVariableValue); err != nil {
		errList = append(errList, err)
		resource.CustomVariableValue = nil
	} else {
		resource.CustomVariableValue = &CustomVariableValue
		return nil
	}

	// ArrayOfVariableValueSingle
	var ArrayOfVariableValueSingle []VariableValueSingle
	if err := json.Unmarshal(raw, &ArrayOfVariableValueSingle); err != nil {
		errList = append(errList, err)
		resource.ArrayOfVariableValueSingle = nil
	} else {
		resource.ArrayOfVariableValueSingle = ArrayOfVariableValueSingle
		return nil
	}

	return errors.Join(errList...)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
	} else {
		resource.Float64 = &Float64
		return nil
	}

	// CustomVariableValue
	var CustomVariableValue CustomVariableValue
	customVariableValuedec := json.NewDecoder(bytes.NewReader(raw))
	customVariableValuedec.DisallowUnknownFields()
	if err := customVariableValuedec.Decode(&CustomVariableValue); err != nil {
		errList = append(errList, err)
	} else {
		resource.CustomVariableValue = &CustomVariableValue
		return nil
	}

	// ArrayOfVariableValueSingle
	var ArrayOfVariableValueSingle []VariableValueSingle
	if err := json.Unmarshal(raw, &ArrayOfVariableValueSingle); err != nil {
		errList = append(errList, err)
	} else {
		resource.ArrayOfVariableValueSingle = ArrayOfVariableValueSingle
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` objects.
func (resource StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) Equals(other StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.Bool == nil && other.Bool != nil || resource.Bool != nil && other.Bool == nil {
		return false
	}

	if resource.Bool != nil {
		if *resource.Bool != *other.Bool {
			return false
		}
	}
	if resource.Float64 == nil && other.Float64 != nil || resource.Float64 != nil && other.Float64 == nil {
		return false
	}

	if resource.Float64 != nil {
		if *resource.Float64 != *other.Float64 {
			return false
		}
	}
	if resource.CustomVariableValue == nil && other.CustomVariableValue != nil || resource.CustomVariableValue != nil && other.CustomVariableValue == nil {
		return false
	}

	if resource.CustomVariableValue != nil {
		if !resource.CustomVariableValue.Equals(*other.CustomVariableValue) {
			return false
		}
	}

	if len(resource.ArrayOfVariableValueSingle) != len(other.ArrayOfVariableValueSingle) {
		return false
	}

	for i1 := range resource.ArrayOfVariableValueSingle {
		if !resource.ArrayOfVariableValueSingle[i1].Equals(other.ArrayOfVariableValueSingle[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` fields for violations and returns them.
func (resource StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) Validate() error {
	var errs cog.BuildErrors
	if resource.CustomVariableValue != nil {
		if err := resource.CustomVariableValue.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CustomVariableValue", err)...)
		}
	}

	for i1 := range resource.ArrayOfVariableValueSingle {
		if err := resource.ArrayOfVariableValueSingle[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ArrayOfVariableValueSingle["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type StringOrBoolOrFloat64OrCustomVariableValue struct {
	String              *string              `json:"String,omitempty"`
	Bool                *bool                `json:"Bool,omitempty"`
	Float64             *float64             `json:"Float64,omitempty"`
	CustomVariableValue *CustomVariableValue `json:"CustomVariableValue,omitempty"`
}

// NewStringOrBoolOrFloat64OrCustomVariableValue creates a new StringOrBoolOrFloat64OrCustomVariableValue object.
func NewStringOrBoolOrFloat64OrCustomVariableValue() *StringOrBoolOrFloat64OrCustomVariableValue {
	return &StringOrBoolOrFloat64OrCustomVariableValue{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrFloat64OrCustomVariableValue` as JSON.
func (resource StringOrBoolOrFloat64OrCustomVariableValue) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}
	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}
	if resource.Float64 != nil {
		return json.Marshal(resource.Float64)
	}
	if resource.CustomVariableValue != nil {
		return json.Marshal(resource.CustomVariableValue)
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValue` from JSON.
func (resource *StringOrBoolOrFloat64OrCustomVariableValue) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
		resource.Float64 = nil
	} else {
		resource.Float64 = &Float64
		return nil
	}

	// CustomVariableValue
	var CustomVariableValue CustomVariableValue
	customVariableValuedec := json.NewDecoder(bytes.NewReader(raw))
	customVariableValuedec.DisallowUnknownFields()
	if err := customVariableValuedec.Decode(&CustomVariableValue); err != nil {
		errList = append(errList, err)
		resource.CustomVariableValue = nil
	} else {
		resource.CustomVariableValue = &CustomVariableValue
		return nil
	}

	return errors.Join(errList...)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValue` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrBoolOrFloat64OrCustomVariableValue) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
	} else {
		resource.Float64 = &Float64
		return nil
	}

	// CustomVariableValue
	var CustomVariableValue CustomVariableValue
	customVariableValuedec := json.NewDecoder(bytes.NewReader(raw))
	customVariableValuedec.DisallowUnknownFields()
	if err := customVariableValuedec.Decode(&CustomVariableValue); err != nil {
		errList = append(errList, err)
	} else {
		resource.CustomVariableValue = &CustomVariableValue
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrBoolOrFloat64OrCustomVariableValue", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrBoolOrFloat64OrCustomVariableValue` objects.
func (resource StringOrBoolOrFloat64OrCustomVariableValue) Equals(other StringOrBoolOrFloat64OrCustomVariableValue) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.Bool == nil && other.Bool != nil || resource.Bool != nil && other.Bool == nil {
		return false
	}

	if resource.Bool != nil {
		if *resource.Bool != *other.Bool {
			return false
		}
	}
	if resource.Float64 == nil && other.Float64 != nil || resource.Float64 != nil && other.Float64 == nil {
		return false
	}

	if resource.Float64 != nil {
		if *resource.Float64 != *other.Float64 {
			return false
		}
	}
	if resource.CustomVariableValue == nil && other.CustomVariableValue != nil || resource.CustomVariableValue != nil && other.CustomVariableValue == nil {
		return false
	}

	if resource.CustomVariableValue != nil {
		if !resource.CustomVariableValue.Equals(*other.CustomVariableValue) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrBoolOrFloat64OrCustomVariableValue` fields for violations and returns them.
func (resource StringOrBoolOrFloat64OrCustomVariableValue) Validate() error {
	var errs cog.BuildErrors
	if resource.CustomVariableValue != nil {
		if err := resource.CustomVariableValue.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CustomVariableValue", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
