---
title: <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder
---
# <span class="badge builder"></span> HeatmapCalculationBucketConfigBuilder

## Constructor

```java
new HeatmapCalculationBucketConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public HeatmapCalculationBucketConfig build()
```

### <span class="badge object-method"></span> mode

Sets the bucket calculation mode

```java
public HeatmapCalculationBucketConfigBuilder mode(HeatmapCalculationMode mode)
```

### <span class="badge object-method"></span> scale

Controls the scale of the buckets

```java
public HeatmapCalculationBucketConfigBuilder scale(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scale)
```

### <span class="badge object-method"></span> value

The number of buckets to use for the axis in the heatmap

```java
public HeatmapCalculationBucketConfigBuilder value(String value)
```

## See also

 * <span class="badge object-type-class"></span> [HeatmapCalculationBucketConfig](./object-HeatmapCalculationBucketConfig.md)
