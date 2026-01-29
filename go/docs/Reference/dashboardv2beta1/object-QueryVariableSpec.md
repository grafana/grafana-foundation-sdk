---
title: <span class="badge object-type-struct"></span> QueryVariableSpec
---
# <span class="badge object-type-struct"></span> QueryVariableSpec

Query variable specification

## Definition

```go
type QueryVariableSpec struct {
    Name string `json:"name"`
    Current dashboardv2beta1.VariableOption `json:"current"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    Refresh dashboardv2beta1.VariableRefresh `json:"refresh"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
    Query dashboardv2beta1.DataQueryKind `json:"query"`
    Regex string `json:"regex"`
    RegexApplyTo *dashboardv2beta1.VariableRegexApplyTo `json:"regexApplyTo,omitempty"`
    Sort dashboardv2beta1.VariableSort `json:"sort"`
    Definition *string `json:"definition,omitempty"`
    Options []dashboardv2beta1.VariableOption `json:"options"`
    Multi bool `json:"multi"`
    IncludeAll bool `json:"includeAll"`
    AllValue *string `json:"allValue,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    AllowCustomValue bool `json:"allowCustomValue"`
    StaticOptions []dashboardv2beta1.VariableOption `json:"staticOptions,omitempty"`
    StaticOptionsOrder *dashboardv2beta1.QueryVariableSpecStaticOptionsOrder `json:"staticOptionsOrder,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryVariableSpec *QueryVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryVariableSpec` objects.

```go
func (queryVariableSpec *QueryVariableSpec) Equals(other QueryVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryVariableSpec` fields for violations and returns them.

```go
func (queryVariableSpec *QueryVariableSpec) Validate() error
```

