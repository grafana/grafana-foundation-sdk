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

### <span class="badge object-method"></span> dims

Table Mode (auto)

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XYDimensionConfig> $dims

```php
dims(\Grafana\Foundation\Cog\Builder $dims)
```

### <span class="badge object-method"></span> legend

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> series

Manual Mode

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\ScatterSeriesConfig>> $series

```php
series(array $series)
```

### <span class="badge object-method"></span> seriesMapping

```php
seriesMapping(\Grafana\Foundation\Xychart\SeriesMapping $seriesMapping)
```

### <span class="badge object-method"></span> tooltip

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip

```php
tooltip(\Grafana\Foundation\Cog\Builder $tooltip)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
