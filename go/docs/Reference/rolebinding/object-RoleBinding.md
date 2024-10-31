---
title: <span class="badge object-type-struct"></span> RoleBinding
---
# <span class="badge object-type-struct"></span> RoleBinding

## Definition

```go
type RoleBinding struct {
    // The role we are discussing
    Role rolebinding.BuiltinRoleRefOrCustomRoleRef `json:"role"`
    // The team or user that has the specified role
    Subject rolebinding.RoleBindingSubject `json:"subject"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RoleBinding` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (roleBinding *RoleBinding) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RoleBinding` objects.

```go
func (roleBinding *RoleBinding) Equals(other RoleBinding) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RoleBinding` fields for violations and returns them.

```go
func (roleBinding *RoleBinding) Validate() error
```

## See also

 * <span class="badge builder"></span> [RoleBindingBuilder](./builder-RoleBindingBuilder.md)
