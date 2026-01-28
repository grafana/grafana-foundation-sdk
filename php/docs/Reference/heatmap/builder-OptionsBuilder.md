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

### <span class="badge object-method"></span> calculate

Controls if the heatmap should be calculated from data

```php
calculate(bool $calculate)
```

### <span class="badge object-method"></span> calculation

Calculation options for the heatmap

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HeatmapCalculationOptions> $calculation

```php
calculation(\Grafana\Foundation\Cog\Builder $calculation)
```

### <span class="badge object-method"></span> cellGap

Controls gap between cells

```php
cellGap(int $cellGap)
```

### <span class="badge object-method"></span> cellRadius

Controls cell radius

```php
cellRadius(float $cellRadius)
```

### <span class="badge object-method"></span> cellValues

Controls cell value unit

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\CellValues> $cellValues

```php
cellValues(\Grafana\Foundation\Cog\Builder $cellValues)
```

### <span class="badge object-method"></span> color

Controls the color options

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapColorOptions> $color

```php
color(\Grafana\Foundation\Cog\Builder $color)
```

### <span class="badge object-method"></span> exemplars

Controls exemplar options

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\ExemplarConfig> $exemplars

```php
exemplars(\Grafana\Foundation\Cog\Builder $exemplars)
```

### <span class="badge object-method"></span> filterValues

Filters values between a given range

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\FilterValueRange> $filterValues

```php
filterValues(\Grafana\Foundation\Cog\Builder $filterValues)
```

### <span class="badge object-method"></span> legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapLegend> $legend

```php
legend(\Grafana\Foundation\Cog\Builder $legend)
```

### <span class="badge object-method"></span> rowsFrame

Controls tick alignment and value name when not calculating from data

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\RowsHeatmapOptions> $rowsFrame

```php
rowsFrame(\Grafana\Foundation\Cog\Builder $rowsFrame)
```

### <span class="badge object-method"></span> showValue

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```php
showValue(\Grafana\Foundation\Common\VisibilityMode $showValue)
```

### <span class="badge object-method"></span> tooltip

Controls tooltip options

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapTooltip> $tooltip

```php
tooltip(\Grafana\Foundation\Cog\Builder $tooltip)
```

### <span class="badge object-method"></span> yAxis

Controls yAxis placement

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\YAxisConfig> $yAxis

```php
yAxis(\Grafana\Foundation\Cog\Builder $yAxis)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
