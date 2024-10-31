---
title: <span class="badge builder"></span> TempoQueryBuilder
---
# <span class="badge builder"></span> TempoQueryBuilder

## Constructor

```php
new TempoQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> filters

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

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

### <span class="badge object-method"></span> minDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```php
minDuration(string $minDuration)
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

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}

```php
serviceMapQuery(string $serviceMapQuery)
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

## See also

 * <span class="badge object-type-class"></span> [TempoQuery](./object-TempoQuery.md)