---
title: <span class="badge object-type-class"></span> CanvasOptionsRoot
---
# <span class="badge object-type-class"></span> CanvasOptionsRoot

## Definition

```python
class CanvasOptionsRoot:
    # Name of the root element
    name: str
    # Type of root element (frame)
    type_val: typing.Literal["frame"]
    # The list of canvas elements attached to the root element
    elements: list[canvas.CanvasElementOptions]
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

 * <span class="badge builder"></span> [CanvasOptionsRoot](./builder-CanvasOptionsRoot.md)
