---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    graph_mode: common.BigValueGraphMode
    color_mode: common.BigValueColorMode
    justify_mode: common.BigValueJustifyMode
    text_mode: common.BigValueTextMode
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    wide_layout: bool
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

