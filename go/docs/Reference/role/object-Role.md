---
title: <span class="badge object-type-struct"></span> Role
---
# <span class="badge object-type-struct"></span> Role

## Definition

```go
type Role struct {
    // The role identifier `managed:builtins:editor:permissions`
    Name string `json:"name"`
    // Optional display
    DisplayName *string `json:"displayName,omitempty"`
    // Name of the team.
    GroupName *string `json:"groupName,omitempty"`
    // Role description
    Description *string `json:"description,omitempty"`
    // Do not show this role
    Hidden bool `json:"hidden"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Role` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (role *Role) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Role` objects.

```go
func (role *Role) Equals(other Role) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Role` fields for violations and returns them.

```go
func (role *Role) Validate() error
```

## See also

 * <span class="badge builder"></span> [RoleBuilder](./builder-RoleBuilder.md)
