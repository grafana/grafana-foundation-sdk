---
title: <span class="badge object-type-struct"></span> VariableValueOption
---
# <span class="badge object-type-struct"></span> VariableValueOption

FIXME: should we introduce this? --- Variable value option

## Definition

```go
type VariableValueOption struct {
    Label string `json:"label"`
    Value dashboardv2beta1.VariableValueSingle `json:"value"`
    Group *string `json:"group,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VariableValueOption` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (variableValueOption *VariableValueOption) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `VariableValueOption` objects.

```go
func (variableValueOption *VariableValueOption) Equals(other VariableValueOption) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VariableValueOption` fields for violations and returns them.

```go
func (variableValueOption *VariableValueOption) Validate() error
```

## See also

 * <span class="badge builder"></span> [VariableValueOptionBuilder](./builder-VariableValueOptionBuilder.md)
