---
title: <span class="badge object-type-struct"></span> VariableOption
---
# <span class="badge object-type-struct"></span> VariableOption

Option to be selected in a variable.

## Definition

```go
type VariableOption struct {
    // Whether the option is selected or not
    Selected *bool `json:"selected,omitempty"`
    // Text to be displayed for the option
    Text dashboard.StringOrArrayOfString `json:"text"`
    // Value of the option
    Value dashboard.StringOrArrayOfString `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VariableOption` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (variableOption *VariableOption) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `VariableOption` objects.

```go
func (variableOption *VariableOption) Equals(other VariableOption) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VariableOption` fields for violations and returns them.

```go
func (variableOption *VariableOption) Validate() error
```

