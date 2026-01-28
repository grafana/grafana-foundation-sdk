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

### <span class="badge object-method"></span> Calculate

Controls if the heatmap should be calculated from data

```go
func (builder *OptionsBuilder) Calculate(calculate bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Calculation

Calculation options for the heatmap

```go
func (builder *OptionsBuilder) Calculation(calculation cog.Builder[common.HeatmapCalculationOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> CellGap

Controls gap between cells

```go
func (builder *OptionsBuilder) CellGap(cellGap uint8) *OptionsBuilder
```

### <span class="badge object-method"></span> CellRadius

Controls cell radius

```go
func (builder *OptionsBuilder) CellRadius(cellRadius float32) *OptionsBuilder
```

### <span class="badge object-method"></span> CellValues

Controls cell value unit

```go
func (builder *OptionsBuilder) CellValues(cellValues cog.Builder[heatmap.CellValues]) *OptionsBuilder
```

### <span class="badge object-method"></span> Color

Controls the color options

```go
func (builder *OptionsBuilder) Color(color cog.Builder[heatmap.HeatmapColorOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Exemplars

Controls exemplar options

```go
func (builder *OptionsBuilder) Exemplars(exemplars cog.Builder[heatmap.ExemplarConfig]) *OptionsBuilder
```

### <span class="badge object-method"></span> FilterValues

Filters values between a given range

```go
func (builder *OptionsBuilder) FilterValues(filterValues cog.Builder[heatmap.FilterValueRange]) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[heatmap.HeatmapLegend]) *OptionsBuilder
```

### <span class="badge object-method"></span> RowsFrame

Controls tick alignment and value name when not calculating from data

```go
func (builder *OptionsBuilder) RowsFrame(rowsFrame cog.Builder[heatmap.RowsHeatmapOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> SelectionMode

Controls which axis to allow selection on

```go
func (builder *OptionsBuilder) SelectionMode(selectionMode heatmap.HeatmapSelectionMode) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowValue

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```go
func (builder *OptionsBuilder) ShowValue(showValue common.VisibilityMode) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

Controls tooltip options

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[heatmap.HeatmapTooltip]) *OptionsBuilder
```

### <span class="badge object-method"></span> YAxis

Controls yAxis placement

```go
func (builder *OptionsBuilder) YAxis(yAxis cog.Builder[heatmap.YAxisConfig]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
