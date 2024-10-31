---
title: <span class="badge object-type-struct"></span> Dataquery
---
# <span class="badge object-type-struct"></span> Dataquery

## Definition

```go
type Dataquery struct {
    Alias *string `json:"alias,omitempty"`
    // Used for live query
    Channel *string `json:"channel,omitempty"`
    CsvContent *string `json:"csvContent,omitempty"`
    CsvFileName *string `json:"csvFileName,omitempty"`
    CsvWave []testdata.CSVWave `json:"csvWave,omitempty"`
    // The datasource
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
    // Drop percentage (the chance we will lose a point 0-100)
    DropPercent *float64 `json:"dropPercent,omitempty"`
    // Possible enum values:
    //  - `"frontend_exception"` 
    //  - `"frontend_observable"` 
    //  - `"server_panic"` 
    ErrorType *testdata.DataqueryErrorType `json:"errorType,omitempty"`
    FlamegraphDiff *bool `json:"flamegraphDiff,omitempty"`
    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    Hide *bool `json:"hide,omitempty"`
    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    IntervalMs *float64 `json:"intervalMs,omitempty"`
    Labels *string `json:"labels,omitempty"`
    LevelColumn *bool `json:"levelColumn,omitempty"`
    Lines *int64 `json:"lines,omitempty"`
    Max *float64 `json:"max,omitempty"`
    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
    Min *float64 `json:"min,omitempty"`
    Nodes *testdata.NodesQuery `json:"nodes,omitempty"`
    Noise *float64 `json:"noise,omitempty"`
    Points [][]any `json:"points,omitempty"`
    PulseWave *testdata.PulseWaveQuery `json:"pulseWave,omitempty"`
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    QueryType *string `json:"queryType,omitempty"`
    RawFrameContent *string `json:"rawFrameContent,omitempty"`
    // RefID is the unique identifier of the query, set by the frontend call.
    RefId *string `json:"refId,omitempty"`
    // Optionally define expected query result behavior
    ResultAssertions *testdata.ResultAssertions `json:"resultAssertions,omitempty"`
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
    ScenarioId *testdata.DataqueryScenarioId `json:"scenarioId,omitempty"`
    SeriesCount *int64 `json:"seriesCount,omitempty"`
    Sim *testdata.SimulationQuery `json:"sim,omitempty"`
    SpanCount *int64 `json:"spanCount,omitempty"`
    Spread *float64 `json:"spread,omitempty"`
    StartValue *float64 `json:"startValue,omitempty"`
    Stream *testdata.StreamingQuery `json:"stream,omitempty"`
    // common parameter used by many query types
    StringInput *string `json:"stringInput,omitempty"`
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    TimeRange *testdata.TimeRange `json:"timeRange,omitempty"`
    Usa *testdata.USAQuery `json:"usa,omitempty"`
    WithNil *bool `json:"withNil,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dataquery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dataquery *Dataquery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dataquery` objects.

```go
func (dataquery *Dataquery) Equals(other Dataquery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.

```go
func (dataquery *Dataquery) Validate() error
```

## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
