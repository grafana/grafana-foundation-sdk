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

### <span class="badge object-method"></span> alignValue

Controls value alignment on the timelines

```php
alignValue(\Grafana\Foundation\Common\TimelineValueAlignment $alignValue)
```

### <span class="badge object-method"></span> legend

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> mergeValues

Merge equal consecutive values

```php
mergeValues(bool $mergeValues)
```

### <span class="badge object-method"></span> rowHeight

Controls the row height

```php
rowHeight(float $rowHeight)
```

### <span class="badge object-method"></span> showValue

Show timeline values on chart

```php
showValue(\Grafana\Foundation\Common\VisibilityMode $showValue)
```

### <span class="badge object-method"></span> timezone

@param array<string> $timezone

```php
timezone(array $timezone)
```

### <span class="badge object-method"></span> tooltip

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip

```php
tooltip(\Grafana\Foundation\Cog\Builder $tooltip)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
