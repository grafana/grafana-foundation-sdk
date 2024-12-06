---
title: <span class="badge builder"></span> AzureMetricDimensionBuilder
---
# <span class="badge builder"></span> AzureMetricDimensionBuilder

## Constructor

```typescript
new AzureMetricDimensionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> dimension

Name of Dimension to be filtered on.

```typescript
dimension(dimension: string)
```

### <span class="badge object-method"></span> filter

@deprecated filter is deprecated in favour of filters to support multiselect.

```typescript
filter(filter: string)
```

### <span class="badge object-method"></span> filters

Values to match with the filter.

```typescript
filters(filters: string[])
```

### <span class="badge object-method"></span> operator

String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.

```typescript
operator(operator: string)
```

## See also

 * <span class="badge object-type-interface"></span> [AzureMetricDimension](./object-AzureMetricDimension.md)
