---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    // Controls if the heatmap should be calculated from data
    Calculate *bool `json:"calculate,omitempty"`
    // Calculation options for the heatmap
    Calculation *common.HeatmapCalculationOptions `json:"calculation,omitempty"`
    // Controls the color options
    Color heatmap.HeatmapColorOptions `json:"color"`
    // Filters values between a given range
    FilterValues *heatmap.FilterValueRange `json:"filterValues,omitempty"`
    // Controls tick alignment and value name when not calculating from data
    RowsFrame *heatmap.RowsHeatmapOptions `json:"rowsFrame,omitempty"`
    // | *{
    // 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls the display of the value in the cell
    ShowValue common.VisibilityMode `json:"showValue"`
    // Controls gap between cells
    CellGap *uint8 `json:"cellGap,omitempty"`
    // Controls cell radius
    CellRadius *float32 `json:"cellRadius,omitempty"`
    // Controls cell value unit
    CellValues *heatmap.CellValues `json:"cellValues,omitempty"`
    // Controls yAxis placement
    YAxis heatmap.YAxisConfig `json:"yAxis"`
    // | *{
    // 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls legend options
    Legend heatmap.HeatmapLegend `json:"legend"`
    // Controls tooltip options
    Tooltip heatmap.HeatmapTooltip `json:"tooltip"`
    // Controls exemplar options
    Exemplars heatmap.ExemplarConfig `json:"exemplars"`
    // Controls which axis to allow selection on
    SelectionMode *heatmap.HeatmapSelectionMode `json:"selectionMode,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (options *Options) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Options` objects.

```go
func (options *Options) Equals(other Options) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.

```go
func (options *Options) Validate() error
```

