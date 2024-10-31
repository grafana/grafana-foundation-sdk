---
title: <span class="badge builder"></span> AzureTracesFilterBuilder
---
# <span class="badge builder"></span> AzureTracesFilterBuilder

## Constructor

```typescript
new AzureTracesFilterBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> filters

Values to filter by.

```typescript
filters(filters: string[])
```

### <span class="badge object-method"></span> operation

Comparison operator to use. Either equals or not equals.

```typescript
operation(operation: string)
```

### <span class="badge object-method"></span> property

Property name, auto-populated based on available traces.

```typescript
property(property: string)
```

## See also

 * <span class="badge object-type-interface"></span> [AzureTracesFilter](./object-AzureTracesFilter.md)
