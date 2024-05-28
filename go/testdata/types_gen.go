// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Dataquery struct {
	Alias *string `json:"alias,omitempty"`
	// Used for live query
	Channel     *string   `json:"channel,omitempty"`
	CsvContent  *string   `json:"csvContent,omitempty"`
	CsvFileName *string   `json:"csvFileName,omitempty"`
	CsvWave     []CSVWave `json:"csvWave,omitempty"`
	// The datasource
	Datasource *Datasource `json:"datasource,omitempty"`
	// Drop percentage (the chance we will lose a point 0-100)
	DropPercent *float64 `json:"dropPercent,omitempty"`
	// Possible enum values:
	//  - `"frontend_exception"`
	//  - `"frontend_observable"`
	//  - `"server_panic"`
	ErrorType      *DataqueryErrorType `json:"errorType,omitempty"`
	FlamegraphDiff *bool               `json:"flamegraphDiff,omitempty"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs  *float64 `json:"intervalMs,omitempty"`
	Labels      *string  `json:"labels,omitempty"`
	LevelColumn *bool    `json:"levelColumn,omitempty"`
	Lines       *int64   `json:"lines,omitempty"`
	Max         *float64 `json:"max,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64          `json:"maxDataPoints,omitempty"`
	Min           *float64        `json:"min,omitempty"`
	Nodes         *NodesQuery     `json:"nodes,omitempty"`
	Noise         *float64        `json:"noise,omitempty"`
	Points        [][]any         `json:"points,omitempty"`
	PulseWave     *PulseWaveQuery `json:"pulseWave,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType       *string `json:"queryType,omitempty"`
	RawFrameContent *string `json:"rawFrameContent,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId *string `json:"refId,omitempty"`
	// Optionally define expected query result behavior
	ResultAssertions *ResultAssertions `json:"resultAssertions,omitempty"`
	// Possible enum values:
	//  - `"annotations"`
	//  - `"arrow"`
	//  - `"csv_content"`
	//  - `"csv_file"`
	//  - `"csv_metric_values"`
	//  - `"datapoints_outside_range"`
	//  - `"exponential_heatmap_bucket_data"`
	//  - `"flame_graph"`
	//  - `"grafana_api"`
	//  - `"linear_heatmap_bucket_data"`
	//  - `"live"`
	//  - `"logs"`
	//  - `"manual_entry"`
	//  - `"no_data_points"`
	//  - `"node_graph"`
	//  - `"predictable_csv_wave"`
	//  - `"predictable_pulse"`
	//  - `"random_walk"`
	//  - `"random_walk_table"`
	//  - `"random_walk_with_error"`
	//  - `"raw_frame"`
	//  - `"server_error_500"`
	//  - `"simulation"`
	//  - `"slow_query"`
	//  - `"streaming_client"`
	//  - `"table_static"`
	//  - `"trace"`
	//  - `"usa"`
	//  - `"variables-query"`
	ScenarioId  *DataqueryScenarioId `json:"scenarioId,omitempty"`
	SeriesCount *int64               `json:"seriesCount,omitempty"`
	Sim         *SimulationQuery     `json:"sim,omitempty"`
	SpanCount   *int64               `json:"spanCount,omitempty"`
	Spread      *float64             `json:"spread,omitempty"`
	StartValue  *float64             `json:"startValue,omitempty"`
	Stream      *StreamingQuery      `json:"stream,omitempty"`
	// common parameter used by many query types
	StringInput *string `json:"stringInput,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *TimeRange `json:"timeRange,omitempty"`
	Usa       *USAQuery  `json:"usa,omitempty"`
	WithNil   *bool      `json:"withNil,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := Dataquery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

type CSVWave struct {
	Labels    *string `json:"labels,omitempty"`
	Name      *string `json:"name,omitempty"`
	TimeStep  *int64  `json:"timeStep,omitempty"`
	ValuesCSV *string `json:"valuesCSV,omitempty"`
}

type Datasource struct {
	// The datasource plugin type
	Type string `json:"type"`
	// Datasource UID
	Uid *string `json:"uid,omitempty"`
}

type NodesQuery struct {
	Count *int64 `json:"count,omitempty"`
	Seed  *int64 `json:"seed,omitempty"`
	// Possible enum values:
	//  - `"random"`
	//  - `"random edges"`
	//  - `"response_medium"`
	//  - `"response_small"`
	//  - `"feature_showcase"`
	Type *NodesQueryType `json:"type,omitempty"`
}

type PulseWaveQuery struct {
	OffCount *int64   `json:"offCount,omitempty"`
	OffValue *float64 `json:"offValue,omitempty"`
	OnCount  *int64   `json:"onCount,omitempty"`
	OnValue  *float64 `json:"onValue,omitempty"`
	TimeStep *int64   `json:"timeStep,omitempty"`
}

type ResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *ResultAssertionsType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

type Key struct {
	Tick float64 `json:"tick"`
	Type string  `json:"type"`
	Uid  *string `json:"uid,omitempty"`
}

type SimulationQuery struct {
	Config any   `json:"config,omitempty"`
	Key    Key   `json:"key"`
	Last   *bool `json:"last,omitempty"`
	Stream *bool `json:"stream,omitempty"`
}

type StreamingQuery struct {
	Bands  *int64  `json:"bands,omitempty"`
	Noise  float64 `json:"noise"`
	Speed  float64 `json:"speed"`
	Spread float64 `json:"spread"`
	// Possible enum values:
	//  - `"fetch"`
	//  - `"logs"`
	//  - `"signal"`
	//  - `"traces"`
	Type StreamingQueryType `json:"type"`
	Url  *string            `json:"url,omitempty"`
}

type TimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

type USAQuery struct {
	Fields []string `json:"fields,omitempty"`
	Mode   *string  `json:"mode,omitempty"`
	Period *string  `json:"period,omitempty"`
	States []string `json:"states,omitempty"`
}

type DataqueryErrorType string

const (
	DataqueryErrorTypeFrontendException  DataqueryErrorType = "frontend_exception"
	DataqueryErrorTypeFrontendObservable DataqueryErrorType = "frontend_observable"
	DataqueryErrorTypeServerPanic        DataqueryErrorType = "server_panic"
)

type DataqueryScenarioId string

const (
	DataqueryScenarioIdAnnotations                  DataqueryScenarioId = "annotations"
	DataqueryScenarioIdArrow                        DataqueryScenarioId = "arrow"
	DataqueryScenarioIdCsvContent                   DataqueryScenarioId = "csv_content"
	DataqueryScenarioIdCsvFile                      DataqueryScenarioId = "csv_file"
	DataqueryScenarioIdCsvMetricValues              DataqueryScenarioId = "csv_metric_values"
	DataqueryScenarioIdDatapointsOutsideRange       DataqueryScenarioId = "datapoints_outside_range"
	DataqueryScenarioIdExponentialHeatmapBucketData DataqueryScenarioId = "exponential_heatmap_bucket_data"
	DataqueryScenarioIdFlameGraph                   DataqueryScenarioId = "flame_graph"
	DataqueryScenarioIdGrafanaApi                   DataqueryScenarioId = "grafana_api"
	DataqueryScenarioIdLinearHeatmapBucketData      DataqueryScenarioId = "linear_heatmap_bucket_data"
	DataqueryScenarioIdLive                         DataqueryScenarioId = "live"
	DataqueryScenarioIdLogs                         DataqueryScenarioId = "logs"
	DataqueryScenarioIdManualEntry                  DataqueryScenarioId = "manual_entry"
	DataqueryScenarioIdNoDataPoints                 DataqueryScenarioId = "no_data_points"
	DataqueryScenarioIdNodeGraph                    DataqueryScenarioId = "node_graph"
	DataqueryScenarioIdPredictableCsvWave           DataqueryScenarioId = "predictable_csv_wave"
	DataqueryScenarioIdPredictablePulse             DataqueryScenarioId = "predictable_pulse"
	DataqueryScenarioIdRandomWalk                   DataqueryScenarioId = "random_walk"
	DataqueryScenarioIdRandomWalkTable              DataqueryScenarioId = "random_walk_table"
	DataqueryScenarioIdRandomWalkWithError          DataqueryScenarioId = "random_walk_with_error"
	DataqueryScenarioIdRawFrame                     DataqueryScenarioId = "raw_frame"
	DataqueryScenarioIdServerError500               DataqueryScenarioId = "server_error_500"
	DataqueryScenarioIdSimulation                   DataqueryScenarioId = "simulation"
	DataqueryScenarioIdSlowQuery                    DataqueryScenarioId = "slow_query"
	DataqueryScenarioIdStreamingClient              DataqueryScenarioId = "streaming_client"
	DataqueryScenarioIdTableStatic                  DataqueryScenarioId = "table_static"
	DataqueryScenarioIdTrace                        DataqueryScenarioId = "trace"
	DataqueryScenarioIdUsa                          DataqueryScenarioId = "usa"
	DataqueryScenarioIdVariablesQuery               DataqueryScenarioId = "variables-query"
)

type NodesQueryType string

const (
	NodesQueryTypeRandom          NodesQueryType = "random"
	NodesQueryTypeRandomEdges     NodesQueryType = "random edges"
	NodesQueryTypeResponseMedium  NodesQueryType = "response_medium"
	NodesQueryTypeResponseSmall   NodesQueryType = "response_small"
	NodesQueryTypeFeatureShowcase NodesQueryType = "feature_showcase"
)

type ResultAssertionsType string

const (
	ResultAssertionsTypeNone             ResultAssertionsType = ""
	ResultAssertionsTypeTimeseriesWide   ResultAssertionsType = "timeseries-wide"
	ResultAssertionsTypeTimeseriesLong   ResultAssertionsType = "timeseries-long"
	ResultAssertionsTypeTimeseriesMany   ResultAssertionsType = "timeseries-many"
	ResultAssertionsTypeTimeseriesMulti  ResultAssertionsType = "timeseries-multi"
	ResultAssertionsTypeDirectoryListing ResultAssertionsType = "directory-listing"
	ResultAssertionsTypeTable            ResultAssertionsType = "table"
	ResultAssertionsTypeNumericWide      ResultAssertionsType = "numeric-wide"
	ResultAssertionsTypeNumericMulti     ResultAssertionsType = "numeric-multi"
	ResultAssertionsTypeNumericLong      ResultAssertionsType = "numeric-long"
	ResultAssertionsTypeLogLines         ResultAssertionsType = "log-lines"
)

type StreamingQueryType string

const (
	StreamingQueryTypeFetch  StreamingQueryType = "fetch"
	StreamingQueryTypeLogs   StreamingQueryType = "logs"
	StreamingQueryTypeSignal StreamingQueryType = "signal"
	StreamingQueryTypeTraces StreamingQueryType = "traces"
)
