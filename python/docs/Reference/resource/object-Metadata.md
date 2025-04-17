---
title: <span class="badge object-type-class"></span> Metadata
---
# <span class="badge object-type-class"></span> Metadata

## Definition

```python
class Metadata:
    name: str
    namespace: typing.Optional[str]
    labels: typing.Optional[dict[str, str]]
    annotations: typing.Optional[dict[str, str]]
    uid: typing.Optional[str]
    resource_version: typing.Optional[str]
    generation: typing.Optional[int]
    creation_timestamp: typing.Optional[str]
    update_timestamp: typing.Optional[str]
    deletion_timestamp: typing.Optional[str]
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

 * <span class="badge builder"></span> [Metadata](./builder-Metadata.md)
