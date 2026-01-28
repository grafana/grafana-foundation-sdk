---
title: <span class="badge builder"></span> Options
---
# <span class="badge builder"></span> Options

## Constructor

```python
Options()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> barchart.Options
```

### <span class="badge object-method"></span> bar_radius

Controls the radius of each bar.

```python
def bar_radius(bar_radius: float) -> typing.Self
```

### <span class="badge object-method"></span> bar_width

Controls the width of bars. 1 = Max width, 0 = Min width.

```python
def bar_width(bar_width: float) -> typing.Self
```

### <span class="badge object-method"></span> color_by_field

Use the color value for a sibling field to color each bar value.

```python
def color_by_field(color_by_field: str) -> typing.Self
```

### <span class="badge object-method"></span> full_highlight

Enables mode which highlights the entire bar area and shows tooltip when cursor

hovers over highlighted area

```python
def full_highlight(full_highlight: bool) -> typing.Self
```

### <span class="badge object-method"></span> group_width

Controls the width of groups. 1 = max with, 0 = min width.

```python
def group_width(group_width: float) -> typing.Self
```

### <span class="badge object-method"></span> legend

```python
def legend(legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self
```

### <span class="badge object-method"></span> orientation

Controls the orientation of the bar chart, either vertical or horizontal.

```python
def orientation(orientation: common.VizOrientation) -> typing.Self
```

### <span class="badge object-method"></span> show_value

This controls whether values are shown on top or to the left of bars.

```python
def show_value(show_value: common.VisibilityMode) -> typing.Self
```

### <span class="badge object-method"></span> stacking

Controls whether bars are stacked or not, either normally or in percent mode.

```python
def stacking(stacking: common.StackingMode) -> typing.Self
```

### <span class="badge object-method"></span> text

```python
def text(text: cogbuilder.Builder[common.VizTextDisplayOptions]) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

```python
def tooltip(tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self
```

### <span class="badge object-method"></span> x_field

Manually select which field from the dataset to represent the x field.

```python
def x_field(x_field: str) -> typing.Self
```

### <span class="badge object-method"></span> x_tick_label_max_length

Sets the max length that a label can have before it is truncated.

```python
def x_tick_label_max_length(x_tick_label_max_length: int) -> typing.Self
```

### <span class="badge object-method"></span> x_tick_label_rotation

Controls the rotation of the x axis labels.

```python
def x_tick_label_rotation(x_tick_label_rotation: int) -> typing.Self
```

### <span class="badge object-method"></span> x_tick_label_spacing

Controls the spacing between x axis labels.

negative values indicate backwards skipping behavior

```python
def x_tick_label_spacing(x_tick_label_spacing: int) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
