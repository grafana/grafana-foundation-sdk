---
title: <span class="badge object-type-class"></span> GroupByVariableSpec
---
# <span class="badge object-type-class"></span> GroupByVariableSpec

GroupBy variable specification

## Definition

```python
class GroupByVariableSpec:
    """
    GroupBy variable specification
    """

    name: str
    default_value: typing.Optional[dashboardv2beta1.VariableOption]
    current: dashboardv2beta1.VariableOption
    options: list[dashboardv2beta1.VariableOption]
    multi: bool
    label: typing.Optional[str]
    hide: dashboardv2beta1.VariableHide
    skip_url_sync: bool
    description: typing.Optional[str]
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

