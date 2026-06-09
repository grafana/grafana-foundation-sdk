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

### <span class="badge object-method"></span> annotationQuery

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\AnnotationQuery> $annotationQuery

```php
annotationQuery(\Grafana\Foundation\Cog\Builder $annotationQuery)
```

### <span class="badge object-method"></span> logsQuery

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\LogsQuery> $logsQuery

```php
logsQuery(\Grafana\Foundation\Cog\Builder $logsQuery)
```

### <span class="badge object-method"></span> metricsQuery

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\MetricsQuery> $metricsQuery

```php
metricsQuery(\Grafana\Foundation\Cog\Builder $metricsQuery)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> labels

@param array<string, string> $labels

```php
labels(array $labels)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
