---
title: <span class="badge builder"></span> Options
---
# <span class="badge builder"></span> Options

## Constructor

```python
Options()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> candlestick.Options
```

### <span class="badge object-method"></span> candle_style

Sets the style of the candlesticks

```python
def candle_style(candle_style: candlestick.CandleStyle) -> typing.Self
```

### <span class="badge object-method"></span> color_strategy

Sets the color strategy for the candlesticks

```python
def color_strategy(color_strategy: candlestick.ColorStrategy) -> typing.Self
```

### <span class="badge object-method"></span> colors

Set which colors are used when the price movement is up or down

```python
def colors(colors: cogbuilder.Builder[candlestick.CandlestickColors]) -> typing.Self
```

### <span class="badge object-method"></span> fields

Map fields to appropriate dimension

```python
def fields(fields: cogbuilder.Builder[candlestick.CandlestickFieldMap]) -> typing.Self
```

### <span class="badge object-method"></span> include_all_fields

When enabled, all fields will be sent to the graph

```python
def include_all_fields(include_all_fields: bool) -> typing.Self
```

### <span class="badge object-method"></span> legend

```python
def legend(legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self
```

### <span class="badge object-method"></span> mode

Sets which dimensions are used for the visualization

```python
def mode(mode: candlestick.VizDisplayMode) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
