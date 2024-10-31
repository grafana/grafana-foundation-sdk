---
title: <span class="badge builder"></span> RangeMapBuilder
---
# <span class="badge builder"></span> RangeMapBuilder

## Constructor

```typescript
new RangeMapBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> options

Range to match against and the result to apply when the value is within the range

```typescript
options(options: {
	// Min value of the range. It can be null which means -Infinity
	from: number | null;
	// Max value of the range. It can be null which means +Infinity
	to: number | null;
	// Config to apply when the value is within the range
	result: dashboard.ValueMappingResult;
})
```

## See also

 * <span class="badge object-type-interface"></span> [RangeMap](./object-RangeMap.md)
