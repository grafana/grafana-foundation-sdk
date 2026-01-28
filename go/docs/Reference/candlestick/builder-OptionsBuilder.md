---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> CandleStyle

Sets the style of the candlesticks

```go
func (builder *OptionsBuilder) CandleStyle(candleStyle candlestick.CandleStyle) *OptionsBuilder
```

### <span class="badge object-method"></span> ColorStrategy

Sets the color strategy for the candlesticks

```go
func (builder *OptionsBuilder) ColorStrategy(colorStrategy candlestick.ColorStrategy) *OptionsBuilder
```

### <span class="badge object-method"></span> Colors

Set which colors are used when the price movement is up or down

```go
func (builder *OptionsBuilder) Colors(colors cog.Builder[candlestick.CandlestickColors]) *OptionsBuilder
```

### <span class="badge object-method"></span> Fields

Map fields to appropriate dimension

```go
func (builder *OptionsBuilder) Fields(fields cog.Builder[candlestick.CandlestickFieldMap]) *OptionsBuilder
```

### <span class="badge object-method"></span> IncludeAllFields

When enabled, all fields will be sent to the graph

```go
func (builder *OptionsBuilder) IncludeAllFields(includeAllFields bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Mode

Sets which dimensions are used for the visualization

```go
func (builder *OptionsBuilder) Mode(mode candlestick.VizDisplayMode) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
