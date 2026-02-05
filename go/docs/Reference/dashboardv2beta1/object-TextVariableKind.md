---
title: <span class="badge object-type-struct"></span> TextVariableKind
---
# <span class="badge object-type-struct"></span> TextVariableKind

Text variable kind

## Definition

```go
type TextVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.TextVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TextVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (textVariableKind *TextVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TextVariableKind` objects.

```go
func (textVariableKind *TextVariableKind) Equals(other TextVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TextVariableKind` fields for violations and returns them.

```go
func (textVariableKind *TextVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [TextVariableBuilder](./builder-TextVariableBuilder.md)
