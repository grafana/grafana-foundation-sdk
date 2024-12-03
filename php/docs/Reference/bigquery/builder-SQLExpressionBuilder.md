---
title: <span class="badge builder"></span> SQLExpressionBuilder
---
# <span class="badge builder"></span> SQLExpressionBuilder

## Constructor

```php
new SQLExpressionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> columns

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorFunctionExpression>> $columns

```php
columns(array $columns)
```

### <span class="badge object-method"></span> from

```php
from(string $from)
```

### <span class="badge object-method"></span> groupBy

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorGroupByExpression>> $groupBy

```php
groupBy(array $groupBy)
```

### <span class="badge object-method"></span> limit

```php
limit(int $limit)
```

### <span class="badge object-method"></span> offset

```php
offset(int $offset)
```

### <span class="badge object-method"></span> orderBy

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorPropertyExpression> $orderBy

```php
orderBy(\Grafana\Foundation\Cog\Builder $orderBy)
```

### <span class="badge object-method"></span> orderByDirection

```php
orderByDirection(\Grafana\Foundation\Bigquery\OrderByDirection $orderByDirection)
```

### <span class="badge object-method"></span> whereString

whereJsonTree?: _

```php
whereString(string $whereString)
```

## See also

 * <span class="badge object-type-class"></span> [SQLExpression](./object-SQLExpression.md)
