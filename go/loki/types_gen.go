// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package loki

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type QueryEditorMode string

const (
	QueryEditorModeCode    QueryEditorMode = "code"
	QueryEditorModeBuilder QueryEditorMode = "builder"
)

type LokiQueryType string

const (
	LokiQueryTypeRange   LokiQueryType = "range"
	LokiQueryTypeInstant LokiQueryType = "instant"
	LokiQueryTypeStream  LokiQueryType = "stream"
)

type SupportingQueryType string

const (
	SupportingQueryTypeLogsVolume     SupportingQueryType = "logsVolume"
	SupportingQueryTypeLogsSample     SupportingQueryType = "logsSample"
	SupportingQueryTypeDataSample     SupportingQueryType = "dataSample"
	SupportingQueryTypeInfiniteScroll SupportingQueryType = "infiniteScroll"
)

type LokiQueryDirection string

const (
	LokiQueryDirectionForward  LokiQueryDirection = "forward"
	LokiQueryDirectionBackward LokiQueryDirection = "backward"
)

type Dataquery struct {
	// The LogQL query.
	Expr string `json:"expr"`
	// Used to override the name of the series.
	LegendFormat *string `json:"legendFormat,omitempty"`
	// Used to limit the number of log rows returned.
	MaxLines *int64 `json:"maxLines,omitempty"`
	// @deprecated, now use step.
	Resolution *int64           `json:"resolution,omitempty"`
	EditorMode *QueryEditorMode `json:"editorMode,omitempty"`
	// @deprecated, now use queryType.
	Range *bool `json:"range,omitempty"`
	// @deprecated, now use queryType.
	Instant *bool `json:"instant,omitempty"`
	// Used to set step value for range queries.
	Step *string `json:"step,omitempty"`
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
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return "loki"
}

// VariantConfig returns the configuration related to loki dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "loki",
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
	// Field "expr"
	if fields["expr"] != nil {
		if string(fields["expr"]) != "null" {
			if err := json.Unmarshal(fields["expr"], &resource.Expr); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expr", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expr", errors.New("required field is null"))...)

		}
		delete(fields, "expr")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expr", errors.New("required field is missing from input"))...)
	}
	// Field "legendFormat"
	if fields["legendFormat"] != nil {
		if string(fields["legendFormat"]) != "null" {
			if err := json.Unmarshal(fields["legendFormat"], &resource.LegendFormat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legendFormat", err)...)
			}

		}
		delete(fields, "legendFormat")

	}
	// Field "maxLines"
	if fields["maxLines"] != nil {
		if string(fields["maxLines"]) != "null" {
			if err := json.Unmarshal(fields["maxLines"], &resource.MaxLines); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxLines", err)...)
			}

		}
		delete(fields, "maxLines")

	}
	// Field "resolution"
	if fields["resolution"] != nil {
		if string(fields["resolution"]) != "null" {
			if err := json.Unmarshal(fields["resolution"], &resource.Resolution); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resolution", err)...)
			}

		}
		delete(fields, "resolution")

	}
	// Field "editorMode"
	if fields["editorMode"] != nil {
		if string(fields["editorMode"]) != "null" {
			if err := json.Unmarshal(fields["editorMode"], &resource.EditorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("editorMode", err)...)
			}

		}
		delete(fields, "editorMode")

	}
	// Field "range"
	if fields["range"] != nil {
		if string(fields["range"]) != "null" {
			if err := json.Unmarshal(fields["range"], &resource.Range); err != nil {
				errs = append(errs, cog.MakeBuildErrors("range", err)...)
			}

		}
		delete(fields, "range")

	}
	// Field "instant"
	if fields["instant"] != nil {
		if string(fields["instant"]) != "null" {
			if err := json.Unmarshal(fields["instant"], &resource.Instant); err != nil {
				errs = append(errs, cog.MakeBuildErrors("instant", err)...)
			}

		}
		delete(fields, "instant")

	}
	// Field "step"
	if fields["step"] != nil {
		if string(fields["step"]) != "null" {
			if err := json.Unmarshal(fields["step"], &resource.Step); err != nil {
				errs = append(errs, cog.MakeBuildErrors("step", err)...)
			}

		}
		delete(fields, "step")

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
	if resource.Expr != other.Expr {
		return false
	}
	if resource.LegendFormat == nil && other.LegendFormat != nil || resource.LegendFormat != nil && other.LegendFormat == nil {
		return false
	}

	if resource.LegendFormat != nil {
		if *resource.LegendFormat != *other.LegendFormat {
			return false
		}
	}
	if resource.MaxLines == nil && other.MaxLines != nil || resource.MaxLines != nil && other.MaxLines == nil {
		return false
	}

	if resource.MaxLines != nil {
		if *resource.MaxLines != *other.MaxLines {
			return false
		}
	}
	if resource.Resolution == nil && other.Resolution != nil || resource.Resolution != nil && other.Resolution == nil {
		return false
	}

	if resource.Resolution != nil {
		if *resource.Resolution != *other.Resolution {
			return false
		}
	}
	if resource.EditorMode == nil && other.EditorMode != nil || resource.EditorMode != nil && other.EditorMode == nil {
		return false
	}

	if resource.EditorMode != nil {
		if *resource.EditorMode != *other.EditorMode {
			return false
		}
	}
	if resource.Range == nil && other.Range != nil || resource.Range != nil && other.Range == nil {
		return false
	}

	if resource.Range != nil {
		if *resource.Range != *other.Range {
			return false
		}
	}
	if resource.Instant == nil && other.Instant != nil || resource.Instant != nil && other.Instant == nil {
		return false
	}

	if resource.Instant != nil {
		if *resource.Instant != *other.Instant {
			return false
		}
	}
	if resource.Step == nil && other.Step != nil || resource.Step != nil && other.Step == nil {
		return false
	}

	if resource.Step != nil {
		if *resource.Step != *other.Step {
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
