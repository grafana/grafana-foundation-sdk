---
title: <span class="badge object-type-class"></span> ArcOption
---
# <span class="badge object-type-class"></span> ArcOption

## Definition

```python
class ArcOption:
    # Field from which to get the value. Values should be less than 1, representing fraction of a circle.
    field: typing.Optional[str]
    # The color of the arc.
    color: typing.Optional[str]
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

 * <span class="badge builder"></span> [ArcOption](./builder-ArcOption.md)
