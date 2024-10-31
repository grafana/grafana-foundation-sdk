---
title: <span class="badge builder"></span> RegexMapBuilder
---
# <span class="badge builder"></span> RegexMapBuilder

## Constructor

```typescript
new RegexMapBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> options

Regular expression to match against and the result to apply when the value matches the regex

```typescript
options(options: {
	// Regular expression to match against
	pattern: string;
	// Config to apply when the value matches the regex
	result: dashboard.ValueMappingResult;
})
```

## See also

 * <span class="badge object-type-interface"></span> [RegexMap](./object-RegexMap.md)
