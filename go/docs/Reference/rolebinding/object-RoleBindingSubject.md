---
title: <span class="badge object-type-struct"></span> RoleBindingSubject
---
# <span class="badge object-type-struct"></span> RoleBindingSubject

## Definition

```go
type RoleBindingSubject struct {
    Kind rolebinding.RoleBindingSubjectKind `json:"kind"`
    // The team/user identifier name
    Name string `json:"name"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RoleBindingSubject` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (roleBindingSubject *RoleBindingSubject) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RoleBindingSubject` objects.

```go
func (roleBindingSubject *RoleBindingSubject) Equals(other RoleBindingSubject) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RoleBindingSubject` fields for violations and returns them.

```go
func (roleBindingSubject *RoleBindingSubject) Validate() error
```

## See also

 * <span class="badge builder"></span> [RoleBindingSubjectBuilder](./builder-RoleBindingSubjectBuilder.md)
