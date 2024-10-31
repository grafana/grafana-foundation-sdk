---
title: <span class="badge object-type-class"></span> AccessPolicy
---
# <span class="badge object-type-class"></span> AccessPolicy

## Definition

```python
class AccessPolicy:
    # The scope where these policies should apply
    scope: accesspolicy.ResourceRef
    # The role that must apply this policy
    role: accesspolicy.RoleRef
    # The set of rules to apply.  Note that * is required to modify
    # access policy rules, and that "none" will reject all actions
    rules: list[accesspolicy.AccessRule]
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

 * <span class="badge builder"></span> [AccessPolicy](./builder-AccessPolicy.md)
