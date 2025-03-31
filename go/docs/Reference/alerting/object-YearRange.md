---
title: <span class="badge object-type-struct"></span> YearRange
---
# <span class="badge object-type-struct"></span> YearRange

## Definition

```go
type YearRange struct {
    Begin *int32 `json:"begin,omitempty"`
    End *int32 `json:"end,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `YearRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (yearRange *YearRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `YearRange` objects.

```go
func (yearRange *YearRange) Equals(other YearRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `YearRange` fields for violations and returns them.

```go
func (yearRange *YearRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [YearRangeBuilder](./builder-YearRangeBuilder.md)
