---
title: <span class="badge object-type-class"></span> HeatmapCalculationOptions
---
# <span class="badge object-type-class"></span> HeatmapCalculationOptions

## Definition

```python
class HeatmapCalculationOptions:
    # The number of buckets to use for the xAxis in the heatmap
    x_buckets: typing.Optional[common.HeatmapCalculationBucketConfig]
    # The number of buckets to use for the yAxis in the heatmap
    y_buckets: typing.Optional[common.HeatmapCalculationBucketConfig]
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

 * <span class="badge builder"></span> [HeatmapCalculationOptions](./builder-HeatmapCalculationOptions.md)
