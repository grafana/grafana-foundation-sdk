---
title: <span class="badge builder"></span> RoleRefBuilder
---
# <span class="badge builder"></span> RoleRefBuilder

## Constructor

```go
func NewRoleRefBuilder() *RoleRefBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RoleRefBuilder) Build() (RoleRef, error)
```

### <span class="badge object-method"></span> Kind

Policies can apply to roles, teams, or users

Applying policies to individual users is supported, but discouraged

```go
func (builder *RoleRefBuilder) Kind(kind accesspolicy.RoleRefKind) *RoleRefBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *RoleRefBuilder) Name(name string) *RoleRefBuilder
```

### <span class="badge object-method"></span> Xname

```go
func (builder *RoleRefBuilder) Xname(xname string) *RoleRefBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RoleRef](./object-RoleRef.md)
