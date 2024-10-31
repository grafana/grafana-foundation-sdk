---
title: <span class="badge object-type-class"></span> Role
---
# <span class="badge object-type-class"></span> Role

## Definition

```python
class Role:
    # The role identifier `managed:builtins:editor:permissions`
    name: str
    # Optional display
    display_name: typing.Optional[str]
    # Name of the team.
    group_name: typing.Optional[str]
    # Role description
    description: typing.Optional[str]
    # Do not show this role
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

## See also

 * <span class="badge builder"></span> [Role](./builder-Role.md)
