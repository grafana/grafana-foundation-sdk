// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TODO docs
type DataSourceJsonData struct {
	AuthType        *string `json:"authType,omitempty"`
	DefaultRegion   *string `json:"defaultRegion,omitempty"`
	Profile         *string `json:"profile,omitempty"`
	ManageAlerts    *bool   `json:"manageAlerts,omitempty"`
	AlertmanagerUid *string `json:"alertmanagerUid,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataSourceJsonData` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DataSourceJsonData) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "authType"
	if fields["authType"] != nil {
		if string(fields["authType"]) != "null" {
			if err := json.Unmarshal(fields["authType"], &resource.AuthType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("authType", err)...)
			}

		}
		delete(fields, "authType")

	}
	// Field "defaultRegion"
	if fields["defaultRegion"] != nil {
		if string(fields["defaultRegion"]) != "null" {
			if err := json.Unmarshal(fields["defaultRegion"], &resource.DefaultRegion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("defaultRegion", err)...)
			}

		}
		delete(fields, "defaultRegion")

	}
	// Field "profile"
	if fields["profile"] != nil {
		if string(fields["profile"]) != "null" {
			if err := json.Unmarshal(fields["profile"], &resource.Profile); err != nil {
				errs = append(errs, cog.MakeBuildErrors("profile", err)...)
			}

		}
		delete(fields, "profile")

	}
	// Field "manageAlerts"
	if fields["manageAlerts"] != nil {
		if string(fields["manageAlerts"]) != "null" {
			if err := json.Unmarshal(fields["manageAlerts"], &resource.ManageAlerts); err != nil {
				errs = append(errs, cog.MakeBuildErrors("manageAlerts", err)...)
			}

		}
		delete(fields, "manageAlerts")

	}
	// Field "alertmanagerUid"
	if fields["alertmanagerUid"] != nil {
		if string(fields["alertmanagerUid"]) != "null" {
			if err := json.Unmarshal(fields["alertmanagerUid"], &resource.AlertmanagerUid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alertmanagerUid", err)...)
			}

		}
		delete(fields, "alertmanagerUid")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DataSourceJsonData", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DataSourceJsonData` objects.
func (resource DataSourceJsonData) Equals(other DataSourceJsonData) bool {
	if resource.AuthType == nil && other.AuthType != nil || resource.AuthType != nil && other.AuthType == nil {
		return false
	}

	if resource.AuthType != nil {
		if *resource.AuthType != *other.AuthType {
			return false
		}
	}
	if resource.DefaultRegion == nil && other.DefaultRegion != nil || resource.DefaultRegion != nil && other.DefaultRegion == nil {
		return false
	}

	if resource.DefaultRegion != nil {
		if *resource.DefaultRegion != *other.DefaultRegion {
			return false
		}
	}
	if resource.Profile == nil && other.Profile != nil || resource.Profile != nil && other.Profile == nil {
		return false
	}

	if resource.Profile != nil {
		if *resource.Profile != *other.Profile {
			return false
		}
	}
	if resource.ManageAlerts == nil && other.ManageAlerts != nil || resource.ManageAlerts != nil && other.ManageAlerts == nil {
		return false
	}

	if resource.ManageAlerts != nil {
		if *resource.ManageAlerts != *other.ManageAlerts {
			return false
		}
	}
	if resource.AlertmanagerUid == nil && other.AlertmanagerUid != nil || resource.AlertmanagerUid != nil && other.AlertmanagerUid == nil {
		return false
	}

	if resource.AlertmanagerUid != nil {
		if *resource.AlertmanagerUid != *other.AlertmanagerUid {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DataSourceJsonData` fields for violations and returns them.
func (resource DataSourceJsonData) Validate() error {
	return nil
}

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DataQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {
			if err := json.Unmarshal(fields["datasource"], &resource.Datasource); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DataQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DataQuery` objects.
func (resource DataQuery) Equals(other DataQuery) bool {
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Datasource, other.Datasource) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DataQuery` fields for violations and returns them.
func (resource DataQuery) Validate() error {
	return nil
}

type BaseDimensionConfig struct {
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseDimensionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BaseDimensionConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BaseDimensionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BaseDimensionConfig` objects.
func (resource BaseDimensionConfig) Equals(other BaseDimensionConfig) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BaseDimensionConfig` fields for violations and returns them.
func (resource BaseDimensionConfig) Validate() error {
	return nil
}

type ScaleDimensionMode string

const (
	ScaleDimensionModeLinear ScaleDimensionMode = "linear"
	ScaleDimensionModeQuad   ScaleDimensionMode = "quad"
)

type ScaleDimensionConfig struct {
	Min   float64  `json:"min"`
	Max   float64  `json:"max"`
	Fixed *float64 `json:"fixed,omitempty"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
	// | *"linear"
	Mode *ScaleDimensionMode `json:"mode,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScaleDimensionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ScaleDimensionConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "min"
	if fields["min"] != nil {
		if string(fields["min"]) != "null" {
			if err := json.Unmarshal(fields["min"], &resource.Min); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("min", errors.New("required field is null"))...)

		}
		delete(fields, "min")
	} else {
		errs = append(errs, cog.MakeBuildErrors("min", errors.New("required field is missing from input"))...)
	}
	// Field "max"
	if fields["max"] != nil {
		if string(fields["max"]) != "null" {
			if err := json.Unmarshal(fields["max"], &resource.Max); err != nil {
				errs = append(errs, cog.MakeBuildErrors("max", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("max", errors.New("required field is null"))...)

		}
		delete(fields, "max")
	} else {
		errs = append(errs, cog.MakeBuildErrors("max", errors.New("required field is missing from input"))...)
	}
	// Field "fixed"
	if fields["fixed"] != nil {
		if string(fields["fixed"]) != "null" {
			if err := json.Unmarshal(fields["fixed"], &resource.Fixed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixed", err)...)
			}

		}
		delete(fields, "fixed")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}

		}
		delete(fields, "mode")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ScaleDimensionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ScaleDimensionConfig` objects.
func (resource ScaleDimensionConfig) Equals(other ScaleDimensionConfig) bool {
	if resource.Min != other.Min {
		return false
	}
	if resource.Max != other.Max {
		return false
	}
	if resource.Fixed == nil && other.Fixed != nil || resource.Fixed != nil && other.Fixed == nil {
		return false
	}

	if resource.Fixed != nil {
		if *resource.Fixed != *other.Fixed {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ScaleDimensionConfig` fields for violations and returns them.
func (resource ScaleDimensionConfig) Validate() error {
	return nil
}

type ColorDimensionConfig struct {
	// color value
	Fixed *string `json:"fixed,omitempty"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ColorDimensionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ColorDimensionConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "fixed"
	if fields["fixed"] != nil {
		if string(fields["fixed"]) != "null" {
			if err := json.Unmarshal(fields["fixed"], &resource.Fixed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixed", err)...)
			}

		}
		delete(fields, "fixed")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ColorDimensionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ColorDimensionConfig` objects.
func (resource ColorDimensionConfig) Equals(other ColorDimensionConfig) bool {
	if resource.Fixed == nil && other.Fixed != nil || resource.Fixed != nil && other.Fixed == nil {
		return false
	}

	if resource.Fixed != nil {
		if *resource.Fixed != *other.Fixed {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ColorDimensionConfig` fields for violations and returns them.
func (resource ColorDimensionConfig) Validate() error {
	return nil
}

type ScalarDimensionMode string

const (
	ScalarDimensionModeMod     ScalarDimensionMode = "mod"
	ScalarDimensionModeClamped ScalarDimensionMode = "clamped"
)

type ScalarDimensionConfig struct {
	Min   float64  `json:"min"`
	Max   float64  `json:"max"`
	Fixed *float64 `json:"fixed,omitempty"`
	// fixed: T -- will be added by each element
	Field *string              `json:"field,omitempty"`
	Mode  *ScalarDimensionMode `json:"mode,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScalarDimensionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ScalarDimensionConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "min"
	if fields["min"] != nil {
		if string(fields["min"]) != "null" {
			if err := json.Unmarshal(fields["min"], &resource.Min); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("min", errors.New("required field is null"))...)

		}
		delete(fields, "min")
	} else {
		errs = append(errs, cog.MakeBuildErrors("min", errors.New("required field is missing from input"))...)
	}
	// Field "max"
	if fields["max"] != nil {
		if string(fields["max"]) != "null" {
			if err := json.Unmarshal(fields["max"], &resource.Max); err != nil {
				errs = append(errs, cog.MakeBuildErrors("max", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("max", errors.New("required field is null"))...)

		}
		delete(fields, "max")
	} else {
		errs = append(errs, cog.MakeBuildErrors("max", errors.New("required field is missing from input"))...)
	}
	// Field "fixed"
	if fields["fixed"] != nil {
		if string(fields["fixed"]) != "null" {
			if err := json.Unmarshal(fields["fixed"], &resource.Fixed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixed", err)...)
			}

		}
		delete(fields, "fixed")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}

		}
		delete(fields, "mode")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ScalarDimensionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ScalarDimensionConfig` objects.
func (resource ScalarDimensionConfig) Equals(other ScalarDimensionConfig) bool {
	if resource.Min != other.Min {
		return false
	}
	if resource.Max != other.Max {
		return false
	}
	if resource.Fixed == nil && other.Fixed != nil || resource.Fixed != nil && other.Fixed == nil {
		return false
	}

	if resource.Fixed != nil {
		if *resource.Fixed != *other.Fixed {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ScalarDimensionConfig` fields for violations and returns them.
func (resource ScalarDimensionConfig) Validate() error {
	return nil
}

type TextDimensionMode string

const (
	TextDimensionModeFixed    TextDimensionMode = "fixed"
	TextDimensionModeField    TextDimensionMode = "field"
	TextDimensionModeTemplate TextDimensionMode = "template"
)

type TextDimensionConfig struct {
	Mode TextDimensionMode `json:"mode"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
	Fixed *string `json:"fixed,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TextDimensionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TextDimensionConfig) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}
	// Field "fixed"
	if fields["fixed"] != nil {
		if string(fields["fixed"]) != "null" {
			if err := json.Unmarshal(fields["fixed"], &resource.Fixed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixed", err)...)
			}

		}
		delete(fields, "fixed")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TextDimensionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TextDimensionConfig` objects.
func (resource TextDimensionConfig) Equals(other TextDimensionConfig) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Fixed == nil && other.Fixed != nil || resource.Fixed != nil && other.Fixed == nil {
		return false
	}

	if resource.Fixed != nil {
		if *resource.Fixed != *other.Fixed {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TextDimensionConfig` fields for violations and returns them.
func (resource TextDimensionConfig) Validate() error {
	return nil
}

type ResourceDimensionMode string

const (
	ResourceDimensionModeFixed   ResourceDimensionMode = "fixed"
	ResourceDimensionModeField   ResourceDimensionMode = "field"
	ResourceDimensionModeMapping ResourceDimensionMode = "mapping"
)

type MapLayerOptions struct {
	Type string `json:"type"`
	// configured unique display name
	Name string `json:"name"`
	// Custom options depending on the type
	Config any `json:"config,omitempty"`
	// Common method to define geometry fields
	Location *FrameGeometrySource `json:"location,omitempty"`
	// Defines a frame MatcherConfig that may filter data for the given layer
	FilterData any `json:"filterData,omitempty"`
	// Common properties:
	// https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
	// Layer opacity (0-1)
	Opacity *int64 `json:"opacity,omitempty"`
	// Check tooltip (defaults to true)
	Tooltip *bool `json:"tooltip,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MapLayerOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MapLayerOptions) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "config"
	if fields["config"] != nil {
		if string(fields["config"]) != "null" {
			if err := json.Unmarshal(fields["config"], &resource.Config); err != nil {
				errs = append(errs, cog.MakeBuildErrors("config", err)...)
			}

		}
		delete(fields, "config")

	}
	// Field "location"
	if fields["location"] != nil {
		if string(fields["location"]) != "null" {

			resource.Location = &FrameGeometrySource{}
			if err := resource.Location.UnmarshalJSONStrict(fields["location"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("location", err)...)
			}

		}
		delete(fields, "location")

	}
	// Field "filterData"
	if fields["filterData"] != nil {
		if string(fields["filterData"]) != "null" {
			if err := json.Unmarshal(fields["filterData"], &resource.FilterData); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filterData", err)...)
			}

		}
		delete(fields, "filterData")

	}
	// Field "opacity"
	if fields["opacity"] != nil {
		if string(fields["opacity"]) != "null" {
			if err := json.Unmarshal(fields["opacity"], &resource.Opacity); err != nil {
				errs = append(errs, cog.MakeBuildErrors("opacity", err)...)
			}

		}
		delete(fields, "opacity")

	}
	// Field "tooltip"
	if fields["tooltip"] != nil {
		if string(fields["tooltip"]) != "null" {
			if err := json.Unmarshal(fields["tooltip"], &resource.Tooltip); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
			}

		}
		delete(fields, "tooltip")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MapLayerOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MapLayerOptions` objects.
func (resource MapLayerOptions) Equals(other MapLayerOptions) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Config, other.Config) {
		return false
	}
	if resource.Location == nil && other.Location != nil || resource.Location != nil && other.Location == nil {
		return false
	}

	if resource.Location != nil {
		if !resource.Location.Equals(*other.Location) {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.FilterData, other.FilterData) {
		return false
	}
	if resource.Opacity == nil && other.Opacity != nil || resource.Opacity != nil && other.Opacity == nil {
		return false
	}

	if resource.Opacity != nil {
		if *resource.Opacity != *other.Opacity {
			return false
		}
	}
	if resource.Tooltip == nil && other.Tooltip != nil || resource.Tooltip != nil && other.Tooltip == nil {
		return false
	}

	if resource.Tooltip != nil {
		if *resource.Tooltip != *other.Tooltip {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MapLayerOptions` fields for violations and returns them.
func (resource MapLayerOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.Location != nil {
		if err := resource.Location.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("location", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type FrameGeometrySourceMode string

const (
	FrameGeometrySourceModeAuto    FrameGeometrySourceMode = "auto"
	FrameGeometrySourceModeGeohash FrameGeometrySourceMode = "geohash"
	FrameGeometrySourceModeCoords  FrameGeometrySourceMode = "coords"
	FrameGeometrySourceModeLookup  FrameGeometrySourceMode = "lookup"
)

type HeatmapCalculationMode string

const (
	HeatmapCalculationModeSize  HeatmapCalculationMode = "size"
	HeatmapCalculationModeCount HeatmapCalculationMode = "count"
)

type HeatmapCellLayout string

const (
	HeatmapCellLayoutLe      HeatmapCellLayout = "le"
	HeatmapCellLayoutGe      HeatmapCellLayout = "ge"
	HeatmapCellLayoutUnknown HeatmapCellLayout = "unknown"
	HeatmapCellLayoutAuto    HeatmapCellLayout = "auto"
)

type HeatmapCalculationBucketConfig struct {
	// Sets the bucket calculation mode
	Mode *HeatmapCalculationMode `json:"mode,omitempty"`
	// The number of buckets to use for the axis in the heatmap
	Value *string `json:"value,omitempty"`
	// Controls the scale of the buckets
	Scale *ScaleDistributionConfig `json:"scale,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapCalculationBucketConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HeatmapCalculationBucketConfig) UnmarshalJSONStrict(raw []byte) error {
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

		}
		delete(fields, "mode")

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
	// Field "scale"
	if fields["scale"] != nil {
		if string(fields["scale"]) != "null" {

			resource.Scale = &ScaleDistributionConfig{}
			if err := resource.Scale.UnmarshalJSONStrict(fields["scale"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scale", err)...)
			}

		}
		delete(fields, "scale")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HeatmapCalculationBucketConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HeatmapCalculationBucketConfig` objects.
func (resource HeatmapCalculationBucketConfig) Equals(other HeatmapCalculationBucketConfig) bool {
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
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
	if resource.Scale == nil && other.Scale != nil || resource.Scale != nil && other.Scale == nil {
		return false
	}

	if resource.Scale != nil {
		if !resource.Scale.Equals(*other.Scale) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `HeatmapCalculationBucketConfig` fields for violations and returns them.
func (resource HeatmapCalculationBucketConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.Scale != nil {
		if err := resource.Scale.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scale", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type LogsSortOrder string

const (
	LogsSortOrderDescending LogsSortOrder = "Descending"
	LogsSortOrderAscending  LogsSortOrder = "Ascending"
)

// TODO docs
type AxisPlacement string

const (
	AxisPlacementAuto   AxisPlacement = "auto"
	AxisPlacementTop    AxisPlacement = "top"
	AxisPlacementRight  AxisPlacement = "right"
	AxisPlacementBottom AxisPlacement = "bottom"
	AxisPlacementLeft   AxisPlacement = "left"
	AxisPlacementHidden AxisPlacement = "hidden"
)

// TODO docs
type AxisColorMode string

const (
	AxisColorModeText   AxisColorMode = "text"
	AxisColorModeSeries AxisColorMode = "series"
)

// TODO docs
type VisibilityMode string

const (
	VisibilityModeAuto   VisibilityMode = "auto"
	VisibilityModeNever  VisibilityMode = "never"
	VisibilityModeAlways VisibilityMode = "always"
)

// TODO docs
type GraphDrawStyle string

const (
	GraphDrawStyleLine   GraphDrawStyle = "line"
	GraphDrawStyleBars   GraphDrawStyle = "bars"
	GraphDrawStylePoints GraphDrawStyle = "points"
)

// TODO docs
type GraphTransform string

const (
	GraphTransformConstant  GraphTransform = "constant"
	GraphTransformNegativeY GraphTransform = "negative-Y"
)

// TODO docs
type LineInterpolation string

const (
	LineInterpolationLinear     LineInterpolation = "linear"
	LineInterpolationSmooth     LineInterpolation = "smooth"
	LineInterpolationStepBefore LineInterpolation = "stepBefore"
	LineInterpolationStepAfter  LineInterpolation = "stepAfter"
)

// TODO docs
type ScaleDistribution string

const (
	ScaleDistributionLinear  ScaleDistribution = "linear"
	ScaleDistributionLog     ScaleDistribution = "log"
	ScaleDistributionOrdinal ScaleDistribution = "ordinal"
	ScaleDistributionSymlog  ScaleDistribution = "symlog"
)

// TODO docs
type GraphGradientMode string

const (
	GraphGradientModeNone    GraphGradientMode = "none"
	GraphGradientModeOpacity GraphGradientMode = "opacity"
	GraphGradientModeHue     GraphGradientMode = "hue"
	GraphGradientModeScheme  GraphGradientMode = "scheme"
)

// TODO docs
type StackingMode string

const (
	StackingModeNone    StackingMode = "none"
	StackingModeNormal  StackingMode = "normal"
	StackingModePercent StackingMode = "percent"
)

// TODO docs
type BarAlignment int64

const (
	BarAlignmentBefore BarAlignment = -1
	BarAlignmentCenter BarAlignment = 0
	BarAlignmentAfter  BarAlignment = 1
)

// TODO docs
type ScaleOrientation int64

const (
	ScaleOrientationHorizontal ScaleOrientation = 0
	ScaleOrientationVertical   ScaleOrientation = 1
)

// TODO docs
type ScaleDirection int64

const (
	ScaleDirectionUp    ScaleDirection = 1
	ScaleDirectionRight ScaleDirection = 1
	ScaleDirectionDown  ScaleDirection = -1
	ScaleDirectionLeft  ScaleDirection = -1
)

// TODO docs
type LineStyle struct {
	Fill *LineStyleFill `json:"fill,omitempty"`
	Dash []float64      `json:"dash,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LineStyle` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LineStyle) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "fill"
	if fields["fill"] != nil {
		if string(fields["fill"]) != "null" {
			if err := json.Unmarshal(fields["fill"], &resource.Fill); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fill", err)...)
			}

		}
		delete(fields, "fill")

	}
	// Field "dash"
	if fields["dash"] != nil {
		if string(fields["dash"]) != "null" {

			if err := json.Unmarshal(fields["dash"], &resource.Dash); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dash", err)...)
			}

		}
		delete(fields, "dash")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LineStyle", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LineStyle` objects.
func (resource LineStyle) Equals(other LineStyle) bool {
	if resource.Fill == nil && other.Fill != nil || resource.Fill != nil && other.Fill == nil {
		return false
	}

	if resource.Fill != nil {
		if *resource.Fill != *other.Fill {
			return false
		}
	}

	if len(resource.Dash) != len(other.Dash) {
		return false
	}

	for i1 := range resource.Dash {
		if resource.Dash[i1] != other.Dash[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `LineStyle` fields for violations and returns them.
func (resource LineStyle) Validate() error {
	return nil
}

// TODO docs
type LineConfig struct {
	LineColor         *string            `json:"lineColor,omitempty"`
	LineWidth         *float64           `json:"lineWidth,omitempty"`
	LineInterpolation *LineInterpolation `json:"lineInterpolation,omitempty"`
	LineStyle         *LineStyle         `json:"lineStyle,omitempty"`
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls *BoolOrFloat64 `json:"spanNulls,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LineConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LineConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "lineColor"
	if fields["lineColor"] != nil {
		if string(fields["lineColor"]) != "null" {
			if err := json.Unmarshal(fields["lineColor"], &resource.LineColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineColor", err)...)
			}

		}
		delete(fields, "lineColor")

	}
	// Field "lineWidth"
	if fields["lineWidth"] != nil {
		if string(fields["lineWidth"]) != "null" {
			if err := json.Unmarshal(fields["lineWidth"], &resource.LineWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineWidth", err)...)
			}

		}
		delete(fields, "lineWidth")

	}
	// Field "lineInterpolation"
	if fields["lineInterpolation"] != nil {
		if string(fields["lineInterpolation"]) != "null" {
			if err := json.Unmarshal(fields["lineInterpolation"], &resource.LineInterpolation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineInterpolation", err)...)
			}

		}
		delete(fields, "lineInterpolation")

	}
	// Field "lineStyle"
	if fields["lineStyle"] != nil {
		if string(fields["lineStyle"]) != "null" {

			resource.LineStyle = &LineStyle{}
			if err := resource.LineStyle.UnmarshalJSONStrict(fields["lineStyle"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
			}

		}
		delete(fields, "lineStyle")

	}
	// Field "spanNulls"
	if fields["spanNulls"] != nil {
		if string(fields["spanNulls"]) != "null" {

			resource.SpanNulls = &BoolOrFloat64{}
			if err := resource.SpanNulls.UnmarshalJSONStrict(fields["spanNulls"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spanNulls", err)...)
			}

		}
		delete(fields, "spanNulls")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LineConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LineConfig` objects.
func (resource LineConfig) Equals(other LineConfig) bool {
	if resource.LineColor == nil && other.LineColor != nil || resource.LineColor != nil && other.LineColor == nil {
		return false
	}

	if resource.LineColor != nil {
		if *resource.LineColor != *other.LineColor {
			return false
		}
	}
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
			return false
		}
	}
	if resource.LineInterpolation == nil && other.LineInterpolation != nil || resource.LineInterpolation != nil && other.LineInterpolation == nil {
		return false
	}

	if resource.LineInterpolation != nil {
		if *resource.LineInterpolation != *other.LineInterpolation {
			return false
		}
	}
	if resource.LineStyle == nil && other.LineStyle != nil || resource.LineStyle != nil && other.LineStyle == nil {
		return false
	}

	if resource.LineStyle != nil {
		if !resource.LineStyle.Equals(*other.LineStyle) {
			return false
		}
	}
	if resource.SpanNulls == nil && other.SpanNulls != nil || resource.SpanNulls != nil && other.SpanNulls == nil {
		return false
	}

	if resource.SpanNulls != nil {
		if !resource.SpanNulls.Equals(*other.SpanNulls) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `LineConfig` fields for violations and returns them.
func (resource LineConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.LineStyle != nil {
		if err := resource.LineStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
		}
	}
	if resource.SpanNulls != nil {
		if err := resource.SpanNulls.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("spanNulls", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type BarConfig struct {
	BarAlignment   *BarAlignment `json:"barAlignment,omitempty"`
	BarWidthFactor *float64      `json:"barWidthFactor,omitempty"`
	BarMaxWidth    *float64      `json:"barMaxWidth,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BarConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BarConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "barAlignment"
	if fields["barAlignment"] != nil {
		if string(fields["barAlignment"]) != "null" {
			if err := json.Unmarshal(fields["barAlignment"], &resource.BarAlignment); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barAlignment", err)...)
			}

		}
		delete(fields, "barAlignment")

	}
	// Field "barWidthFactor"
	if fields["barWidthFactor"] != nil {
		if string(fields["barWidthFactor"]) != "null" {
			if err := json.Unmarshal(fields["barWidthFactor"], &resource.BarWidthFactor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barWidthFactor", err)...)
			}

		}
		delete(fields, "barWidthFactor")

	}
	// Field "barMaxWidth"
	if fields["barMaxWidth"] != nil {
		if string(fields["barMaxWidth"]) != "null" {
			if err := json.Unmarshal(fields["barMaxWidth"], &resource.BarMaxWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barMaxWidth", err)...)
			}

		}
		delete(fields, "barMaxWidth")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BarConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BarConfig` objects.
func (resource BarConfig) Equals(other BarConfig) bool {
	if resource.BarAlignment == nil && other.BarAlignment != nil || resource.BarAlignment != nil && other.BarAlignment == nil {
		return false
	}

	if resource.BarAlignment != nil {
		if *resource.BarAlignment != *other.BarAlignment {
			return false
		}
	}
	if resource.BarWidthFactor == nil && other.BarWidthFactor != nil || resource.BarWidthFactor != nil && other.BarWidthFactor == nil {
		return false
	}

	if resource.BarWidthFactor != nil {
		if *resource.BarWidthFactor != *other.BarWidthFactor {
			return false
		}
	}
	if resource.BarMaxWidth == nil && other.BarMaxWidth != nil || resource.BarMaxWidth != nil && other.BarMaxWidth == nil {
		return false
	}

	if resource.BarMaxWidth != nil {
		if *resource.BarMaxWidth != *other.BarMaxWidth {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BarConfig` fields for violations and returns them.
func (resource BarConfig) Validate() error {
	return nil
}

// TODO docs
type FillConfig struct {
	FillColor   *string  `json:"fillColor,omitempty"`
	FillOpacity *float64 `json:"fillOpacity,omitempty"`
	FillBelowTo *string  `json:"fillBelowTo,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FillConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FillConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "fillColor"
	if fields["fillColor"] != nil {
		if string(fields["fillColor"]) != "null" {
			if err := json.Unmarshal(fields["fillColor"], &resource.FillColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillColor", err)...)
			}

		}
		delete(fields, "fillColor")

	}
	// Field "fillOpacity"
	if fields["fillOpacity"] != nil {
		if string(fields["fillOpacity"]) != "null" {
			if err := json.Unmarshal(fields["fillOpacity"], &resource.FillOpacity); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillOpacity", err)...)
			}

		}
		delete(fields, "fillOpacity")

	}
	// Field "fillBelowTo"
	if fields["fillBelowTo"] != nil {
		if string(fields["fillBelowTo"]) != "null" {
			if err := json.Unmarshal(fields["fillBelowTo"], &resource.FillBelowTo); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillBelowTo", err)...)
			}

		}
		delete(fields, "fillBelowTo")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FillConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FillConfig` objects.
func (resource FillConfig) Equals(other FillConfig) bool {
	if resource.FillColor == nil && other.FillColor != nil || resource.FillColor != nil && other.FillColor == nil {
		return false
	}

	if resource.FillColor != nil {
		if *resource.FillColor != *other.FillColor {
			return false
		}
	}
	if resource.FillOpacity == nil && other.FillOpacity != nil || resource.FillOpacity != nil && other.FillOpacity == nil {
		return false
	}

	if resource.FillOpacity != nil {
		if *resource.FillOpacity != *other.FillOpacity {
			return false
		}
	}
	if resource.FillBelowTo == nil && other.FillBelowTo != nil || resource.FillBelowTo != nil && other.FillBelowTo == nil {
		return false
	}

	if resource.FillBelowTo != nil {
		if *resource.FillBelowTo != *other.FillBelowTo {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FillConfig` fields for violations and returns them.
func (resource FillConfig) Validate() error {
	return nil
}

// TODO docs
type PointsConfig struct {
	ShowPoints  *VisibilityMode `json:"showPoints,omitempty"`
	PointSize   *float64        `json:"pointSize,omitempty"`
	PointColor  *string         `json:"pointColor,omitempty"`
	PointSymbol *string         `json:"pointSymbol,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PointsConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PointsConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "showPoints"
	if fields["showPoints"] != nil {
		if string(fields["showPoints"]) != "null" {
			if err := json.Unmarshal(fields["showPoints"], &resource.ShowPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showPoints", err)...)
			}

		}
		delete(fields, "showPoints")

	}
	// Field "pointSize"
	if fields["pointSize"] != nil {
		if string(fields["pointSize"]) != "null" {
			if err := json.Unmarshal(fields["pointSize"], &resource.PointSize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
			}

		}
		delete(fields, "pointSize")

	}
	// Field "pointColor"
	if fields["pointColor"] != nil {
		if string(fields["pointColor"]) != "null" {
			if err := json.Unmarshal(fields["pointColor"], &resource.PointColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointColor", err)...)
			}

		}
		delete(fields, "pointColor")

	}
	// Field "pointSymbol"
	if fields["pointSymbol"] != nil {
		if string(fields["pointSymbol"]) != "null" {
			if err := json.Unmarshal(fields["pointSymbol"], &resource.PointSymbol); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSymbol", err)...)
			}

		}
		delete(fields, "pointSymbol")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PointsConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PointsConfig` objects.
func (resource PointsConfig) Equals(other PointsConfig) bool {
	if resource.ShowPoints == nil && other.ShowPoints != nil || resource.ShowPoints != nil && other.ShowPoints == nil {
		return false
	}

	if resource.ShowPoints != nil {
		if *resource.ShowPoints != *other.ShowPoints {
			return false
		}
	}
	if resource.PointSize == nil && other.PointSize != nil || resource.PointSize != nil && other.PointSize == nil {
		return false
	}

	if resource.PointSize != nil {
		if *resource.PointSize != *other.PointSize {
			return false
		}
	}
	if resource.PointColor == nil && other.PointColor != nil || resource.PointColor != nil && other.PointColor == nil {
		return false
	}

	if resource.PointColor != nil {
		if *resource.PointColor != *other.PointColor {
			return false
		}
	}
	if resource.PointSymbol == nil && other.PointSymbol != nil || resource.PointSymbol != nil && other.PointSymbol == nil {
		return false
	}

	if resource.PointSymbol != nil {
		if *resource.PointSymbol != *other.PointSymbol {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PointsConfig` fields for violations and returns them.
func (resource PointsConfig) Validate() error {
	return nil
}

// TODO docs
type ScaleDistributionConfig struct {
	Type            ScaleDistribution `json:"type"`
	Log             *float64          `json:"log,omitempty"`
	LinearThreshold *float64          `json:"linearThreshold,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScaleDistributionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ScaleDistributionConfig) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "log"
	if fields["log"] != nil {
		if string(fields["log"]) != "null" {
			if err := json.Unmarshal(fields["log"], &resource.Log); err != nil {
				errs = append(errs, cog.MakeBuildErrors("log", err)...)
			}

		}
		delete(fields, "log")

	}
	// Field "linearThreshold"
	if fields["linearThreshold"] != nil {
		if string(fields["linearThreshold"]) != "null" {
			if err := json.Unmarshal(fields["linearThreshold"], &resource.LinearThreshold); err != nil {
				errs = append(errs, cog.MakeBuildErrors("linearThreshold", err)...)
			}

		}
		delete(fields, "linearThreshold")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ScaleDistributionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ScaleDistributionConfig` objects.
func (resource ScaleDistributionConfig) Equals(other ScaleDistributionConfig) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Log == nil && other.Log != nil || resource.Log != nil && other.Log == nil {
		return false
	}

	if resource.Log != nil {
		if *resource.Log != *other.Log {
			return false
		}
	}
	if resource.LinearThreshold == nil && other.LinearThreshold != nil || resource.LinearThreshold != nil && other.LinearThreshold == nil {
		return false
	}

	if resource.LinearThreshold != nil {
		if *resource.LinearThreshold != *other.LinearThreshold {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ScaleDistributionConfig` fields for violations and returns them.
func (resource ScaleDistributionConfig) Validate() error {
	return nil
}

// TODO docs
type AxisConfig struct {
	AxisPlacement     *AxisPlacement           `json:"axisPlacement,omitempty"`
	AxisColorMode     *AxisColorMode           `json:"axisColorMode,omitempty"`
	AxisLabel         *string                  `json:"axisLabel,omitempty"`
	AxisWidth         *float64                 `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                 `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                 `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                    `json:"axisGridShow,omitempty"`
	ScaleDistribution *ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                    `json:"axisCenteredZero,omitempty"`
	AxisBorderShow    *bool                    `json:"axisBorderShow,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AxisConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AxisConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "axisPlacement"
	if fields["axisPlacement"] != nil {
		if string(fields["axisPlacement"]) != "null" {
			if err := json.Unmarshal(fields["axisPlacement"], &resource.AxisPlacement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisPlacement", err)...)
			}

		}
		delete(fields, "axisPlacement")

	}
	// Field "axisColorMode"
	if fields["axisColorMode"] != nil {
		if string(fields["axisColorMode"]) != "null" {
			if err := json.Unmarshal(fields["axisColorMode"], &resource.AxisColorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisColorMode", err)...)
			}

		}
		delete(fields, "axisColorMode")

	}
	// Field "axisLabel"
	if fields["axisLabel"] != nil {
		if string(fields["axisLabel"]) != "null" {
			if err := json.Unmarshal(fields["axisLabel"], &resource.AxisLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisLabel", err)...)
			}

		}
		delete(fields, "axisLabel")

	}
	// Field "axisWidth"
	if fields["axisWidth"] != nil {
		if string(fields["axisWidth"]) != "null" {
			if err := json.Unmarshal(fields["axisWidth"], &resource.AxisWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisWidth", err)...)
			}

		}
		delete(fields, "axisWidth")

	}
	// Field "axisSoftMin"
	if fields["axisSoftMin"] != nil {
		if string(fields["axisSoftMin"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMin"], &resource.AxisSoftMin); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMin", err)...)
			}

		}
		delete(fields, "axisSoftMin")

	}
	// Field "axisSoftMax"
	if fields["axisSoftMax"] != nil {
		if string(fields["axisSoftMax"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMax"], &resource.AxisSoftMax); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMax", err)...)
			}

		}
		delete(fields, "axisSoftMax")

	}
	// Field "axisGridShow"
	if fields["axisGridShow"] != nil {
		if string(fields["axisGridShow"]) != "null" {
			if err := json.Unmarshal(fields["axisGridShow"], &resource.AxisGridShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisGridShow", err)...)
			}

		}
		delete(fields, "axisGridShow")

	}
	// Field "scaleDistribution"
	if fields["scaleDistribution"] != nil {
		if string(fields["scaleDistribution"]) != "null" {

			resource.ScaleDistribution = &ScaleDistributionConfig{}
			if err := resource.ScaleDistribution.UnmarshalJSONStrict(fields["scaleDistribution"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
			}

		}
		delete(fields, "scaleDistribution")

	}
	// Field "axisCenteredZero"
	if fields["axisCenteredZero"] != nil {
		if string(fields["axisCenteredZero"]) != "null" {
			if err := json.Unmarshal(fields["axisCenteredZero"], &resource.AxisCenteredZero); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisCenteredZero", err)...)
			}

		}
		delete(fields, "axisCenteredZero")

	}
	// Field "axisBorderShow"
	if fields["axisBorderShow"] != nil {
		if string(fields["axisBorderShow"]) != "null" {
			if err := json.Unmarshal(fields["axisBorderShow"], &resource.AxisBorderShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisBorderShow", err)...)
			}

		}
		delete(fields, "axisBorderShow")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AxisConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AxisConfig` objects.
func (resource AxisConfig) Equals(other AxisConfig) bool {
	if resource.AxisPlacement == nil && other.AxisPlacement != nil || resource.AxisPlacement != nil && other.AxisPlacement == nil {
		return false
	}

	if resource.AxisPlacement != nil {
		if *resource.AxisPlacement != *other.AxisPlacement {
			return false
		}
	}
	if resource.AxisColorMode == nil && other.AxisColorMode != nil || resource.AxisColorMode != nil && other.AxisColorMode == nil {
		return false
	}

	if resource.AxisColorMode != nil {
		if *resource.AxisColorMode != *other.AxisColorMode {
			return false
		}
	}
	if resource.AxisLabel == nil && other.AxisLabel != nil || resource.AxisLabel != nil && other.AxisLabel == nil {
		return false
	}

	if resource.AxisLabel != nil {
		if *resource.AxisLabel != *other.AxisLabel {
			return false
		}
	}
	if resource.AxisWidth == nil && other.AxisWidth != nil || resource.AxisWidth != nil && other.AxisWidth == nil {
		return false
	}

	if resource.AxisWidth != nil {
		if *resource.AxisWidth != *other.AxisWidth {
			return false
		}
	}
	if resource.AxisSoftMin == nil && other.AxisSoftMin != nil || resource.AxisSoftMin != nil && other.AxisSoftMin == nil {
		return false
	}

	if resource.AxisSoftMin != nil {
		if *resource.AxisSoftMin != *other.AxisSoftMin {
			return false
		}
	}
	if resource.AxisSoftMax == nil && other.AxisSoftMax != nil || resource.AxisSoftMax != nil && other.AxisSoftMax == nil {
		return false
	}

	if resource.AxisSoftMax != nil {
		if *resource.AxisSoftMax != *other.AxisSoftMax {
			return false
		}
	}
	if resource.AxisGridShow == nil && other.AxisGridShow != nil || resource.AxisGridShow != nil && other.AxisGridShow == nil {
		return false
	}

	if resource.AxisGridShow != nil {
		if *resource.AxisGridShow != *other.AxisGridShow {
			return false
		}
	}
	if resource.ScaleDistribution == nil && other.ScaleDistribution != nil || resource.ScaleDistribution != nil && other.ScaleDistribution == nil {
		return false
	}

	if resource.ScaleDistribution != nil {
		if !resource.ScaleDistribution.Equals(*other.ScaleDistribution) {
			return false
		}
	}
	if resource.AxisCenteredZero == nil && other.AxisCenteredZero != nil || resource.AxisCenteredZero != nil && other.AxisCenteredZero == nil {
		return false
	}

	if resource.AxisCenteredZero != nil {
		if *resource.AxisCenteredZero != *other.AxisCenteredZero {
			return false
		}
	}
	if resource.AxisBorderShow == nil && other.AxisBorderShow != nil || resource.AxisBorderShow != nil && other.AxisBorderShow == nil {
		return false
	}

	if resource.AxisBorderShow != nil {
		if *resource.AxisBorderShow != *other.AxisBorderShow {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AxisConfig` fields for violations and returns them.
func (resource AxisConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.ScaleDistribution != nil {
		if err := resource.ScaleDistribution.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type HideSeriesConfig struct {
	Tooltip bool `json:"tooltip"`
	Legend  bool `json:"legend"`
	Viz     bool `json:"viz"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HideSeriesConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HideSeriesConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {
			if err := json.Unmarshal(fields["legend"], &resource.Legend); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is null"))...)

		}
		delete(fields, "legend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is missing from input"))...)
	}
	// Field "viz"
	if fields["viz"] != nil {
		if string(fields["viz"]) != "null" {
			if err := json.Unmarshal(fields["viz"], &resource.Viz); err != nil {
				errs = append(errs, cog.MakeBuildErrors("viz", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("viz", errors.New("required field is null"))...)

		}
		delete(fields, "viz")
	} else {
		errs = append(errs, cog.MakeBuildErrors("viz", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HideSeriesConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HideSeriesConfig` objects.
func (resource HideSeriesConfig) Equals(other HideSeriesConfig) bool {
	if resource.Tooltip != other.Tooltip {
		return false
	}
	if resource.Legend != other.Legend {
		return false
	}
	if resource.Viz != other.Viz {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `HideSeriesConfig` fields for violations and returns them.
func (resource HideSeriesConfig) Validate() error {
	return nil
}

// TODO docs
type StackingConfig struct {
	Mode  *StackingMode `json:"mode,omitempty"`
	Group *string       `json:"group,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StackingConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StackingConfig) UnmarshalJSONStrict(raw []byte) error {
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

		}
		delete(fields, "mode")

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
		errs = append(errs, cog.MakeBuildErrors("StackingConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StackingConfig` objects.
func (resource StackingConfig) Equals(other StackingConfig) bool {
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `StackingConfig` fields for violations and returns them.
func (resource StackingConfig) Validate() error {
	return nil
}

// TODO docs
type StackableFieldConfig struct {
	Stacking *StackingConfig `json:"stacking,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StackableFieldConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StackableFieldConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "stacking"
	if fields["stacking"] != nil {
		if string(fields["stacking"]) != "null" {

			resource.Stacking = &StackingConfig{}
			if err := resource.Stacking.UnmarshalJSONStrict(fields["stacking"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stacking", err)...)
			}

		}
		delete(fields, "stacking")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("StackableFieldConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StackableFieldConfig` objects.
func (resource StackableFieldConfig) Equals(other StackableFieldConfig) bool {
	if resource.Stacking == nil && other.Stacking != nil || resource.Stacking != nil && other.Stacking == nil {
		return false
	}

	if resource.Stacking != nil {
		if !resource.Stacking.Equals(*other.Stacking) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StackableFieldConfig` fields for violations and returns them.
func (resource StackableFieldConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.Stacking != nil {
		if err := resource.Stacking.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("stacking", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type HideableFieldConfig struct {
	HideFrom *HideSeriesConfig `json:"hideFrom,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HideableFieldConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HideableFieldConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "hideFrom"
	if fields["hideFrom"] != nil {
		if string(fields["hideFrom"]) != "null" {

			resource.HideFrom = &HideSeriesConfig{}
			if err := resource.HideFrom.UnmarshalJSONStrict(fields["hideFrom"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
			}

		}
		delete(fields, "hideFrom")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HideableFieldConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HideableFieldConfig` objects.
func (resource HideableFieldConfig) Equals(other HideableFieldConfig) bool {
	if resource.HideFrom == nil && other.HideFrom != nil || resource.HideFrom != nil && other.HideFrom == nil {
		return false
	}

	if resource.HideFrom != nil {
		if !resource.HideFrom.Equals(*other.HideFrom) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `HideableFieldConfig` fields for violations and returns them.
func (resource HideableFieldConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.HideFrom != nil {
		if err := resource.HideFrom.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type GraphTresholdsStyleMode string

const (
	GraphTresholdsStyleModeOff           GraphTresholdsStyleMode = "off"
	GraphTresholdsStyleModeLine          GraphTresholdsStyleMode = "line"
	GraphTresholdsStyleModeDashed        GraphTresholdsStyleMode = "dashed"
	GraphTresholdsStyleModeArea          GraphTresholdsStyleMode = "area"
	GraphTresholdsStyleModeLineAndArea   GraphTresholdsStyleMode = "line+area"
	GraphTresholdsStyleModeDashedAndArea GraphTresholdsStyleMode = "dashed+area"
	GraphTresholdsStyleModeSeries        GraphTresholdsStyleMode = "series"
)

// TODO docs
type GraphThresholdsStyleConfig struct {
	Mode GraphTresholdsStyleMode `json:"mode"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GraphThresholdsStyleConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GraphThresholdsStyleConfig) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GraphThresholdsStyleConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GraphThresholdsStyleConfig` objects.
func (resource GraphThresholdsStyleConfig) Equals(other GraphThresholdsStyleConfig) bool {
	if resource.Mode != other.Mode {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GraphThresholdsStyleConfig` fields for violations and returns them.
func (resource GraphThresholdsStyleConfig) Validate() error {
	return nil
}

// TODO docs
type LegendPlacement string

const (
	LegendPlacementBottom LegendPlacement = "bottom"
	LegendPlacementRight  LegendPlacement = "right"
)

// TODO docs
// Note: "hidden" needs to remain as an option for plugins compatibility
type LegendDisplayMode string

const (
	LegendDisplayModeList   LegendDisplayMode = "list"
	LegendDisplayModeTable  LegendDisplayMode = "table"
	LegendDisplayModeHidden LegendDisplayMode = "hidden"
)

// TODO docs
type SingleStatBaseOptions struct {
	ReduceOptions ReduceDataOptions      `json:"reduceOptions"`
	Text          *VizTextDisplayOptions `json:"text,omitempty"`
	Orientation   VizOrientation         `json:"orientation"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SingleStatBaseOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SingleStatBaseOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "reduceOptions"
	if fields["reduceOptions"] != nil {
		if string(fields["reduceOptions"]) != "null" {

			resource.ReduceOptions = ReduceDataOptions{}
			if err := resource.ReduceOptions.UnmarshalJSONStrict(fields["reduceOptions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("reduceOptions", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("reduceOptions", errors.New("required field is null"))...)

		}
		delete(fields, "reduceOptions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("reduceOptions", errors.New("required field is missing from input"))...)
	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {

			resource.Text = &VizTextDisplayOptions{}
			if err := resource.Text.UnmarshalJSONStrict(fields["text"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}

		}
		delete(fields, "text")

	}
	// Field "orientation"
	if fields["orientation"] != nil {
		if string(fields["orientation"]) != "null" {
			if err := json.Unmarshal(fields["orientation"], &resource.Orientation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orientation", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("orientation", errors.New("required field is null"))...)

		}
		delete(fields, "orientation")
	} else {
		errs = append(errs, cog.MakeBuildErrors("orientation", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SingleStatBaseOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SingleStatBaseOptions` objects.
func (resource SingleStatBaseOptions) Equals(other SingleStatBaseOptions) bool {
	if !resource.ReduceOptions.Equals(other.ReduceOptions) {
		return false
	}
	if resource.Text == nil && other.Text != nil || resource.Text != nil && other.Text == nil {
		return false
	}

	if resource.Text != nil {
		if !resource.Text.Equals(*other.Text) {
			return false
		}
	}
	if resource.Orientation != other.Orientation {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `SingleStatBaseOptions` fields for violations and returns them.
func (resource SingleStatBaseOptions) Validate() error {
	var errs cog.BuildErrors
	if err := resource.ReduceOptions.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("reduceOptions", err)...)
	}
	if resource.Text != nil {
		if err := resource.Text.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("text", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type ReduceDataOptions struct {
	// If true show each row value
	Values *bool `json:"values,omitempty"`
	// if showing all values limit
	Limit *float64 `json:"limit,omitempty"`
	// When !values, pick one value for the whole field
	Calcs []string `json:"calcs"`
	// Which fields to show.  By default this is only numeric fields
	Fields *string `json:"fields,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ReduceDataOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ReduceDataOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "limit"
	if fields["limit"] != nil {
		if string(fields["limit"]) != "null" {
			if err := json.Unmarshal(fields["limit"], &resource.Limit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("limit", err)...)
			}

		}
		delete(fields, "limit")

	}
	// Field "calcs"
	if fields["calcs"] != nil {
		if string(fields["calcs"]) != "null" {

			if err := json.Unmarshal(fields["calcs"], &resource.Calcs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("calcs", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("calcs", errors.New("required field is null"))...)

		}
		delete(fields, "calcs")
	} else {
		errs = append(errs, cog.MakeBuildErrors("calcs", errors.New("required field is missing from input"))...)
	}
	// Field "fields"
	if fields["fields"] != nil {
		if string(fields["fields"]) != "null" {
			if err := json.Unmarshal(fields["fields"], &resource.Fields); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fields", err)...)
			}

		}
		delete(fields, "fields")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ReduceDataOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ReduceDataOptions` objects.
func (resource ReduceDataOptions) Equals(other ReduceDataOptions) bool {
	if resource.Values == nil && other.Values != nil || resource.Values != nil && other.Values == nil {
		return false
	}

	if resource.Values != nil {
		if *resource.Values != *other.Values {
			return false
		}
	}
	if resource.Limit == nil && other.Limit != nil || resource.Limit != nil && other.Limit == nil {
		return false
	}

	if resource.Limit != nil {
		if *resource.Limit != *other.Limit {
			return false
		}
	}

	if len(resource.Calcs) != len(other.Calcs) {
		return false
	}

	for i1 := range resource.Calcs {
		if resource.Calcs[i1] != other.Calcs[i1] {
			return false
		}
	}
	if resource.Fields == nil && other.Fields != nil || resource.Fields != nil && other.Fields == nil {
		return false
	}

	if resource.Fields != nil {
		if *resource.Fields != *other.Fields {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ReduceDataOptions` fields for violations and returns them.
func (resource ReduceDataOptions) Validate() error {
	return nil
}

// TODO docs
type VizOrientation string

const (
	VizOrientationAuto       VizOrientation = "auto"
	VizOrientationVertical   VizOrientation = "vertical"
	VizOrientationHorizontal VizOrientation = "horizontal"
)

// TODO docs
type OptionsWithTooltip struct {
	Tooltip VizTooltipOptions `json:"tooltip"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithTooltip` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *OptionsWithTooltip) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "tooltip"
	if fields["tooltip"] != nil {
		if string(fields["tooltip"]) != "null" {

			resource.Tooltip = VizTooltipOptions{}
			if err := resource.Tooltip.UnmarshalJSONStrict(fields["tooltip"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is null"))...)

		}
		delete(fields, "tooltip")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("OptionsWithTooltip", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `OptionsWithTooltip` objects.
func (resource OptionsWithTooltip) Equals(other OptionsWithTooltip) bool {
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `OptionsWithTooltip` fields for violations and returns them.
func (resource OptionsWithTooltip) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type OptionsWithLegend struct {
	Legend VizLegendOptions `json:"legend"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithLegend` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *OptionsWithLegend) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {

			resource.Legend = VizLegendOptions{}
			if err := resource.Legend.UnmarshalJSONStrict(fields["legend"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is null"))...)

		}
		delete(fields, "legend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("OptionsWithLegend", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `OptionsWithLegend` objects.
func (resource OptionsWithLegend) Equals(other OptionsWithLegend) bool {
	if !resource.Legend.Equals(other.Legend) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `OptionsWithLegend` fields for violations and returns them.
func (resource OptionsWithLegend) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type OptionsWithTimezones struct {
	Timezone []TimeZone `json:"timezone,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithTimezones` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *OptionsWithTimezones) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("OptionsWithTimezones", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `OptionsWithTimezones` objects.
func (resource OptionsWithTimezones) Equals(other OptionsWithTimezones) bool {

	if len(resource.Timezone) != len(other.Timezone) {
		return false
	}

	for i1 := range resource.Timezone {
		if resource.Timezone[i1] != other.Timezone[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `OptionsWithTimezones` fields for violations and returns them.
func (resource OptionsWithTimezones) Validate() error {
	return nil
}

// TODO docs
type OptionsWithTextFormatting struct {
	Text *VizTextDisplayOptions `json:"text,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithTextFormatting` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *OptionsWithTextFormatting) UnmarshalJSONStrict(raw []byte) error {
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

			resource.Text = &VizTextDisplayOptions{}
			if err := resource.Text.UnmarshalJSONStrict(fields["text"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}

		}
		delete(fields, "text")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("OptionsWithTextFormatting", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `OptionsWithTextFormatting` objects.
func (resource OptionsWithTextFormatting) Equals(other OptionsWithTextFormatting) bool {
	if resource.Text == nil && other.Text != nil || resource.Text != nil && other.Text == nil {
		return false
	}

	if resource.Text != nil {
		if !resource.Text.Equals(*other.Text) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `OptionsWithTextFormatting` fields for violations and returns them.
func (resource OptionsWithTextFormatting) Validate() error {
	var errs cog.BuildErrors
	if resource.Text != nil {
		if err := resource.Text.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("text", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type BigValueColorMode string

const (
	BigValueColorModeValue           BigValueColorMode = "value"
	BigValueColorModeBackground      BigValueColorMode = "background"
	BigValueColorModeBackgroundSolid BigValueColorMode = "background_solid"
	BigValueColorModeNone            BigValueColorMode = "none"
)

// TODO docs
type BigValueGraphMode string

const (
	BigValueGraphModeNone BigValueGraphMode = "none"
	BigValueGraphModeLine BigValueGraphMode = "line"
	BigValueGraphModeArea BigValueGraphMode = "area"
)

// TODO docs
type BigValueJustifyMode string

const (
	BigValueJustifyModeAuto   BigValueJustifyMode = "auto"
	BigValueJustifyModeCenter BigValueJustifyMode = "center"
)

// TODO docs
type BigValueTextMode string

const (
	BigValueTextModeAuto         BigValueTextMode = "auto"
	BigValueTextModeValue        BigValueTextMode = "value"
	BigValueTextModeValueAndName BigValueTextMode = "value_and_name"
	BigValueTextModeName         BigValueTextMode = "name"
	BigValueTextModeNone         BigValueTextMode = "none"
)

// TODO -- should not be table specific!
// TODO docs
type FieldTextAlignment string

const (
	FieldTextAlignmentAuto   FieldTextAlignment = "auto"
	FieldTextAlignmentLeft   FieldTextAlignment = "left"
	FieldTextAlignmentRight  FieldTextAlignment = "right"
	FieldTextAlignmentCenter FieldTextAlignment = "center"
)

// Controls the value alignment in the TimelineChart component
type TimelineValueAlignment string

const (
	TimelineValueAlignmentCenter TimelineValueAlignment = "center"
	TimelineValueAlignmentLeft   TimelineValueAlignment = "left"
	TimelineValueAlignmentRight  TimelineValueAlignment = "right"
)

// TODO docs
type VizTextDisplayOptions struct {
	// Explicit title text size
	TitleSize *float64 `json:"titleSize,omitempty"`
	// Explicit value text size
	ValueSize *float64 `json:"valueSize,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizTextDisplayOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *VizTextDisplayOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "titleSize"
	if fields["titleSize"] != nil {
		if string(fields["titleSize"]) != "null" {
			if err := json.Unmarshal(fields["titleSize"], &resource.TitleSize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("titleSize", err)...)
			}

		}
		delete(fields, "titleSize")

	}
	// Field "valueSize"
	if fields["valueSize"] != nil {
		if string(fields["valueSize"]) != "null" {
			if err := json.Unmarshal(fields["valueSize"], &resource.ValueSize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valueSize", err)...)
			}

		}
		delete(fields, "valueSize")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("VizTextDisplayOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `VizTextDisplayOptions` objects.
func (resource VizTextDisplayOptions) Equals(other VizTextDisplayOptions) bool {
	if resource.TitleSize == nil && other.TitleSize != nil || resource.TitleSize != nil && other.TitleSize == nil {
		return false
	}

	if resource.TitleSize != nil {
		if *resource.TitleSize != *other.TitleSize {
			return false
		}
	}
	if resource.ValueSize == nil && other.ValueSize != nil || resource.ValueSize != nil && other.ValueSize == nil {
		return false
	}

	if resource.ValueSize != nil {
		if *resource.ValueSize != *other.ValueSize {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `VizTextDisplayOptions` fields for violations and returns them.
func (resource VizTextDisplayOptions) Validate() error {
	return nil
}

// TODO docs
type TooltipDisplayMode string

const (
	TooltipDisplayModeSingle TooltipDisplayMode = "single"
	TooltipDisplayModeMulti  TooltipDisplayMode = "multi"
	TooltipDisplayModeNone   TooltipDisplayMode = "none"
)

// TODO docs
type SortOrder string

const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
	SortOrderNone       SortOrder = "none"
)

// TODO docs
type GraphFieldConfig struct {
	DrawStyle         *GraphDrawStyle             `json:"drawStyle,omitempty"`
	GradientMode      *GraphGradientMode          `json:"gradientMode,omitempty"`
	ThresholdsStyle   *GraphThresholdsStyleConfig `json:"thresholdsStyle,omitempty"`
	LineColor         *string                     `json:"lineColor,omitempty"`
	LineWidth         *float64                    `json:"lineWidth,omitempty"`
	LineInterpolation *LineInterpolation          `json:"lineInterpolation,omitempty"`
	LineStyle         *LineStyle                  `json:"lineStyle,omitempty"`
	FillColor         *string                     `json:"fillColor,omitempty"`
	FillOpacity       *float64                    `json:"fillOpacity,omitempty"`
	ShowPoints        *VisibilityMode             `json:"showPoints,omitempty"`
	PointSize         *float64                    `json:"pointSize,omitempty"`
	PointColor        *string                     `json:"pointColor,omitempty"`
	AxisPlacement     *AxisPlacement              `json:"axisPlacement,omitempty"`
	AxisColorMode     *AxisColorMode              `json:"axisColorMode,omitempty"`
	AxisLabel         *string                     `json:"axisLabel,omitempty"`
	AxisWidth         *float64                    `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                    `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                    `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                       `json:"axisGridShow,omitempty"`
	ScaleDistribution *ScaleDistributionConfig    `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                       `json:"axisCenteredZero,omitempty"`
	BarAlignment      *BarAlignment               `json:"barAlignment,omitempty"`
	BarWidthFactor    *float64                    `json:"barWidthFactor,omitempty"`
	Stacking          *StackingConfig             `json:"stacking,omitempty"`
	HideFrom          *HideSeriesConfig           `json:"hideFrom,omitempty"`
	Transform         *GraphTransform             `json:"transform,omitempty"`
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls      *BoolOrFloat64 `json:"spanNulls,omitempty"`
	FillBelowTo    *string        `json:"fillBelowTo,omitempty"`
	PointSymbol    *string        `json:"pointSymbol,omitempty"`
	AxisBorderShow *bool          `json:"axisBorderShow,omitempty"`
	BarMaxWidth    *float64       `json:"barMaxWidth,omitempty"`
	InsertNulls    *BoolOrUint32  `json:"insertNulls,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GraphFieldConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GraphFieldConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "drawStyle"
	if fields["drawStyle"] != nil {
		if string(fields["drawStyle"]) != "null" {
			if err := json.Unmarshal(fields["drawStyle"], &resource.DrawStyle); err != nil {
				errs = append(errs, cog.MakeBuildErrors("drawStyle", err)...)
			}

		}
		delete(fields, "drawStyle")

	}
	// Field "gradientMode"
	if fields["gradientMode"] != nil {
		if string(fields["gradientMode"]) != "null" {
			if err := json.Unmarshal(fields["gradientMode"], &resource.GradientMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("gradientMode", err)...)
			}

		}
		delete(fields, "gradientMode")

	}
	// Field "thresholdsStyle"
	if fields["thresholdsStyle"] != nil {
		if string(fields["thresholdsStyle"]) != "null" {

			resource.ThresholdsStyle = &GraphThresholdsStyleConfig{}
			if err := resource.ThresholdsStyle.UnmarshalJSONStrict(fields["thresholdsStyle"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("thresholdsStyle", err)...)
			}

		}
		delete(fields, "thresholdsStyle")

	}
	// Field "lineColor"
	if fields["lineColor"] != nil {
		if string(fields["lineColor"]) != "null" {
			if err := json.Unmarshal(fields["lineColor"], &resource.LineColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineColor", err)...)
			}

		}
		delete(fields, "lineColor")

	}
	// Field "lineWidth"
	if fields["lineWidth"] != nil {
		if string(fields["lineWidth"]) != "null" {
			if err := json.Unmarshal(fields["lineWidth"], &resource.LineWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineWidth", err)...)
			}

		}
		delete(fields, "lineWidth")

	}
	// Field "lineInterpolation"
	if fields["lineInterpolation"] != nil {
		if string(fields["lineInterpolation"]) != "null" {
			if err := json.Unmarshal(fields["lineInterpolation"], &resource.LineInterpolation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineInterpolation", err)...)
			}

		}
		delete(fields, "lineInterpolation")

	}
	// Field "lineStyle"
	if fields["lineStyle"] != nil {
		if string(fields["lineStyle"]) != "null" {

			resource.LineStyle = &LineStyle{}
			if err := resource.LineStyle.UnmarshalJSONStrict(fields["lineStyle"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
			}

		}
		delete(fields, "lineStyle")

	}
	// Field "fillColor"
	if fields["fillColor"] != nil {
		if string(fields["fillColor"]) != "null" {
			if err := json.Unmarshal(fields["fillColor"], &resource.FillColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillColor", err)...)
			}

		}
		delete(fields, "fillColor")

	}
	// Field "fillOpacity"
	if fields["fillOpacity"] != nil {
		if string(fields["fillOpacity"]) != "null" {
			if err := json.Unmarshal(fields["fillOpacity"], &resource.FillOpacity); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillOpacity", err)...)
			}

		}
		delete(fields, "fillOpacity")

	}
	// Field "showPoints"
	if fields["showPoints"] != nil {
		if string(fields["showPoints"]) != "null" {
			if err := json.Unmarshal(fields["showPoints"], &resource.ShowPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showPoints", err)...)
			}

		}
		delete(fields, "showPoints")

	}
	// Field "pointSize"
	if fields["pointSize"] != nil {
		if string(fields["pointSize"]) != "null" {
			if err := json.Unmarshal(fields["pointSize"], &resource.PointSize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
			}

		}
		delete(fields, "pointSize")

	}
	// Field "pointColor"
	if fields["pointColor"] != nil {
		if string(fields["pointColor"]) != "null" {
			if err := json.Unmarshal(fields["pointColor"], &resource.PointColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointColor", err)...)
			}

		}
		delete(fields, "pointColor")

	}
	// Field "axisPlacement"
	if fields["axisPlacement"] != nil {
		if string(fields["axisPlacement"]) != "null" {
			if err := json.Unmarshal(fields["axisPlacement"], &resource.AxisPlacement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisPlacement", err)...)
			}

		}
		delete(fields, "axisPlacement")

	}
	// Field "axisColorMode"
	if fields["axisColorMode"] != nil {
		if string(fields["axisColorMode"]) != "null" {
			if err := json.Unmarshal(fields["axisColorMode"], &resource.AxisColorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisColorMode", err)...)
			}

		}
		delete(fields, "axisColorMode")

	}
	// Field "axisLabel"
	if fields["axisLabel"] != nil {
		if string(fields["axisLabel"]) != "null" {
			if err := json.Unmarshal(fields["axisLabel"], &resource.AxisLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisLabel", err)...)
			}

		}
		delete(fields, "axisLabel")

	}
	// Field "axisWidth"
	if fields["axisWidth"] != nil {
		if string(fields["axisWidth"]) != "null" {
			if err := json.Unmarshal(fields["axisWidth"], &resource.AxisWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisWidth", err)...)
			}

		}
		delete(fields, "axisWidth")

	}
	// Field "axisSoftMin"
	if fields["axisSoftMin"] != nil {
		if string(fields["axisSoftMin"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMin"], &resource.AxisSoftMin); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMin", err)...)
			}

		}
		delete(fields, "axisSoftMin")

	}
	// Field "axisSoftMax"
	if fields["axisSoftMax"] != nil {
		if string(fields["axisSoftMax"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMax"], &resource.AxisSoftMax); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMax", err)...)
			}

		}
		delete(fields, "axisSoftMax")

	}
	// Field "axisGridShow"
	if fields["axisGridShow"] != nil {
		if string(fields["axisGridShow"]) != "null" {
			if err := json.Unmarshal(fields["axisGridShow"], &resource.AxisGridShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisGridShow", err)...)
			}

		}
		delete(fields, "axisGridShow")

	}
	// Field "scaleDistribution"
	if fields["scaleDistribution"] != nil {
		if string(fields["scaleDistribution"]) != "null" {

			resource.ScaleDistribution = &ScaleDistributionConfig{}
			if err := resource.ScaleDistribution.UnmarshalJSONStrict(fields["scaleDistribution"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
			}

		}
		delete(fields, "scaleDistribution")

	}
	// Field "axisCenteredZero"
	if fields["axisCenteredZero"] != nil {
		if string(fields["axisCenteredZero"]) != "null" {
			if err := json.Unmarshal(fields["axisCenteredZero"], &resource.AxisCenteredZero); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisCenteredZero", err)...)
			}

		}
		delete(fields, "axisCenteredZero")

	}
	// Field "barAlignment"
	if fields["barAlignment"] != nil {
		if string(fields["barAlignment"]) != "null" {
			if err := json.Unmarshal(fields["barAlignment"], &resource.BarAlignment); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barAlignment", err)...)
			}

		}
		delete(fields, "barAlignment")

	}
	// Field "barWidthFactor"
	if fields["barWidthFactor"] != nil {
		if string(fields["barWidthFactor"]) != "null" {
			if err := json.Unmarshal(fields["barWidthFactor"], &resource.BarWidthFactor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barWidthFactor", err)...)
			}

		}
		delete(fields, "barWidthFactor")

	}
	// Field "stacking"
	if fields["stacking"] != nil {
		if string(fields["stacking"]) != "null" {

			resource.Stacking = &StackingConfig{}
			if err := resource.Stacking.UnmarshalJSONStrict(fields["stacking"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stacking", err)...)
			}

		}
		delete(fields, "stacking")

	}
	// Field "hideFrom"
	if fields["hideFrom"] != nil {
		if string(fields["hideFrom"]) != "null" {

			resource.HideFrom = &HideSeriesConfig{}
			if err := resource.HideFrom.UnmarshalJSONStrict(fields["hideFrom"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
			}

		}
		delete(fields, "hideFrom")

	}
	// Field "transform"
	if fields["transform"] != nil {
		if string(fields["transform"]) != "null" {
			if err := json.Unmarshal(fields["transform"], &resource.Transform); err != nil {
				errs = append(errs, cog.MakeBuildErrors("transform", err)...)
			}

		}
		delete(fields, "transform")

	}
	// Field "spanNulls"
	if fields["spanNulls"] != nil {
		if string(fields["spanNulls"]) != "null" {

			resource.SpanNulls = &BoolOrFloat64{}
			if err := resource.SpanNulls.UnmarshalJSONStrict(fields["spanNulls"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spanNulls", err)...)
			}

		}
		delete(fields, "spanNulls")

	}
	// Field "fillBelowTo"
	if fields["fillBelowTo"] != nil {
		if string(fields["fillBelowTo"]) != "null" {
			if err := json.Unmarshal(fields["fillBelowTo"], &resource.FillBelowTo); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillBelowTo", err)...)
			}

		}
		delete(fields, "fillBelowTo")

	}
	// Field "pointSymbol"
	if fields["pointSymbol"] != nil {
		if string(fields["pointSymbol"]) != "null" {
			if err := json.Unmarshal(fields["pointSymbol"], &resource.PointSymbol); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSymbol", err)...)
			}

		}
		delete(fields, "pointSymbol")

	}
	// Field "axisBorderShow"
	if fields["axisBorderShow"] != nil {
		if string(fields["axisBorderShow"]) != "null" {
			if err := json.Unmarshal(fields["axisBorderShow"], &resource.AxisBorderShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisBorderShow", err)...)
			}

		}
		delete(fields, "axisBorderShow")

	}
	// Field "barMaxWidth"
	if fields["barMaxWidth"] != nil {
		if string(fields["barMaxWidth"]) != "null" {
			if err := json.Unmarshal(fields["barMaxWidth"], &resource.BarMaxWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barMaxWidth", err)...)
			}

		}
		delete(fields, "barMaxWidth")

	}
	// Field "insertNulls"
	if fields["insertNulls"] != nil {
		if string(fields["insertNulls"]) != "null" {

			resource.InsertNulls = &BoolOrUint32{}
			if err := resource.InsertNulls.UnmarshalJSONStrict(fields["insertNulls"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("insertNulls", err)...)
			}

		}
		delete(fields, "insertNulls")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GraphFieldConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GraphFieldConfig` objects.
func (resource GraphFieldConfig) Equals(other GraphFieldConfig) bool {
	if resource.DrawStyle == nil && other.DrawStyle != nil || resource.DrawStyle != nil && other.DrawStyle == nil {
		return false
	}

	if resource.DrawStyle != nil {
		if *resource.DrawStyle != *other.DrawStyle {
			return false
		}
	}
	if resource.GradientMode == nil && other.GradientMode != nil || resource.GradientMode != nil && other.GradientMode == nil {
		return false
	}

	if resource.GradientMode != nil {
		if *resource.GradientMode != *other.GradientMode {
			return false
		}
	}
	if resource.ThresholdsStyle == nil && other.ThresholdsStyle != nil || resource.ThresholdsStyle != nil && other.ThresholdsStyle == nil {
		return false
	}

	if resource.ThresholdsStyle != nil {
		if !resource.ThresholdsStyle.Equals(*other.ThresholdsStyle) {
			return false
		}
	}
	if resource.LineColor == nil && other.LineColor != nil || resource.LineColor != nil && other.LineColor == nil {
		return false
	}

	if resource.LineColor != nil {
		if *resource.LineColor != *other.LineColor {
			return false
		}
	}
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
			return false
		}
	}
	if resource.LineInterpolation == nil && other.LineInterpolation != nil || resource.LineInterpolation != nil && other.LineInterpolation == nil {
		return false
	}

	if resource.LineInterpolation != nil {
		if *resource.LineInterpolation != *other.LineInterpolation {
			return false
		}
	}
	if resource.LineStyle == nil && other.LineStyle != nil || resource.LineStyle != nil && other.LineStyle == nil {
		return false
	}

	if resource.LineStyle != nil {
		if !resource.LineStyle.Equals(*other.LineStyle) {
			return false
		}
	}
	if resource.FillColor == nil && other.FillColor != nil || resource.FillColor != nil && other.FillColor == nil {
		return false
	}

	if resource.FillColor != nil {
		if *resource.FillColor != *other.FillColor {
			return false
		}
	}
	if resource.FillOpacity == nil && other.FillOpacity != nil || resource.FillOpacity != nil && other.FillOpacity == nil {
		return false
	}

	if resource.FillOpacity != nil {
		if *resource.FillOpacity != *other.FillOpacity {
			return false
		}
	}
	if resource.ShowPoints == nil && other.ShowPoints != nil || resource.ShowPoints != nil && other.ShowPoints == nil {
		return false
	}

	if resource.ShowPoints != nil {
		if *resource.ShowPoints != *other.ShowPoints {
			return false
		}
	}
	if resource.PointSize == nil && other.PointSize != nil || resource.PointSize != nil && other.PointSize == nil {
		return false
	}

	if resource.PointSize != nil {
		if *resource.PointSize != *other.PointSize {
			return false
		}
	}
	if resource.PointColor == nil && other.PointColor != nil || resource.PointColor != nil && other.PointColor == nil {
		return false
	}

	if resource.PointColor != nil {
		if *resource.PointColor != *other.PointColor {
			return false
		}
	}
	if resource.AxisPlacement == nil && other.AxisPlacement != nil || resource.AxisPlacement != nil && other.AxisPlacement == nil {
		return false
	}

	if resource.AxisPlacement != nil {
		if *resource.AxisPlacement != *other.AxisPlacement {
			return false
		}
	}
	if resource.AxisColorMode == nil && other.AxisColorMode != nil || resource.AxisColorMode != nil && other.AxisColorMode == nil {
		return false
	}

	if resource.AxisColorMode != nil {
		if *resource.AxisColorMode != *other.AxisColorMode {
			return false
		}
	}
	if resource.AxisLabel == nil && other.AxisLabel != nil || resource.AxisLabel != nil && other.AxisLabel == nil {
		return false
	}

	if resource.AxisLabel != nil {
		if *resource.AxisLabel != *other.AxisLabel {
			return false
		}
	}
	if resource.AxisWidth == nil && other.AxisWidth != nil || resource.AxisWidth != nil && other.AxisWidth == nil {
		return false
	}

	if resource.AxisWidth != nil {
		if *resource.AxisWidth != *other.AxisWidth {
			return false
		}
	}
	if resource.AxisSoftMin == nil && other.AxisSoftMin != nil || resource.AxisSoftMin != nil && other.AxisSoftMin == nil {
		return false
	}

	if resource.AxisSoftMin != nil {
		if *resource.AxisSoftMin != *other.AxisSoftMin {
			return false
		}
	}
	if resource.AxisSoftMax == nil && other.AxisSoftMax != nil || resource.AxisSoftMax != nil && other.AxisSoftMax == nil {
		return false
	}

	if resource.AxisSoftMax != nil {
		if *resource.AxisSoftMax != *other.AxisSoftMax {
			return false
		}
	}
	if resource.AxisGridShow == nil && other.AxisGridShow != nil || resource.AxisGridShow != nil && other.AxisGridShow == nil {
		return false
	}

	if resource.AxisGridShow != nil {
		if *resource.AxisGridShow != *other.AxisGridShow {
			return false
		}
	}
	if resource.ScaleDistribution == nil && other.ScaleDistribution != nil || resource.ScaleDistribution != nil && other.ScaleDistribution == nil {
		return false
	}

	if resource.ScaleDistribution != nil {
		if !resource.ScaleDistribution.Equals(*other.ScaleDistribution) {
			return false
		}
	}
	if resource.AxisCenteredZero == nil && other.AxisCenteredZero != nil || resource.AxisCenteredZero != nil && other.AxisCenteredZero == nil {
		return false
	}

	if resource.AxisCenteredZero != nil {
		if *resource.AxisCenteredZero != *other.AxisCenteredZero {
			return false
		}
	}
	if resource.BarAlignment == nil && other.BarAlignment != nil || resource.BarAlignment != nil && other.BarAlignment == nil {
		return false
	}

	if resource.BarAlignment != nil {
		if *resource.BarAlignment != *other.BarAlignment {
			return false
		}
	}
	if resource.BarWidthFactor == nil && other.BarWidthFactor != nil || resource.BarWidthFactor != nil && other.BarWidthFactor == nil {
		return false
	}

	if resource.BarWidthFactor != nil {
		if *resource.BarWidthFactor != *other.BarWidthFactor {
			return false
		}
	}
	if resource.Stacking == nil && other.Stacking != nil || resource.Stacking != nil && other.Stacking == nil {
		return false
	}

	if resource.Stacking != nil {
		if !resource.Stacking.Equals(*other.Stacking) {
			return false
		}
	}
	if resource.HideFrom == nil && other.HideFrom != nil || resource.HideFrom != nil && other.HideFrom == nil {
		return false
	}

	if resource.HideFrom != nil {
		if !resource.HideFrom.Equals(*other.HideFrom) {
			return false
		}
	}
	if resource.Transform == nil && other.Transform != nil || resource.Transform != nil && other.Transform == nil {
		return false
	}

	if resource.Transform != nil {
		if *resource.Transform != *other.Transform {
			return false
		}
	}
	if resource.SpanNulls == nil && other.SpanNulls != nil || resource.SpanNulls != nil && other.SpanNulls == nil {
		return false
	}

	if resource.SpanNulls != nil {
		if !resource.SpanNulls.Equals(*other.SpanNulls) {
			return false
		}
	}
	if resource.FillBelowTo == nil && other.FillBelowTo != nil || resource.FillBelowTo != nil && other.FillBelowTo == nil {
		return false
	}

	if resource.FillBelowTo != nil {
		if *resource.FillBelowTo != *other.FillBelowTo {
			return false
		}
	}
	if resource.PointSymbol == nil && other.PointSymbol != nil || resource.PointSymbol != nil && other.PointSymbol == nil {
		return false
	}

	if resource.PointSymbol != nil {
		if *resource.PointSymbol != *other.PointSymbol {
			return false
		}
	}
	if resource.AxisBorderShow == nil && other.AxisBorderShow != nil || resource.AxisBorderShow != nil && other.AxisBorderShow == nil {
		return false
	}

	if resource.AxisBorderShow != nil {
		if *resource.AxisBorderShow != *other.AxisBorderShow {
			return false
		}
	}
	if resource.BarMaxWidth == nil && other.BarMaxWidth != nil || resource.BarMaxWidth != nil && other.BarMaxWidth == nil {
		return false
	}

	if resource.BarMaxWidth != nil {
		if *resource.BarMaxWidth != *other.BarMaxWidth {
			return false
		}
	}
	if resource.InsertNulls == nil && other.InsertNulls != nil || resource.InsertNulls != nil && other.InsertNulls == nil {
		return false
	}

	if resource.InsertNulls != nil {
		if !resource.InsertNulls.Equals(*other.InsertNulls) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GraphFieldConfig` fields for violations and returns them.
func (resource GraphFieldConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.ThresholdsStyle != nil {
		if err := resource.ThresholdsStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("thresholdsStyle", err)...)
		}
	}
	if resource.LineStyle != nil {
		if err := resource.LineStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
		}
	}
	if resource.ScaleDistribution != nil {
		if err := resource.ScaleDistribution.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
		}
	}
	if resource.Stacking != nil {
		if err := resource.Stacking.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("stacking", err)...)
		}
	}
	if resource.HideFrom != nil {
		if err := resource.HideFrom.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
		}
	}
	if resource.SpanNulls != nil {
		if err := resource.SpanNulls.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("spanNulls", err)...)
		}
	}
	if resource.InsertNulls != nil {
		if err := resource.InsertNulls.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("insertNulls", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TODO docs
type VizLegendOptions struct {
	DisplayMode LegendDisplayMode `json:"displayMode"`
	Placement   LegendPlacement   `json:"placement"`
	ShowLegend  bool              `json:"showLegend"`
	AsTable     *bool             `json:"asTable,omitempty"`
	IsVisible   *bool             `json:"isVisible,omitempty"`
	SortBy      *string           `json:"sortBy,omitempty"`
	SortDesc    *bool             `json:"sortDesc,omitempty"`
	Width       *float64          `json:"width,omitempty"`
	Calcs       []string          `json:"calcs"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizLegendOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *VizLegendOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "displayMode"
	if fields["displayMode"] != nil {
		if string(fields["displayMode"]) != "null" {
			if err := json.Unmarshal(fields["displayMode"], &resource.DisplayMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("displayMode", errors.New("required field is null"))...)

		}
		delete(fields, "displayMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("displayMode", errors.New("required field is missing from input"))...)
	}
	// Field "placement"
	if fields["placement"] != nil {
		if string(fields["placement"]) != "null" {
			if err := json.Unmarshal(fields["placement"], &resource.Placement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("placement", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("placement", errors.New("required field is null"))...)

		}
		delete(fields, "placement")
	} else {
		errs = append(errs, cog.MakeBuildErrors("placement", errors.New("required field is missing from input"))...)
	}
	// Field "showLegend"
	if fields["showLegend"] != nil {
		if string(fields["showLegend"]) != "null" {
			if err := json.Unmarshal(fields["showLegend"], &resource.ShowLegend); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showLegend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showLegend", errors.New("required field is null"))...)

		}
		delete(fields, "showLegend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showLegend", errors.New("required field is missing from input"))...)
	}
	// Field "asTable"
	if fields["asTable"] != nil {
		if string(fields["asTable"]) != "null" {
			if err := json.Unmarshal(fields["asTable"], &resource.AsTable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("asTable", err)...)
			}

		}
		delete(fields, "asTable")

	}
	// Field "isVisible"
	if fields["isVisible"] != nil {
		if string(fields["isVisible"]) != "null" {
			if err := json.Unmarshal(fields["isVisible"], &resource.IsVisible); err != nil {
				errs = append(errs, cog.MakeBuildErrors("isVisible", err)...)
			}

		}
		delete(fields, "isVisible")

	}
	// Field "sortBy"
	if fields["sortBy"] != nil {
		if string(fields["sortBy"]) != "null" {
			if err := json.Unmarshal(fields["sortBy"], &resource.SortBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sortBy", err)...)
			}

		}
		delete(fields, "sortBy")

	}
	// Field "sortDesc"
	if fields["sortDesc"] != nil {
		if string(fields["sortDesc"]) != "null" {
			if err := json.Unmarshal(fields["sortDesc"], &resource.SortDesc); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sortDesc", err)...)
			}

		}
		delete(fields, "sortDesc")

	}
	// Field "width"
	if fields["width"] != nil {
		if string(fields["width"]) != "null" {
			if err := json.Unmarshal(fields["width"], &resource.Width); err != nil {
				errs = append(errs, cog.MakeBuildErrors("width", err)...)
			}

		}
		delete(fields, "width")

	}
	// Field "calcs"
	if fields["calcs"] != nil {
		if string(fields["calcs"]) != "null" {

			if err := json.Unmarshal(fields["calcs"], &resource.Calcs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("calcs", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("calcs", errors.New("required field is null"))...)

		}
		delete(fields, "calcs")
	} else {
		errs = append(errs, cog.MakeBuildErrors("calcs", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("VizLegendOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `VizLegendOptions` objects.
func (resource VizLegendOptions) Equals(other VizLegendOptions) bool {
	if resource.DisplayMode != other.DisplayMode {
		return false
	}
	if resource.Placement != other.Placement {
		return false
	}
	if resource.ShowLegend != other.ShowLegend {
		return false
	}
	if resource.AsTable == nil && other.AsTable != nil || resource.AsTable != nil && other.AsTable == nil {
		return false
	}

	if resource.AsTable != nil {
		if *resource.AsTable != *other.AsTable {
			return false
		}
	}
	if resource.IsVisible == nil && other.IsVisible != nil || resource.IsVisible != nil && other.IsVisible == nil {
		return false
	}

	if resource.IsVisible != nil {
		if *resource.IsVisible != *other.IsVisible {
			return false
		}
	}
	if resource.SortBy == nil && other.SortBy != nil || resource.SortBy != nil && other.SortBy == nil {
		return false
	}

	if resource.SortBy != nil {
		if *resource.SortBy != *other.SortBy {
			return false
		}
	}
	if resource.SortDesc == nil && other.SortDesc != nil || resource.SortDesc != nil && other.SortDesc == nil {
		return false
	}

	if resource.SortDesc != nil {
		if *resource.SortDesc != *other.SortDesc {
			return false
		}
	}
	if resource.Width == nil && other.Width != nil || resource.Width != nil && other.Width == nil {
		return false
	}

	if resource.Width != nil {
		if *resource.Width != *other.Width {
			return false
		}
	}

	if len(resource.Calcs) != len(other.Calcs) {
		return false
	}

	for i1 := range resource.Calcs {
		if resource.Calcs[i1] != other.Calcs[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `VizLegendOptions` fields for violations and returns them.
func (resource VizLegendOptions) Validate() error {
	return nil
}

// Enum expressing the possible display modes
// for the bar gauge component of Grafana UI
type BarGaugeDisplayMode string

const (
	BarGaugeDisplayModeBasic    BarGaugeDisplayMode = "basic"
	BarGaugeDisplayModeLcd      BarGaugeDisplayMode = "lcd"
	BarGaugeDisplayModeGradient BarGaugeDisplayMode = "gradient"
)

// Allows for the table cell gauge display type to set the gauge mode.
type BarGaugeValueMode string

const (
	BarGaugeValueModeColor  BarGaugeValueMode = "color"
	BarGaugeValueModeText   BarGaugeValueMode = "text"
	BarGaugeValueModeHidden BarGaugeValueMode = "hidden"
)

// Allows for the bar gauge name to be placed explicitly
type BarGaugeNamePlacement string

const (
	BarGaugeNamePlacementAuto BarGaugeNamePlacement = "auto"
	BarGaugeNamePlacementTop  BarGaugeNamePlacement = "top"
	BarGaugeNamePlacementLeft BarGaugeNamePlacement = "left"
)

// Allows for the bar gauge size to be set explicitly
type BarGaugeSizing string

const (
	BarGaugeSizingAuto   BarGaugeSizing = "auto"
	BarGaugeSizingManual BarGaugeSizing = "manual"
)

// TODO docs
type VizTooltipOptions struct {
	Mode TooltipDisplayMode `json:"mode"`
	Sort SortOrder          `json:"sort"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizTooltipOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *VizTooltipOptions) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("VizTooltipOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `VizTooltipOptions` objects.
func (resource VizTooltipOptions) Equals(other VizTooltipOptions) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Sort != other.Sort {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `VizTooltipOptions` fields for violations and returns them.
func (resource VizTooltipOptions) Validate() error {
	return nil
}

type Labels map[string]string

// Internally, this is the "type" of cell that's being displayed
// in the table such as colored text, JSON, gauge, etc.
// The color-background-solid, gradient-gauge, and lcd-gauge
// modes are deprecated in favor of new cell subOptions
type TableCellDisplayMode string

const (
	TableCellDisplayModeAuto                 TableCellDisplayMode = "auto"
	TableCellDisplayModeColorText            TableCellDisplayMode = "color-text"
	TableCellDisplayModeColorBackground      TableCellDisplayMode = "color-background"
	TableCellDisplayModeColorBackgroundSolid TableCellDisplayMode = "color-background-solid"
	TableCellDisplayModeGradientGauge        TableCellDisplayMode = "gradient-gauge"
	TableCellDisplayModeLcdGauge             TableCellDisplayMode = "lcd-gauge"
	TableCellDisplayModeJSONView             TableCellDisplayMode = "json-view"
	TableCellDisplayModeBasicGauge           TableCellDisplayMode = "basic"
	TableCellDisplayModeImage                TableCellDisplayMode = "image"
	TableCellDisplayModeGauge                TableCellDisplayMode = "gauge"
	TableCellDisplayModeSparkline            TableCellDisplayMode = "sparkline"
	TableCellDisplayModeCustom               TableCellDisplayMode = "custom"
)

// Display mode to the "Colored Background" display
// mode for table cells. Either displays a solid color (basic mode)
// or a gradient.
type TableCellBackgroundDisplayMode string

const (
	TableCellBackgroundDisplayModeBasic    TableCellBackgroundDisplayMode = "basic"
	TableCellBackgroundDisplayModeGradient TableCellBackgroundDisplayMode = "gradient"
)

// Sort by field state
type TableSortByFieldState struct {
	// Sets the display name of the field to sort by
	DisplayName string `json:"displayName"`
	// Flag used to indicate descending sort order
	Desc *bool `json:"desc,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableSortByFieldState` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableSortByFieldState) UnmarshalJSONStrict(raw []byte) error {
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
		} else {
			errs = append(errs, cog.MakeBuildErrors("displayName", errors.New("required field is null"))...)

		}
		delete(fields, "displayName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("displayName", errors.New("required field is missing from input"))...)
	}
	// Field "desc"
	if fields["desc"] != nil {
		if string(fields["desc"]) != "null" {
			if err := json.Unmarshal(fields["desc"], &resource.Desc); err != nil {
				errs = append(errs, cog.MakeBuildErrors("desc", err)...)
			}

		}
		delete(fields, "desc")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableSortByFieldState", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableSortByFieldState` objects.
func (resource TableSortByFieldState) Equals(other TableSortByFieldState) bool {
	if resource.DisplayName != other.DisplayName {
		return false
	}
	if resource.Desc == nil && other.Desc != nil || resource.Desc != nil && other.Desc == nil {
		return false
	}

	if resource.Desc != nil {
		if *resource.Desc != *other.Desc {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableSortByFieldState` fields for violations and returns them.
func (resource TableSortByFieldState) Validate() error {
	return nil
}

// Footer options
type TableFooterOptions struct {
	Show bool `json:"show"`
	// actually 1 value
	Reducer          []string `json:"reducer"`
	Fields           []string `json:"fields,omitempty"`
	EnablePagination *bool    `json:"enablePagination,omitempty"`
	CountRows        *bool    `json:"countRows,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableFooterOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableFooterOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "show"
	if fields["show"] != nil {
		if string(fields["show"]) != "null" {
			if err := json.Unmarshal(fields["show"], &resource.Show); err != nil {
				errs = append(errs, cog.MakeBuildErrors("show", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("show", errors.New("required field is null"))...)

		}
		delete(fields, "show")
	} else {
		errs = append(errs, cog.MakeBuildErrors("show", errors.New("required field is missing from input"))...)
	}
	// Field "reducer"
	if fields["reducer"] != nil {
		if string(fields["reducer"]) != "null" {

			if err := json.Unmarshal(fields["reducer"], &resource.Reducer); err != nil {
				errs = append(errs, cog.MakeBuildErrors("reducer", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("reducer", errors.New("required field is null"))...)

		}
		delete(fields, "reducer")
	} else {
		errs = append(errs, cog.MakeBuildErrors("reducer", errors.New("required field is missing from input"))...)
	}
	// Field "fields"
	if fields["fields"] != nil {
		if string(fields["fields"]) != "null" {

			if err := json.Unmarshal(fields["fields"], &resource.Fields); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fields", err)...)
			}

		}
		delete(fields, "fields")

	}
	// Field "enablePagination"
	if fields["enablePagination"] != nil {
		if string(fields["enablePagination"]) != "null" {
			if err := json.Unmarshal(fields["enablePagination"], &resource.EnablePagination); err != nil {
				errs = append(errs, cog.MakeBuildErrors("enablePagination", err)...)
			}

		}
		delete(fields, "enablePagination")

	}
	// Field "countRows"
	if fields["countRows"] != nil {
		if string(fields["countRows"]) != "null" {
			if err := json.Unmarshal(fields["countRows"], &resource.CountRows); err != nil {
				errs = append(errs, cog.MakeBuildErrors("countRows", err)...)
			}

		}
		delete(fields, "countRows")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableFooterOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableFooterOptions` objects.
func (resource TableFooterOptions) Equals(other TableFooterOptions) bool {
	if resource.Show != other.Show {
		return false
	}

	if len(resource.Reducer) != len(other.Reducer) {
		return false
	}

	for i1 := range resource.Reducer {
		if resource.Reducer[i1] != other.Reducer[i1] {
			return false
		}
	}

	if len(resource.Fields) != len(other.Fields) {
		return false
	}

	for i1 := range resource.Fields {
		if resource.Fields[i1] != other.Fields[i1] {
			return false
		}
	}
	if resource.EnablePagination == nil && other.EnablePagination != nil || resource.EnablePagination != nil && other.EnablePagination == nil {
		return false
	}

	if resource.EnablePagination != nil {
		if *resource.EnablePagination != *other.EnablePagination {
			return false
		}
	}
	if resource.CountRows == nil && other.CountRows != nil || resource.CountRows != nil && other.CountRows == nil {
		return false
	}

	if resource.CountRows != nil {
		if *resource.CountRows != *other.CountRows {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableFooterOptions` fields for violations and returns them.
func (resource TableFooterOptions) Validate() error {
	return nil
}

// Auto mode table cell options
type TableAutoCellOptions struct {
	Type string `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableAutoCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableAutoCellOptions) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableAutoCellOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableAutoCellOptions` objects.
func (resource TableAutoCellOptions) Equals(other TableAutoCellOptions) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableAutoCellOptions` fields for violations and returns them.
func (resource TableAutoCellOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "auto") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == auto"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Colored text cell options
type TableColorTextCellOptions struct {
	Type string `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableColorTextCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableColorTextCellOptions) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableColorTextCellOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableColorTextCellOptions` objects.
func (resource TableColorTextCellOptions) Equals(other TableColorTextCellOptions) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableColorTextCellOptions` fields for violations and returns them.
func (resource TableColorTextCellOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "color-text") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == color-text"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Json view cell options
type TableJsonViewCellOptions struct {
	Type string `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableJsonViewCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableJsonViewCellOptions) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableJsonViewCellOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableJsonViewCellOptions` objects.
func (resource TableJsonViewCellOptions) Equals(other TableJsonViewCellOptions) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableJsonViewCellOptions` fields for violations and returns them.
func (resource TableJsonViewCellOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "json-view") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == json-view"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Json view cell options
type TableImageCellOptions struct {
	Type string `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableImageCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableImageCellOptions) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableImageCellOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableImageCellOptions` objects.
func (resource TableImageCellOptions) Equals(other TableImageCellOptions) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableImageCellOptions` fields for violations and returns them.
func (resource TableImageCellOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "image") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == image"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Gauge cell options
type TableBarGaugeCellOptions struct {
	Type             string               `json:"type"`
	Mode             *BarGaugeDisplayMode `json:"mode,omitempty"`
	ValueDisplayMode *BarGaugeValueMode   `json:"valueDisplayMode,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableBarGaugeCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableBarGaugeCellOptions) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}

		}
		delete(fields, "mode")

	}
	// Field "valueDisplayMode"
	if fields["valueDisplayMode"] != nil {
		if string(fields["valueDisplayMode"]) != "null" {
			if err := json.Unmarshal(fields["valueDisplayMode"], &resource.ValueDisplayMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valueDisplayMode", err)...)
			}

		}
		delete(fields, "valueDisplayMode")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableBarGaugeCellOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableBarGaugeCellOptions` objects.
func (resource TableBarGaugeCellOptions) Equals(other TableBarGaugeCellOptions) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
			return false
		}
	}
	if resource.ValueDisplayMode == nil && other.ValueDisplayMode != nil || resource.ValueDisplayMode != nil && other.ValueDisplayMode == nil {
		return false
	}

	if resource.ValueDisplayMode != nil {
		if *resource.ValueDisplayMode != *other.ValueDisplayMode {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableBarGaugeCellOptions` fields for violations and returns them.
func (resource TableBarGaugeCellOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "gauge") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == gauge"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Sparkline cell options
type TableSparklineCellOptions struct {
	Type              string                      `json:"type"`
	DrawStyle         *GraphDrawStyle             `json:"drawStyle,omitempty"`
	GradientMode      *GraphGradientMode          `json:"gradientMode,omitempty"`
	ThresholdsStyle   *GraphThresholdsStyleConfig `json:"thresholdsStyle,omitempty"`
	LineColor         *string                     `json:"lineColor,omitempty"`
	LineWidth         *float64                    `json:"lineWidth,omitempty"`
	LineInterpolation *LineInterpolation          `json:"lineInterpolation,omitempty"`
	LineStyle         *LineStyle                  `json:"lineStyle,omitempty"`
	FillColor         *string                     `json:"fillColor,omitempty"`
	FillOpacity       *float64                    `json:"fillOpacity,omitempty"`
	ShowPoints        *VisibilityMode             `json:"showPoints,omitempty"`
	PointSize         *float64                    `json:"pointSize,omitempty"`
	PointColor        *string                     `json:"pointColor,omitempty"`
	AxisPlacement     *AxisPlacement              `json:"axisPlacement,omitempty"`
	AxisColorMode     *AxisColorMode              `json:"axisColorMode,omitempty"`
	AxisLabel         *string                     `json:"axisLabel,omitempty"`
	AxisWidth         *float64                    `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                    `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                    `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                       `json:"axisGridShow,omitempty"`
	ScaleDistribution *ScaleDistributionConfig    `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                       `json:"axisCenteredZero,omitempty"`
	BarAlignment      *BarAlignment               `json:"barAlignment,omitempty"`
	BarWidthFactor    *float64                    `json:"barWidthFactor,omitempty"`
	Stacking          *StackingConfig             `json:"stacking,omitempty"`
	HideFrom          *HideSeriesConfig           `json:"hideFrom,omitempty"`
	HideValue         *bool                       `json:"hideValue,omitempty"`
	Transform         *GraphTransform             `json:"transform,omitempty"`
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls      *BoolOrFloat64 `json:"spanNulls,omitempty"`
	FillBelowTo    *string        `json:"fillBelowTo,omitempty"`
	PointSymbol    *string        `json:"pointSymbol,omitempty"`
	AxisBorderShow *bool          `json:"axisBorderShow,omitempty"`
	BarMaxWidth    *float64       `json:"barMaxWidth,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableSparklineCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableSparklineCellOptions) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "drawStyle"
	if fields["drawStyle"] != nil {
		if string(fields["drawStyle"]) != "null" {
			if err := json.Unmarshal(fields["drawStyle"], &resource.DrawStyle); err != nil {
				errs = append(errs, cog.MakeBuildErrors("drawStyle", err)...)
			}

		}
		delete(fields, "drawStyle")

	}
	// Field "gradientMode"
	if fields["gradientMode"] != nil {
		if string(fields["gradientMode"]) != "null" {
			if err := json.Unmarshal(fields["gradientMode"], &resource.GradientMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("gradientMode", err)...)
			}

		}
		delete(fields, "gradientMode")

	}
	// Field "thresholdsStyle"
	if fields["thresholdsStyle"] != nil {
		if string(fields["thresholdsStyle"]) != "null" {

			resource.ThresholdsStyle = &GraphThresholdsStyleConfig{}
			if err := resource.ThresholdsStyle.UnmarshalJSONStrict(fields["thresholdsStyle"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("thresholdsStyle", err)...)
			}

		}
		delete(fields, "thresholdsStyle")

	}
	// Field "lineColor"
	if fields["lineColor"] != nil {
		if string(fields["lineColor"]) != "null" {
			if err := json.Unmarshal(fields["lineColor"], &resource.LineColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineColor", err)...)
			}

		}
		delete(fields, "lineColor")

	}
	// Field "lineWidth"
	if fields["lineWidth"] != nil {
		if string(fields["lineWidth"]) != "null" {
			if err := json.Unmarshal(fields["lineWidth"], &resource.LineWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineWidth", err)...)
			}

		}
		delete(fields, "lineWidth")

	}
	// Field "lineInterpolation"
	if fields["lineInterpolation"] != nil {
		if string(fields["lineInterpolation"]) != "null" {
			if err := json.Unmarshal(fields["lineInterpolation"], &resource.LineInterpolation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineInterpolation", err)...)
			}

		}
		delete(fields, "lineInterpolation")

	}
	// Field "lineStyle"
	if fields["lineStyle"] != nil {
		if string(fields["lineStyle"]) != "null" {

			resource.LineStyle = &LineStyle{}
			if err := resource.LineStyle.UnmarshalJSONStrict(fields["lineStyle"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
			}

		}
		delete(fields, "lineStyle")

	}
	// Field "fillColor"
	if fields["fillColor"] != nil {
		if string(fields["fillColor"]) != "null" {
			if err := json.Unmarshal(fields["fillColor"], &resource.FillColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillColor", err)...)
			}

		}
		delete(fields, "fillColor")

	}
	// Field "fillOpacity"
	if fields["fillOpacity"] != nil {
		if string(fields["fillOpacity"]) != "null" {
			if err := json.Unmarshal(fields["fillOpacity"], &resource.FillOpacity); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillOpacity", err)...)
			}

		}
		delete(fields, "fillOpacity")

	}
	// Field "showPoints"
	if fields["showPoints"] != nil {
		if string(fields["showPoints"]) != "null" {
			if err := json.Unmarshal(fields["showPoints"], &resource.ShowPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showPoints", err)...)
			}

		}
		delete(fields, "showPoints")

	}
	// Field "pointSize"
	if fields["pointSize"] != nil {
		if string(fields["pointSize"]) != "null" {
			if err := json.Unmarshal(fields["pointSize"], &resource.PointSize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
			}

		}
		delete(fields, "pointSize")

	}
	// Field "pointColor"
	if fields["pointColor"] != nil {
		if string(fields["pointColor"]) != "null" {
			if err := json.Unmarshal(fields["pointColor"], &resource.PointColor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointColor", err)...)
			}

		}
		delete(fields, "pointColor")

	}
	// Field "axisPlacement"
	if fields["axisPlacement"] != nil {
		if string(fields["axisPlacement"]) != "null" {
			if err := json.Unmarshal(fields["axisPlacement"], &resource.AxisPlacement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisPlacement", err)...)
			}

		}
		delete(fields, "axisPlacement")

	}
	// Field "axisColorMode"
	if fields["axisColorMode"] != nil {
		if string(fields["axisColorMode"]) != "null" {
			if err := json.Unmarshal(fields["axisColorMode"], &resource.AxisColorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisColorMode", err)...)
			}

		}
		delete(fields, "axisColorMode")

	}
	// Field "axisLabel"
	if fields["axisLabel"] != nil {
		if string(fields["axisLabel"]) != "null" {
			if err := json.Unmarshal(fields["axisLabel"], &resource.AxisLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisLabel", err)...)
			}

		}
		delete(fields, "axisLabel")

	}
	// Field "axisWidth"
	if fields["axisWidth"] != nil {
		if string(fields["axisWidth"]) != "null" {
			if err := json.Unmarshal(fields["axisWidth"], &resource.AxisWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisWidth", err)...)
			}

		}
		delete(fields, "axisWidth")

	}
	// Field "axisSoftMin"
	if fields["axisSoftMin"] != nil {
		if string(fields["axisSoftMin"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMin"], &resource.AxisSoftMin); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMin", err)...)
			}

		}
		delete(fields, "axisSoftMin")

	}
	// Field "axisSoftMax"
	if fields["axisSoftMax"] != nil {
		if string(fields["axisSoftMax"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMax"], &resource.AxisSoftMax); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMax", err)...)
			}

		}
		delete(fields, "axisSoftMax")

	}
	// Field "axisGridShow"
	if fields["axisGridShow"] != nil {
		if string(fields["axisGridShow"]) != "null" {
			if err := json.Unmarshal(fields["axisGridShow"], &resource.AxisGridShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisGridShow", err)...)
			}

		}
		delete(fields, "axisGridShow")

	}
	// Field "scaleDistribution"
	if fields["scaleDistribution"] != nil {
		if string(fields["scaleDistribution"]) != "null" {

			resource.ScaleDistribution = &ScaleDistributionConfig{}
			if err := resource.ScaleDistribution.UnmarshalJSONStrict(fields["scaleDistribution"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
			}

		}
		delete(fields, "scaleDistribution")

	}
	// Field "axisCenteredZero"
	if fields["axisCenteredZero"] != nil {
		if string(fields["axisCenteredZero"]) != "null" {
			if err := json.Unmarshal(fields["axisCenteredZero"], &resource.AxisCenteredZero); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisCenteredZero", err)...)
			}

		}
		delete(fields, "axisCenteredZero")

	}
	// Field "barAlignment"
	if fields["barAlignment"] != nil {
		if string(fields["barAlignment"]) != "null" {
			if err := json.Unmarshal(fields["barAlignment"], &resource.BarAlignment); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barAlignment", err)...)
			}

		}
		delete(fields, "barAlignment")

	}
	// Field "barWidthFactor"
	if fields["barWidthFactor"] != nil {
		if string(fields["barWidthFactor"]) != "null" {
			if err := json.Unmarshal(fields["barWidthFactor"], &resource.BarWidthFactor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barWidthFactor", err)...)
			}

		}
		delete(fields, "barWidthFactor")

	}
	// Field "stacking"
	if fields["stacking"] != nil {
		if string(fields["stacking"]) != "null" {

			resource.Stacking = &StackingConfig{}
			if err := resource.Stacking.UnmarshalJSONStrict(fields["stacking"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stacking", err)...)
			}

		}
		delete(fields, "stacking")

	}
	// Field "hideFrom"
	if fields["hideFrom"] != nil {
		if string(fields["hideFrom"]) != "null" {

			resource.HideFrom = &HideSeriesConfig{}
			if err := resource.HideFrom.UnmarshalJSONStrict(fields["hideFrom"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
			}

		}
		delete(fields, "hideFrom")

	}
	// Field "hideValue"
	if fields["hideValue"] != nil {
		if string(fields["hideValue"]) != "null" {
			if err := json.Unmarshal(fields["hideValue"], &resource.HideValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideValue", err)...)
			}

		}
		delete(fields, "hideValue")

	}
	// Field "transform"
	if fields["transform"] != nil {
		if string(fields["transform"]) != "null" {
			if err := json.Unmarshal(fields["transform"], &resource.Transform); err != nil {
				errs = append(errs, cog.MakeBuildErrors("transform", err)...)
			}

		}
		delete(fields, "transform")

	}
	// Field "spanNulls"
	if fields["spanNulls"] != nil {
		if string(fields["spanNulls"]) != "null" {

			resource.SpanNulls = &BoolOrFloat64{}
			if err := resource.SpanNulls.UnmarshalJSONStrict(fields["spanNulls"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spanNulls", err)...)
			}

		}
		delete(fields, "spanNulls")

	}
	// Field "fillBelowTo"
	if fields["fillBelowTo"] != nil {
		if string(fields["fillBelowTo"]) != "null" {
			if err := json.Unmarshal(fields["fillBelowTo"], &resource.FillBelowTo); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillBelowTo", err)...)
			}

		}
		delete(fields, "fillBelowTo")

	}
	// Field "pointSymbol"
	if fields["pointSymbol"] != nil {
		if string(fields["pointSymbol"]) != "null" {
			if err := json.Unmarshal(fields["pointSymbol"], &resource.PointSymbol); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSymbol", err)...)
			}

		}
		delete(fields, "pointSymbol")

	}
	// Field "axisBorderShow"
	if fields["axisBorderShow"] != nil {
		if string(fields["axisBorderShow"]) != "null" {
			if err := json.Unmarshal(fields["axisBorderShow"], &resource.AxisBorderShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisBorderShow", err)...)
			}

		}
		delete(fields, "axisBorderShow")

	}
	// Field "barMaxWidth"
	if fields["barMaxWidth"] != nil {
		if string(fields["barMaxWidth"]) != "null" {
			if err := json.Unmarshal(fields["barMaxWidth"], &resource.BarMaxWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barMaxWidth", err)...)
			}

		}
		delete(fields, "barMaxWidth")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableSparklineCellOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableSparklineCellOptions` objects.
func (resource TableSparklineCellOptions) Equals(other TableSparklineCellOptions) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.DrawStyle == nil && other.DrawStyle != nil || resource.DrawStyle != nil && other.DrawStyle == nil {
		return false
	}

	if resource.DrawStyle != nil {
		if *resource.DrawStyle != *other.DrawStyle {
			return false
		}
	}
	if resource.GradientMode == nil && other.GradientMode != nil || resource.GradientMode != nil && other.GradientMode == nil {
		return false
	}

	if resource.GradientMode != nil {
		if *resource.GradientMode != *other.GradientMode {
			return false
		}
	}
	if resource.ThresholdsStyle == nil && other.ThresholdsStyle != nil || resource.ThresholdsStyle != nil && other.ThresholdsStyle == nil {
		return false
	}

	if resource.ThresholdsStyle != nil {
		if !resource.ThresholdsStyle.Equals(*other.ThresholdsStyle) {
			return false
		}
	}
	if resource.LineColor == nil && other.LineColor != nil || resource.LineColor != nil && other.LineColor == nil {
		return false
	}

	if resource.LineColor != nil {
		if *resource.LineColor != *other.LineColor {
			return false
		}
	}
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
			return false
		}
	}
	if resource.LineInterpolation == nil && other.LineInterpolation != nil || resource.LineInterpolation != nil && other.LineInterpolation == nil {
		return false
	}

	if resource.LineInterpolation != nil {
		if *resource.LineInterpolation != *other.LineInterpolation {
			return false
		}
	}
	if resource.LineStyle == nil && other.LineStyle != nil || resource.LineStyle != nil && other.LineStyle == nil {
		return false
	}

	if resource.LineStyle != nil {
		if !resource.LineStyle.Equals(*other.LineStyle) {
			return false
		}
	}
	if resource.FillColor == nil && other.FillColor != nil || resource.FillColor != nil && other.FillColor == nil {
		return false
	}

	if resource.FillColor != nil {
		if *resource.FillColor != *other.FillColor {
			return false
		}
	}
	if resource.FillOpacity == nil && other.FillOpacity != nil || resource.FillOpacity != nil && other.FillOpacity == nil {
		return false
	}

	if resource.FillOpacity != nil {
		if *resource.FillOpacity != *other.FillOpacity {
			return false
		}
	}
	if resource.ShowPoints == nil && other.ShowPoints != nil || resource.ShowPoints != nil && other.ShowPoints == nil {
		return false
	}

	if resource.ShowPoints != nil {
		if *resource.ShowPoints != *other.ShowPoints {
			return false
		}
	}
	if resource.PointSize == nil && other.PointSize != nil || resource.PointSize != nil && other.PointSize == nil {
		return false
	}

	if resource.PointSize != nil {
		if *resource.PointSize != *other.PointSize {
			return false
		}
	}
	if resource.PointColor == nil && other.PointColor != nil || resource.PointColor != nil && other.PointColor == nil {
		return false
	}

	if resource.PointColor != nil {
		if *resource.PointColor != *other.PointColor {
			return false
		}
	}
	if resource.AxisPlacement == nil && other.AxisPlacement != nil || resource.AxisPlacement != nil && other.AxisPlacement == nil {
		return false
	}

	if resource.AxisPlacement != nil {
		if *resource.AxisPlacement != *other.AxisPlacement {
			return false
		}
	}
	if resource.AxisColorMode == nil && other.AxisColorMode != nil || resource.AxisColorMode != nil && other.AxisColorMode == nil {
		return false
	}

	if resource.AxisColorMode != nil {
		if *resource.AxisColorMode != *other.AxisColorMode {
			return false
		}
	}
	if resource.AxisLabel == nil && other.AxisLabel != nil || resource.AxisLabel != nil && other.AxisLabel == nil {
		return false
	}

	if resource.AxisLabel != nil {
		if *resource.AxisLabel != *other.AxisLabel {
			return false
		}
	}
	if resource.AxisWidth == nil && other.AxisWidth != nil || resource.AxisWidth != nil && other.AxisWidth == nil {
		return false
	}

	if resource.AxisWidth != nil {
		if *resource.AxisWidth != *other.AxisWidth {
			return false
		}
	}
	if resource.AxisSoftMin == nil && other.AxisSoftMin != nil || resource.AxisSoftMin != nil && other.AxisSoftMin == nil {
		return false
	}

	if resource.AxisSoftMin != nil {
		if *resource.AxisSoftMin != *other.AxisSoftMin {
			return false
		}
	}
	if resource.AxisSoftMax == nil && other.AxisSoftMax != nil || resource.AxisSoftMax != nil && other.AxisSoftMax == nil {
		return false
	}

	if resource.AxisSoftMax != nil {
		if *resource.AxisSoftMax != *other.AxisSoftMax {
			return false
		}
	}
	if resource.AxisGridShow == nil && other.AxisGridShow != nil || resource.AxisGridShow != nil && other.AxisGridShow == nil {
		return false
	}

	if resource.AxisGridShow != nil {
		if *resource.AxisGridShow != *other.AxisGridShow {
			return false
		}
	}
	if resource.ScaleDistribution == nil && other.ScaleDistribution != nil || resource.ScaleDistribution != nil && other.ScaleDistribution == nil {
		return false
	}

	if resource.ScaleDistribution != nil {
		if !resource.ScaleDistribution.Equals(*other.ScaleDistribution) {
			return false
		}
	}
	if resource.AxisCenteredZero == nil && other.AxisCenteredZero != nil || resource.AxisCenteredZero != nil && other.AxisCenteredZero == nil {
		return false
	}

	if resource.AxisCenteredZero != nil {
		if *resource.AxisCenteredZero != *other.AxisCenteredZero {
			return false
		}
	}
	if resource.BarAlignment == nil && other.BarAlignment != nil || resource.BarAlignment != nil && other.BarAlignment == nil {
		return false
	}

	if resource.BarAlignment != nil {
		if *resource.BarAlignment != *other.BarAlignment {
			return false
		}
	}
	if resource.BarWidthFactor == nil && other.BarWidthFactor != nil || resource.BarWidthFactor != nil && other.BarWidthFactor == nil {
		return false
	}

	if resource.BarWidthFactor != nil {
		if *resource.BarWidthFactor != *other.BarWidthFactor {
			return false
		}
	}
	if resource.Stacking == nil && other.Stacking != nil || resource.Stacking != nil && other.Stacking == nil {
		return false
	}

	if resource.Stacking != nil {
		if !resource.Stacking.Equals(*other.Stacking) {
			return false
		}
	}
	if resource.HideFrom == nil && other.HideFrom != nil || resource.HideFrom != nil && other.HideFrom == nil {
		return false
	}

	if resource.HideFrom != nil {
		if !resource.HideFrom.Equals(*other.HideFrom) {
			return false
		}
	}
	if resource.HideValue == nil && other.HideValue != nil || resource.HideValue != nil && other.HideValue == nil {
		return false
	}

	if resource.HideValue != nil {
		if *resource.HideValue != *other.HideValue {
			return false
		}
	}
	if resource.Transform == nil && other.Transform != nil || resource.Transform != nil && other.Transform == nil {
		return false
	}

	if resource.Transform != nil {
		if *resource.Transform != *other.Transform {
			return false
		}
	}
	if resource.SpanNulls == nil && other.SpanNulls != nil || resource.SpanNulls != nil && other.SpanNulls == nil {
		return false
	}

	if resource.SpanNulls != nil {
		if !resource.SpanNulls.Equals(*other.SpanNulls) {
			return false
		}
	}
	if resource.FillBelowTo == nil && other.FillBelowTo != nil || resource.FillBelowTo != nil && other.FillBelowTo == nil {
		return false
	}

	if resource.FillBelowTo != nil {
		if *resource.FillBelowTo != *other.FillBelowTo {
			return false
		}
	}
	if resource.PointSymbol == nil && other.PointSymbol != nil || resource.PointSymbol != nil && other.PointSymbol == nil {
		return false
	}

	if resource.PointSymbol != nil {
		if *resource.PointSymbol != *other.PointSymbol {
			return false
		}
	}
	if resource.AxisBorderShow == nil && other.AxisBorderShow != nil || resource.AxisBorderShow != nil && other.AxisBorderShow == nil {
		return false
	}

	if resource.AxisBorderShow != nil {
		if *resource.AxisBorderShow != *other.AxisBorderShow {
			return false
		}
	}
	if resource.BarMaxWidth == nil && other.BarMaxWidth != nil || resource.BarMaxWidth != nil && other.BarMaxWidth == nil {
		return false
	}

	if resource.BarMaxWidth != nil {
		if *resource.BarMaxWidth != *other.BarMaxWidth {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableSparklineCellOptions` fields for violations and returns them.
func (resource TableSparklineCellOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "sparkline") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == sparkline"),
		)...)
	}
	if resource.ThresholdsStyle != nil {
		if err := resource.ThresholdsStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("thresholdsStyle", err)...)
		}
	}
	if resource.LineStyle != nil {
		if err := resource.LineStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
		}
	}
	if resource.ScaleDistribution != nil {
		if err := resource.ScaleDistribution.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
		}
	}
	if resource.Stacking != nil {
		if err := resource.Stacking.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("stacking", err)...)
		}
	}
	if resource.HideFrom != nil {
		if err := resource.HideFrom.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
		}
	}
	if resource.SpanNulls != nil {
		if err := resource.SpanNulls.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("spanNulls", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Colored background cell options
type TableColoredBackgroundCellOptions struct {
	Type string                          `json:"type"`
	Mode *TableCellBackgroundDisplayMode `json:"mode,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableColoredBackgroundCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableColoredBackgroundCellOptions) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}

		}
		delete(fields, "mode")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableColoredBackgroundCellOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableColoredBackgroundCellOptions` objects.
func (resource TableColoredBackgroundCellOptions) Equals(other TableColoredBackgroundCellOptions) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableColoredBackgroundCellOptions` fields for violations and returns them.
func (resource TableColoredBackgroundCellOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "color-background") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == color-background"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Height of a table cell
type TableCellHeight string

const (
	TableCellHeightSm TableCellHeight = "sm"
	TableCellHeightMd TableCellHeight = "md"
	TableCellHeightLg TableCellHeight = "lg"
)

// Table cell options. Each cell has a display mode
// and other potential options for that display.
type TableCellOptions = TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions

// Use UTC/GMT timezone
const TimeZoneUtc = "utc"

// Use the timezone defined by end user web browser
const TimeZoneBrowser = "browser"

// Optional formats for the template variable replace functions
// See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options
type VariableFormatID string

const (
	VariableFormatIDLucene        VariableFormatID = "lucene"
	VariableFormatIDRaw           VariableFormatID = "raw"
	VariableFormatIDRegex         VariableFormatID = "regex"
	VariableFormatIDPipe          VariableFormatID = "pipe"
	VariableFormatIDDistributed   VariableFormatID = "distributed"
	VariableFormatIDCSV           VariableFormatID = "csv"
	VariableFormatIDHTML          VariableFormatID = "html"
	VariableFormatIDJSON          VariableFormatID = "json"
	VariableFormatIDPercentEncode VariableFormatID = "percentencode"
	VariableFormatIDUriEncode     VariableFormatID = "uriencode"
	VariableFormatIDSingleQuote   VariableFormatID = "singlequote"
	VariableFormatIDDoubleQuote   VariableFormatID = "doublequote"
	VariableFormatIDSQLString     VariableFormatID = "sqlstring"
	VariableFormatIDDate          VariableFormatID = "date"
	VariableFormatIDGlob          VariableFormatID = "glob"
	VariableFormatIDText          VariableFormatID = "text"
	VariableFormatIDQueryParam    VariableFormatID = "queryparam"
)

// Links to a resource (image/svg path)
type ResourceDimensionConfig struct {
	Mode ResourceDimensionMode `json:"mode"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
	Fixed *string `json:"fixed,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceDimensionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ResourceDimensionConfig) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}
	// Field "fixed"
	if fields["fixed"] != nil {
		if string(fields["fixed"]) != "null" {
			if err := json.Unmarshal(fields["fixed"], &resource.Fixed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixed", err)...)
			}

		}
		delete(fields, "fixed")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ResourceDimensionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ResourceDimensionConfig` objects.
func (resource ResourceDimensionConfig) Equals(other ResourceDimensionConfig) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Fixed == nil && other.Fixed != nil || resource.Fixed != nil && other.Fixed == nil {
		return false
	}

	if resource.Fixed != nil {
		if *resource.Fixed != *other.Fixed {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ResourceDimensionConfig` fields for violations and returns them.
func (resource ResourceDimensionConfig) Validate() error {
	return nil
}

type FrameGeometrySource struct {
	Mode FrameGeometrySourceMode `json:"mode"`
	// Field mappings
	Geohash   *string `json:"geohash,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Wkt       *string `json:"wkt,omitempty"`
	Lookup    *string `json:"lookup,omitempty"`
	// Path to Gazetteer
	Gazetteer *string `json:"gazetteer,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FrameGeometrySource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FrameGeometrySource) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "geohash"
	if fields["geohash"] != nil {
		if string(fields["geohash"]) != "null" {
			if err := json.Unmarshal(fields["geohash"], &resource.Geohash); err != nil {
				errs = append(errs, cog.MakeBuildErrors("geohash", err)...)
			}

		}
		delete(fields, "geohash")

	}
	// Field "latitude"
	if fields["latitude"] != nil {
		if string(fields["latitude"]) != "null" {
			if err := json.Unmarshal(fields["latitude"], &resource.Latitude); err != nil {
				errs = append(errs, cog.MakeBuildErrors("latitude", err)...)
			}

		}
		delete(fields, "latitude")

	}
	// Field "longitude"
	if fields["longitude"] != nil {
		if string(fields["longitude"]) != "null" {
			if err := json.Unmarshal(fields["longitude"], &resource.Longitude); err != nil {
				errs = append(errs, cog.MakeBuildErrors("longitude", err)...)
			}

		}
		delete(fields, "longitude")

	}
	// Field "wkt"
	if fields["wkt"] != nil {
		if string(fields["wkt"]) != "null" {
			if err := json.Unmarshal(fields["wkt"], &resource.Wkt); err != nil {
				errs = append(errs, cog.MakeBuildErrors("wkt", err)...)
			}

		}
		delete(fields, "wkt")

	}
	// Field "lookup"
	if fields["lookup"] != nil {
		if string(fields["lookup"]) != "null" {
			if err := json.Unmarshal(fields["lookup"], &resource.Lookup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lookup", err)...)
			}

		}
		delete(fields, "lookup")

	}
	// Field "gazetteer"
	if fields["gazetteer"] != nil {
		if string(fields["gazetteer"]) != "null" {
			if err := json.Unmarshal(fields["gazetteer"], &resource.Gazetteer); err != nil {
				errs = append(errs, cog.MakeBuildErrors("gazetteer", err)...)
			}

		}
		delete(fields, "gazetteer")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FrameGeometrySource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FrameGeometrySource` objects.
func (resource FrameGeometrySource) Equals(other FrameGeometrySource) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Geohash == nil && other.Geohash != nil || resource.Geohash != nil && other.Geohash == nil {
		return false
	}

	if resource.Geohash != nil {
		if *resource.Geohash != *other.Geohash {
			return false
		}
	}
	if resource.Latitude == nil && other.Latitude != nil || resource.Latitude != nil && other.Latitude == nil {
		return false
	}

	if resource.Latitude != nil {
		if *resource.Latitude != *other.Latitude {
			return false
		}
	}
	if resource.Longitude == nil && other.Longitude != nil || resource.Longitude != nil && other.Longitude == nil {
		return false
	}

	if resource.Longitude != nil {
		if *resource.Longitude != *other.Longitude {
			return false
		}
	}
	if resource.Wkt == nil && other.Wkt != nil || resource.Wkt != nil && other.Wkt == nil {
		return false
	}

	if resource.Wkt != nil {
		if *resource.Wkt != *other.Wkt {
			return false
		}
	}
	if resource.Lookup == nil && other.Lookup != nil || resource.Lookup != nil && other.Lookup == nil {
		return false
	}

	if resource.Lookup != nil {
		if *resource.Lookup != *other.Lookup {
			return false
		}
	}
	if resource.Gazetteer == nil && other.Gazetteer != nil || resource.Gazetteer != nil && other.Gazetteer == nil {
		return false
	}

	if resource.Gazetteer != nil {
		if *resource.Gazetteer != *other.Gazetteer {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FrameGeometrySource` fields for violations and returns them.
func (resource FrameGeometrySource) Validate() error {
	return nil
}

type HeatmapCalculationOptions struct {
	// The number of buckets to use for the xAxis in the heatmap
	XBuckets *HeatmapCalculationBucketConfig `json:"xBuckets,omitempty"`
	// The number of buckets to use for the yAxis in the heatmap
	YBuckets *HeatmapCalculationBucketConfig `json:"yBuckets,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapCalculationOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HeatmapCalculationOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "xBuckets"
	if fields["xBuckets"] != nil {
		if string(fields["xBuckets"]) != "null" {

			resource.XBuckets = &HeatmapCalculationBucketConfig{}
			if err := resource.XBuckets.UnmarshalJSONStrict(fields["xBuckets"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("xBuckets", err)...)
			}

		}
		delete(fields, "xBuckets")

	}
	// Field "yBuckets"
	if fields["yBuckets"] != nil {
		if string(fields["yBuckets"]) != "null" {

			resource.YBuckets = &HeatmapCalculationBucketConfig{}
			if err := resource.YBuckets.UnmarshalJSONStrict(fields["yBuckets"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("yBuckets", err)...)
			}

		}
		delete(fields, "yBuckets")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HeatmapCalculationOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HeatmapCalculationOptions` objects.
func (resource HeatmapCalculationOptions) Equals(other HeatmapCalculationOptions) bool {
	if resource.XBuckets == nil && other.XBuckets != nil || resource.XBuckets != nil && other.XBuckets == nil {
		return false
	}

	if resource.XBuckets != nil {
		if !resource.XBuckets.Equals(*other.XBuckets) {
			return false
		}
	}
	if resource.YBuckets == nil && other.YBuckets != nil || resource.YBuckets != nil && other.YBuckets == nil {
		return false
	}

	if resource.YBuckets != nil {
		if !resource.YBuckets.Equals(*other.YBuckets) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `HeatmapCalculationOptions` fields for violations and returns them.
func (resource HeatmapCalculationOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.XBuckets != nil {
		if err := resource.XBuckets.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("xBuckets", err)...)
		}
	}
	if resource.YBuckets != nil {
		if err := resource.YBuckets.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("yBuckets", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type LogsDedupStrategy string

const (
	LogsDedupStrategyNone      LogsDedupStrategy = "none"
	LogsDedupStrategyExact     LogsDedupStrategy = "exact"
	LogsDedupStrategyNumbers   LogsDedupStrategy = "numbers"
	LogsDedupStrategySignature LogsDedupStrategy = "signature"
)

// Compare two values
type ComparisonOperation string

const (
	ComparisonOperationEQ  ComparisonOperation = "eq"
	ComparisonOperationNEQ ComparisonOperation = "neq"
	ComparisonOperationLT  ComparisonOperation = "lt"
	ComparisonOperationLTE ComparisonOperation = "lte"
	ComparisonOperationGT  ComparisonOperation = "gt"
	ComparisonOperationGTE ComparisonOperation = "gte"
)

// Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
// Generally defines alignment, filtering capabilties, display options, etc.
type TableFieldOptions struct {
	Width    *float64           `json:"width,omitempty"`
	MinWidth *float64           `json:"minWidth,omitempty"`
	Align    FieldTextAlignment `json:"align"`
	// This field is deprecated in favor of using cellOptions
	DisplayMode *TableCellDisplayMode `json:"displayMode,omitempty"`
	CellOptions *TableCellOptions     `json:"cellOptions,omitempty"`
	// ?? default is missing or false ??
	Hidden     *bool `json:"hidden,omitempty"`
	Inspect    bool  `json:"inspect"`
	Filterable *bool `json:"filterable,omitempty"`
	// Hides any header for a column, useful for columns that show some static content or buttons.
	HideHeader *bool `json:"hideHeader,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableFieldOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableFieldOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "width"
	if fields["width"] != nil {
		if string(fields["width"]) != "null" {
			if err := json.Unmarshal(fields["width"], &resource.Width); err != nil {
				errs = append(errs, cog.MakeBuildErrors("width", err)...)
			}

		}
		delete(fields, "width")

	}
	// Field "minWidth"
	if fields["minWidth"] != nil {
		if string(fields["minWidth"]) != "null" {
			if err := json.Unmarshal(fields["minWidth"], &resource.MinWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minWidth", err)...)
			}

		}
		delete(fields, "minWidth")

	}
	// Field "align"
	if fields["align"] != nil {
		if string(fields["align"]) != "null" {
			if err := json.Unmarshal(fields["align"], &resource.Align); err != nil {
				errs = append(errs, cog.MakeBuildErrors("align", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("align", errors.New("required field is null"))...)

		}
		delete(fields, "align")
	} else {
		errs = append(errs, cog.MakeBuildErrors("align", errors.New("required field is missing from input"))...)
	}
	// Field "displayMode"
	if fields["displayMode"] != nil {
		if string(fields["displayMode"]) != "null" {
			if err := json.Unmarshal(fields["displayMode"], &resource.DisplayMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayMode", err)...)
			}

		}
		delete(fields, "displayMode")

	}
	// Field "cellOptions"
	if fields["cellOptions"] != nil {
		if string(fields["cellOptions"]) != "null" {

			resource.CellOptions = &TableCellOptions{}
			if err := resource.CellOptions.UnmarshalJSONStrict(fields["cellOptions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cellOptions", err)...)
			}

		}
		delete(fields, "cellOptions")

	}
	// Field "hidden"
	if fields["hidden"] != nil {
		if string(fields["hidden"]) != "null" {
			if err := json.Unmarshal(fields["hidden"], &resource.Hidden); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hidden", err)...)
			}

		}
		delete(fields, "hidden")

	}
	// Field "inspect"
	if fields["inspect"] != nil {
		if string(fields["inspect"]) != "null" {
			if err := json.Unmarshal(fields["inspect"], &resource.Inspect); err != nil {
				errs = append(errs, cog.MakeBuildErrors("inspect", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("inspect", errors.New("required field is null"))...)

		}
		delete(fields, "inspect")
	} else {
		errs = append(errs, cog.MakeBuildErrors("inspect", errors.New("required field is missing from input"))...)
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
	// Field "hideHeader"
	if fields["hideHeader"] != nil {
		if string(fields["hideHeader"]) != "null" {
			if err := json.Unmarshal(fields["hideHeader"], &resource.HideHeader); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideHeader", err)...)
			}

		}
		delete(fields, "hideHeader")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TableFieldOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TableFieldOptions` objects.
func (resource TableFieldOptions) Equals(other TableFieldOptions) bool {
	if resource.Width == nil && other.Width != nil || resource.Width != nil && other.Width == nil {
		return false
	}

	if resource.Width != nil {
		if *resource.Width != *other.Width {
			return false
		}
	}
	if resource.MinWidth == nil && other.MinWidth != nil || resource.MinWidth != nil && other.MinWidth == nil {
		return false
	}

	if resource.MinWidth != nil {
		if *resource.MinWidth != *other.MinWidth {
			return false
		}
	}
	if resource.Align != other.Align {
		return false
	}
	if resource.DisplayMode == nil && other.DisplayMode != nil || resource.DisplayMode != nil && other.DisplayMode == nil {
		return false
	}

	if resource.DisplayMode != nil {
		if *resource.DisplayMode != *other.DisplayMode {
			return false
		}
	}
	if resource.CellOptions == nil && other.CellOptions != nil || resource.CellOptions != nil && other.CellOptions == nil {
		return false
	}

	if resource.CellOptions != nil {
		if !resource.CellOptions.Equals(*other.CellOptions) {
			return false
		}
	}
	if resource.Hidden == nil && other.Hidden != nil || resource.Hidden != nil && other.Hidden == nil {
		return false
	}

	if resource.Hidden != nil {
		if *resource.Hidden != *other.Hidden {
			return false
		}
	}
	if resource.Inspect != other.Inspect {
		return false
	}
	if resource.Filterable == nil && other.Filterable != nil || resource.Filterable != nil && other.Filterable == nil {
		return false
	}

	if resource.Filterable != nil {
		if *resource.Filterable != *other.Filterable {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `TableFieldOptions` fields for violations and returns them.
func (resource TableFieldOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.CellOptions != nil {
		if err := resource.CellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("cellOptions", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// A specific timezone from https://en.wikipedia.org/wiki/Tz_database
type TimeZone string

type LineStyleFill string

const (
	LineStyleFillSolid  LineStyleFill = "solid"
	LineStyleFillDash   LineStyleFill = "dash"
	LineStyleFillDot    LineStyleFill = "dot"
	LineStyleFillSquare LineStyleFill = "square"
)

type BoolOrFloat64 struct {
	Bool    *bool    `json:"Bool,omitempty"`
	Float64 *float64 `json:"Float64,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `BoolOrFloat64` as JSON.
func (resource BoolOrFloat64) MarshalJSON() ([]byte, error) {
	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}

	if resource.Float64 != nil {
		return json.Marshal(resource.Float64)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `BoolOrFloat64` from JSON.
func (resource *BoolOrFloat64) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

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

	return errors.Join(errList...)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BoolOrFloat64` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BoolOrFloat64) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

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

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("BoolOrFloat64", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BoolOrFloat64` objects.
func (resource BoolOrFloat64) Equals(other BoolOrFloat64) bool {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `BoolOrFloat64` fields for violations and returns them.
func (resource BoolOrFloat64) Validate() error {
	return nil
}

type BoolOrUint32 struct {
	Bool   *bool   `json:"Bool,omitempty"`
	Uint32 *uint32 `json:"Uint32,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `BoolOrUint32` as JSON.
func (resource BoolOrUint32) MarshalJSON() ([]byte, error) {
	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}

	if resource.Uint32 != nil {
		return json.Marshal(resource.Uint32)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `BoolOrUint32` from JSON.
func (resource *BoolOrUint32) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Uint32
	var Uint32 uint32
	if err := json.Unmarshal(raw, &Uint32); err != nil {
		errList = append(errList, err)
		resource.Uint32 = nil
	} else {
		resource.Uint32 = &Uint32
		return nil
	}

	return errors.Join(errList...)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BoolOrUint32` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BoolOrUint32) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// Bool
	var Bool bool

	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Uint32
	var Uint32 uint32

	if err := json.Unmarshal(raw, &Uint32); err != nil {
		errList = append(errList, err)
	} else {
		resource.Uint32 = &Uint32
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("BoolOrUint32", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BoolOrUint32` objects.
func (resource BoolOrUint32) Equals(other BoolOrUint32) bool {
	if resource.Bool == nil && other.Bool != nil || resource.Bool != nil && other.Bool == nil {
		return false
	}

	if resource.Bool != nil {
		if *resource.Bool != *other.Bool {
			return false
		}
	}
	if resource.Uint32 == nil && other.Uint32 != nil || resource.Uint32 != nil && other.Uint32 == nil {
		return false
	}

	if resource.Uint32 != nil {
		if *resource.Uint32 != *other.Uint32 {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BoolOrUint32` fields for violations and returns them.
func (resource BoolOrUint32) Validate() error {
	return nil
}

type TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions struct {
	TableAutoCellOptions              *TableAutoCellOptions              `json:"TableAutoCellOptions,omitempty"`
	TableSparklineCellOptions         *TableSparklineCellOptions         `json:"TableSparklineCellOptions,omitempty"`
	TableBarGaugeCellOptions          *TableBarGaugeCellOptions          `json:"TableBarGaugeCellOptions,omitempty"`
	TableColoredBackgroundCellOptions *TableColoredBackgroundCellOptions `json:"TableColoredBackgroundCellOptions,omitempty"`
	TableColorTextCellOptions         *TableColorTextCellOptions         `json:"TableColorTextCellOptions,omitempty"`
	TableImageCellOptions             *TableImageCellOptions             `json:"TableImageCellOptions,omitempty"`
	TableJsonViewCellOptions          *TableJsonViewCellOptions          `json:"TableJsonViewCellOptions,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions` as JSON.
func (resource TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions) MarshalJSON() ([]byte, error) {
	if resource.TableAutoCellOptions != nil {
		return json.Marshal(resource.TableAutoCellOptions)
	}
	if resource.TableSparklineCellOptions != nil {
		return json.Marshal(resource.TableSparklineCellOptions)
	}
	if resource.TableBarGaugeCellOptions != nil {
		return json.Marshal(resource.TableBarGaugeCellOptions)
	}
	if resource.TableColoredBackgroundCellOptions != nil {
		return json.Marshal(resource.TableColoredBackgroundCellOptions)
	}
	if resource.TableColorTextCellOptions != nil {
		return json.Marshal(resource.TableColorTextCellOptions)
	}
	if resource.TableImageCellOptions != nil {
		return json.Marshal(resource.TableImageCellOptions)
	}
	if resource.TableJsonViewCellOptions != nil {
		return json.Marshal(resource.TableJsonViewCellOptions)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions` from JSON.
func (resource *TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions) UnmarshalJSON(raw []byte) error {
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
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "auto":
		var tableAutoCellOptions TableAutoCellOptions
		if err := json.Unmarshal(raw, &tableAutoCellOptions); err != nil {
			return err
		}

		resource.TableAutoCellOptions = &tableAutoCellOptions
		return nil
	case "color-background":
		var tableColoredBackgroundCellOptions TableColoredBackgroundCellOptions
		if err := json.Unmarshal(raw, &tableColoredBackgroundCellOptions); err != nil {
			return err
		}

		resource.TableColoredBackgroundCellOptions = &tableColoredBackgroundCellOptions
		return nil
	case "color-text":
		var tableColorTextCellOptions TableColorTextCellOptions
		if err := json.Unmarshal(raw, &tableColorTextCellOptions); err != nil {
			return err
		}

		resource.TableColorTextCellOptions = &tableColorTextCellOptions
		return nil
	case "gauge":
		var tableBarGaugeCellOptions TableBarGaugeCellOptions
		if err := json.Unmarshal(raw, &tableBarGaugeCellOptions); err != nil {
			return err
		}

		resource.TableBarGaugeCellOptions = &tableBarGaugeCellOptions
		return nil
	case "image":
		var tableImageCellOptions TableImageCellOptions
		if err := json.Unmarshal(raw, &tableImageCellOptions); err != nil {
			return err
		}

		resource.TableImageCellOptions = &tableImageCellOptions
		return nil
	case "json-view":
		var tableJsonViewCellOptions TableJsonViewCellOptions
		if err := json.Unmarshal(raw, &tableJsonViewCellOptions); err != nil {
			return err
		}

		resource.TableJsonViewCellOptions = &tableJsonViewCellOptions
		return nil
	case "sparkline":
		var tableSparklineCellOptions TableSparklineCellOptions
		if err := json.Unmarshal(raw, &tableSparklineCellOptions); err != nil {
			return err
		}

		resource.TableSparklineCellOptions = &tableSparklineCellOptions
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions) UnmarshalJSONStrict(raw []byte) error {
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
	case "auto":
		tableAutoCellOptions := &TableAutoCellOptions{}
		if err := tableAutoCellOptions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TableAutoCellOptions = tableAutoCellOptions
		return nil
	case "color-background":
		tableColoredBackgroundCellOptions := &TableColoredBackgroundCellOptions{}
		if err := tableColoredBackgroundCellOptions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TableColoredBackgroundCellOptions = tableColoredBackgroundCellOptions
		return nil
	case "color-text":
		tableColorTextCellOptions := &TableColorTextCellOptions{}
		if err := tableColorTextCellOptions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TableColorTextCellOptions = tableColorTextCellOptions
		return nil
	case "gauge":
		tableBarGaugeCellOptions := &TableBarGaugeCellOptions{}
		if err := tableBarGaugeCellOptions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TableBarGaugeCellOptions = tableBarGaugeCellOptions
		return nil
	case "image":
		tableImageCellOptions := &TableImageCellOptions{}
		if err := tableImageCellOptions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TableImageCellOptions = tableImageCellOptions
		return nil
	case "json-view":
		tableJsonViewCellOptions := &TableJsonViewCellOptions{}
		if err := tableJsonViewCellOptions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TableJsonViewCellOptions = tableJsonViewCellOptions
		return nil
	case "sparkline":
		tableSparklineCellOptions := &TableSparklineCellOptions{}
		if err := tableSparklineCellOptions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TableSparklineCellOptions = tableSparklineCellOptions
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions` objects.
func (resource TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions) Equals(other TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions) bool {
	if resource.TableAutoCellOptions == nil && other.TableAutoCellOptions != nil || resource.TableAutoCellOptions != nil && other.TableAutoCellOptions == nil {
		return false
	}

	if resource.TableAutoCellOptions != nil {
		if !resource.TableAutoCellOptions.Equals(*other.TableAutoCellOptions) {
			return false
		}
	}
	if resource.TableSparklineCellOptions == nil && other.TableSparklineCellOptions != nil || resource.TableSparklineCellOptions != nil && other.TableSparklineCellOptions == nil {
		return false
	}

	if resource.TableSparklineCellOptions != nil {
		if !resource.TableSparklineCellOptions.Equals(*other.TableSparklineCellOptions) {
			return false
		}
	}
	if resource.TableBarGaugeCellOptions == nil && other.TableBarGaugeCellOptions != nil || resource.TableBarGaugeCellOptions != nil && other.TableBarGaugeCellOptions == nil {
		return false
	}

	if resource.TableBarGaugeCellOptions != nil {
		if !resource.TableBarGaugeCellOptions.Equals(*other.TableBarGaugeCellOptions) {
			return false
		}
	}
	if resource.TableColoredBackgroundCellOptions == nil && other.TableColoredBackgroundCellOptions != nil || resource.TableColoredBackgroundCellOptions != nil && other.TableColoredBackgroundCellOptions == nil {
		return false
	}

	if resource.TableColoredBackgroundCellOptions != nil {
		if !resource.TableColoredBackgroundCellOptions.Equals(*other.TableColoredBackgroundCellOptions) {
			return false
		}
	}
	if resource.TableColorTextCellOptions == nil && other.TableColorTextCellOptions != nil || resource.TableColorTextCellOptions != nil && other.TableColorTextCellOptions == nil {
		return false
	}

	if resource.TableColorTextCellOptions != nil {
		if !resource.TableColorTextCellOptions.Equals(*other.TableColorTextCellOptions) {
			return false
		}
	}
	if resource.TableImageCellOptions == nil && other.TableImageCellOptions != nil || resource.TableImageCellOptions != nil && other.TableImageCellOptions == nil {
		return false
	}

	if resource.TableImageCellOptions != nil {
		if !resource.TableImageCellOptions.Equals(*other.TableImageCellOptions) {
			return false
		}
	}
	if resource.TableJsonViewCellOptions == nil && other.TableJsonViewCellOptions != nil || resource.TableJsonViewCellOptions != nil && other.TableJsonViewCellOptions == nil {
		return false
	}

	if resource.TableJsonViewCellOptions != nil {
		if !resource.TableJsonViewCellOptions.Equals(*other.TableJsonViewCellOptions) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions` fields for violations and returns them.
func (resource TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableJsonViewCellOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.TableAutoCellOptions != nil {
		if err := resource.TableAutoCellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TableAutoCellOptions", err)...)
		}
	}
	if resource.TableSparklineCellOptions != nil {
		if err := resource.TableSparklineCellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TableSparklineCellOptions", err)...)
		}
	}
	if resource.TableBarGaugeCellOptions != nil {
		if err := resource.TableBarGaugeCellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TableBarGaugeCellOptions", err)...)
		}
	}
	if resource.TableColoredBackgroundCellOptions != nil {
		if err := resource.TableColoredBackgroundCellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TableColoredBackgroundCellOptions", err)...)
		}
	}
	if resource.TableColorTextCellOptions != nil {
		if err := resource.TableColorTextCellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TableColorTextCellOptions", err)...)
		}
	}
	if resource.TableImageCellOptions != nil {
		if err := resource.TableImageCellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TableImageCellOptions", err)...)
		}
	}
	if resource.TableJsonViewCellOptions != nil {
		if err := resource.TableJsonViewCellOptions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TableJsonViewCellOptions", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
