---
title: <span class="badge object-type-class"></span> AdhocVariableSpec
---
# <span class="badge object-type-class"></span> AdhocVariableSpec

Adhoc variable specification

## Definition

```python
class AdhocVariableSpec:
    """
    Adhoc variable specification
    """

    name: str
    base_filters: list[dashboardv2beta1.AdHocFilterWithLabels]
    filters: list[dashboardv2beta1.AdHocFilterWithLabels]
    default_keys: list[dashboardv2beta1.MetricFindValue]
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

