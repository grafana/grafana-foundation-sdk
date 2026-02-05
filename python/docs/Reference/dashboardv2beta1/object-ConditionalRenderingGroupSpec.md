---
title: <span class="badge object-type-class"></span> ConditionalRenderingGroupSpec
---
# <span class="badge object-type-class"></span> ConditionalRenderingGroupSpec

## Definition

```python
class ConditionalRenderingGroupSpec:
    visibility: typing.Literal["show", "hide"]
    condition: typing.Literal["and", "or"]
    items: list[typing.Union[dashboardv2beta1.ConditionalRenderingVariableKind, dashboardv2beta1.ConditionalRenderingDataKind, dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind]]
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

 * <span class="badge builder"></span> [ConditionalRenderingGroupSpec](./builder-ConditionalRenderingGroupSpec.md)
