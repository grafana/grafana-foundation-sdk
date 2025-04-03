---
title: <span class="badge object-type-class"></span> FilterValueRange
---
# <span class="badge object-type-class"></span> FilterValueRange

Controls the value filter range

## Definition

```php
class FilterValueRange implements \JsonSerializable
{
    /**
     * Sets the filter range to values less than or equal to the given value
     */
    public ?float $le;

    /**
     * Sets the filter range to values greater than or equal to the given value
     */
    public ?float $ge;

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

 * <span class="badge builder"></span> [FilterValueRangeBuilder](./builder-FilterValueRangeBuilder.md)
