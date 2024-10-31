---
title: <span class="badge object-type-class"></span> FieldConfig
---
# <span class="badge object-type-class"></span> FieldConfig

## Definition

```php
class FieldConfig implements \JsonSerializable
{
    /**
     * Controls line width of the bars.
     */
    public ?int $lineWidth;

    /**
     * Controls the fill opacity of the bars.
     */
    public ?int $fillOpacity;

    /**
     * Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
     * Gradient appearance is influenced by the Fill opacity setting.
     */
    public ?\Grafana\Foundation\Common\GraphGradientMode $gradientMode;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?bool $axisCenteredZero;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    /**
     * Threshold rendering
     */
    public ?\Grafana\Foundation\Common\GraphThresholdsStyleConfig $thresholdsStyle;

    public ?bool $axisBorderShow;

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

