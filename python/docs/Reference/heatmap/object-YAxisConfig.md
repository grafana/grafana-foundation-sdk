---
title: <span class="badge object-type-class"></span> YAxisConfig
---
# <span class="badge object-type-class"></span> YAxisConfig

Configuration options for the yAxis

## Definition

```python
class YAxisConfig:
    """
    Configuration options for the yAxis
    """

    # Sets the yAxis unit
    unit: typing.Optional[str]
    # Reverses the yAxis
    reverse: typing.Optional[bool]
    # Controls the number of decimals for yAxis values
    decimals: typing.Optional[float]
    # Sets the minimum value for the yAxis
    min_val: typing.Optional[float]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    # Sets the maximum value for the yAxis
    max_val: typing.Optional[float]
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

 * <span class="badge builder"></span> [YAxisConfig](./builder-YAxisConfig.md)
