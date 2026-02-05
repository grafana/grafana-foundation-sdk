---
title: <span class="badge object-type-class"></span> AutoGridLayoutSpec
---
# <span class="badge object-type-class"></span> AutoGridLayoutSpec

## Definition

```python
class AutoGridLayoutSpec:
    max_column_count: typing.Optional[float]
    column_width_mode: typing.Literal["narrow", "standard", "wide", "custom"]
    column_width: typing.Optional[float]
    row_height_mode: typing.Literal["short", "standard", "tall", "custom"]
    row_height: typing.Optional[float]
    fill_screen: typing.Optional[bool]
    items: list[dashboardv2beta1.AutoGridLayoutItemKind]
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

 * <span class="badge builder"></span> [AutoGridLayoutSpec](./builder-AutoGridLayoutSpec.md)
