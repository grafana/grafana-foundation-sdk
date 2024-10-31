---
title: <span class="badge object-type-class"></span> ScatterSeriesConfig
---
# <span class="badge object-type-class"></span> ScatterSeriesConfig

## Definition

```php
class ScatterSeriesConfig implements \JsonSerializable
{
    public ?string $x;

    public ?string $y;

    public ?\Grafana\Foundation\Xychart\ScatterShow $show;

    public ?\Grafana\Foundation\Common\ScaleDimensionConfig $pointSize;

    public ?\Grafana\Foundation\Common\ColorDimensionConfig $pointColor;

    public ?\Grafana\Foundation\Common\ColorDimensionConfig $lineColor;

    public ?int $lineWidth;

    public ?\Grafana\Foundation\Common\LineStyle $lineStyle;

    public ?\Grafana\Foundation\Common\VisibilityMode $label;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?string $name;

    public ?\Grafana\Foundation\Common\TextDimensionConfig $labelValue;

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

 * <span class="badge builder"></span> [ScatterSeriesConfigBuilder](./builder-ScatterSeriesConfigBuilder.md)
