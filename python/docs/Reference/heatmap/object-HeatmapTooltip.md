---
title: <span class="badge object-type-class"></span> HeatmapTooltip
---
# <span class="badge object-type-class"></span> HeatmapTooltip

Controls tooltip options

## Definition

```python
class HeatmapTooltip:
    """
    Controls tooltip options
    """

    # Controls how the tooltip is shown
    mode: common.TooltipDisplayMode
    max_height: typing.Optional[float]
    max_width: typing.Optional[float]
    # Controls if the tooltip shows a histogram of the y-axis values
    y_histogram: typing.Optional[bool]
    # Controls if the tooltip shows a color scale in header
    show_color_scale: typing.Optional[bool]
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

 * <span class="badge builder"></span> [HeatmapTooltip](./builder-HeatmapTooltip.md)
