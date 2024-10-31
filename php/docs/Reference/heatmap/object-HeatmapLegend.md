---
title: <span class="badge object-type-class"></span> HeatmapLegend
---
# <span class="badge object-type-class"></span> HeatmapLegend

Controls legend options

## Definition

```php
class HeatmapLegend implements \JsonSerializable
{
    /**
     * Controls if the legend is shown
     */
    public bool $show;

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

 * <span class="badge builder"></span> [HeatmapLegendBuilder](./builder-HeatmapLegendBuilder.md)
