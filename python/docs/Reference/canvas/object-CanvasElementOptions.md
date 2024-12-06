---
title: <span class="badge object-type-class"></span> CanvasElementOptions
---
# <span class="badge object-type-class"></span> CanvasElementOptions

## Definition

```python
class CanvasElementOptions:
    name: str
    type_val: str
    # TODO: figure out how to define this (element config(s))
    config: typing.Optional[object]
    constraint: typing.Optional[canvas.Constraint]
    placement: typing.Optional[canvas.Placement]
    background: typing.Optional[canvas.BackgroundConfig]
    border: typing.Optional[canvas.LineConfig]
    connections: typing.Optional[list[canvas.CanvasConnection]]
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

 * <span class="badge builder"></span> [CanvasElementOptions](./builder-CanvasElementOptions.md)
