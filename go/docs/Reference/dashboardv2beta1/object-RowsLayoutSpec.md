---
title: <span class="badge object-type-struct"></span> RowsLayoutSpec
---
# <span class="badge object-type-struct"></span> RowsLayoutSpec

## Definition

```go
type RowsLayoutSpec struct {
    Rows []dashboardv2beta1.RowsLayoutRowKind `json:"rows"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsLayoutSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rowsLayoutSpec *RowsLayoutSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RowsLayoutSpec` objects.

```go
func (rowsLayoutSpec *RowsLayoutSpec) Equals(other RowsLayoutSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RowsLayoutSpec` fields for violations and returns them.

```go
func (rowsLayoutSpec *RowsLayoutSpec) Validate() error
```

