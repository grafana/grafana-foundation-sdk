---
title: <span class="badge object-type-struct"></span> DatasourceVariableSpec
---
# <span class="badge object-type-struct"></span> DatasourceVariableSpec

Datasource variable specification

## Definition

```go
type DatasourceVariableSpec struct {
    Name string `json:"name"`
    PluginId string `json:"pluginId"`
    Refresh dashboardv2beta1.VariableRefresh `json:"refresh"`
    Regex string `json:"regex"`
    Current dashboardv2beta1.VariableOption `json:"current"`
    Options []dashboardv2beta1.VariableOption `json:"options"`
    Multi bool `json:"multi"`
    IncludeAll bool `json:"includeAll"`
    AllValue *string `json:"allValue,omitempty"`
    Label *string `json:"label,omitempty"`
    Hide dashboardv2beta1.VariableHide `json:"hide"`
    SkipUrlSync bool `json:"skipUrlSync"`
    Description *string `json:"description,omitempty"`
    AllowCustomValue bool `json:"allowCustomValue"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DatasourceVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (datasourceVariableSpec *DatasourceVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DatasourceVariableSpec` objects.

```go
func (datasourceVariableSpec *DatasourceVariableSpec) Equals(other DatasourceVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DatasourceVariableSpec` fields for violations and returns them.

```go
func (datasourceVariableSpec *DatasourceVariableSpec) Validate() error
```

