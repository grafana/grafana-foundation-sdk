---
title: <span class="badge object-type-class"></span> RowsHeatmapOptions
---
# <span class="badge object-type-class"></span> RowsHeatmapOptions

Controls frame rows options

## Definition

```php
class RowsHeatmapOptions implements \JsonSerializable
{
    /**
     * Sets the name of the cell when not calculating from data
     */
    public ?string $value;

    /**
     * Controls tick alignment when not calculating from data
     */
    public ?\Grafana\Foundation\Common\HeatmapCellLayout $layout;

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

 * <span class="badge builder"></span> [RowsHeatmapOptionsBuilder](./builder-RowsHeatmapOptionsBuilder.md)
