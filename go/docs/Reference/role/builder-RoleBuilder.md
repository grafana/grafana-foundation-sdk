---
title: <span class="badge builder"></span> RoleBuilder
---
# <span class="badge builder"></span> RoleBuilder

## Constructor

```go
func NewRoleBuilder() *RoleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RoleBuilder) Build() (Role, error)
```

### <span class="badge object-method"></span> Description

Role description

```go
func (builder *RoleBuilder) Description(description string) *RoleBuilder
```

### <span class="badge object-method"></span> DisplayName

Optional display

```go
func (builder *RoleBuilder) DisplayName(displayName string) *RoleBuilder
```

### <span class="badge object-method"></span> GroupName

Name of the team.

```go
func (builder *RoleBuilder) GroupName(groupName string) *RoleBuilder
```

### <span class="badge object-method"></span> Hidden

Do not show this role

```go
func (builder *RoleBuilder) Hidden(hidden bool) *RoleBuilder
```

### <span class="badge object-method"></span> Name

The role identifier `managed:builtins:editor:permissions`

```go
func (builder *RoleBuilder) Name(name string) *RoleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Role](./object-Role.md)
