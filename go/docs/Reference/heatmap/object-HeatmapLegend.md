---
title: <span class="badge object-type-struct"></span> HeatmapLegend
---
# <span class="badge object-type-struct"></span> HeatmapLegend

Controls legend options

## Definition

```go
type HeatmapLegend struct {
    // Controls if the legend is shown
    Show bool `json:"show"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapLegend` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (heatmapLegend *HeatmapLegend) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HeatmapLegend` objects.

```go
func (heatmapLegend *HeatmapLegend) Equals(other HeatmapLegend) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HeatmapLegend` fields for violations and returns them.

```go
func (heatmapLegend *HeatmapLegend) Validate() error
```

## See also

 * <span class="badge builder"></span> [HeatmapLegendBuilder](./builder-HeatmapLegendBuilder.md)
