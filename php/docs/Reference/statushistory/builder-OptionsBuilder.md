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

### <span class="badge object-method"></span> colWidth

Controls the column width

```php
colWidth(float $colWidth)
```

### <span class="badge object-method"></span> legend

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> perPage

Enables pagination when > 0

```php
perPage(float $perPage)
```

### <span class="badge object-method"></span> rowHeight

Set the height of the rows

```php
rowHeight(float $rowHeight)
```

### <span class="badge object-method"></span> showValue

Show values on the columns

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
