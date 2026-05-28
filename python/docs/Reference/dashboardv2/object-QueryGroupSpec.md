---
title: <span class="badge object-type-class"></span> QueryGroupSpec
---
# <span class="badge object-type-class"></span> QueryGroupSpec

## Definition

```python
class QueryGroupSpec:
    queries: list[dashboardv2.PanelQueryKind]
    transformations: list[dashboardv2.TransformationKind]
    query_options: dashboardv2.QueryOptionsSpec
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

