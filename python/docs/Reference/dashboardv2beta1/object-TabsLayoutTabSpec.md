---
title: <span class="badge object-type-class"></span> TabsLayoutTabSpec
---
# <span class="badge object-type-class"></span> TabsLayoutTabSpec

## Definition

```python
class TabsLayoutTabSpec:
    title: typing.Optional[str]
    layout: typing.Union[dashboardv2beta1.GridLayoutKind, dashboardv2beta1.RowsLayoutKind, dashboardv2beta1.AutoGridLayoutKind, dashboardv2beta1.TabsLayoutKind]
    conditional_rendering: typing.Optional[dashboardv2beta1.ConditionalRenderingGroupKind]
    repeat: typing.Optional[dashboardv2beta1.TabRepeatOptions]
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

 * <span class="badge builder"></span> [TabsLayoutTabSpec](./builder-TabsLayoutTabSpec.md)
