---
title: <span class="badge object-type-class"></span> ResourceGroupsQuery
---
# <span class="badge object-type-class"></span> ResourceGroupsQuery

## Definition

```python
class ResourceGroupsQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["ResourceGroupsQuery"]
    subscription: str
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

 * <span class="badge builder"></span> [ResourceGroupsQuery](./builder-ResourceGroupsQuery.md)
