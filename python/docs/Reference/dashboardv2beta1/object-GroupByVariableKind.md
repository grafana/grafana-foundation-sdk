---
title: <span class="badge object-type-class"></span> GroupByVariableKind
---
# <span class="badge object-type-class"></span> GroupByVariableKind

Group variable kind

## Definition

```python
class GroupByVariableKind:
    """
    Group variable kind
    """

    kind: typing.Literal["GroupByVariable"]
    group: str
    datasource: typing.Optional[dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource]
    spec: dashboardv2beta1.GroupByVariableSpec
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

 * <span class="badge builder"></span> [GroupByVariable](./builder-GroupByVariable.md)
