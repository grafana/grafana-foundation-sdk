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

### <span class="badge object-method"></span> from

FROM part of the SQL expression

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression> $from

```php
from($from)
```

### <span class="badge object-method"></span> groupBy

GROUP BY part of the SQL expression

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression> $groupBy

```php
groupBy(\Grafana\Foundation\Cog\Builder $groupBy)
```

### <span class="badge object-method"></span> limit

LIMIT part of the SQL expression

```php
limit(int $limit)
```

### <span class="badge object-method"></span> orderBy

ORDER BY part of the SQL expression

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression> $orderBy

```php
orderBy(\Grafana\Foundation\Cog\Builder $orderBy)
```

### <span class="badge object-method"></span> orderByDirection

The sort order of the SQL expression, `ASC` or `DESC`

```php
orderByDirection(string $orderByDirection)
```

### <span class="badge object-method"></span> select

SELECT part of the SQL expression

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression> $select

```php
select(\Grafana\Foundation\Cog\Builder $select)
```

### <span class="badge object-method"></span> where

WHERE part of the SQL expression

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression> $where

```php
where(\Grafana\Foundation\Cog\Builder $where)
```

## See also

 * <span class="badge object-type-class"></span> [SQLExpression](./object-SQLExpression.md)
