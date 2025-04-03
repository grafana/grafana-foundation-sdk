---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    /**
     * Sets which dimensions are used for the visualization
     */
    public \Grafana\Foundation\Candlestick\VizDisplayMode $mode;

    /**
     * Sets the style of the candlesticks
     */
    public \Grafana\Foundation\Candlestick\CandleStyle $candleStyle;

    /**
     * Sets the color strategy for the candlesticks
     */
    public \Grafana\Foundation\Candlestick\ColorStrategy $colorStrategy;

    /**
     * Map fields to appropriate dimension
     */
    public \Grafana\Foundation\Candlestick\CandlestickFieldMap $fields;

    /**
     * Set which colors are used when the price movement is up or down
     */
    public \Grafana\Foundation\Candlestick\CandlestickColors $colors;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * When enabled, all fields will be sent to the graph
     */
    public ?bool $includeAllFields;

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

