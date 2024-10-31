---
title: <span class="badge object-type-class"></span> FieldConfig
---
# <span class="badge object-type-class"></span> FieldConfig

## Definition

```python
class FieldConfig:
    show: typing.Optional[xychart.ScatterShow]
    point_size: typing.Optional[common.ScaleDimensionConfig]
    point_color: typing.Optional[common.ColorDimensionConfig]
    line_color: typing.Optional[common.ColorDimensionConfig]
    line_width: typing.Optional[int]
    line_style: typing.Optional[common.LineStyle]
    label: typing.Optional[common.VisibilityMode]
    hide_from: typing.Optional[common.HideSeriesConfig]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    axis_centered_zero: typing.Optional[bool]
    label_value: typing.Optional[common.TextDimensionConfig]
    axis_border_show: typing.Optional[bool]
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

