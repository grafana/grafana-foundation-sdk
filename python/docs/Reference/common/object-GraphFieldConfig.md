---
title: <span class="badge object-type-class"></span> GraphFieldConfig
---
# <span class="badge object-type-class"></span> GraphFieldConfig

TODO docs

## Definition

```python
class GraphFieldConfig:
    """
    TODO docs
    """

    draw_style: typing.Optional[common.GraphDrawStyle]
    gradient_mode: typing.Optional[common.GraphGradientMode]
    thresholds_style: typing.Optional[common.GraphThresholdsStyleConfig]
    line_color: typing.Optional[str]
    line_width: typing.Optional[float]
    line_interpolation: typing.Optional[common.LineInterpolation]
    line_style: typing.Optional[common.LineStyle]
    fill_color: typing.Optional[str]
    fill_opacity: typing.Optional[float]
    show_points: typing.Optional[common.VisibilityMode]
    point_size: typing.Optional[float]
    point_color: typing.Optional[str]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    bar_alignment: typing.Optional[common.BarAlignment]
    bar_width_factor: typing.Optional[float]
    stacking: typing.Optional[common.StackingConfig]
    hide_from: typing.Optional[common.HideSeriesConfig]
    transform: typing.Optional[common.GraphTransform]
    # Indicate if null values should be treated as gaps or connected.
    # When the value is a number, it represents the maximum delta in the
    # X axis that should be considered connected.  For timeseries, this is milliseconds
    span_nulls: typing.Optional[typing.Union[bool, float]]
    fill_below_to: typing.Optional[str]
    point_symbol: typing.Optional[str]
    axis_centered_zero: typing.Optional[bool]
    bar_max_width: typing.Optional[float]
    insert_nulls: typing.Optional[typing.Union[bool, int]]
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

 * <span class="badge builder"></span> [GraphFieldConfig](./builder-GraphFieldConfig.md)
