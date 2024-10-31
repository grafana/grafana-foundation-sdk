---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Controls if the heatmap should be calculated from data
    calculate: typing.Optional[bool]
    # Calculation options for the heatmap
    calculation: typing.Optional[common.HeatmapCalculationOptions]
    # Controls the color options
    color: heatmap.HeatmapColorOptions
    # Filters values between a given range
    filter_values: typing.Optional[heatmap.FilterValueRange]
    # Controls tick alignment and value name when not calculating from data
    rows_frame: typing.Optional[heatmap.RowsHeatmapOptions]
    # | *{
    # 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    # }
    # Controls the display of the value in the cell
    show_value: common.VisibilityMode
    # Controls gap between cells
    cell_gap: typing.Optional[int]
    # Controls cell radius
    cell_radius: typing.Optional[float]
    # Controls cell value unit
    cell_values: typing.Optional[heatmap.CellValues]
    # Controls yAxis placement
    y_axis: heatmap.YAxisConfig
    # | *{
    # 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    # }
    # Controls legend options
    legend: heatmap.HeatmapLegend
    # Controls tooltip options
    tooltip: heatmap.HeatmapTooltip
    # Controls exemplar options
    exemplars: heatmap.ExemplarConfig
    # Controls which axis to allow selection on
    selection_mode: typing.Optional[heatmap.HeatmapSelectionMode]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

