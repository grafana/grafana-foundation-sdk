---
title: <span class="badge builder"></span> AzureMetricDimensionBuilder
---
# <span class="badge builder"></span> AzureMetricDimensionBuilder

## Constructor

```php
new AzureMetricDimensionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> dimension

Name of Dimension to be filtered on.

```php
dimension(string $dimension)
```

### <span class="badge object-method"></span> filter

@deprecated filter is deprecated in favour of filters to support multiselect.

```php
filter(string $filter)
```

### <span class="badge object-method"></span> filters

Values to match with the filter.

@param array<string> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> operator

String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.

```php
operator(string $operator)
```

## See also

 * <span class="badge object-type-class"></span> [AzureMetricDimension](./object-AzureMetricDimension.md)
