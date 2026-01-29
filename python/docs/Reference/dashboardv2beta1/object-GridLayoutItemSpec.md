---
title: <span class="badge object-type-class"></span> GridLayoutItemSpec
---
# <span class="badge object-type-class"></span> GridLayoutItemSpec

## Definition

```python
class GridLayoutItemSpec:
    x: int
    y: int
    width: int
    height: int
    # reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
    element: dashboardv2beta1.ElementReference
    repeat: typing.Optional[dashboardv2beta1.RepeatOptions]
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

