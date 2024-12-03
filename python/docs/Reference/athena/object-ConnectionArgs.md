---
title: <span class="badge object-type-class"></span> ConnectionArgs
---
# <span class="badge object-type-class"></span> ConnectionArgs

## Definition

```python
class ConnectionArgs:
    region: typing.Optional[str]
    catalog: typing.Optional[str]
    database: typing.Optional[str]
    result_reuse_enabled: typing.Optional[bool]
    result_reuse_max_age_in_minutes: typing.Optional[float]
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

 * <span class="badge builder"></span> [ConnectionArgs](./builder-ConnectionArgs.md)
