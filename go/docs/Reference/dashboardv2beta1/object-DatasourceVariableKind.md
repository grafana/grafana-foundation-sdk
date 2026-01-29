---
title: <span class="badge object-type-struct"></span> DatasourceVariableKind
---
# <span class="badge object-type-struct"></span> DatasourceVariableKind

Datasource variable kind

## Definition

```go
type DatasourceVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.DatasourceVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DatasourceVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (datasourceVariableKind *DatasourceVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DatasourceVariableKind` objects.

```go
func (datasourceVariableKind *DatasourceVariableKind) Equals(other DatasourceVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DatasourceVariableKind` fields for violations and returns them.

```go
func (datasourceVariableKind *DatasourceVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [DatasourceVariableBuilder](./builder-DatasourceVariableBuilder.md)
