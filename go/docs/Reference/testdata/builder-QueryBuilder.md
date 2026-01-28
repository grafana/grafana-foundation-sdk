---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```go
func NewQueryBuilder() *QueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error)
```

### <span class="badge object-method"></span> Alias

```go
func (builder *QueryBuilder) Alias(alias string) *QueryBuilder
```

### <span class="badge object-method"></span> Channel

Used for live query

```go
func (builder *QueryBuilder) Channel(channel string) *QueryBuilder
```

### <span class="badge object-method"></span> CsvContent

```go
func (builder *QueryBuilder) CsvContent(csvContent string) *QueryBuilder
```

### <span class="badge object-method"></span> CsvFileName

```go
func (builder *QueryBuilder) CsvFileName(csvFileName string) *QueryBuilder
```

### <span class="badge object-method"></span> CsvWave

```go
func (builder *QueryBuilder) CsvWave(csvWave []cog.Builder[testdata.CSVWave]) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> DropPercent

Drop percentage (the chance we will lose a point 0-100)

```go
func (builder *QueryBuilder) DropPercent(dropPercent float64) *QueryBuilder
```

### <span class="badge object-method"></span> ErrorSource

Possible enum values:

 - `"plugin"` 

 - `"downstream"` 

```go
func (builder *QueryBuilder) ErrorSource(errorSource testdata.DataqueryErrorSource) *QueryBuilder
```

### <span class="badge object-method"></span> ErrorType

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```go
func (builder *QueryBuilder) ErrorType(errorType testdata.DataqueryErrorType) *QueryBuilder
```

### <span class="badge object-method"></span> FlamegraphDiff

```go
func (builder *QueryBuilder) FlamegraphDiff(flamegraphDiff bool) *QueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *QueryBuilder) IntervalMs(intervalMs float64) *QueryBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryBuilder) Labels(labels string) *QueryBuilder
```

### <span class="badge object-method"></span> LevelColumn

```go
func (builder *QueryBuilder) LevelColumn(levelColumn bool) *QueryBuilder
```

### <span class="badge object-method"></span> Lines

```go
func (builder *QueryBuilder) Lines(lines int64) *QueryBuilder
```

### <span class="badge object-method"></span> Max

```go
func (builder *QueryBuilder) Max(max float64) *QueryBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *QueryBuilder) MaxDataPoints(maxDataPoints int64) *QueryBuilder
```

### <span class="badge object-method"></span> Min

```go
func (builder *QueryBuilder) Min(min float64) *QueryBuilder
```

### <span class="badge object-method"></span> Nodes

```go
func (builder *QueryBuilder) Nodes(nodes cog.Builder[testdata.NodesQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> Noise

```go
func (builder *QueryBuilder) Noise(noise float64) *QueryBuilder
```

### <span class="badge object-method"></span> OldDatasource

The datasource

```go
func (builder *QueryBuilder) OldDatasource(datasource common.DataSourceRef) *QueryBuilder
```

### <span class="badge object-method"></span> Points

```go
func (builder *QueryBuilder) Points(points [][]any) *QueryBuilder
```

### <span class="badge object-method"></span> PulseWave

```go
func (builder *QueryBuilder) PulseWave(pulseWave cog.Builder[testdata.PulseWaveQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> RawFrameContent

```go
func (builder *QueryBuilder) RawFrameContent(rawFrameContent string) *QueryBuilder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *QueryBuilder) ResultAssertions(resultAssertions cog.Builder[testdata.ResultAssertions]) *QueryBuilder
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

 - `"query_meta"` 

 - `"random_walk"` 

 - `"random_walk_table"` 

 - `"random_walk_with_error"` 

 - `"raw_frame"` 

 - `"server_error_500"` 

 - `"steps"` 

 - `"simulation"` 

 - `"slow_query"` 

 - `"streaming_client"` 

 - `"table_static"` 

 - `"trace"` 

 - `"usa"` 

 - `"variables-query"` 

```go
func (builder *QueryBuilder) ScenarioId(scenarioId testdata.DataqueryScenarioId) *QueryBuilder
```

### <span class="badge object-method"></span> SeriesCount

```go
func (builder *QueryBuilder) SeriesCount(seriesCount int64) *QueryBuilder
```

### <span class="badge object-method"></span> Sim

```go
func (builder *QueryBuilder) Sim(sim cog.Builder[testdata.SimulationQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> SpanCount

```go
func (builder *QueryBuilder) SpanCount(spanCount int64) *QueryBuilder
```

### <span class="badge object-method"></span> Spread

```go
func (builder *QueryBuilder) Spread(spread float64) *QueryBuilder
```

### <span class="badge object-method"></span> StartValue

```go
func (builder *QueryBuilder) StartValue(startValue float64) *QueryBuilder
```

### <span class="badge object-method"></span> Stream

```go
func (builder *QueryBuilder) Stream(stream cog.Builder[testdata.StreamingQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> StringInput

common parameter used by many query types

```go
func (builder *QueryBuilder) StringInput(stringInput string) *QueryBuilder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *QueryBuilder) TimeRange(timeRange cog.Builder[testdata.TimeRange]) *QueryBuilder
```

### <span class="badge object-method"></span> Usa

```go
func (builder *QueryBuilder) Usa(usa cog.Builder[testdata.USAQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

### <span class="badge object-method"></span> WithNil

```go
func (builder *QueryBuilder) WithNil(withNil bool) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
