---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```typescript
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> alias

```typescript
alias(alias: string)
```

### <span class="badge object-method"></span> channel

Used for live query

```typescript
channel(channel: string)
```

### <span class="badge object-method"></span> csvContent

```typescript
csvContent(csvContent: string)
```

### <span class="badge object-method"></span> csvFileName

```typescript
csvFileName(csvFileName: string)
```

### <span class="badge object-method"></span> csvWave

```typescript
csvWave(csvWave: cog.Builder<testdata.CSVWave>[])
```

### <span class="badge object-method"></span> datasource

The datasource

```typescript
datasource(datasource: dashboard.DataSourceRef)
```

### <span class="badge object-method"></span> dropPercent

Drop percentage (the chance we will lose a point 0-100)

```typescript
dropPercent(dropPercent: number)
```

### <span class="badge object-method"></span> errorSource

Possible enum values:

 - `"plugin"` 

 - `"downstream"` 

```typescript
errorSource(errorSource: "plugin" | "downstream")
```

### <span class="badge object-method"></span> errorType

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```typescript
errorType(errorType: "frontend_exception" | "frontend_observable" | "server_panic")
```

### <span class="badge object-method"></span> flamegraphDiff

```typescript
flamegraphDiff(flamegraphDiff: boolean)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```typescript
intervalMs(intervalMs: number)
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: string)
```

### <span class="badge object-method"></span> levelColumn

```typescript
levelColumn(levelColumn: boolean)
```

### <span class="badge object-method"></span> lines

```typescript
lines(lines: number)
```

### <span class="badge object-method"></span> max

```typescript
max(max: number)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```typescript
maxDataPoints(maxDataPoints: number)
```

### <span class="badge object-method"></span> min

```typescript
min(min: number)
```

### <span class="badge object-method"></span> nodes

```typescript
nodes(nodes: cog.Builder<testdata.NodesQuery>)
```

### <span class="badge object-method"></span> noise

```typescript
noise(noise: number)
```

### <span class="badge object-method"></span> points

```typescript
points(points: any[][])
```

### <span class="badge object-method"></span> pulseWave

```typescript
pulseWave(pulseWave: cog.Builder<testdata.PulseWaveQuery>)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> rawFrameContent

```typescript
rawFrameContent(rawFrameContent: string)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```typescript
resultAssertions(resultAssertions: cog.Builder<testdata.ResultAssertions>)
```

### <span class="badge object-method"></span> scenarioId

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

```typescript
scenarioId(scenarioId: "annotations" | "arrow" | "csv_content" | "csv_file" | "csv_metric_values" | "datapoints_outside_range" | "error_with_source" | "exponential_heatmap_bucket_data" | "flame_graph" | "grafana_api" | "linear_heatmap_bucket_data" | "live" | "logs" | "manual_entry" | "no_data_points" | "node_graph" | "predictable_csv_wave" | "predictable_pulse" | "random_walk" | "random_walk_table" | "random_walk_with_error" | "raw_frame" | "server_error_500" | "simulation" | "slow_query" | "streaming_client" | "table_static" | "trace" | "usa" | "variables-query")
```

### <span class="badge object-method"></span> seriesCount

```typescript
seriesCount(seriesCount: number)
```

### <span class="badge object-method"></span> sim

```typescript
sim(sim: cog.Builder<testdata.SimulationQuery>)
```

### <span class="badge object-method"></span> spanCount

```typescript
spanCount(spanCount: number)
```

### <span class="badge object-method"></span> spread

```typescript
spread(spread: number)
```

### <span class="badge object-method"></span> startValue

```typescript
startValue(startValue: number)
```

### <span class="badge object-method"></span> stream

```typescript
stream(stream: cog.Builder<testdata.StreamingQuery>)
```

### <span class="badge object-method"></span> stringInput

common parameter used by many query types

```typescript
stringInput(stringInput: string)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```typescript
timeRange(timeRange: cog.Builder<testdata.TimeRange>)
```

### <span class="badge object-method"></span> usa

```typescript
usa(usa: cog.Builder<testdata.USAQuery>)
```

### <span class="badge object-method"></span> withNil

```typescript
withNil(withNil: boolean)
```

## See also

 * <span class="badge object-type-interface"></span> [dataquery](./object-dataquery.md)
