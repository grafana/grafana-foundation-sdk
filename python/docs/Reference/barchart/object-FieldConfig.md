---
title: <span class="badge object-type-class"></span> FieldConfig
---
# <span class="badge object-type-class"></span> FieldConfig

## Definition

```python
class FieldConfig:
    # Controls line width of the bars.
    line_width: typing.Optional[int]
    # Controls the fill opacity of the bars.
    fill_opacity: typing.Optional[int]
    # Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
    # Gradient appearance is influenced by the Fill opacity setting.
    gradient_mode: typing.Optional[common.GraphGradientMode]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    hide_from: typing.Optional[common.HideSeriesConfig]
    # Threshold rendering
    thresholds_style: typing.Optional[common.GraphThresholdsStyleConfig]
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

