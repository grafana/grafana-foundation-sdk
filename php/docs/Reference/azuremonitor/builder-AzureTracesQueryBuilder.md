---
title: <span class="badge builder"></span> AzureTracesQueryBuilder
---
# <span class="badge builder"></span> AzureTracesQueryBuilder

## Constructor

```php
new AzureTracesQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> filters

Filters for property values.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureTracesFilter>> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> operationId

Operation ID. Used only for Traces queries.

```php
operationId(string $operationId)
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```php
query(string $query)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

@param array<string> $resources

```php
resources(array $resources)
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as.

```php
resultFormat(\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat)
```

### <span class="badge object-method"></span> traceTypes

Types of events to filter by.

@param array<string> $traceTypes

```php
traceTypes(array $traceTypes)
```

## See also

 * <span class="badge object-type-class"></span> [AzureTracesQuery](./object-AzureTracesQuery.md)
