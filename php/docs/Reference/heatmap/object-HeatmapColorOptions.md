---
title: <span class="badge object-type-class"></span> HeatmapColorOptions
---
# <span class="badge object-type-class"></span> HeatmapColorOptions

Controls various color options

## Definition

```php
class HeatmapColorOptions implements \JsonSerializable
{
    /**
     * Sets the color mode
     */
    public ?\Grafana\Foundation\Heatmap\HeatmapColorMode $mode;

    /**
     * Controls the color scheme used
     */
    public string $scheme;

    /**
     * Controls the color fill when in opacity mode
     */
    public string $fill;

    /**
     * Controls the color scale
     */
    public ?\Grafana\Foundation\Heatmap\HeatmapColorScale $scale;

    /**
     * Controls the exponent when scale is set to exponential
     */
    public float $exponent;

    /**
     * Controls the number of color steps
     */
    public int $steps;

    /**
     * Reverses the color scheme
     */
    public bool $reverse;

    /**
     * Sets the minimum value for the color scale
     */
    public ?float $min;

    /**
     * Sets the maximum value for the color scale
     */
    public ?float $max;

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

 * <span class="badge builder"></span> [HeatmapColorOptionsBuilder](./builder-HeatmapColorOptionsBuilder.md)
