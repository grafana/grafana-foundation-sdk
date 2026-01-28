---
title: <span class="badge object-type-struct"></span> TextVariableSpec
---
# <span class="badge object-type-struct"></span> TextVariableSpec

Text variable specification

## Definition

```go
type TextVariableSpec struct {
    Name string `json:"name"`
    Current dashboardv2beta1.VariableOption `json:"current"`
    Query string `json:"query"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TextVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (textVariableSpec *TextVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TextVariableSpec` objects.

```go
func (textVariableSpec *TextVariableSpec) Equals(other TextVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TextVariableSpec` fields for violations and returns them.

```go
func (textVariableSpec *TextVariableSpec) Validate() error
```

