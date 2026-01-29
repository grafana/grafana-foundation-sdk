---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```typescript
new QueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> typeClassicConditions

```typescript
typeClassicConditions(typeClassicConditions: cog.Builder<expr.TypeClassicConditions>)
```

### <span class="badge object-method"></span> typeMath

```typescript
typeMath(typeMath: cog.Builder<expr.TypeMath>)
```

### <span class="badge object-method"></span> typeReduce

```typescript
typeReduce(typeReduce: cog.Builder<expr.TypeReduce>)
```

### <span class="badge object-method"></span> typeResample

```typescript
typeResample(typeResample: cog.Builder<expr.TypeResample>)
```

### <span class="badge object-method"></span> typeSql

```typescript
typeSql(typeSql: cog.Builder<expr.TypeSql>)
```

### <span class="badge object-method"></span> typeThreshold

```typescript
typeThreshold(typeThreshold: cog.Builder<expr.TypeThreshold>)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
