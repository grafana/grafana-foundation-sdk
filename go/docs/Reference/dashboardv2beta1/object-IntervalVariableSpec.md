---
title: <span class="badge object-type-struct"></span> IntervalVariableSpec
---
# <span class="badge object-type-struct"></span> IntervalVariableSpec

Interval variable specification

## Definition

```go
type IntervalVariableSpec struct {
    Name string `json:"name"`
    Query string `json:"query"`
    Current dashboardv2beta1.VariableOption `json:"current"`
    Options []dashboardv2beta1.VariableOption `json:"options"`
    Auto bool `json:"auto"`
    AutoMin string `json:"auto_min"`
    AutoCount int64 `json:"auto_count"`
    Refresh dashboardv2beta1.VariableRefresh `json:"refresh"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `IntervalVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (intervalVariableSpec *IntervalVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `IntervalVariableSpec` objects.

```go
func (intervalVariableSpec *IntervalVariableSpec) Equals(other IntervalVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `IntervalVariableSpec` fields for violations and returns them.

```go
func (intervalVariableSpec *IntervalVariableSpec) Validate() error
```

