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

### <span class="badge object-method"></span> adhocFilters

Additional Ad-hoc filters that take precedence over Scope on conflict.

```java
public QueryV2Builder adhocFilters(List<com.grafana.foundation.cog.Builder<AdhocFilters>> adhocFilters)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> editorMode

what we should show in the editor

Possible enum values:

 - `"builder"` 

 - `"code"` 

```java
public QueryV2Builder editorMode(QueryEditorMode editorMode)
```

### <span class="badge object-method"></span> exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```java
public QueryV2Builder exemplar(Boolean exemplar)
```

### <span class="badge object-method"></span> expr

The actual expression/query that will be evaluated by Prometheus

```java
public QueryV2Builder expr(String expr)
```

### <span class="badge object-method"></span> format

The response format

Possible enum values:

 - `"time_series"` 

 - `"table"` 

 - `"heatmap"` 

```java
public QueryV2Builder format(PromQueryFormat format)
```

### <span class="badge object-method"></span> groupByKeys

Group By parameters to apply to aggregate expressions in the query

```java
public QueryV2Builder groupByKeys(List<String> groupByKeys)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public QueryV2Builder hide(Boolean hide)
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```java
public QueryV2Builder instant(Boolean instant)
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```java
public QueryV2Builder interval(String interval)
```

### <span class="badge object-method"></span> intervalFactor

Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

Deprecated: use interval

```java
public QueryV2Builder intervalFactor(Long intervalFactor)
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

### <span class="badge object-method"></span> legendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```java
public QueryV2Builder legendFormat(String legendFormat)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```java
public QueryV2Builder maxDataPoints(Long maxDataPoints)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```java
public QueryV2Builder queryType(String queryType)
```

### <span class="badge object-method"></span> range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```java
public QueryV2Builder range(Boolean range)
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

### <span class="badge object-method"></span> scopes

A set of filters applied to apply to the query

```java
public QueryV2Builder scopes(List<com.grafana.foundation.cog.Builder<Scopes>> scopes)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```java
public QueryV2Builder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
