---
title: <span class="badge object-type-struct"></span> HeatmapCalculationBucketConfig
---
# <span class="badge object-type-struct"></span> HeatmapCalculationBucketConfig

## Definition

```go
type HeatmapCalculationBucketConfig struct {
    // Sets the bucket calculation mode
    Mode *common.HeatmapCalculationMode `json:"mode,omitempty"`
    // The number of buckets to use for the axis in the heatmap
    Value *string `json:"value,omitempty"`
    // Controls the scale of the buckets
    Scale *common.ScaleDistributionConfig `json:"scale,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapCalculationBucketConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (heatmapCalculationBucketConfig *HeatmapCalculationBucketConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HeatmapCalculationBucketConfig` objects.

```go
func (heatmapCalculationBucketConfig *HeatmapCalculationBucketConfig) Equals(other HeatmapCalculationBucketConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HeatmapCalculationBucketConfig` fields for violations and returns them.

```go
func (heatmapCalculationBucketConfig *HeatmapCalculationBucketConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [HeatmapCalculationBucketConfigBuilder](./builder-HeatmapCalculationBucketConfigBuilder.md)
