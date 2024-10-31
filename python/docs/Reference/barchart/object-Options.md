---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Manually select which field from the dataset to represent the x field.
    x_field: typing.Optional[str]
    # Use the color value for a sibling field to color each bar value.
    color_by_field: typing.Optional[str]
    # Controls the orientation of the bar chart, either vertical or horizontal.
    orientation: common.VizOrientation
    # Controls the radius of each bar.
    bar_radius: typing.Optional[float]
    # Controls the rotation of the x axis labels.
    x_tick_label_rotation: int
    # Sets the max length that a label can have before it is truncated.
    x_tick_label_max_length: int
    # Controls the spacing between x axis labels.
    # negative values indicate backwards skipping behavior
    x_tick_label_spacing: typing.Optional[int]
    # Controls whether bars are stacked or not, either normally or in percent mode.
    stacking: common.StackingMode
    # This controls whether values are shown on top or to the left of bars.
    show_value: common.VisibilityMode
    # Controls the width of bars. 1 = Max width, 0 = Min width.
    bar_width: float
    # Controls the width of groups. 1 = max with, 0 = min width.
    group_width: float
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    # Enables mode which highlights the entire bar area and shows tooltip when cursor
    # hovers over highlighted area
    full_highlight: bool
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

