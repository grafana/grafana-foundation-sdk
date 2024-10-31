---
title: <span class="badge builder"></span> RoleBinding
---
# <span class="badge builder"></span> RoleBinding

## Constructor

```python
RoleBinding()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> rolebinding.RoleBinding
```

### <span class="badge object-method"></span> role

The role we are discussing

```python
def role(role: typing.Union[cogbuilder.Builder[rolebinding.BuiltinRoleRef], cogbuilder.Builder[rolebinding.CustomRoleRef]]) -> typing.Self
```

### <span class="badge object-method"></span> subject

The team or user that has the specified role

```python
def subject(subject: cogbuilder.Builder[rolebinding.RoleBindingSubject]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RoleBinding](./object-RoleBinding.md)
