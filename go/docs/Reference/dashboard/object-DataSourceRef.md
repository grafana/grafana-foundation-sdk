---
title: <span class="badge object-type-struct"></span> DataSourceRef
---
# <span class="badge object-type-struct"></span> DataSourceRef

Ref to a DataSource instance

## Definition

```go
type DataSourceRef struct {
    // The plugin type-id
    Type *string `json:"type,omitempty"`
    // Specific datasource instance
    Uid *string `json:"uid,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DataSourceRef` from JSON.

```go
func (dataSourceRef *DataSourceRef) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataSourceRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dataSourceRef *DataSourceRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DataSourceRef` objects.

```go
func (dataSourceRef *DataSourceRef) Equals(other DataSourceRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DataSourceRef` fields for violations and returns them.

```go
func (dataSourceRef *DataSourceRef) Validate() error
```

