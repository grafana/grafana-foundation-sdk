---
title: <span class="badge object-type-struct"></span> Dashboardv2beta1DataQueryKindDatasource
---
# <span class="badge object-type-struct"></span> Dashboardv2beta1DataQueryKindDatasource

## Definition

```go
type Dashboardv2beta1DataQueryKindDatasource struct {
    Name *string `json:"name,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1DataQueryKindDatasource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2beta1DataQueryKindDatasource *Dashboardv2beta1DataQueryKindDatasource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2beta1DataQueryKindDatasource` objects.

```go
func (dashboardv2beta1DataQueryKindDatasource *Dashboardv2beta1DataQueryKindDatasource) Equals(other Dashboardv2beta1DataQueryKindDatasource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2beta1DataQueryKindDatasource` fields for violations and returns them.

```go
func (dashboardv2beta1DataQueryKindDatasource *Dashboardv2beta1DataQueryKindDatasource) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2beta1DataQueryKindDatasourceBuilder](./builder-Dashboardv2beta1DataQueryKindDatasourceBuilder.md)
