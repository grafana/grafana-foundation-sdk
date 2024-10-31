---
title: <span class="badge object-type-struct"></span> BuiltinRoleRefOrCustomRoleRef
---
# <span class="badge object-type-struct"></span> BuiltinRoleRefOrCustomRoleRef

## Definition

```go
type BuiltinRoleRefOrCustomRoleRef struct {
    BuiltinRoleRef *rolebinding.BuiltinRoleRef `json:"BuiltinRoleRef,omitempty"`
    CustomRoleRef *rolebinding.CustomRoleRef `json:"CustomRoleRef,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `BuiltinRoleRefOrCustomRoleRef` as JSON.

```go
func (builtinRoleRefOrCustomRoleRef *BuiltinRoleRefOrCustomRoleRef) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `BuiltinRoleRefOrCustomRoleRef` from JSON.

```go
func (builtinRoleRefOrCustomRoleRef *BuiltinRoleRefOrCustomRoleRef) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuiltinRoleRefOrCustomRoleRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builtinRoleRefOrCustomRoleRef *BuiltinRoleRefOrCustomRoleRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuiltinRoleRefOrCustomRoleRef` objects.

```go
func (builtinRoleRefOrCustomRoleRef *BuiltinRoleRefOrCustomRoleRef) Equals(other BuiltinRoleRefOrCustomRoleRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuiltinRoleRefOrCustomRoleRef` fields for violations and returns them.

```go
func (builtinRoleRefOrCustomRoleRef *BuiltinRoleRefOrCustomRoleRef) Validate() error
```

