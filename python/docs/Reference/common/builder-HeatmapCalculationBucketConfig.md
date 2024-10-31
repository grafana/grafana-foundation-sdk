---
title: <span class="badge builder"></span> HeatmapCalculationBucketConfig
---
# <span class="badge builder"></span> HeatmapCalculationBucketConfig

## Constructor

```python
HeatmapCalculationBucketConfig()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> common.HeatmapCalculationBucketConfig
```

### <span class="badge object-method"></span> mode

Sets the bucket calculation mode

```python
def mode(mode: common.HeatmapCalculationMode) -> typing.Self
```

### <span class="badge object-method"></span> scale

Controls the scale of the buckets

```python
def scale(scale: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self
```

### <span class="badge object-method"></span> value

The number of buckets to use for the axis in the heatmap

```python
def value(value: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [HeatmapCalculationBucketConfig](./object-HeatmapCalculationBucketConfig.md)
