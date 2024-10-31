---
title: <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder
---
# <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder

## Constructor

```go
func NewHeatmapCalculationBucketConfigBuilder() *HeatmapCalculationBucketConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *HeatmapCalculationBucketConfigBuilder) Build() (HeatmapCalculationBucketConfig, error)
```

### <span class="badge object-method"></span> Mode

Sets the bucket calculation mode

```go
func (builder *HeatmapCalculationBucketConfigBuilder) Mode(mode common.HeatmapCalculationMode) *HeatmapCalculationBucketConfigBuilder
```

### <span class="badge object-method"></span> Scale

Controls the scale of the buckets

```go
func (builder *HeatmapCalculationBucketConfigBuilder) Scale(scale cog.Builder[common.ScaleDistributionConfig]) *HeatmapCalculationBucketConfigBuilder
```

### <span class="badge object-method"></span> Value

The number of buckets to use for the axis in the heatmap

```go
func (builder *HeatmapCalculationBucketConfigBuilder) Value(value string) *HeatmapCalculationBucketConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [HeatmapCalculationBucketConfig](./object-HeatmapCalculationBucketConfig.md)
