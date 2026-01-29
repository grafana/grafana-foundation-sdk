---
title: <span class="badge object-type-class"></span> ConditionalRenderingGroupKind
---
# <span class="badge object-type-class"></span> ConditionalRenderingGroupKind

## Definition

```python
class ConditionalRenderingGroupKind:
    kind: typing.Literal["ConditionalRenderingGroup"]
    spec: dashboardv2beta1.ConditionalRenderingGroupSpec
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

 * <span class="badge builder"></span> [ConditionalRenderingGroup](./builder-ConditionalRenderingGroup.md)
