---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```java
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Dataquery build()
```

### <span class="badge object-method"></span> alias

```java
public dataqueryBuilder alias(String alias)
```

### <span class="badge object-method"></span> channel

Used for live query

```java
public dataqueryBuilder channel(String channel)
```

### <span class="badge object-method"></span> csvContent

```java
public dataqueryBuilder csvContent(String csvContent)
```

### <span class="badge object-method"></span> csvFileName

```java
public dataqueryBuilder csvFileName(String csvFileName)
```

### <span class="badge object-method"></span> csvWave

```java
public dataqueryBuilder csvWave(List<CSVWave> csvWave)
```

### <span class="badge object-method"></span> datasource

The datasource

```java
public dataqueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> dropPercent

Drop percentage (the chance we will lose a point 0-100)

```java
public dataqueryBuilder dropPercent(Double dropPercent)
```

### <span class="badge object-method"></span> errorType

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```java
public dataqueryBuilder errorType(DataqueryErrorType errorType)
```

### <span class="badge object-method"></span> flamegraphDiff

```java
public dataqueryBuilder flamegraphDiff(Boolean flamegraphDiff)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public dataqueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```java
public dataqueryBuilder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> labels

```java
public dataqueryBuilder labels(String labels)
```

### <span class="badge object-method"></span> levelColumn

```java
public dataqueryBuilder levelColumn(Boolean levelColumn)
```

### <span class="badge object-method"></span> lines

```java
public dataqueryBuilder lines(Long lines)
```

### <span class="badge object-method"></span> max

```java
public dataqueryBuilder max(Double max)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```java
public dataqueryBuilder maxDataPoints(Long maxDataPoints)
```

### <span class="badge object-method"></span> min

```java
public dataqueryBuilder min(Double min)
```

### <span class="badge object-method"></span> nodes

```java
public dataqueryBuilder nodes(NodesQuery nodes)
```

### <span class="badge object-method"></span> noise

```java
public dataqueryBuilder noise(Double noise)
```

### <span class="badge object-method"></span> points

```java
public dataqueryBuilder points(List<List<Object>> points)
```

### <span class="badge object-method"></span> pulseWave

```java
public dataqueryBuilder pulseWave(PulseWaveQuery pulseWave)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```java
public dataqueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> rawFrameContent

```java
public dataqueryBuilder rawFrameContent(String rawFrameContent)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```java
public dataqueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```java
public dataqueryBuilder resultAssertions(ResultAssertions resultAssertions)
```

### <span class="badge object-method"></span> scenarioId

Possible enum values:

 - `"annotations"` 

 - `"arrow"` 

 - `"csv_content"` 

 - `"csv_file"` 

 - `"csv_metric_values"` 

 - `"datapoints_outside_range"` 

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

```java
public dataqueryBuilder scenarioId(DataqueryScenarioId scenarioId)
```

### <span class="badge object-method"></span> seriesCount

```java
public dataqueryBuilder seriesCount(Long seriesCount)
```

### <span class="badge object-method"></span> sim

```java
public dataqueryBuilder sim(SimulationQuery sim)
```

### <span class="badge object-method"></span> spanCount

```java
public dataqueryBuilder spanCount(Long spanCount)
```

### <span class="badge object-method"></span> spread

```java
public dataqueryBuilder spread(Double spread)
```

### <span class="badge object-method"></span> startValue

```java
public dataqueryBuilder startValue(Double startValue)
```

### <span class="badge object-method"></span> stream

```java
public dataqueryBuilder stream(StreamingQuery stream)
```

### <span class="badge object-method"></span> stringInput

common parameter used by many query types

```java
public dataqueryBuilder stringInput(String stringInput)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```java
public dataqueryBuilder timeRange(TimeRange timeRange)
```

### <span class="badge object-method"></span> usa

```java
public dataqueryBuilder usa(USAQuery usa)
```

### <span class="badge object-method"></span> withNil

```java
public dataqueryBuilder withNil(Boolean withNil)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
