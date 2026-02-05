---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```php
new QueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```php
aliasBy(string $aliasBy)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> intervalMs

Time interval in milliseconds.

```php
intervalMs(float $intervalMs)
```

### <span class="badge object-method"></span> promQLQuery

PromQL sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery> $promQLQuery

```php
promQLQuery(\Grafana\Foundation\Cog\Builder $promQLQuery)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> sloQuery

SLO sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\SLOQuery> $sloQuery

```php
sloQuery(\Grafana\Foundation\Cog\Builder $sloQuery)
```

### <span class="badge object-method"></span> timeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList> $timeSeriesList

```php
timeSeriesList(\Grafana\Foundation\Cog\Builder $timeSeriesList)
```

### <span class="badge object-method"></span> timeSeriesQuery

Time Series sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery> $timeSeriesQuery

```php
timeSeriesQuery(\Grafana\Foundation\Cog\Builder $timeSeriesQuery)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
