---
title: <span class="badge object-type-struct"></span> HeatmapTooltip
---
# <span class="badge object-type-struct"></span> HeatmapTooltip

Controls tooltip options

## Definition

```go
type HeatmapTooltip struct {
    // Controls if the tooltip is shown
    Show bool `json:"show"`
    // Controls if the tooltip shows a histogram of the y-axis values
    YHistogram *bool `json:"yHistogram,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapTooltip` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (heatmapTooltip *HeatmapTooltip) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HeatmapTooltip` objects.

```go
func (heatmapTooltip *HeatmapTooltip) Equals(other HeatmapTooltip) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HeatmapTooltip` fields for violations and returns them.

```go
func (heatmapTooltip *HeatmapTooltip) Validate() error
```

## See also

 * <span class="badge builder"></span> [HeatmapTooltipBuilder](./builder-HeatmapTooltipBuilder.md)
