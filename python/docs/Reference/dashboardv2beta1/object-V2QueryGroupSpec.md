---
title: <span class="badge object-type-class"></span> V2QueryGroupSpec
---
# <span class="badge object-type-class"></span> V2QueryGroupSpec

## Definition

```python
class V2QueryGroupSpec:
    queries: list[dashboardv2beta1.PanelQueryKind]
    transformations: list[dashboardv2beta1.V2TransformationKind]
    query_options: dashboardv2beta1.QueryOptionsSpec
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

 * <span class="badge builder"></span> [V2QueryGroupSpec](./builder-V2QueryGroupSpec.md)
