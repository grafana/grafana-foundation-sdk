---
title: <span class="badge object-type-struct"></span> RowsLayoutKind
---
# <span class="badge object-type-struct"></span> RowsLayoutKind

## Definition

```go
type RowsLayoutKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.RowsLayoutSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsLayoutKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rowsLayoutKind *RowsLayoutKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RowsLayoutKind` objects.

```go
func (rowsLayoutKind *RowsLayoutKind) Equals(other RowsLayoutKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RowsLayoutKind` fields for violations and returns them.

```go
func (rowsLayoutKind *RowsLayoutKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [RowsBuilder](./builder-RowsBuilder.md)
