---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```php
new QueryV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> adhocFilters

Additional Ad-hoc filters that take precedence over Scope on conflict.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\AdhocFilters>> $adhocFilters

```php
adhocFilters(array $adhocFilters)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> editorMode

what we should show in the editor

Possible enum values:

 - `"builder"` 

 - `"code"` 

```php
editorMode(\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode)
```

### <span class="badge object-method"></span> exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```php
exemplar(bool $exemplar)
```

### <span class="badge object-method"></span> expr

The actual expression/query that will be evaluated by Prometheus

```php
expr(string $expr)
```

### <span class="badge object-method"></span> format

The response format

Possible enum values:

 - `"time_series"` 

 - `"table"` 

 - `"heatmap"` 

```php
format(\Grafana\Foundation\Prometheus\PromQueryFormat $format)
```

### <span class="badge object-method"></span> groupByKeys

Group By parameters to apply to aggregate expressions in the query

@param array<string> $groupByKeys

```php
groupByKeys(array $groupByKeys)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```php
instant(bool $instant)
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```php
interval(string $interval)
```

### <span class="badge object-method"></span> intervalFactor

Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

Deprecated: use interval

```php
intervalFactor(int $intervalFactor)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```php
intervalMs(float $intervalMs)
```

### <span class="badge object-method"></span> labels

@param array<string, string> $labels

```php
labels(array $labels)
```

### <span class="badge object-method"></span> legendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```php
legendFormat(string $legendFormat)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```php
maxDataPoints(int $maxDataPoints)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```php
range(bool $range)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\ResultAssertions> $resultAssertions

```php
resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions)
```

### <span class="badge object-method"></span> scopes

A set of filters applied to apply to the query

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\Scopes>> $scopes

```php
scopes(array $scopes)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\TimeRange> $timeRange

```php
timeRange(\Grafana\Foundation\Cog\Builder $timeRange)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
