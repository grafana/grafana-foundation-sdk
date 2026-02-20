---
title: <span class="badge object-type-class"></span> Action
---
# <span class="badge object-type-class"></span> Action

Dashboard action

## Definition

```python
class Action:
    """
    Dashboard action
    """

    type_val: dashboard.ActionType
    title: str
    fetch: typing.Optional[dashboard.FetchOptions]
    infinity: typing.Optional[dashboard.InfinityOptions]
    confirmation: typing.Optional[str]
    one_click: typing.Optional[bool]
    variables: typing.Optional[list[dashboard.ActionVariable]]
    style: typing.Optional[dashboard.DashboardActionStyle]
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
