---
title: <span class="badge object-type-class"></span> AccessRule
---
# <span class="badge object-type-class"></span> AccessRule

## Definition

```python
class AccessRule:
    # The kind this rule applies to (dashboards, alert, etc)
    kind: str
    # READ, WRITE, CREATE, DELETE, ...
    # should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
    verb: typing.Union[typing.Literal["*"]]
    # Specific sub-elements like "alert.rules" or "dashboard.permissions"????
    target: typing.Optional[str]
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

 * <span class="badge builder"></span> [AccessRule](./builder-AccessRule.md)
