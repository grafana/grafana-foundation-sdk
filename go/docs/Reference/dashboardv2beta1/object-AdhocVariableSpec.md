---
title: <span class="badge object-type-struct"></span> AdhocVariableSpec
---
# <span class="badge object-type-struct"></span> AdhocVariableSpec

Adhoc variable specification

## Definition

```go
type AdhocVariableSpec struct {
    Name string `json:"name"`
    BaseFilters []dashboardv2beta1.AdHocFilterWithLabels `json:"baseFilters"`
    Filters []dashboardv2beta1.AdHocFilterWithLabels `json:"filters"`
    DefaultKeys []dashboardv2beta1.MetricFindValue `json:"defaultKeys"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
    AllowCustomValue bool `json:"allowCustomValue"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdhocVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (adhocVariableSpec *AdhocVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AdhocVariableSpec` objects.

```go
func (adhocVariableSpec *AdhocVariableSpec) Equals(other AdhocVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AdhocVariableSpec` fields for violations and returns them.

```go
func (adhocVariableSpec *AdhocVariableSpec) Validate() error
```

