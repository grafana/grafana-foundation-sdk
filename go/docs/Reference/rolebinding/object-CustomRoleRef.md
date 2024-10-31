---
title: <span class="badge object-type-struct"></span> CustomRoleRef
---
# <span class="badge object-type-struct"></span> CustomRoleRef

## Definition

```go
type CustomRoleRef struct {
    Kind string `json:"kind"`
    Name string `json:"name"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomRoleRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (customRoleRef *CustomRoleRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CustomRoleRef` objects.

```go
func (customRoleRef *CustomRoleRef) Equals(other CustomRoleRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CustomRoleRef` fields for violations and returns them.

```go
func (customRoleRef *CustomRoleRef) Validate() error
```

## See also

 * <span class="badge builder"></span> [CustomRoleRefBuilder](./builder-CustomRoleRefBuilder.md)
