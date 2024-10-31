---
title: <span class="badge object-type-class"></span> Placement
---
# <span class="badge object-type-class"></span> Placement

## Definition

```python
class Placement:
    top: typing.Optional[float]
    left: typing.Optional[float]
    right: typing.Optional[float]
    bottom: typing.Optional[float]
    width: typing.Optional[float]
    height: typing.Optional[float]
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

 * <span class="badge builder"></span> [Placement](./builder-Placement.md)
