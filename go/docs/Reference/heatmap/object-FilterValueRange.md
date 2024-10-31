---
title: <span class="badge object-type-struct"></span> FilterValueRange
---
# <span class="badge object-type-struct"></span> FilterValueRange

Controls the value filter range

## Definition

```go
type FilterValueRange struct {
    // Sets the filter range to values less than or equal to the given value
    Le *float32 `json:"le,omitempty"`
    // Sets the filter range to values greater than or equal to the given value
    Ge *float32 `json:"ge,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FilterValueRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (filterValueRange *FilterValueRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FilterValueRange` objects.

```go
func (filterValueRange *FilterValueRange) Equals(other FilterValueRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FilterValueRange` fields for violations and returns them.

```go
func (filterValueRange *FilterValueRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [FilterValueRangeBuilder](./builder-FilterValueRangeBuilder.md)
