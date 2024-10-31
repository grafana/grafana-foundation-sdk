---
title: <span class="badge object-type-class"></span> CellValues
---
# <span class="badge object-type-class"></span> CellValues

Controls cell value options

## Definition

```php
class CellValues implements \JsonSerializable
{
    /**
     * Controls the cell value unit
     */
    public ?string $unit;

    /**
     * Controls the number of decimals for cell values
     */
    public ?float $decimals;

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

 * <span class="badge builder"></span> [CellValuesBuilder](./builder-CellValuesBuilder.md)