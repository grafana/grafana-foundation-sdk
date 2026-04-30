---
title: <span class="badge object-type-struct"></span> DatasourceControlSourceRef
---
# <span class="badge object-type-struct"></span> DatasourceControlSourceRef

Source information for controls (e.g. variables or links)

## Definition

```go
type DatasourceControlSourceRef struct {
    Type string `json:"type"`
    // The plugin type-id
    Group string `json:"group"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DatasourceControlSourceRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (datasourceControlSourceRef *DatasourceControlSourceRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DatasourceControlSourceRef` objects.

```go
func (datasourceControlSourceRef *DatasourceControlSourceRef) Equals(other DatasourceControlSourceRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DatasourceControlSourceRef` fields for violations and returns them.

```go
func (datasourceControlSourceRef *DatasourceControlSourceRef) Validate() error
```

## See also

 * <span class="badge builder"></span> [DatasourceControlSourceRefBuilder](./builder-DatasourceControlSourceRefBuilder.md)
