---
title: <span class="badge object-type-struct"></span> Dashboardv2GroupByVariableKindDatasource
---
# <span class="badge object-type-struct"></span> Dashboardv2GroupByVariableKindDatasource

## Definition

```go
type Dashboardv2GroupByVariableKindDatasource struct {
    Name *string `json:"name,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2GroupByVariableKindDatasource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2GroupByVariableKindDatasource *Dashboardv2GroupByVariableKindDatasource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2GroupByVariableKindDatasource` objects.

```go
func (dashboardv2GroupByVariableKindDatasource *Dashboardv2GroupByVariableKindDatasource) Equals(other Dashboardv2GroupByVariableKindDatasource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2GroupByVariableKindDatasource` fields for violations and returns them.

```go
func (dashboardv2GroupByVariableKindDatasource *Dashboardv2GroupByVariableKindDatasource) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2GroupByVariableKindDatasourceBuilder](./builder-Dashboardv2GroupByVariableKindDatasourceBuilder.md)
