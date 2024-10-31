---
title: <span class="badge object-type-class"></span> AzureMetricDimension
---
# <span class="badge object-type-class"></span> AzureMetricDimension

## Definition

```php
class AzureMetricDimension implements \JsonSerializable
{
    /**
     * Name of Dimension to be filtered on.
     */
    public ?string $dimension;

    /**
     * String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
     */
    public ?string $operator;

    /**
     * Values to match with the filter.
     * @var array<string>|null
     */
    public ?array $filters;

    /**
     * @deprecated filter is deprecated in favour of filters to support multiselect.
     */
    public ?string $filter;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [AzureMetricDimensionBuilder](./builder-AzureMetricDimensionBuilder.md)
