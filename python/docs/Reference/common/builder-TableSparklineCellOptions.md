---
title: <span class="badge builder"></span> TableSparklineCellOptions
---
# <span class="badge builder"></span> TableSparklineCellOptions

## Constructor

```python
TableSparklineCellOptions()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> common.TableSparklineCellOptions
```

### <span class="badge object-method"></span> axis_border_show

```python
def axis_border_show(axis_border_show: bool) -> typing.Self
```

### <span class="badge object-method"></span> axis_centered_zero

```python
def axis_centered_zero(axis_centered_zero: bool) -> typing.Self
```

### <span class="badge object-method"></span> axis_color_mode

```python
def axis_color_mode(axis_color_mode: common.AxisColorMode) -> typing.Self
```

### <span class="badge object-method"></span> axis_grid_show

```python
def axis_grid_show(axis_grid_show: bool) -> typing.Self
```

### <span class="badge object-method"></span> axis_label

```python
def axis_label(axis_label: str) -> typing.Self
```

### <span class="badge object-method"></span> axis_placement

```python
def axis_placement(axis_placement: common.AxisPlacement) -> typing.Self
```

### <span class="badge object-method"></span> axis_soft_max

```python
def axis_soft_max(axis_soft_max: float) -> typing.Self
```

### <span class="badge object-method"></span> axis_soft_min

```python
def axis_soft_min(axis_soft_min: float) -> typing.Self
```

### <span class="badge object-method"></span> axis_width

```python
def axis_width(axis_width: float) -> typing.Self
```

### <span class="badge object-method"></span> bar_alignment

```python
def bar_alignment(bar_alignment: common.BarAlignment) -> typing.Self
```

### <span class="badge object-method"></span> bar_max_width

```python
def bar_max_width(bar_max_width: float) -> typing.Self
```

### <span class="badge object-method"></span> bar_width_factor

```python
def bar_width_factor(bar_width_factor: float) -> typing.Self
```

### <span class="badge object-method"></span> draw_style

```python
def draw_style(draw_style: common.GraphDrawStyle) -> typing.Self
```

### <span class="badge object-method"></span> fill_below_to

```python
def fill_below_to(fill_below_to: str) -> typing.Self
```

### <span class="badge object-method"></span> fill_color

```python
def fill_color(fill_color: str) -> typing.Self
```

### <span class="badge object-method"></span> fill_opacity

```python
def fill_opacity(fill_opacity: float) -> typing.Self
```

### <span class="badge object-method"></span> gradient_mode

```python
def gradient_mode(gradient_mode: common.GraphGradientMode) -> typing.Self
```

### <span class="badge object-method"></span> hide_from

```python
def hide_from(hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self
```

### <span class="badge object-method"></span> hide_value

```python
def hide_value(hide_value: bool) -> typing.Self
```

### <span class="badge object-method"></span> line_color

```python
def line_color(line_color: str) -> typing.Self
```

### <span class="badge object-method"></span> line_interpolation

```python
def line_interpolation(line_interpolation: common.LineInterpolation) -> typing.Self
```

### <span class="badge object-method"></span> line_style

```python
def line_style(line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self
```

### <span class="badge object-method"></span> line_width

```python
def line_width(line_width: float) -> typing.Self
```

### <span class="badge object-method"></span> point_color

```python
def point_color(point_color: str) -> typing.Self
```

### <span class="badge object-method"></span> point_size

```python
def point_size(point_size: float) -> typing.Self
```

### <span class="badge object-method"></span> point_symbol

```python
def point_symbol(point_symbol: str) -> typing.Self
```

### <span class="badge object-method"></span> scale_distribution

```python
def scale_distribution(scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self
```

### <span class="badge object-method"></span> show_points

```python
def show_points(show_points: common.VisibilityMode) -> typing.Self
```

### <span class="badge object-method"></span> span_nulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```python
def span_nulls(span_nulls: typing.Union[bool, float]) -> typing.Self
```

### <span class="badge object-method"></span> stacking

```python
def stacking(stacking: cogbuilder.Builder[common.StackingConfig]) -> typing.Self
```

### <span class="badge object-method"></span> thresholds_style

```python
def thresholds_style(thresholds_style: cogbuilder.Builder[common.GraphThresholdsStyleConfig]) -> typing.Self
```

### <span class="badge object-method"></span> transform

```python
def transform(transform: common.GraphTransform) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TableSparklineCellOptions](./object-TableSparklineCellOptions.md)
