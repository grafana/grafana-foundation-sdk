---
title: <span class="badge object-type-struct"></span> PieChartLegendOptions
---
# <span class="badge object-type-struct"></span> PieChartLegendOptions

## Definition

```go
type PieChartLegendOptions struct {
    Values []piechart.PieChartLegendValues `json:"values"`
    DisplayMode common.LegendDisplayMode `json:"displayMode"`
    Placement common.LegendPlacement `json:"placement"`
    ShowLegend bool `json:"showLegend"`
    AsTable *bool `json:"asTable,omitempty"`
    IsVisible *bool `json:"isVisible,omitempty"`
    SortBy *string `json:"sortBy,omitempty"`
    SortDesc *bool `json:"sortDesc,omitempty"`
    Width *float64 `json:"width,omitempty"`
    Calcs []string `json:"calcs"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PieChartLegendOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (pieChartLegendOptions *PieChartLegendOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PieChartLegendOptions` objects.

```go
func (pieChartLegendOptions *PieChartLegendOptions) Equals(other PieChartLegendOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PieChartLegendOptions` fields for violations and returns them.

```go
func (pieChartLegendOptions *PieChartLegendOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [PieChartLegendOptionsBuilder](./builder-PieChartLegendOptionsBuilder.md)
