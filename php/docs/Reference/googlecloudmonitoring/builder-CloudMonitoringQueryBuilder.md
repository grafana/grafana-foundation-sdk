---
title: <span class="badge builder"></span> CloudMonitoringQueryBuilder
---
# <span class="badge builder"></span> CloudMonitoringQueryBuilder

## Constructor

```php
new CloudMonitoringQueryBuilder()
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

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

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

## See also

 * <span class="badge object-type-class"></span> [CloudMonitoringQuery](./object-CloudMonitoringQuery.md)
