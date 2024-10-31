---
title: <span class="badge object-type-class"></span> Average
---
# <span class="badge object-type-class"></span> Average

## Definition

```python
class Average:
    type_val: typing.Literal["avg"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional[elasticsearch.ElasticsearchAverageSettings]
    hide: typing.Optional[bool]
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

 * <span class="badge builder"></span> [Average](./builder-Average.md)
