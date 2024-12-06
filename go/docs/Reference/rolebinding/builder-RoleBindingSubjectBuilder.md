---
title: <span class="badge builder"></span> RoleBindingSubjectBuilder
---
# <span class="badge builder"></span> RoleBindingSubjectBuilder

## Constructor

```go
func NewRoleBindingSubjectBuilder() *RoleBindingSubjectBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RoleBindingSubjectBuilder) Build() (RoleBindingSubject, error)
```

### <span class="badge object-method"></span> Kind

```go
func (builder *RoleBindingSubjectBuilder) Kind(kind rolebinding.RoleBindingSubjectKind) *RoleBindingSubjectBuilder
```

### <span class="badge object-method"></span> Name

The team/user identifier name

```go
func (builder *RoleBindingSubjectBuilder) Name(name string) *RoleBindingSubjectBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RoleBindingSubject](./object-RoleBindingSubject.md)
