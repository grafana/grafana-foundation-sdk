---
title: <span class="badge object-type-class"></span> DataSourceRef
---
# <span class="badge object-type-class"></span> DataSourceRef

Ref to a DataSource instance

## Definition

```python
class DataSourceRef:
    """
    Ref to a DataSource instance
    """

    # The plugin type-id
    type_val: typing.Optional[str]
    # Specific datasource instance
    uid: typing.Optional[str]
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

