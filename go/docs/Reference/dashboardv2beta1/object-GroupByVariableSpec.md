---
title: <span class="badge object-type-struct"></span> GroupByVariableSpec
---
# <span class="badge object-type-struct"></span> GroupByVariableSpec

GroupBy variable specification

## Definition

```go
type GroupByVariableSpec struct {
    Name string `json:"name"`
    DefaultValue *dashboardv2beta1.VariableOption `json:"defaultValue,omitempty"`
    Current dashboardv2beta1.VariableOption `json:"current"`
    Options []dashboardv2beta1.VariableOption `json:"options"`
    Multi bool `json:"multi"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GroupByVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (groupByVariableSpec *GroupByVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GroupByVariableSpec` objects.

```go
func (groupByVariableSpec *GroupByVariableSpec) Equals(other GroupByVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GroupByVariableSpec` fields for violations and returns them.

```go
func (groupByVariableSpec *GroupByVariableSpec) Validate() error
```

