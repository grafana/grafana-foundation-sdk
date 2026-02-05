---
title: <span class="badge builder"></span> RowsLayoutRowBuilder
---
# <span class="badge builder"></span> RowsLayoutRowBuilder

## Constructor

```php
new RowsLayoutRowBuilder(string $title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> autoGridLayout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind> $autoGridLayoutKind

```php
autoGridLayout(\Grafana\Foundation\Cog\Builder $autoGridLayoutKind)
```

### <span class="badge object-method"></span> rowsLayout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind> $rowsLayoutKind

```php
rowsLayout(\Grafana\Foundation\Cog\Builder $rowsLayoutKind)
```

### <span class="badge object-method"></span> tabsLayout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind> $tabsLayoutKind

```php
tabsLayout(\Grafana\Foundation\Cog\Builder $tabsLayoutKind)
```

### <span class="badge object-method"></span> collapse

```php
collapse(bool $collapse)
```

### <span class="badge object-method"></span> conditionalRendering

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind> $conditionalRendering

```php
conditionalRendering(\Grafana\Foundation\Cog\Builder $conditionalRendering)
```

### <span class="badge object-method"></span> fillScreen

```php
fillScreen(bool $fillScreen)
```

### <span class="badge object-method"></span> gridLayout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind> $gridLayoutKind

```php
gridLayout(\Grafana\Foundation\Cog\Builder $gridLayoutKind)
```

### <span class="badge object-method"></span> hideHeader

```php
hideHeader(bool $hideHeader)
```

### <span class="badge object-method"></span> repeat

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions> $repeat

```php
repeat(\Grafana\Foundation\Cog\Builder $repeat)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

## See also

 * <span class="badge object-type-class"></span> [RowsLayoutRowKind](./object-RowsLayoutRowKind.md)
