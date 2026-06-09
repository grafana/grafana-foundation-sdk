---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```typescript
new QueryV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> classicConditions

```typescript
classicConditions(typeClassicConditions: cog.Builder<expr.TypeClassicConditions>)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: Record<string, string>)
```

### <span class="badge object-method"></span> math

```typescript
math(typeMath: cog.Builder<expr.TypeMath>)
```

### <span class="badge object-method"></span> reduce

```typescript
reduce(typeReduce: cog.Builder<expr.TypeReduce>)
```

### <span class="badge object-method"></span> resample

```typescript
resample(typeResample: cog.Builder<expr.TypeResample>)
```

### <span class="badge object-method"></span> sql

```typescript
sql(typeSql: cog.Builder<expr.TypeSql>)
```

### <span class="badge object-method"></span> threshold

```typescript
threshold(typeThreshold: cog.Builder<expr.TypeThreshold>)
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
