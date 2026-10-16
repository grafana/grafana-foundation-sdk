---
title: <span class="badge object-type-struct"></span> CellSpec
---
# <span class="badge object-type-struct"></span> CellSpec

## Definition

```go
type CellSpec struct {
    Content dashboardv2beta1.CellContentKind `json:"content"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CellSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cellSpec *CellSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CellSpec` objects.

```go
func (cellSpec *CellSpec) Equals(other CellSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CellSpec` fields for violations and returns them.

```go
func (cellSpec *CellSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [CellSpecBuilder](./builder-CellSpecBuilder.md)
