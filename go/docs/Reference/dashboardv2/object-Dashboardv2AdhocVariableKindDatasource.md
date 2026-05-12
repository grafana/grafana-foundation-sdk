---
title: <span class="badge object-type-struct"></span> Dashboardv2AdhocVariableKindDatasource
---
# <span class="badge object-type-struct"></span> Dashboardv2AdhocVariableKindDatasource

## Definition

```go
type Dashboardv2AdhocVariableKindDatasource struct {
    Name *string `json:"name,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2AdhocVariableKindDatasource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2AdhocVariableKindDatasource *Dashboardv2AdhocVariableKindDatasource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2AdhocVariableKindDatasource` objects.

```go
func (dashboardv2AdhocVariableKindDatasource *Dashboardv2AdhocVariableKindDatasource) Equals(other Dashboardv2AdhocVariableKindDatasource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2AdhocVariableKindDatasource` fields for violations and returns them.

```go
func (dashboardv2AdhocVariableKindDatasource *Dashboardv2AdhocVariableKindDatasource) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2AdhocVariableKindDatasourceBuilder](./builder-Dashboardv2AdhocVariableKindDatasourceBuilder.md)
