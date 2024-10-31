---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    public ?\Grafana\Foundation\Xychart\SeriesMapping $seriesMapping;

    /**
     * Table Mode (auto)
     */
    public \Grafana\Foundation\Xychart\XYDimensionConfig $dims;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * Manual Mode
     * @var array<\Grafana\Foundation\Xychart\ScatterSeriesConfig>
     */
    public array $series;

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

