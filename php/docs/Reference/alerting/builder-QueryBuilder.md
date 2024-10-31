---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```php
new QueryBuilder(string $refId)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> datasourceUid

Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.

```php
datasourceUid(string $datasourceUid)
```

### <span class="badge object-method"></span> model

JSON is the raw JSON query and includes the above properties as well as custom properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cog\Dataquery> $model

```php
model(\Grafana\Foundation\Cog\Builder $model)
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

### <span class="badge object-method"></span> relativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

```php
relativeTimeRange(\Grafana\Foundation\Alerting\RelativeTimeRange $relativeTimeRange)
```

## See also

 * <span class="badge object-type-class"></span> [Query](./object-Query.md)
