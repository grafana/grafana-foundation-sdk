---
title: <span class="badge object-type-class"></span> TimeOption
---
# <span class="badge object-type-class"></span> TimeOption

Counterpart for TypeScript's TimeOption type.

## Definition

```python
class TimeOption:
    """
    Counterpart for TypeScript's TimeOption type.
    """

    display: str
    from_val: str
    to: str
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

 * <span class="badge builder"></span> [TimeOption](./builder-TimeOption.md)
