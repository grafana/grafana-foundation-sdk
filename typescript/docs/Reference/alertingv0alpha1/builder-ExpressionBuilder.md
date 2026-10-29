---
title: <span class="badge builder"></span> ExpressionBuilder
---
# <span class="badge builder"></span> ExpressionBuilder

## Constructor

```typescript
new ExpressionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> datasourceUID

The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource

```typescript
datasourceUID(datasourceUID: alertingv0alpha1.DatasourceUID)
```

### <span class="badge object-method"></span> model

```typescript
model(model: any)
```

### <span class="badge object-method"></span> queryType

The type of query if this is a query expression

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> relativeTimeRange

```typescript
relativeTimeRange(relativeTimeRange: cog.Builder<alertingv0alpha1.RelativeTimeRange>)
```

### <span class="badge object-method"></span> source

Used to mark the expression to be used as the final source for the rule evaluation

Only one expression in a rule can be marked as the source

For AlertRules, this is the expression that will be evaluated against the alerting condition

For RecordingRules, this is the expression that will be recorded

```typescript
source(source: boolean)
```

## See also

 * <span class="badge object-type-interface"></span> [Expression](./object-Expression.md)
