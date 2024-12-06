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

### <span class="badge object-method"></span> columns

```typescript
columns(columns: cog.Builder<bigquery.QueryEditorFunctionExpression>[])
```

### <span class="badge object-method"></span> from

```typescript
from(from: string)
```

### <span class="badge object-method"></span> groupBy

```typescript
groupBy(groupBy: cog.Builder<bigquery.QueryEditorGroupByExpression>[])
```

### <span class="badge object-method"></span> limit

```typescript
limit(limit: number)
```

### <span class="badge object-method"></span> offset

```typescript
offset(offset: number)
```

### <span class="badge object-method"></span> orderBy

```typescript
orderBy(orderBy: cog.Builder<bigquery.QueryEditorPropertyExpression>)
```

### <span class="badge object-method"></span> orderByDirection

```typescript
orderByDirection(orderByDirection: bigquery.OrderByDirection)
```

### <span class="badge object-method"></span> whereString

whereJsonTree?: _

```typescript
whereString(whereString: string)
```

## See also

 * <span class="badge object-type-interface"></span> [SQLExpression](./object-SQLExpression.md)
