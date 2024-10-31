---
title: <span class="badge object-type-class"></span> YAxisConfig
---
# <span class="badge object-type-class"></span> YAxisConfig

Configuration options for the yAxis

## Definition

```php
class YAxisConfig implements \JsonSerializable
{
    /**
     * Sets the yAxis unit
     */
    public ?string $unit;

    /**
     * Reverses the yAxis
     */
    public ?bool $reverse;

    /**
     * Controls the number of decimals for yAxis values
     */
    public ?float $decimals;

    /**
     * Sets the minimum value for the yAxis
     */
    public ?float $min;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    /**
     * Sets the maximum value for the yAxis
     */
    public ?float $max;

    public ?bool $axisCenteredZero;

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

 * <span class="badge builder"></span> [YAxisConfigBuilder](./builder-YAxisConfigBuilder.md)
