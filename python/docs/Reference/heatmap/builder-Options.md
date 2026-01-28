---
title: <span class="badge builder"></span> Options
---
# <span class="badge builder"></span> Options

## Constructor

```python
Options()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> heatmap.Options
```

### <span class="badge object-method"></span> calculate

Controls if the heatmap should be calculated from data

```python
def calculate(calculate: bool) -> typing.Self
```

### <span class="badge object-method"></span> calculation

Calculation options for the heatmap

```python
def calculation(calculation: cogbuilder.Builder[common.HeatmapCalculationOptions]) -> typing.Self
```

### <span class="badge object-method"></span> cell_gap

Controls gap between cells

```python
def cell_gap(cell_gap: int) -> typing.Self
```

### <span class="badge object-method"></span> cell_radius

Controls cell radius

```python
def cell_radius(cell_radius: float) -> typing.Self
```

### <span class="badge object-method"></span> cell_values

Controls cell value unit

```python
def cell_values(cell_values: cogbuilder.Builder[heatmap.CellValues]) -> typing.Self
```

### <span class="badge object-method"></span> color

Controls the color options

```python
def color(color: cogbuilder.Builder[heatmap.HeatmapColorOptions]) -> typing.Self
```

### <span class="badge object-method"></span> exemplars

Controls exemplar options

```python
def exemplars(exemplars: cogbuilder.Builder[heatmap.ExemplarConfig]) -> typing.Self
```

### <span class="badge object-method"></span> filter_values

Filters values between a given range

```python
def filter_values(filter_values: cogbuilder.Builder[heatmap.FilterValueRange]) -> typing.Self
```

### <span class="badge object-method"></span> legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

```python
def legend(legend: cogbuilder.Builder[heatmap.HeatmapLegend]) -> typing.Self
```

### <span class="badge object-method"></span> rows_frame

Controls tick alignment and value name when not calculating from data

```python
def rows_frame(rows_frame: cogbuilder.Builder[heatmap.RowsHeatmapOptions]) -> typing.Self
```

### <span class="badge object-method"></span> selection_mode

Controls which axis to allow selection on

```python
def selection_mode(selection_mode: heatmap.HeatmapSelectionMode) -> typing.Self
```

### <span class="badge object-method"></span> show_value

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```python
def show_value(show_value: common.VisibilityMode) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

Controls tooltip options

```python
def tooltip(tooltip: cogbuilder.Builder[heatmap.HeatmapTooltip]) -> typing.Self
```

### <span class="badge object-method"></span> y_axis

Controls yAxis placement

```python
def y_axis(y_axis: cogbuilder.Builder[heatmap.YAxisConfig]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
