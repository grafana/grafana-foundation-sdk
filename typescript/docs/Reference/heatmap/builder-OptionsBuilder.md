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

### <span class="badge object-method"></span> calculate

Controls if the heatmap should be calculated from data

```typescript
calculate(calculate: boolean)
```

### <span class="badge object-method"></span> calculation

Calculation options for the heatmap

```typescript
calculation(calculation: cog.Builder<common.HeatmapCalculationOptions>)
```

### <span class="badge object-method"></span> cellGap

Controls gap between cells

```typescript
cellGap(cellGap: number)
```

### <span class="badge object-method"></span> cellRadius

Controls cell radius

```typescript
cellRadius(cellRadius: number)
```

### <span class="badge object-method"></span> cellValues

Controls cell value unit

```typescript
cellValues(cellValues: cog.Builder<heatmap.CellValues>)
```

### <span class="badge object-method"></span> color

Controls the color options

```typescript
color(color: cog.Builder<heatmap.HeatmapColorOptions>)
```

### <span class="badge object-method"></span> exemplars

Controls exemplar options

```typescript
exemplars(exemplars: cog.Builder<heatmap.ExemplarConfig>)
```

### <span class="badge object-method"></span> filterValues

Filters values between a given range

```typescript
filterValues(filterValues: cog.Builder<heatmap.FilterValueRange>)
```

### <span class="badge object-method"></span> legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

```typescript
legend(legend: cog.Builder<heatmap.HeatmapLegend>)
```

### <span class="badge object-method"></span> rowsFrame

Controls tick alignment and value name when not calculating from data

```typescript
rowsFrame(rowsFrame: cog.Builder<heatmap.RowsHeatmapOptions>)
```

### <span class="badge object-method"></span> selectionMode

Controls which axis to allow selection on

```typescript
selectionMode(selectionMode: heatmap.HeatmapSelectionMode)
```

### <span class="badge object-method"></span> showValue

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```typescript
showValue(showValue: common.VisibilityMode)
```

### <span class="badge object-method"></span> tooltip

Controls tooltip options

```typescript
tooltip(tooltip: cog.Builder<heatmap.HeatmapTooltip>)
```

### <span class="badge object-method"></span> yAxis

Controls yAxis placement

```typescript
yAxis(yAxis: cog.Builder<heatmap.YAxisConfig>)
```

## See also

 * <span class="badge object-type-interface"></span> [Options](./object-Options.md)
