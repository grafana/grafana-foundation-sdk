---
title: <span class="badge object-type-class"></span> DatasourceVariableKind
---
# <span class="badge object-type-class"></span> DatasourceVariableKind

Datasource variable kind

## Definition

```python
class DatasourceVariableKind:
    """
    Datasource variable kind
    """

    kind: typing.Literal["DatasourceVariable"]
    spec: dashboardv2beta1.DatasourceVariableSpec
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

 * <span class="badge builder"></span> [DatasourceVariable](./builder-DatasourceVariable.md)
