---
title: <span class="badge object-type-class"></span> RepeatOptions
---
# <span class="badge object-type-class"></span> RepeatOptions

## Definition

```python
class RepeatOptions:
    mode: str
    value: str
    direction: typing.Optional[typing.Literal["h", "v"]]
    max_per_row: typing.Optional[int]
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

 * <span class="badge builder"></span> [RepeatOptions](./builder-RepeatOptions.md)
