---
title: <span class="badge object-type-struct"></span> CustomVariableKind
---
# <span class="badge object-type-struct"></span> CustomVariableKind

Custom variable kind

## Definition

```go
type CustomVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.CustomVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (customVariableKind *CustomVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CustomVariableKind` objects.

```go
func (customVariableKind *CustomVariableKind) Equals(other CustomVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CustomVariableKind` fields for violations and returns them.

```go
func (customVariableKind *CustomVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [CustomVariableBuilder](./builder-CustomVariableBuilder.md)
