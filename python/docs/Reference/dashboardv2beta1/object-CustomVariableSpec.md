---
title: <span class="badge object-type-class"></span> CustomVariableSpec
---
# <span class="badge object-type-class"></span> CustomVariableSpec

Custom variable specification

## Definition

```python
class CustomVariableSpec:
    """
    Custom variable specification
    """

    name: str
    query: str
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
    values_format: typing.Optional[typing.Literal["csv", "json"]]
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

