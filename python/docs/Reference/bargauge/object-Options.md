---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    display_mode: common.BarGaugeDisplayMode
    value_mode: common.BarGaugeValueMode
    name_placement: common.BarGaugeNamePlacement
    show_unfilled: bool
    sizing: common.BarGaugeSizing
    min_viz_width: int
    min_viz_height: int
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    max_viz_height: int
    orientation: common.VizOrientation
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

