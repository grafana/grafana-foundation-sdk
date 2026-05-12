---
title: <span class="badge object-type-class"></span> TabsLayoutTabSpec
---
# <span class="badge object-type-class"></span> TabsLayoutTabSpec

## Definition

```python
class TabsLayoutTabSpec:
    title: typing.Optional[str]
    layout: typing.Union[dashboardv2.GridLayoutKind, dashboardv2.RowsLayoutKind, dashboardv2.AutoGridLayoutKind, dashboardv2.TabsLayoutKind]
    conditional_rendering: typing.Optional[dashboardv2.ConditionalRenderingGroupKind]
    repeat: typing.Optional[dashboardv2.TabRepeatOptions]
    variables: typing.Optional[list[dashboardv2.VariableKind]]
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

