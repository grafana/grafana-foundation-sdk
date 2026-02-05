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

### <span class="badge object-method"></span> cloudWatchAnnotationQuery

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchAnnotationQuery> $cloudWatchAnnotationQuery

```php
cloudWatchAnnotationQuery(\Grafana\Foundation\Cog\Builder $cloudWatchAnnotationQuery)
```

### <span class="badge object-method"></span> cloudWatchLogsQuery

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery> $cloudWatchLogsQuery

```php
cloudWatchLogsQuery(\Grafana\Foundation\Cog\Builder $cloudWatchLogsQuery)
```

### <span class="badge object-method"></span> cloudWatchMetricsQuery

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchMetricsQuery> $cloudWatchMetricsQuery

```php
cloudWatchMetricsQuery(\Grafana\Foundation\Cog\Builder $cloudWatchMetricsQuery)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
