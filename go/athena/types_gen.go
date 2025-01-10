// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package athena

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
type Dataquery struct {
	Format         FormatOptions  `json:"format"`
	ConnectionArgs ConnectionArgs `json:"connectionArgs"`
	Table          *string        `json:"table,omitempty"`
	Column         *string        `json:"column,omitempty"`
	QueryID        *string        `json:"queryID,omitempty"`
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
	RawSQL    string  `json:"rawSQL"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return "grafana-athena-datasource"
}

// NewDataquery creates a new Dataquery object.
func NewDataquery() *Dataquery {
	return &Dataquery{
		ConnectionArgs: *NewConnectionArgs(),
		RawSQL:         "",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dataquery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dataquery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "format"
	if fields["format"] != nil {
		if string(fields["format"]) != "null" {
			if err := json.Unmarshal(fields["format"], &resource.Format); err != nil {
				errs = append(errs, cog.MakeBuildErrors("format", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("format", errors.New("required field is null"))...)

		}
		delete(fields, "format")
	} else {
		errs = append(errs, cog.MakeBuildErrors("format", errors.New("required field is missing from input"))...)
	}
	// Field "connectionArgs"
	if fields["connectionArgs"] != nil {
		if string(fields["connectionArgs"]) != "null" {

			resource.ConnectionArgs = ConnectionArgs{}
			if err := resource.ConnectionArgs.UnmarshalJSONStrict(fields["connectionArgs"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("connectionArgs", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("connectionArgs", errors.New("required field is null"))...)

		}
		delete(fields, "connectionArgs")
	} else {
		errs = append(errs, cog.MakeBuildErrors("connectionArgs", errors.New("required field is missing from input"))...)
	}
	// Field "table"
	if fields["table"] != nil {
		if string(fields["table"]) != "null" {
			if err := json.Unmarshal(fields["table"], &resource.Table); err != nil {
				errs = append(errs, cog.MakeBuildErrors("table", err)...)
			}

		}
		delete(fields, "table")

	}
	// Field "column"
	if fields["column"] != nil {
		if string(fields["column"]) != "null" {
			if err := json.Unmarshal(fields["column"], &resource.Column); err != nil {
				errs = append(errs, cog.MakeBuildErrors("column", err)...)
			}

		}
		delete(fields, "column")

	}
	// Field "queryID"
	if fields["queryID"] != nil {
		if string(fields["queryID"]) != "null" {
			if err := json.Unmarshal(fields["queryID"], &resource.QueryID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryID", err)...)
			}

		}
		delete(fields, "queryID")

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
	// Field "rawSQL"
	if fields["rawSQL"] != nil {
		if string(fields["rawSQL"]) != "null" {
			if err := json.Unmarshal(fields["rawSQL"], &resource.RawSQL); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawSQL", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("rawSQL", errors.New("required field is null"))...)

		}
		delete(fields, "rawSQL")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dataquery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
func (resource Dataquery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(Dataquery)
	if !ok {
		return false
	}
	if resource.Format != other.Format {
		return false
	}
	if !resource.ConnectionArgs.Equals(other.ConnectionArgs) {
		return false
	}
	if resource.Table == nil && other.Table != nil || resource.Table != nil && other.Table == nil {
		return false
	}

	if resource.Table != nil {
		if *resource.Table != *other.Table {
			return false
		}
	}
	if resource.Column == nil && other.Column != nil || resource.Column != nil && other.Column == nil {
		return false
	}

	if resource.Column != nil {
		if *resource.Column != *other.Column {
			return false
		}
	}
	if resource.QueryID == nil && other.QueryID != nil || resource.QueryID != nil && other.QueryID == nil {
		return false
	}

	if resource.QueryID != nil {
		if *resource.QueryID != *other.QueryID {
			return false
		}
	}
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
	if resource.RawSQL != other.RawSQL {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.
func (resource Dataquery) Validate() error {
	var errs cog.BuildErrors
	if err := resource.ConnectionArgs.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("connectionArgs", err)...)
	}
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

const DefaultKey = "__default"

type ConnectionArgs struct {
	Region                     *string  `json:"region,omitempty"`
	Catalog                    *string  `json:"catalog,omitempty"`
	Database                   *string  `json:"database,omitempty"`
	ResultReuseEnabled         *bool    `json:"resultReuseEnabled,omitempty"`
	ResultReuseMaxAgeInMinutes *float64 `json:"resultReuseMaxAgeInMinutes,omitempty"`
}

// NewConnectionArgs creates a new ConnectionArgs object.
func NewConnectionArgs() *ConnectionArgs {
	return &ConnectionArgs{
		Region:                     (func(input string) *string { return &input })("__default"),
		Catalog:                    (func(input string) *string { return &input })("__default"),
		Database:                   (func(input string) *string { return &input })("__default"),
		ResultReuseEnabled:         (func(input bool) *bool { return &input })(false),
		ResultReuseMaxAgeInMinutes: (func(input float64) *float64 { return &input })(60),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConnectionArgs` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConnectionArgs) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}

		}
		delete(fields, "region")

	}
	// Field "catalog"
	if fields["catalog"] != nil {
		if string(fields["catalog"]) != "null" {
			if err := json.Unmarshal(fields["catalog"], &resource.Catalog); err != nil {
				errs = append(errs, cog.MakeBuildErrors("catalog", err)...)
			}

		}
		delete(fields, "catalog")

	}
	// Field "database"
	if fields["database"] != nil {
		if string(fields["database"]) != "null" {
			if err := json.Unmarshal(fields["database"], &resource.Database); err != nil {
				errs = append(errs, cog.MakeBuildErrors("database", err)...)
			}

		}
		delete(fields, "database")

	}
	// Field "resultReuseEnabled"
	if fields["resultReuseEnabled"] != nil {
		if string(fields["resultReuseEnabled"]) != "null" {
			if err := json.Unmarshal(fields["resultReuseEnabled"], &resource.ResultReuseEnabled); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultReuseEnabled", err)...)
			}

		}
		delete(fields, "resultReuseEnabled")

	}
	// Field "resultReuseMaxAgeInMinutes"
	if fields["resultReuseMaxAgeInMinutes"] != nil {
		if string(fields["resultReuseMaxAgeInMinutes"]) != "null" {
			if err := json.Unmarshal(fields["resultReuseMaxAgeInMinutes"], &resource.ResultReuseMaxAgeInMinutes); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultReuseMaxAgeInMinutes", err)...)
			}

		}
		delete(fields, "resultReuseMaxAgeInMinutes")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConnectionArgs", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConnectionArgs` objects.
func (resource ConnectionArgs) Equals(other ConnectionArgs) bool {
	if resource.Region == nil && other.Region != nil || resource.Region != nil && other.Region == nil {
		return false
	}

	if resource.Region != nil {
		if *resource.Region != *other.Region {
			return false
		}
	}
	if resource.Catalog == nil && other.Catalog != nil || resource.Catalog != nil && other.Catalog == nil {
		return false
	}

	if resource.Catalog != nil {
		if *resource.Catalog != *other.Catalog {
			return false
		}
	}
	if resource.Database == nil && other.Database != nil || resource.Database != nil && other.Database == nil {
		return false
	}

	if resource.Database != nil {
		if *resource.Database != *other.Database {
			return false
		}
	}
	if resource.ResultReuseEnabled == nil && other.ResultReuseEnabled != nil || resource.ResultReuseEnabled != nil && other.ResultReuseEnabled == nil {
		return false
	}

	if resource.ResultReuseEnabled != nil {
		if *resource.ResultReuseEnabled != *other.ResultReuseEnabled {
			return false
		}
	}
	if resource.ResultReuseMaxAgeInMinutes == nil && other.ResultReuseMaxAgeInMinutes != nil || resource.ResultReuseMaxAgeInMinutes != nil && other.ResultReuseMaxAgeInMinutes == nil {
		return false
	}

	if resource.ResultReuseMaxAgeInMinutes != nil {
		if *resource.ResultReuseMaxAgeInMinutes != *other.ResultReuseMaxAgeInMinutes {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConnectionArgs` fields for violations and returns them.
func (resource ConnectionArgs) Validate() error {
	return nil
}

type FormatOptions int64

const (
	FormatOptionsTimeSeries FormatOptions = 0
	FormatOptionsTable      FormatOptions = 1
	FormatOptionsLogs       FormatOptions = 2
)

// VariantConfig returns the configuration related to grafana-athena-datasource dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "grafana-athena-datasource",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery Dataquery
			if cast, ok := input.(*Dataquery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(Dataquery)
			}
			return DataqueryConverter(dataquery)
		},
	}
}
