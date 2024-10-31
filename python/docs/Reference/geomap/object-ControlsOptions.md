---
title: <span class="badge object-type-class"></span> ControlsOptions
---
# <span class="badge object-type-class"></span> ControlsOptions

## Definition

```python
class ControlsOptions:
    # Zoom (upper left)
    show_zoom: typing.Optional[bool]
    # let the mouse wheel zoom
    mouse_wheel_zoom: typing.Optional[bool]
    # Lower right
    show_attribution: typing.Optional[bool]
    # Scale options
    show_scale: typing.Optional[bool]
    # Show debug
    show_debug: typing.Optional[bool]
    # Show measure
    show_measure: typing.Optional[bool]
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

 * <span class="badge builder"></span> [ControlsOptions](./builder-ControlsOptions.md)
