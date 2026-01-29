---
title: <span class="badge object-type-class"></span> PanelQuerySpec
---
# <span class="badge object-type-class"></span> PanelQuerySpec

## Definition

```python
class PanelQuerySpec:
    query: dashboardv2beta1.DataQueryKind
    ref_id: str
    hidden: bool
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

