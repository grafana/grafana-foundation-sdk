---
title: <span class="badge builder"></span> DataQueryKindBuilder
---
# <span class="badge builder"></span> DataQueryKindBuilder

## Constructor

```typescript
new DataQueryKindBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> group

```typescript
group(group: string)
```

### <span class="badge object-method"></span> spec

```typescript
spec(spec: any)
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [DataQueryKind](./object-DataQueryKind.md)
