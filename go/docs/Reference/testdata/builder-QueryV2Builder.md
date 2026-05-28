---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```go
func NewQueryV2Builder() *QueryV2Builder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error)
```

### <span class="badge object-method"></span> Alias

```go
func (builder *QueryV2Builder) Alias(alias string) *QueryV2Builder
```

### <span class="badge object-method"></span> Channel

Used for live query

```go
func (builder *QueryV2Builder) Channel(channel string) *QueryV2Builder
```

### <span class="badge object-method"></span> CsvContent

```go
func (builder *QueryV2Builder) CsvContent(csvContent string) *QueryV2Builder
```

### <span class="badge object-method"></span> CsvFileName

```go
func (builder *QueryV2Builder) CsvFileName(csvFileName string) *QueryV2Builder
```

### <span class="badge object-method"></span> CsvWave

```go
func (builder *QueryV2Builder) CsvWave(csvWave []cog.Builder[testdata.CSVWave]) *QueryV2Builder
```

### <span class="badge object-method"></span> DataqueryLabels

```go
func (builder *QueryV2Builder) DataqueryLabels(labels string) *QueryV2Builder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder
```

### <span class="badge object-method"></span> DropPercent

Drop percentage (the chance we will lose a point 0-100)

```go
func (builder *QueryV2Builder) DropPercent(dropPercent float64) *QueryV2Builder
```

### <span class="badge object-method"></span> ErrorSource

Possible enum values:

 - `"plugin"` 

 - `"downstream"` 

```go
func (builder *QueryV2Builder) ErrorSource(errorSource testdata.DataqueryErrorSource) *QueryV2Builder
```

### <span class="badge object-method"></span> ErrorType

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```go
func (builder *QueryV2Builder) ErrorType(errorType testdata.DataqueryErrorType) *QueryV2Builder
```

### <span class="badge object-method"></span> FlamegraphDiff

```go
func (builder *QueryV2Builder) FlamegraphDiff(flamegraphDiff bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder
```

### <span class="badge object-method"></span> IntervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```go
func (builder *QueryV2Builder) IntervalMs(intervalMs float64) *QueryV2Builder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder
```

### <span class="badge object-method"></span> LevelColumn

```go
func (builder *QueryV2Builder) LevelColumn(levelColumn bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Lines

```go
func (builder *QueryV2Builder) Lines(lines int64) *QueryV2Builder
```

### <span class="badge object-method"></span> Max

```go
func (builder *QueryV2Builder) Max(max float64) *QueryV2Builder
```

### <span class="badge object-method"></span> MaxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```go
func (builder *QueryV2Builder) MaxDataPoints(maxDataPoints int64) *QueryV2Builder
```

### <span class="badge object-method"></span> Min

```go
func (builder *QueryV2Builder) Min(min float64) *QueryV2Builder
```

### <span class="badge object-method"></span> Nodes

```go
func (builder *QueryV2Builder) Nodes(nodes cog.Builder[testdata.NodesQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> Noise

```go
func (builder *QueryV2Builder) Noise(noise float64) *QueryV2Builder
```

### <span class="badge object-method"></span> Points

```go
func (builder *QueryV2Builder) Points(points [][]any) *QueryV2Builder
```

### <span class="badge object-method"></span> PulseWave

```go
func (builder *QueryV2Builder) PulseWave(pulseWave cog.Builder[testdata.PulseWaveQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> QueryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```go
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder
```

### <span class="badge object-method"></span> RawFrameContent

```go
func (builder *QueryV2Builder) RawFrameContent(rawFrameContent string) *QueryV2Builder
```

### <span class="badge object-method"></span> RefId

RefID is the unique identifier of the query, set by the frontend call.

```go
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder
```

### <span class="badge object-method"></span> ResultAssertions

Optionally define expected query result behavior

```go
func (builder *QueryV2Builder) ResultAssertions(resultAssertions cog.Builder[testdata.ResultAssertions]) *QueryV2Builder
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
func (builder *QueryV2Builder) ScenarioId(scenarioId testdata.DataqueryScenarioId) *QueryV2Builder
```

### <span class="badge object-method"></span> SeriesCount

```go
func (builder *QueryV2Builder) SeriesCount(seriesCount int64) *QueryV2Builder
```

### <span class="badge object-method"></span> Sim

```go
func (builder *QueryV2Builder) Sim(sim cog.Builder[testdata.SimulationQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> SpanCount

```go
func (builder *QueryV2Builder) SpanCount(spanCount int64) *QueryV2Builder
```

### <span class="badge object-method"></span> Spread

```go
func (builder *QueryV2Builder) Spread(spread float64) *QueryV2Builder
```

### <span class="badge object-method"></span> StartValue

```go
func (builder *QueryV2Builder) StartValue(startValue float64) *QueryV2Builder
```

### <span class="badge object-method"></span> Stream

```go
func (builder *QueryV2Builder) Stream(stream cog.Builder[testdata.StreamingQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> StringInput

common parameter used by many query types

```go
func (builder *QueryV2Builder) StringInput(stringInput string) *QueryV2Builder
```

### <span class="badge object-method"></span> TimeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```go
func (builder *QueryV2Builder) TimeRange(timeRange cog.Builder[testdata.TimeRange]) *QueryV2Builder
```

### <span class="badge object-method"></span> Usa

```go
func (builder *QueryV2Builder) Usa(usa cog.Builder[testdata.USAQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryV2Builder) Version(version string) *QueryV2Builder
```

### <span class="badge object-method"></span> WithNil

```go
func (builder *QueryV2Builder) WithNil(withNil bool) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
