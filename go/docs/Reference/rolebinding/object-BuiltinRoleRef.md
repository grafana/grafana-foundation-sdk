---
title: <span class="badge object-type-struct"></span> BuiltinRoleRef
---
# <span class="badge object-type-struct"></span> BuiltinRoleRef

## Definition

```go
type BuiltinRoleRef struct {
    Kind string `json:"kind"`
    Name rolebinding.BuiltinRoleRefName `json:"name"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuiltinRoleRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builtinRoleRef *BuiltinRoleRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuiltinRoleRef` objects.

```go
func (builtinRoleRef *BuiltinRoleRef) Equals(other BuiltinRoleRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuiltinRoleRef` fields for violations and returns them.

```go
func (builtinRoleRef *BuiltinRoleRef) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuiltinRoleRefBuilder](./builder-BuiltinRoleRefBuilder.md)
