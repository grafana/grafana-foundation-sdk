---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    /**
     * Manually select which field from the dataset to represent the x field.
     */
    public ?string $xField;

    /**
     * Use the color value for a sibling field to color each bar value.
     */
    public ?string $colorByField;

    /**
     * Controls the orientation of the bar chart, either vertical or horizontal.
     */
    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * Controls the radius of each bar.
     */
    public ?float $barRadius;

    /**
     * Controls the rotation of the x axis labels.
     */
    public int $xTickLabelRotation;

    /**
     * Sets the max length that a label can have before it is truncated.
     */
    public int $xTickLabelMaxLength;

    /**
     * Controls the spacing between x axis labels.
     * negative values indicate backwards skipping behavior
     */
    public ?int $xTickLabelSpacing;

    /**
     * Controls whether bars are stacked or not, either normally or in percent mode.
     */
    public \Grafana\Foundation\Common\StackingMode $stacking;

    /**
     * This controls whether values are shown on top or to the left of bars.
     */
    public \Grafana\Foundation\Common\VisibilityMode $showValue;

    /**
     * Controls the width of bars. 1 = Max width, 0 = Min width.
     */
    public float $barWidth;

    /**
     * Controls the width of groups. 1 = max with, 0 = min width.
     */
    public float $groupWidth;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    /**
     * Enables mode which highlights the entire bar area and shows tooltip when cursor
     * hovers over highlighted area
     */
    public bool $fullHighlight;

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

