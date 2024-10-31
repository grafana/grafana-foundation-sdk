---
title: <span class="badge object-type-class"></span> FilterValueRange
---
# <span class="badge object-type-class"></span> FilterValueRange

Controls the value filter range

## Definition

```python
class FilterValueRange:
    """
    Controls the value filter range
    """

    # Sets the filter range to values less than or equal to the given value
    le: typing.Optional[float]
    # Sets the filter range to values greater than or equal to the given value
    ge: typing.Optional[float]
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

 * <span class="badge builder"></span> [FilterValueRange](./builder-FilterValueRange.md)
