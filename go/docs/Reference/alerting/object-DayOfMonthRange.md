---
title: <span class="badge object-type-struct"></span> DayOfMonthRange
---
# <span class="badge object-type-struct"></span> DayOfMonthRange

## Definition

```go
type DayOfMonthRange struct {
    Begin *int32 `json:"begin,omitempty"`
    End *int32 `json:"end,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DayOfMonthRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dayOfMonthRange *DayOfMonthRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DayOfMonthRange` objects.

```go
func (dayOfMonthRange *DayOfMonthRange) Equals(other DayOfMonthRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DayOfMonthRange` fields for violations and returns them.

```go
func (dayOfMonthRange *DayOfMonthRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [DayOfMonthRangeBuilder](./builder-DayOfMonthRangeBuilder.md)
