---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    /**
     * Controls if the heatmap should be calculated from data
     */
    public ?bool $calculate;

    /**
     * Calculation options for the heatmap
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationOptions $calculation;

    /**
     * Controls the color options
     */
    public \Grafana\Foundation\Heatmap\HeatmapColorOptions $color;

    /**
     * Filters values between a given range
     */
    public ?\Grafana\Foundation\Heatmap\FilterValueRange $filterValues;

    /**
     * Controls tick alignment and value name when not calculating from data
     */
    public ?\Grafana\Foundation\Heatmap\RowsHeatmapOptions $rowsFrame;

    /**
     * | *{
     * 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
     * }
     * Controls the display of the value in the cell
     */
    public \Grafana\Foundation\Common\VisibilityMode $showValue;

    /**
     * Controls gap between cells
     */
    public ?int $cellGap;

    /**
     * Controls cell radius
     */
    public ?float $cellRadius;

    /**
     * Controls cell value unit
     */
    public ?\Grafana\Foundation\Heatmap\CellValues $cellValues;

    /**
     * Controls yAxis placement
     */
    public \Grafana\Foundation\Heatmap\YAxisConfig $yAxis;

    /**
     * | *{
     * 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
     * }
     * Controls legend options
     */
    public \Grafana\Foundation\Heatmap\HeatmapLegend $legend;

    /**
     * Controls tooltip options
     */
    public \Grafana\Foundation\Heatmap\HeatmapTooltip $tooltip;

    /**
     * Controls exemplar options
     */
    public \Grafana\Foundation\Heatmap\ExemplarConfig $exemplars;

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

