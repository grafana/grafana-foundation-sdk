---
title: <span class="badge object-type-class"></span> Action
---
# <span class="badge object-type-class"></span> Action

## Definition

```python
class Action:
    type_val: dashboardv2beta1.ActionType
    title: str
    fetch: typing.Optional[dashboardv2beta1.FetchOptions]
    infinity: typing.Optional[dashboardv2beta1.InfinityOptions]
    confirmation: typing.Optional[str]
    one_click: typing.Optional[bool]
    variables: typing.Optional[list[dashboardv2beta1.ActionVariable]]
    style: typing.Optional[dashboardv2beta1.Dashboardv2beta1ActionStyle]
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

 * <span class="badge builder"></span> [Action](./builder-Action.md)
