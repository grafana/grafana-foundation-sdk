---
title: <span class="badge object-type-struct"></span> VizLegendOptions
---
# <span class="badge object-type-struct"></span> VizLegendOptions

TODO docs

## Definition

```go
type VizLegendOptions struct {
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

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizLegendOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (vizLegendOptions *VizLegendOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `VizLegendOptions` objects.

```go
func (vizLegendOptions *VizLegendOptions) Equals(other VizLegendOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VizLegendOptions` fields for violations and returns them.

```go
func (vizLegendOptions *VizLegendOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [VizLegendOptionsBuilder](./builder-VizLegendOptionsBuilder.md)
