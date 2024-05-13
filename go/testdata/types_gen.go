// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

type TestDataQueryType string

const (
	TestDataQueryTypeRandomWalk                   TestDataQueryType = "random_walk"
	TestDataQueryTypeSlowQuery                    TestDataQueryType = "slow_query"
	TestDataQueryTypeRandomWalkWithError          TestDataQueryType = "random_walk_with_error"
	TestDataQueryTypeRandomWalkTable              TestDataQueryType = "random_walk_table"
	TestDataQueryTypeExponentialHeatmapBucketData TestDataQueryType = "exponential_heatmap_bucket_data"
	TestDataQueryTypeLinearHeatmapBucketData      TestDataQueryType = "linear_heatmap_bucket_data"
	TestDataQueryTypeNoDataPoints                 TestDataQueryType = "no_data_points"
	TestDataQueryTypeDataPointsOutsideRange       TestDataQueryType = "datapoints_outside_range"
	TestDataQueryTypeCSVMetricValues              TestDataQueryType = "csv_metric_values"
	TestDataQueryTypePredictablePulse             TestDataQueryType = "predictable_pulse"
	TestDataQueryTypePredictableCSVWave           TestDataQueryType = "predictable_csv_wave"
	TestDataQueryTypeStreamingClient              TestDataQueryType = "streaming_client"
	TestDataQueryTypeSimulation                   TestDataQueryType = "simulation"
	TestDataQueryTypeUSA                          TestDataQueryType = "usa"
	TestDataQueryTypeLive                         TestDataQueryType = "live"
	TestDataQueryTypeGrafanaAPI                   TestDataQueryType = "grafana_api"
	TestDataQueryTypeArrow                        TestDataQueryType = "arrow"
	TestDataQueryTypeAnnotations                  TestDataQueryType = "annotations"
	TestDataQueryTypeTableStatic                  TestDataQueryType = "table_static"
	TestDataQueryTypeServerError500               TestDataQueryType = "server_error_500"
	TestDataQueryTypeLogs                         TestDataQueryType = "logs"
	TestDataQueryTypeNodeGraph                    TestDataQueryType = "node_graph"
	TestDataQueryTypeFlameGraph                   TestDataQueryType = "flame_graph"
	TestDataQueryTypeRawFrame                     TestDataQueryType = "raw_frame"
	TestDataQueryTypeCSVFile                      TestDataQueryType = "csv_file"
	TestDataQueryTypeCSVContent                   TestDataQueryType = "csv_content"
	TestDataQueryTypeTrace                        TestDataQueryType = "trace"
	TestDataQueryTypeManualEntry                  TestDataQueryType = "manual_entry"
	TestDataQueryTypeVariablesQuery               TestDataQueryType = "variables-query"
)

type StreamingQuery struct {
	Type   StreamingQueryType `json:"type"`
	Speed  int32              `json:"speed"`
	Spread int32              `json:"spread"`
	Noise  int32              `json:"noise"`
	Bands  *int32             `json:"bands,omitempty"`
	Url    *string            `json:"url,omitempty"`
}

type PulseWaveQuery struct {
	TimeStep *int64   `json:"timeStep,omitempty"`
	OnCount  *int64   `json:"onCount,omitempty"`
	OffCount *int64   `json:"offCount,omitempty"`
	OnValue  *float64 `json:"onValue,omitempty"`
	OffValue *float64 `json:"offValue,omitempty"`
}

type SimulationQuery struct {
	Key    Key   `json:"key"`
	Config any   `json:"config,omitempty"`
	Stream *bool `json:"stream,omitempty"`
	Last   *bool `json:"last,omitempty"`
}

type NodesQuery struct {
	Type  *NodesQueryType `json:"type,omitempty"`
	Count *int64          `json:"count,omitempty"`
}

type USAQuery struct {
	Mode   *string  `json:"mode,omitempty"`
	Period *string  `json:"period,omitempty"`
	Fields []string `json:"fields,omitempty"`
	States []string `json:"states,omitempty"`
}

type CSVWave struct {
	TimeStep  *int64  `json:"timeStep,omitempty"`
	Name      *string `json:"name,omitempty"`
	ValuesCSV *string `json:"valuesCSV,omitempty"`
	Labels    *string `json:"labels,omitempty"`
}

// TODO: Should this live here given it's not used in the dataquery?
type Scenario struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	StringInput    string  `json:"stringInput"`
	Description    *string `json:"description,omitempty"`
	HideAliasField *bool   `json:"hideAliasField,omitempty"`
}

type Dataquery struct {
	Alias           *string             `json:"alias,omitempty"`
	ScenarioId      *TestDataQueryType  `json:"scenarioId,omitempty"`
	StringInput     *string             `json:"stringInput,omitempty"`
	Stream          *StreamingQuery     `json:"stream,omitempty"`
	PulseWave       *PulseWaveQuery     `json:"pulseWave,omitempty"`
	Sim             *SimulationQuery    `json:"sim,omitempty"`
	CsvWave         []CSVWave           `json:"csvWave,omitempty"`
	Labels          *string             `json:"labels,omitempty"`
	Lines           *int64              `json:"lines,omitempty"`
	LevelColumn     *bool               `json:"levelColumn,omitempty"`
	Channel         *string             `json:"channel,omitempty"`
	Nodes           *NodesQuery         `json:"nodes,omitempty"`
	CsvFileName     *string             `json:"csvFileName,omitempty"`
	CsvContent      *string             `json:"csvContent,omitempty"`
	RawFrameContent *string             `json:"rawFrameContent,omitempty"`
	SeriesCount     *int32              `json:"seriesCount,omitempty"`
	Usa             *USAQuery           `json:"usa,omitempty"`
	ErrorType       *DataqueryErrorType `json:"errorType,omitempty"`
	SpanCount       *int32              `json:"spanCount,omitempty"`
	Points          [][]StringOrInt64   `json:"points,omitempty"`
	// Drop percentage (the chance we will lose a point 0-100)
	DropPercent *float64 `json:"dropPercent,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId *string `json:"refId,omitempty"`
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

func (resource Dataquery) ImplementsDataqueryVariant() {}

type Key struct {
	Type string  `json:"type"`
	Tick float64 `json:"tick"`
	Uid  *string `json:"uid,omitempty"`
}

type StreamingQueryType string

const (
	StreamingQueryTypeSignal StreamingQueryType = "signal"
	StreamingQueryTypeLogs   StreamingQueryType = "logs"
	StreamingQueryTypeFetch  StreamingQueryType = "fetch"
)

type NodesQueryType string

const (
	NodesQueryTypeRandom      NodesQueryType = "random"
	NodesQueryTypeResponse    NodesQueryType = "response"
	NodesQueryTypeRandomEdges NodesQueryType = "random edges"
)

type DataqueryErrorType string

const (
	DataqueryErrorTypeServerPanic        DataqueryErrorType = "server_panic"
	DataqueryErrorTypeFrontendException  DataqueryErrorType = "frontend_exception"
	DataqueryErrorTypeFrontendObservable DataqueryErrorType = "frontend_observable"
)

type StringOrInt64 struct {
	String *string `json:"String,omitempty"`
	Int64  *int64  `json:"Int64,omitempty"`
}
