---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```go
func NewDataqueryBuilder() *DataqueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataqueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Alias

```go
func (builder *DataqueryBuilder) Alias(alias string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Channel

Used for live query

```go
func (builder *DataqueryBuilder) Channel(channel string) *DataqueryBuilder
```

### <span class="badge object-method"></span> CsvContent

```go
func (builder *DataqueryBuilder) CsvContent(csvContent string) *DataqueryBuilder
```

### <span class="badge object-method"></span> CsvFileName

```go
func (builder *DataqueryBuilder) CsvFileName(csvFileName string) *DataqueryBuilder
```

### <span class="badge object-method"></span> CsvWave

```go
func (builder *DataqueryBuilder) CsvWave(csvWave []cog.Builder[testdata.CSVWave]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Datasource

The datasource

```go
func (builder *DataqueryBuilder) Datasource(datasource dashboard.DataSourceRef) *DataqueryBuilder
```

### <span class="badge object-method"></span> DropPercent

Drop percentage (the chance we will lose a point 0-100)

```go
func (builder *DataqueryBuilder) DropPercent(dropPercent float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> ErrorSource

Possible enum values:

 - `"plugin"` 

 - `"downstream"` 

```go
func (builder *DataqueryBuilder) ErrorSource(errorSource testdata.DataqueryErrorSource) *DataqueryBuilder
```

### <span class="badge object-method"></span> ErrorType

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```go
func (builder *DataqueryBuilder) ErrorType(errorType testdata.DataqueryErrorType) *DataqueryBuilder
```

### <span class="badge object-method"></span> FlamegraphDiff

```go
func (builder *DataqueryBuilder) FlamegraphDiff(flamegraphDiff bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *DataqueryBuilder) IntervalMs(intervalMs float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *DataqueryBuilder) Labels(labels string) *DataqueryBuilder
```

### <span class="badge object-method"></span> LevelColumn

```go
func (builder *DataqueryBuilder) LevelColumn(levelColumn bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Lines

```go
func (builder *DataqueryBuilder) Lines(lines int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Max

```go
func (builder *DataqueryBuilder) Max(max float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *DataqueryBuilder) MaxDataPoints(maxDataPoints int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Min

```go
func (builder *DataqueryBuilder) Min(min float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Nodes

```go
func (builder *DataqueryBuilder) Nodes(nodes cog.Builder[testdata.NodesQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Noise

```go
func (builder *DataqueryBuilder) Noise(noise float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Points

```go
func (builder *DataqueryBuilder) Points(points [][]any) *DataqueryBuilder
```

### <span class="badge object-method"></span> PulseWave

```go
func (builder *DataqueryBuilder) PulseWave(pulseWave cog.Builder[testdata.PulseWaveQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RawFrameContent

```go
func (builder *DataqueryBuilder) RawFrameContent(rawFrameContent string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *DataqueryBuilder) ResultAssertions(resultAssertions cog.Builder[testdata.ResultAssertions]) *DataqueryBuilder
```

### <span class="badge object-method"></span> ScenarioId

Possible enum values:

 - `"annotations"` 

 - `"arrow"` 

 - `"csv_content"` 

 - `"csv_file"` 

 - `"csv_metric_values"` 

 - `"datapoints_outside_range"` 

 - `"error_with_source"` 

 - `"exponential_heatmap_bucket_data"` 

 - `"flame_graph"` 

 - `"grafana_api"` 

 - `"linear_heatmap_bucket_data"` 

 - `"live"` 

 - `"logs"` 

 - `"manual_entry"` 

 - `"no_data_points"` 

 - `"node_graph"` 

 - `"predictable_csv_wave"` 

 - `"predictable_pulse"` 

 - `"random_walk"` 

 - `"random_walk_table"` 

 - `"random_walk_with_error"` 

 - `"raw_frame"` 

 - `"server_error_500"` 

 - `"simulation"` 

 - `"slow_query"` 

 - `"streaming_client"` 

 - `"table_static"` 

 - `"trace"` 

 - `"usa"` 

 - `"variables-query"` 

```go
func (builder *DataqueryBuilder) ScenarioId(scenarioId testdata.DataqueryScenarioId) *DataqueryBuilder
```

### <span class="badge object-method"></span> SeriesCount

```go
func (builder *DataqueryBuilder) SeriesCount(seriesCount int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Sim

```go
func (builder *DataqueryBuilder) Sim(sim cog.Builder[testdata.SimulationQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> SpanCount

```go
func (builder *DataqueryBuilder) SpanCount(spanCount int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Spread

```go
func (builder *DataqueryBuilder) Spread(spread float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> StartValue

```go
func (builder *DataqueryBuilder) StartValue(startValue float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Stream

```go
func (builder *DataqueryBuilder) Stream(stream cog.Builder[testdata.StreamingQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> StringInput

common parameter used by many query types

```go
func (builder *DataqueryBuilder) StringInput(stringInput string) *DataqueryBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *DataqueryBuilder) TimeRange(timeRange cog.Builder[testdata.TimeRange]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Usa

```go
func (builder *DataqueryBuilder) Usa(usa cog.Builder[testdata.USAQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> WithNil

```go
func (builder *DataqueryBuilder) WithNil(withNil bool) *DataqueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
