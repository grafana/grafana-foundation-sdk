---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```typescript
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> candleStyle

Sets the style of the candlesticks

```typescript
candleStyle(candleStyle: candlestick.CandleStyle)
```

### <span class="badge object-method"></span> colorStrategy

Sets the color strategy for the candlesticks

```typescript
colorStrategy(colorStrategy: candlestick.ColorStrategy)
```

### <span class="badge object-method"></span> colors

Set which colors are used when the price movement is up or down

```typescript
colors(colors: cog.Builder<candlestick.CandlestickColors>)
```

### <span class="badge object-method"></span> fields

Map fields to appropriate dimension

```typescript
fields(fields: cog.Builder<candlestick.CandlestickFieldMap>)
```

### <span class="badge object-method"></span> includeAllFields

When enabled, all fields will be sent to the graph

```typescript
includeAllFields(includeAllFields: boolean)
```

### <span class="badge object-method"></span> legend

```typescript
legend(legend: cog.Builder<common.VizLegendOptions>)
```

### <span class="badge object-method"></span> mode

Sets which dimensions are used for the visualization

```typescript
mode(mode: candlestick.VizDisplayMode)
```

### <span class="badge object-method"></span> tooltip

```typescript
tooltip(tooltip: cog.Builder<common.VizTooltipOptions>)
```

## See also

 * <span class="badge object-type-interface"></span> [Options](./object-Options.md)
