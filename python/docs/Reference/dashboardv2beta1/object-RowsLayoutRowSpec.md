---
title: <span class="badge object-type-class"></span> RowsLayoutRowSpec
---
# <span class="badge object-type-class"></span> RowsLayoutRowSpec

## Definition

```python
class RowsLayoutRowSpec:
    title: typing.Optional[str]
    collapse: typing.Optional[bool]
    hide_header: typing.Optional[bool]
    fill_screen: typing.Optional[bool]
    conditional_rendering: typing.Optional[dashboardv2beta1.ConditionalRenderingGroupKind]
    repeat: typing.Optional[dashboardv2beta1.RowRepeatOptions]
    layout: typing.Union[dashboardv2beta1.GridLayoutKind, dashboardv2beta1.AutoGridLayoutKind, dashboardv2beta1.TabsLayoutKind, dashboardv2beta1.RowsLayoutKind]
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

 * <span class="badge builder"></span> [RowsLayoutRowSpec](./builder-RowsLayoutRowSpec.md)
