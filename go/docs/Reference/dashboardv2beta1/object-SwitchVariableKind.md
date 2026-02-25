---
title: <span class="badge object-type-struct"></span> SwitchVariableKind
---
# <span class="badge object-type-struct"></span> SwitchVariableKind

## Definition

```go
type SwitchVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.SwitchVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SwitchVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (switchVariableKind *SwitchVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SwitchVariableKind` objects.

```go
func (switchVariableKind *SwitchVariableKind) Equals(other SwitchVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SwitchVariableKind` fields for violations and returns them.

```go
func (switchVariableKind *SwitchVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [SwitchVariableBuilder](./builder-SwitchVariableBuilder.md)
