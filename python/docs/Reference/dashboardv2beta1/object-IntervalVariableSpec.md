---
title: <span class="badge object-type-class"></span> IntervalVariableSpec
---
# <span class="badge object-type-class"></span> IntervalVariableSpec

Interval variable specification

## Definition

```python
class IntervalVariableSpec:
    """
    Interval variable specification
    """

    name: str
    query: str
    current: dashboardv2beta1.VariableOption
    options: list[dashboardv2beta1.VariableOption]
    auto: bool
    auto_min: str
    auto_count: int
    refresh: typing.Literal["onTimeRangeChanged"]
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

