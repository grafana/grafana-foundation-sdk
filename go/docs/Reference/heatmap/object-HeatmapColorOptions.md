---
title: <span class="badge object-type-struct"></span> HeatmapColorOptions
---
# <span class="badge object-type-struct"></span> HeatmapColorOptions

Controls various color options

## Definition

```go
type HeatmapColorOptions struct {
    // Sets the color mode
    Mode *heatmap.HeatmapColorMode `json:"mode,omitempty"`
    // Controls the color scheme used
    Scheme string `json:"scheme"`
    // Controls the color fill when in opacity mode
    Fill string `json:"fill"`
    // Controls the color scale
    Scale *heatmap.HeatmapColorScale `json:"scale,omitempty"`
    // Controls the exponent when scale is set to exponential
    Exponent float32 `json:"exponent"`
    // Controls the number of color steps
    Steps uint64 `json:"steps"`
    // Reverses the color scheme
    Reverse bool `json:"reverse"`
    // Sets the minimum value for the color scale
    Min *float32 `json:"min,omitempty"`
    // Sets the maximum value for the color scale
    Max *float32 `json:"max,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapColorOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (heatmapColorOptions *HeatmapColorOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HeatmapColorOptions` objects.

```go
func (heatmapColorOptions *HeatmapColorOptions) Equals(other HeatmapColorOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HeatmapColorOptions` fields for violations and returns them.

```go
func (heatmapColorOptions *HeatmapColorOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [HeatmapColorOptionsBuilder](./builder-HeatmapColorOptionsBuilder.md)
