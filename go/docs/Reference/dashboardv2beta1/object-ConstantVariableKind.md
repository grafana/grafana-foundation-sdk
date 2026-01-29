---
title: <span class="badge object-type-struct"></span> ConstantVariableKind
---
# <span class="badge object-type-struct"></span> ConstantVariableKind

Constant variable kind

## Definition

```go
type ConstantVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.ConstantVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConstantVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (constantVariableKind *ConstantVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConstantVariableKind` objects.

```go
func (constantVariableKind *ConstantVariableKind) Equals(other ConstantVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConstantVariableKind` fields for violations and returns them.

```go
func (constantVariableKind *ConstantVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConstantVariableBuilder](./builder-ConstantVariableBuilder.md)
