---
title: <span class="badge object-type-class"></span> VizLegendOptions
---
# <span class="badge object-type-class"></span> VizLegendOptions

TODO docs

## Definition

```python
class VizLegendOptions:
    """
    TODO docs
    """

    display_mode: common.LegendDisplayMode
    placement: common.LegendPlacement
    show_legend: bool
    as_table: typing.Optional[bool]
    is_visible: typing.Optional[bool]
    sort_by: typing.Optional[str]
    sort_desc: typing.Optional[bool]
    width: typing.Optional[float]
    calcs: list[str]
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

 * <span class="badge builder"></span> [VizLegendOptions](./builder-VizLegendOptions.md)
