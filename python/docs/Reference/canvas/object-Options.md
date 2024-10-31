---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Enable inline editing
    inline_editing: bool
    # Show all available element types
    show_advanced_types: bool
    # Enable pan and zoom
    pan_zoom: bool
    # The root element of canvas (frame), where all canvas elements are nested
    # TODO: Figure out how to define a default value for this
    root: canvas.CanvasOptionsRoot
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

