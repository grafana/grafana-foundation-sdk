---
title: <span class="badge object-type-class"></span> V2TransformationKind
---
# <span class="badge object-type-class"></span> V2TransformationKind

Dashboard v2 shape: the transformation ID moved from `kind` to `group`.

## Definition

```python
class V2TransformationKind:
    """
    Dashboard v2 shape: the transformation ID moved from `kind` to `group`.
    """

    kind: typing.Literal["Transformation"]
    # The group is the transformation ID
    group: str
    spec: dashboardv2beta1.V2TransformationSpec
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

 * <span class="badge builder"></span> [V2TransformationKind](./builder-V2TransformationKind.md)
