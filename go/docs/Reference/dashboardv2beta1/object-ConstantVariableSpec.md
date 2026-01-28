---
title: <span class="badge object-type-struct"></span> ConstantVariableSpec
---
# <span class="badge object-type-struct"></span> ConstantVariableSpec

Constant variable specification

## Definition

```go
type ConstantVariableSpec struct {
    Name string `json:"name"`
    Query string `json:"query"`
    Current dashboardv2beta1.VariableOption `json:"current"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConstantVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (constantVariableSpec *ConstantVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConstantVariableSpec` objects.

```go
func (constantVariableSpec *ConstantVariableSpec) Equals(other ConstantVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConstantVariableSpec` fields for violations and returns them.

```go
func (constantVariableSpec *ConstantVariableSpec) Validate() error
```

