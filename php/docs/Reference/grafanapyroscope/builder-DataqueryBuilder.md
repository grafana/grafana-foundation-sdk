---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```php
new DataqueryBuilder()
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

### <span class="badge object-method"></span> groupBy

Allows to group the results.

@param array<string> $groupBy

```php
groupBy(array $groupBy)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> labelSelector

Specifies the query label selectors.

```php
labelSelector(string $labelSelector)
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

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)