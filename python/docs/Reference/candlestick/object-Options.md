---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Sets which dimensions are used for the visualization
    mode: candlestick.VizDisplayMode
    # Sets the style of the candlesticks
    candle_style: candlestick.CandleStyle
    # Sets the color strategy for the candlesticks
    color_strategy: candlestick.ColorStrategy
    # Map fields to appropriate dimension
    fields: candlestick.CandlestickFieldMap
    # Set which colors are used when the price movement is up or down
    colors: candlestick.CandlestickColors
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    # When enabled, all fields will be sent to the graph
    include_all_fields: typing.Optional[bool]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

