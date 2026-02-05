---
title: <span class="badge object-type-class"></span> PanelSpec
---
# <span class="badge object-type-class"></span> PanelSpec

## Definition

```python
class PanelSpec:
    id_val: float
    title: str
    description: str
    links: list[dashboardv2beta1.DataLink]
    data: dashboardv2beta1.QueryGroupKind
    viz_config: dashboardv2beta1.VizConfigKind
    transparent: typing.Optional[bool]
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

