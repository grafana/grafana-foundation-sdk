---
title: <span class="badge builder"></span> SQLExpressionBuilder
---
# <span class="badge builder"></span> SQLExpressionBuilder

## Constructor

```typescript
new SQLExpressionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> from

FROM part of the SQL expression

```typescript
from(from: cog.Builder<cloudwatch.QueryEditorPropertyExpression> | cog.Builder<cloudwatch.QueryEditorFunctionExpression>)
```

### <span class="badge object-method"></span> groupBy

GROUP BY part of the SQL expression

```typescript
groupBy(groupBy: cog.Builder<cloudwatch.QueryEditorArrayExpression>)
```

### <span class="badge object-method"></span> limit

LIMIT part of the SQL expression

```typescript
limit(limit: number)
```

### <span class="badge object-method"></span> orderBy

ORDER BY part of the SQL expression

```typescript
orderBy(orderBy: cog.Builder<cloudwatch.QueryEditorFunctionExpression>)
```

### <span class="badge object-method"></span> orderByDirection

The sort order of the SQL expression, `ASC` or `DESC`

```typescript
orderByDirection(orderByDirection: string)
```

### <span class="badge object-method"></span> select

SELECT part of the SQL expression

```typescript
select(select: cog.Builder<cloudwatch.QueryEditorFunctionExpression>)
```

### <span class="badge object-method"></span> where

WHERE part of the SQL expression

```typescript
where(where: cog.Builder<cloudwatch.QueryEditorArrayExpression>)
```

## See also

 * <span class="badge object-type-interface"></span> [SQLExpression](./object-SQLExpression.md)
