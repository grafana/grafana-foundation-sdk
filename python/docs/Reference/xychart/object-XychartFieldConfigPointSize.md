---
title: <span class="badge object-type-class"></span> XychartFieldConfigPointSize
---
# <span class="badge object-type-class"></span> XychartFieldConfigPointSize

## Definition

```python
class XychartFieldConfigPointSize:
    fixed: typing.Optional[int]
    min_val: typing.Optional[int]
    max_val: typing.Optional[int]
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

 * <span class="badge builder"></span> [XychartFieldConfigPointSize](./builder-XychartFieldConfigPointSize.md)
