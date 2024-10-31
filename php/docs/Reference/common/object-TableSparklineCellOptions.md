---
title: <span class="badge object-type-class"></span> TableSparklineCellOptions
---
# <span class="badge object-type-class"></span> TableSparklineCellOptions

Sparkline cell options

## Definition

```php
class TableSparklineCellOptions implements \JsonSerializable
{
    public string $type;

    public ?\Grafana\Foundation\Common\GraphDrawStyle $drawStyle;

    public ?\Grafana\Foundation\Common\GraphGradientMode $gradientMode;

    public ?\Grafana\Foundation\Common\GraphThresholdsStyleConfig $thresholdsStyle;

    public ?string $lineColor;

    public ?float $lineWidth;

    public ?\Grafana\Foundation\Common\LineInterpolation $lineInterpolation;

    public ?\Grafana\Foundation\Common\LineStyle $lineStyle;

    public ?string $fillColor;

    public ?float $fillOpacity;

    public ?\Grafana\Foundation\Common\VisibilityMode $showPoints;

    public ?float $pointSize;

    public ?string $pointColor;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?\Grafana\Foundation\Common\BarAlignment $barAlignment;

    public ?float $barWidthFactor;

    public ?\Grafana\Foundation\Common\StackingConfig $stacking;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?\Grafana\Foundation\Common\GraphTransform $transform;

    /**
     * Indicate if null values should be treated as gaps or connected.
     * When the value is a number, it represents the maximum delta in the
     * X axis that should be considered connected.  For timeseries, this is milliseconds
     * @var bool|float|null
     */
    public $spanNulls;

    public ?string $fillBelowTo;

    public ?string $pointSymbol;

    public ?bool $axisCenteredZero;

    public ?float $barMaxWidth;

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

 * <span class="badge builder"></span> [TableSparklineCellOptionsBuilder](./builder-TableSparklineCellOptionsBuilder.md)
