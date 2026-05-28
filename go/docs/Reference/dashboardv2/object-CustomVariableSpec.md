---
title: <span class="badge object-type-struct"></span> CustomVariableSpec
---
# <span class="badge object-type-struct"></span> CustomVariableSpec

Custom variable specification

## Definition

```go
type CustomVariableSpec struct {
    Name string `json:"name"`
    Query string `json:"query"`
    Current dashboardv2.VariableOption `json:"current"`
    Options []dashboardv2.VariableOption `json:"options"`
    Multi bool `json:"multi"`
    IncludeAll bool `json:"includeAll"`
    AllValue *string `json:"allValue,omitempty"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
    AllowCustomValue bool `json:"allowCustomValue"`
    ValuesFormat *dashboardv2.CustomVariableSpecValuesFormat `json:"valuesFormat,omitempty"`
    Origin *dashboardv2.ControlSourceRef `json:"origin,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (customVariableSpec *CustomVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CustomVariableSpec` objects.

```go
func (customVariableSpec *CustomVariableSpec) Equals(other CustomVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CustomVariableSpec` fields for violations and returns them.

```go
func (customVariableSpec *CustomVariableSpec) Validate() error
```

