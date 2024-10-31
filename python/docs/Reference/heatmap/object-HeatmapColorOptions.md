---
title: <span class="badge object-type-class"></span> HeatmapColorOptions
---
# <span class="badge object-type-class"></span> HeatmapColorOptions

Controls various color options

## Definition

```python
class HeatmapColorOptions:
    """
    Controls various color options
    """

    # Sets the color mode
    mode: typing.Optional[heatmap.HeatmapColorMode]
    # Controls the color scheme used
    scheme: str
    # Controls the color fill when in opacity mode
    fill: str
    # Controls the color scale
    scale: typing.Optional[heatmap.HeatmapColorScale]
    # Controls the exponent when scale is set to exponential
    exponent: float
    # Controls the number of color steps
    steps: int
    # Reverses the color scheme
    reverse: bool
    # Sets the minimum value for the color scale
    min_val: typing.Optional[float]
    # Sets the maximum value for the color scale
    max_val: typing.Optional[float]
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

 * <span class="badge builder"></span> [HeatmapColorOptions](./builder-HeatmapColorOptions.md)
