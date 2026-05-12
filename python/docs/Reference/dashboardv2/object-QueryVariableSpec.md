---
title: <span class="badge object-type-class"></span> QueryVariableSpec
---
# <span class="badge object-type-class"></span> QueryVariableSpec

Query variable specification

## Definition

```python
class QueryVariableSpec:
    """
    Query variable specification
    """

    name: str
    current: dashboardv2.VariableOption
    label: typing.Optional[str]
    hide: dashboardv2.VariableHide
    refresh: dashboardv2.VariableRefresh
    skip_url_sync: bool
    description: typing.Optional[str]
    query: dashboardv2.DataQueryKind
    regex: str
    regex_apply_to: typing.Optional[dashboardv2.VariableRegexApplyTo]
    sort: dashboardv2.VariableSort
    definition: typing.Optional[str]
    options: list[dashboardv2.VariableOption]
    multi: bool
    include_all: bool
    all_value: typing.Optional[str]
    placeholder: typing.Optional[str]
    allow_custom_value: bool
    static_options: typing.Optional[list[dashboardv2.VariableOption]]
    static_options_order: typing.Optional[typing.Literal["before", "after", "sorted"]]
    origin: typing.Optional[dashboardv2.ControlSourceRef]
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

