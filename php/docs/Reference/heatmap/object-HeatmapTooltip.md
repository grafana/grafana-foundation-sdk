---
title: <span class="badge object-type-class"></span> HeatmapTooltip
---
# <span class="badge object-type-class"></span> HeatmapTooltip

Controls tooltip options

## Definition

```php
class HeatmapTooltip implements \JsonSerializable
{
    /**
     * Controls how the tooltip is shown
     */
    public \Grafana\Foundation\Common\TooltipDisplayMode $mode;

    /**
     * Controls if the tooltip shows a histogram of the y-axis values
     */
    public ?bool $yHistogram;

    /**
     * Controls if the tooltip shows a color scale in header
     */
    public ?bool $showColorScale;

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

 * <span class="badge builder"></span> [HeatmapTooltipBuilder](./builder-HeatmapTooltipBuilder.md)
