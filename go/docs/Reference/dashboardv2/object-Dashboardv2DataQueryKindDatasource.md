---
title: <span class="badge object-type-struct"></span> Dashboardv2DataQueryKindDatasource
---
# <span class="badge object-type-struct"></span> Dashboardv2DataQueryKindDatasource

## Definition

```go
type Dashboardv2DataQueryKindDatasource struct {
    Name *string `json:"name,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2DataQueryKindDatasource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2DataQueryKindDatasource *Dashboardv2DataQueryKindDatasource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2DataQueryKindDatasource` objects.

```go
func (dashboardv2DataQueryKindDatasource *Dashboardv2DataQueryKindDatasource) Equals(other Dashboardv2DataQueryKindDatasource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2DataQueryKindDatasource` fields for violations and returns them.

```go
func (dashboardv2DataQueryKindDatasource *Dashboardv2DataQueryKindDatasource) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2DataQueryKindDatasourceBuilder](./builder-Dashboardv2DataQueryKindDatasourceBuilder.md)
