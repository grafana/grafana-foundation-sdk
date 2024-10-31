---
title: <span class="badge object-type-struct"></span> CellValues
---
# <span class="badge object-type-struct"></span> CellValues

Controls cell value options

## Definition

```go
type CellValues struct {
    // Controls the cell value unit
    Unit *string `json:"unit,omitempty"`
    // Controls the number of decimals for cell values
    Decimals *float32 `json:"decimals,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CellValues` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cellValues *CellValues) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CellValues` objects.

```go
func (cellValues *CellValues) Equals(other CellValues) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CellValues` fields for violations and returns them.

```go
func (cellValues *CellValues) Validate() error
```

## See also

 * <span class="badge builder"></span> [CellValuesBuilder](./builder-CellValuesBuilder.md)
