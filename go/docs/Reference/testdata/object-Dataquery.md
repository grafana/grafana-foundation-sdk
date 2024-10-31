---
title: <span class="badge object-type-struct"></span> Dataquery
---
# <span class="badge object-type-struct"></span> Dataquery

## Definition

```go
type Dataquery struct {
    Alias *string `json:"alias,omitempty"`
    ScenarioId *testdata.TestDataQueryType `json:"scenarioId,omitempty"`
    StringInput *string `json:"stringInput,omitempty"`
    Stream *testdata.StreamingQuery `json:"stream,omitempty"`
    PulseWave *testdata.PulseWaveQuery `json:"pulseWave,omitempty"`
    Sim *testdata.SimulationQuery `json:"sim,omitempty"`
    CsvWave []testdata.CSVWave `json:"csvWave,omitempty"`
    Labels *string `json:"labels,omitempty"`
    Lines *int64 `json:"lines,omitempty"`
    LevelColumn *bool `json:"levelColumn,omitempty"`
    Channel *string `json:"channel,omitempty"`
    Nodes *testdata.NodesQuery `json:"nodes,omitempty"`
    CsvFileName *string `json:"csvFileName,omitempty"`
    CsvContent *string `json:"csvContent,omitempty"`
    RawFrameContent *string `json:"rawFrameContent,omitempty"`
    SeriesCount *int32 `json:"seriesCount,omitempty"`
    Usa *testdata.USAQuery `json:"usa,omitempty"`
    ErrorType *testdata.DataqueryErrorType `json:"errorType,omitempty"`
    SpanCount *int32 `json:"spanCount,omitempty"`
    Points [][]testdata.StringOrInt64 `json:"points,omitempty"`
    // Drop percentage (the chance we will lose a point 0-100)
    DropPercent *float64 `json:"dropPercent,omitempty"`
    FlamegraphDiff *bool `json:"flamegraphDiff,omitempty"`
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
