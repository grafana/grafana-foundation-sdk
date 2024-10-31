---
title: <span class="badge object-type-class"></span> AxisConfig
---
# <span class="badge object-type-class"></span> AxisConfig

TODO docs

## Definition

```python
class AxisConfig:
    """
    TODO docs
    """

    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    axis_centered_zero: typing.Optional[bool]
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

 * <span class="badge builder"></span> [AxisConfig](./builder-AxisConfig.md)
