---
title: <span class="badge builder"></span> TypeResampleBuilder
---
# <span class="badge builder"></span> TypeResampleBuilder

## Constructor

```php
new TypeResampleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> datasource

The datasource

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> downsampler

The downsample function

Possible enum values:

 - `"sum"` 

 - `"mean"` 

 - `"min"` 

 - `"max"` 

 - `"count"` 

 - `"last"` 

 - `"median"` 

```php
downsampler(\Grafana\Foundation\Expr\TypeResampleDownsampler $downsampler)
```

### <span class="badge object-method"></span> expression

The math expression

```php
expression(string $expression)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```php
intervalMs(float $intervalMs)
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

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeResampleResultAssertions> $resultAssertions

```php
resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeResampleTimeRange> $timeRange

```php
timeRange(\Grafana\Foundation\Cog\Builder $timeRange)
```

### <span class="badge object-method"></span> upsampler

The upsample function

Possible enum values:

 - `"pad"` Use the last seen value

 - `"backfilling"` backfill

 - `"fillna"` Do not fill values (nill)

```php
upsampler(\Grafana\Foundation\Expr\TypeResampleUpsampler $upsampler)
```

### <span class="badge object-method"></span> window

The time duration

```php
window(string $window)
```

## See also

 * <span class="badge object-type-class"></span> [TypeResample](./object-TypeResample.md)
