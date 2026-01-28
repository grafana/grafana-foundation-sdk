---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```php
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> candleStyle

Sets the style of the candlesticks

```php
candleStyle(\Grafana\Foundation\Candlestick\CandleStyle $candleStyle)
```

### <span class="badge object-method"></span> colorStrategy

Sets the color strategy for the candlesticks

```php
colorStrategy(\Grafana\Foundation\Candlestick\ColorStrategy $colorStrategy)
```

### <span class="badge object-method"></span> colors

Set which colors are used when the price movement is up or down

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickColors> $colors

```php
colors(\Grafana\Foundation\Cog\Builder $colors)
```

### <span class="badge object-method"></span> fields

Map fields to appropriate dimension

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickFieldMap> $fields

```php
fields(\Grafana\Foundation\Cog\Builder $fields)
```

### <span class="badge object-method"></span> includeAllFields

When enabled, all fields will be sent to the graph

```php
includeAllFields(bool $includeAllFields)
```

### <span class="badge object-method"></span> legend

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> mode

Sets which dimensions are used for the visualization

```php
mode(\Grafana\Foundation\Candlestick\VizDisplayMode $mode)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
