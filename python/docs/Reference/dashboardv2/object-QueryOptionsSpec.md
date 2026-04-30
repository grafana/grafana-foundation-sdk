---
title: <span class="badge object-type-class"></span> QueryOptionsSpec
---
# <span class="badge object-type-class"></span> QueryOptionsSpec

## Definition

```python
class QueryOptionsSpec:
    time_from: typing.Optional[str]
    max_data_points: typing.Optional[int]
    time_shift: typing.Optional[str]
    query_caching_ttl: typing.Optional[int]
    interval: typing.Optional[str]
    cache_timeout: typing.Optional[str]
    hide_time_override: typing.Optional[bool]
    time_compare: typing.Optional[str]
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

 * <span class="badge builder"></span> [QueryOptionsSpec](./builder-QueryOptionsSpec.md)
