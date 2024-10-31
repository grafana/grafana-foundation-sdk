---
title: <span class="badge object-type-class"></span> RoleRef
---
# <span class="badge object-type-class"></span> RoleRef

## Definition

```python
class RoleRef:
    # Policies can apply to roles, teams, or users
    # Applying policies to individual users is supported, but discouraged
    kind: typing.Literal["Role", "BuiltinRole", "Team", "User"]
    name: str
    xname: str
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

 * <span class="badge builder"></span> [RoleRef](./builder-RoleRef.md)
