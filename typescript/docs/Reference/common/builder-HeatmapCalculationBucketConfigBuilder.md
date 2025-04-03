---
title: <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder
---
# <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder

## Constructor

```typescript
new HeatmapCalculationBucketConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> mode

Sets the bucket calculation mode

```typescript
mode(mode: common.HeatmapCalculationMode)
```

### <span class="badge object-method"></span> scale

Controls the scale of the buckets

```typescript
scale(scale: cog.Builder<common.ScaleDistributionConfig>)
```

### <span class="badge object-method"></span> value

The number of buckets to use for the axis in the heatmap

```typescript
value(value: string)
```

## See also

 * <span class="badge object-type-interface"></span> [HeatmapCalculationBucketConfig](./object-HeatmapCalculationBucketConfig.md)
