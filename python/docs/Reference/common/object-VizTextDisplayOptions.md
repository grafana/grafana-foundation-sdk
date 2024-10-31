---
title: <span class="badge object-type-class"></span> VizTextDisplayOptions
---
# <span class="badge object-type-class"></span> VizTextDisplayOptions

TODO docs

## Definition

```python
class VizTextDisplayOptions:
    """
    TODO docs
    """

    # Explicit title text size
    title_size: typing.Optional[float]
    # Explicit value text size
    value_size: typing.Optional[float]
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

 * <span class="badge builder"></span> [VizTextDisplayOptions](./builder-VizTextDisplayOptions.md)
