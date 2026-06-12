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

### <span class="badge object-method"></span> adhocFilters

Additional Ad-hoc filters that take precedence over Scope on conflict.

```typescript
adhocFilters(adhocFilters: cog.Builder<prometheus.AdhocFilters>[])
```

### <span class="badge object-method"></span> datasource

The datasource

```typescript
datasource(datasource: common.DataSourceRef)
```

### <span class="badge object-method"></span> editorMode

what we should show in the editor

Possible enum values:

 - `"builder"` 

 - `"code"` 

```typescript
editorMode(editorMode: prometheus.QueryEditorMode)
```

### <span class="badge object-method"></span> exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```typescript
exemplar(exemplar: boolean)
```

### <span class="badge object-method"></span> expr

The actual expression/query that will be evaluated by Prometheus

```typescript
expr(expr: string)
```

### <span class="badge object-method"></span> format

The response format

Possible enum values:

 - `"time_series"` 

 - `"table"` 

 - `"heatmap"` 

```typescript
format(format: prometheus.PromQueryFormat)
```

### <span class="badge object-method"></span> groupByKeys

Group By parameters to apply to aggregate expressions in the query

```typescript
groupByKeys(groupByKeys: string[])
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```typescript
instant()
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```typescript
interval(interval: string)
```

### <span class="badge object-method"></span> intervalFactor

Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

Deprecated: use interval

```typescript
intervalFactor(intervalFactor: number)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```typescript
intervalMs(intervalMs: number)
```

### <span class="badge object-method"></span> legendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```typescript
legendFormat(legendFormat: string)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```typescript
maxDataPoints(maxDataPoints: number)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```typescript
range()
```

### <span class="badge object-method"></span> rangeAndInstant

```typescript
rangeAndInstant()
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```typescript
resultAssertions(resultAssertions: cog.Builder<prometheus.ResultAssertions>)
```

### <span class="badge object-method"></span> scopes

A set of filters applied to apply to the query

```typescript
scopes(scopes: cog.Builder<prometheus.Scopes>[])
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```typescript
timeRange(timeRange: cog.Builder<prometheus.TimeRange>)
```

## See also

 * <span class="badge object-type-interface"></span> [Dataquery](./object-Dataquery.md)
