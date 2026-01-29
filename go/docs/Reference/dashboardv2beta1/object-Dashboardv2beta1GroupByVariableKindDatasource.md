---
title: <span class="badge object-type-struct"></span> Dashboardv2beta1GroupByVariableKindDatasource
---
# <span class="badge object-type-struct"></span> Dashboardv2beta1GroupByVariableKindDatasource

## Definition

```go
type Dashboardv2beta1GroupByVariableKindDatasource struct {
    Name *string `json:"name,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1GroupByVariableKindDatasource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2beta1GroupByVariableKindDatasource *Dashboardv2beta1GroupByVariableKindDatasource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2beta1GroupByVariableKindDatasource` objects.

```go
func (dashboardv2beta1GroupByVariableKindDatasource *Dashboardv2beta1GroupByVariableKindDatasource) Equals(other Dashboardv2beta1GroupByVariableKindDatasource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2beta1GroupByVariableKindDatasource` fields for violations and returns them.

```go
func (dashboardv2beta1GroupByVariableKindDatasource *Dashboardv2beta1GroupByVariableKindDatasource) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2beta1GroupByVariableKindDatasourceBuilder](./builder-Dashboardv2beta1GroupByVariableKindDatasourceBuilder.md)
