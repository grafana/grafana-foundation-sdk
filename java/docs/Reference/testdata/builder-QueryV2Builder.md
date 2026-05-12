---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```java
new QueryV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public QueryV2 build()
```

### <span class="badge object-method"></span> alias

```java
public QueryV2Builder alias(String alias)
```

### <span class="badge object-method"></span> channel

Used for live query

```java
public QueryV2Builder channel(String channel)
```

### <span class="badge object-method"></span> csvContent

```java
public QueryV2Builder csvContent(String csvContent)
```

### <span class="badge object-method"></span> csvFileName

```java
public QueryV2Builder csvFileName(String csvFileName)
```

### <span class="badge object-method"></span> csvWave

```java
public QueryV2Builder csvWave(List<com.grafana.foundation.cog.Builder<CSVWave>> csvWave)
```

### <span class="badge object-method"></span> dataqueryLabels

```java
public QueryV2Builder dataqueryLabels(String labels)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> dropPercent

Drop percentage (the chance we will lose a point 0-100)

```java
public QueryV2Builder dropPercent(Double dropPercent)
```

### <span class="badge object-method"></span> errorSource

Possible enum values:

 - `"plugin"` 

 - `"downstream"` 

```java
public QueryV2Builder errorSource(DataqueryErrorSource errorSource)
```

### <span class="badge object-method"></span> errorType

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```java
public QueryV2Builder errorType(DataqueryErrorType errorType)
```

### <span class="badge object-method"></span> flamegraphDiff

```java
public QueryV2Builder flamegraphDiff(Boolean flamegraphDiff)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public QueryV2Builder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```java
public QueryV2Builder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> labels

```java
public QueryV2Builder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> levelColumn

```java
public QueryV2Builder levelColumn(Boolean levelColumn)
```

### <span class="badge object-method"></span> lines

```java
public QueryV2Builder lines(Long lines)
```

### <span class="badge object-method"></span> max

```java
public QueryV2Builder max(Double max)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```java
public QueryV2Builder maxDataPoints(Long maxDataPoints)
```

### <span class="badge object-method"></span> min

```java
public QueryV2Builder min(Double min)
```

### <span class="badge object-method"></span> nodes

```java
public QueryV2Builder nodes(com.grafana.foundation.cog.Builder<NodesQuery> nodes)
```

### <span class="badge object-method"></span> noise

```java
public QueryV2Builder noise(Double noise)
```

### <span class="badge object-method"></span> points

```java
public QueryV2Builder points(List<List<Object>> points)
```

### <span class="badge object-method"></span> pulseWave

```java
public QueryV2Builder pulseWave(com.grafana.foundation.cog.Builder<PulseWaveQuery> pulseWave)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```java
public QueryV2Builder queryType(String queryType)
```

### <span class="badge object-method"></span> rawFrameContent

```java
public QueryV2Builder rawFrameContent(String rawFrameContent)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```java
public QueryV2Builder refId(String refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```java
public QueryV2Builder resultAssertions(com.grafana.foundation.cog.Builder<ResultAssertions> resultAssertions)
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

```java
public QueryV2Builder scenarioId(DataqueryScenarioId scenarioId)
```

### <span class="badge object-method"></span> seriesCount

```java
public QueryV2Builder seriesCount(Long seriesCount)
```

### <span class="badge object-method"></span> sim

```java
public QueryV2Builder sim(com.grafana.foundation.cog.Builder<SimulationQuery> sim)
```

### <span class="badge object-method"></span> spanCount

```java
public QueryV2Builder spanCount(Long spanCount)
```

### <span class="badge object-method"></span> spread

```java
public QueryV2Builder spread(Double spread)
```

### <span class="badge object-method"></span> startValue

```java
public QueryV2Builder startValue(Double startValue)
```

### <span class="badge object-method"></span> stream

```java
public QueryV2Builder stream(com.grafana.foundation.cog.Builder<StreamingQuery> stream)
```

### <span class="badge object-method"></span> stringInput

common parameter used by many query types

```java
public QueryV2Builder stringInput(String stringInput)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```java
public QueryV2Builder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange)
```

### <span class="badge object-method"></span> usa

```java
public QueryV2Builder usa(com.grafana.foundation.cog.Builder<USAQuery> usa)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

### <span class="badge object-method"></span> withNil

```java
public QueryV2Builder withNil(Boolean withNil)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
