// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package loki

import (
	"encoding/json"

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
	SupportingQueryTypeLogsVolume SupportingQueryType = "logsVolume"
	SupportingQueryTypeLogsSample SupportingQueryType = "logsSample"
	SupportingQueryTypeDataSample SupportingQueryType = "dataSample"
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
