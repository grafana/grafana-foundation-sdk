---
title: <span class="badge builder"></span> HeatmapCalculationOptionsBuilder
---
# <span class="badge builder"></span> HeatmapCalculationOptionsBuilder

## Constructor

```go
func NewHeatmapCalculationOptionsBuilder() *HeatmapCalculationOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *HeatmapCalculationOptionsBuilder) Build() (HeatmapCalculationOptions, error)
```

### <span class="badge object-method"></span> XBuckets

The number of buckets to use for the xAxis in the heatmap

```go
func (builder *HeatmapCalculationOptionsBuilder) XBuckets(xBuckets cog.Builder[common.HeatmapCalculationBucketConfig]) *HeatmapCalculationOptionsBuilder
```

### <span class="badge object-method"></span> YBuckets

The number of buckets to use for the yAxis in the heatmap

```go
func (builder *HeatmapCalculationOptionsBuilder) YBuckets(yBuckets cog.Builder[common.HeatmapCalculationBucketConfig]) *HeatmapCalculationOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [HeatmapCalculationOptions](./object-HeatmapCalculationOptions.md)
