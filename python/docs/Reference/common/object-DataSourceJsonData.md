---
title: <span class="badge object-type-class"></span> DataSourceJsonData
---
# <span class="badge object-type-class"></span> DataSourceJsonData

TODO docs

## Definition

```python
class DataSourceJsonData:
    """
    TODO docs
    """

    auth_type: typing.Optional[str]
    default_region: typing.Optional[str]
    profile: typing.Optional[str]
    manage_alerts: typing.Optional[bool]
    alertmanager_uid: typing.Optional[str]
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

 * <span class="badge builder"></span> [DataSourceJsonData](./builder-DataSourceJsonData.md)
