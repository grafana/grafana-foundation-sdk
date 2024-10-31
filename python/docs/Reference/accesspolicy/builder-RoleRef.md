---
title: <span class="badge builder"></span> RoleRef
---
# <span class="badge builder"></span> RoleRef

## Constructor

```python
RoleRef()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> accesspolicy.RoleRef
```

### <span class="badge object-method"></span> kind

Policies can apply to roles, teams, or users

Applying policies to individual users is supported, but discouraged

```python
def kind(kind: typing.Literal["Role", "BuiltinRole", "Team", "User"]) -> typing.Self
```

### <span class="badge object-method"></span> name

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> xname

```python
def xname(xname: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RoleRef](./object-RoleRef.md)
