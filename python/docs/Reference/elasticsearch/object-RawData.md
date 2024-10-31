---
title: <span class="badge object-type-class"></span> RawData
---
# <span class="badge object-type-class"></span> RawData

## Definition

```python
class RawData:
    type_val: typing.Literal["raw_data"]
    id_val: str
    settings: typing.Optional[elasticsearch.ElasticsearchRawDataSettings]
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

 * <span class="badge builder"></span> [RawData](./builder-RawData.md)
