---
title: <span class="badge builder"></span> FieldConfig
---
# <span class="badge builder"></span> FieldConfig

## Constructor

```python
FieldConfig()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> histogram.FieldConfig
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

### <span class="badge object-method"></span> fill_opacity

Controls the fill opacity of the bars.

```python
def fill_opacity(fill_opacity: int) -> typing.Self
```

### <span class="badge object-method"></span> gradient_mode

Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.

Gradient appearance is influenced by the Fill opacity setting.

```python
def gradient_mode(gradient_mode: common.GraphGradientMode) -> typing.Self
```

### <span class="badge object-method"></span> hide_from

```python
def hide_from(hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self
```

### <span class="badge object-method"></span> line_width

Controls line width of the bars.

```python
def line_width(line_width: int) -> typing.Self
```

### <span class="badge object-method"></span> scale_distribution

```python
def scale_distribution(scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [FieldConfig](./object-FieldConfig.md)
