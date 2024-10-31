---
title: <span class="badge object-type-struct"></span> HeatmapCalculationOptions
---
# <span class="badge object-type-struct"></span> HeatmapCalculationOptions

## Definition

```go
type HeatmapCalculationOptions struct {
    // The number of buckets to use for the xAxis in the heatmap
    XBuckets *common.HeatmapCalculationBucketConfig `json:"xBuckets,omitempty"`
    // The number of buckets to use for the yAxis in the heatmap
    YBuckets *common.HeatmapCalculationBucketConfig `json:"yBuckets,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapCalculationOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (heatmapCalculationOptions *HeatmapCalculationOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HeatmapCalculationOptions` objects.

```go
func (heatmapCalculationOptions *HeatmapCalculationOptions) Equals(other HeatmapCalculationOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HeatmapCalculationOptions` fields for violations and returns them.

```go
func (heatmapCalculationOptions *HeatmapCalculationOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [HeatmapCalculationOptionsBuilder](./builder-HeatmapCalculationOptionsBuilder.md)
