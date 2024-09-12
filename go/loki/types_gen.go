// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package loki

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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

func (resource Dataquery) ImplementsDataqueryVariant() {}

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "loki",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := Dataquery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}
