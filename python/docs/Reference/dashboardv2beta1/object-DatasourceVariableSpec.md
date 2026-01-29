---
title: <span class="badge object-type-class"></span> DatasourceVariableSpec
---
# <span class="badge object-type-class"></span> DatasourceVariableSpec

Datasource variable specification

## Definition

```python
class DatasourceVariableSpec:
    """
    Datasource variable specification
    """

    name: str
    plugin_id: str
    refresh: dashboardv2beta1.VariableRefresh
    regex: str
    current: dashboardv2beta1.VariableOption
    options: list[dashboardv2beta1.VariableOption]
    multi: bool
    include_all: bool
    all_value: typing.Optional[str]
    label: typing.Optional[str]
    hide: dashboardv2beta1.VariableHide
    skip_url_sync: bool
    description: typing.Optional[str]
    allow_custom_value: bool
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

