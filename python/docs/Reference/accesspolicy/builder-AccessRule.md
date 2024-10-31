---
title: <span class="badge builder"></span> AccessRule
---
# <span class="badge builder"></span> AccessRule

## Constructor

```python
AccessRule()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> accesspolicy.AccessRule
```

### <span class="badge object-method"></span> kind

The kind this rule applies to (dashboards, alert, etc)

```python
def kind(kind: str) -> typing.Self
```

### <span class="badge object-method"></span> target

Specific sub-elements like "alert.rules" or "dashboard.permissions"????

```python
def target(target: str) -> typing.Self
```

### <span class="badge object-method"></span> verb

READ, WRITE, CREATE, DELETE, ...

should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"

```python
def verb(verb: typing.Union[typing.Literal["*"]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AccessRule](./object-AccessRule.md)
