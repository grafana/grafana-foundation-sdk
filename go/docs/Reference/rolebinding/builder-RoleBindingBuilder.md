---
title: <span class="badge builder"></span> RoleBindingBuilder
---
# <span class="badge builder"></span> RoleBindingBuilder

## Constructor

```go
func NewRoleBindingBuilder() *RoleBindingBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RoleBindingBuilder) Build() (RoleBinding, error)
```

### <span class="badge object-method"></span> Role

The role we are discussing

```go
func (builder *RoleBindingBuilder) Role(role rolebinding.BuiltinRoleRefOrCustomRoleRef) *RoleBindingBuilder
```

### <span class="badge object-method"></span> Subject

The team or user that has the specified role

```go
func (builder *RoleBindingBuilder) Subject(subject cog.Builder[rolebinding.RoleBindingSubject]) *RoleBindingBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RoleBinding](./object-RoleBinding.md)
