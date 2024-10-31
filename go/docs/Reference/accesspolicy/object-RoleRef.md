---
title: <span class="badge object-type-struct"></span> RoleRef
---
# <span class="badge object-type-struct"></span> RoleRef

## Definition

```go
type RoleRef struct {
    // Policies can apply to roles, teams, or users
    // Applying policies to individual users is supported, but discouraged
    Kind accesspolicy.RoleRefKind `json:"kind"`
    Name string `json:"name"`
    Xname string `json:"xname"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RoleRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (roleRef *RoleRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RoleRef` objects.

```go
func (roleRef *RoleRef) Equals(other RoleRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RoleRef` fields for violations and returns them.

```go
func (roleRef *RoleRef) Validate() error
```

## See also

 * <span class="badge builder"></span> [RoleRefBuilder](./builder-RoleRefBuilder.md)
