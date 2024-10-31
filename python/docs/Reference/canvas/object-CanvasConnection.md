---
title: <span class="badge object-type-class"></span> CanvasConnection
---
# <span class="badge object-type-class"></span> CanvasConnection

## Definition

```python
class CanvasConnection:
    source: canvas.ConnectionCoordinates
    target: canvas.ConnectionCoordinates
    target_name: typing.Optional[str]
    path: canvas.ConnectionPath
    color: typing.Optional[common.ColorDimensionConfig]
    size: typing.Optional[common.ScaleDimensionConfig]
    vertices: typing.Optional[list[canvas.ConnectionCoordinates]]
    source_original: typing.Optional[canvas.ConnectionCoordinates]
    target_original: typing.Optional[canvas.ConnectionCoordinates]
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

 * <span class="badge builder"></span> [CanvasConnection](./builder-CanvasConnection.md)
