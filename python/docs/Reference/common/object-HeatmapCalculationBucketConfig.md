---
title: <span class="badge object-type-class"></span> HeatmapCalculationBucketConfig
---
# <span class="badge object-type-class"></span> HeatmapCalculationBucketConfig

## Definition

```python
class HeatmapCalculationBucketConfig:
    # Sets the bucket calculation mode
    mode: typing.Optional[common.HeatmapCalculationMode]
    # The number of buckets to use for the axis in the heatmap
    value: typing.Optional[str]
    # Controls the scale of the buckets
    scale: typing.Optional[common.ScaleDistributionConfig]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [HeatmapCalculationBucketConfig](./builder-HeatmapCalculationBucketConfig.md)
