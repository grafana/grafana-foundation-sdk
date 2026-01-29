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

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource

```php
datasource(\Grafana\Foundation\Cog\Builder $datasource)
```

### <span class="badge object-method"></span> exemplars

For metric queries, how many exemplars to request, 0 means no exemplars

```php
exemplars(int $exemplars)
```

### <span class="badge object-method"></span> filters

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> groupBy

deprecated Filters that are used to query the metrics summary

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>> $groupBy

```php
groupBy(array $groupBy)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> limit

Defines the maximum number of traces that are returned from Tempo

```php
limit(int $limit)
```

### <span class="badge object-method"></span> maxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```php
maxDuration(string $maxDuration)
```

### <span class="badge object-method"></span> metricsQueryType

For metric queries, whether to run instant or range queries

```php
metricsQueryType(\Grafana\Foundation\Tempo\MetricsQueryType $metricsQueryType)
```

### <span class="badge object-method"></span> minDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```php
minDuration(string $minDuration)
```

### <span class="badge object-method"></span> oldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
oldDatasource(\Grafana\Foundation\Common\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> query

TraceQL query or trace ID

```php
query(string $query)
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

### <span class="badge object-method"></span> search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```php
search(string $search)
```

### <span class="badge object-method"></span> serviceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```php
serviceMapIncludeNamespace(bool $serviceMapIncludeNamespace)
```

### <span class="badge object-method"></span> serviceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

@param string|array<string> $serviceMapQuery

```php
serviceMapQuery($serviceMapQuery)
```

### <span class="badge object-method"></span> serviceName

@deprecated Query traces by service name

```php
serviceName(string $serviceName)
```

### <span class="badge object-method"></span> spanName

@deprecated Query traces by span name

```php
spanName(string $spanName)
```

### <span class="badge object-method"></span> spss

Defines the maximum number of spans per spanset that are returned from Tempo

```php
spss(int $spss)
```

### <span class="badge object-method"></span> step

For metric queries, the step size to use

```php
step(string $step)
```

### <span class="badge object-method"></span> tableType

The type of the table that is used to display the search results

```php
tableType(\Grafana\Foundation\Tempo\SearchTableType $tableType)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
