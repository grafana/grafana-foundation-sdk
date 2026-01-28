---
title: <span class="badge object-type-class"></span> AutoGridLayoutItemSpec
---
# <span class="badge object-type-class"></span> AutoGridLayoutItemSpec

## Definition

```python
class AutoGridLayoutItemSpec:
    element: dashboardv2beta1.ElementReference
    repeat: typing.Optional[dashboardv2beta1.AutoGridRepeatOptions]
    conditional_rendering: typing.Optional[dashboardv2beta1.ConditionalRenderingGroupKind]
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

 * <span class="badge builder"></span> [AutoGridLayoutItemSpec](./builder-AutoGridLayoutItemSpec.md)
