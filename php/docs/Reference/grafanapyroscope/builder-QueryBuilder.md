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

### <span class="badge object-method"></span> groupBy

Allows to group the results.

@param array<string> $groupBy

```php
groupBy(array $groupBy)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> labelSelector

Specifies the query label selectors.

```php
labelSelector(string $labelSelector)
```

### <span class="badge object-method"></span> limit

Sets the maximum number of time series.

```php
limit(int $limit)
```

### <span class="badge object-method"></span> maxNodes

Sets the maximum number of nodes in the flamegraph.

```php
maxNodes(int $maxNodes)
```

### <span class="badge object-method"></span> profileTypeId

Specifies the type of profile to query.

```php
profileTypeId(string $profileTypeId)
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

### <span class="badge object-method"></span> spanSelector

Specifies the query span selectors.

@param array<string> $spanSelector

```php
spanSelector(array $spanSelector)
```

### <span class="badge object-method"></span> version

```php
version(string $version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
