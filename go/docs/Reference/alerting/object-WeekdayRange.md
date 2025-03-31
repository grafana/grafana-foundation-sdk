---
title: <span class="badge object-type-struct"></span> WeekdayRange
---
# <span class="badge object-type-struct"></span> WeekdayRange

## Definition

```go
type WeekdayRange struct {
    Begin *int32 `json:"begin,omitempty"`
    End *int32 `json:"end,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `WeekdayRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (weekdayRange *WeekdayRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `WeekdayRange` objects.

```go
func (weekdayRange *WeekdayRange) Equals(other WeekdayRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `WeekdayRange` fields for violations and returns them.

```go
func (weekdayRange *WeekdayRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [WeekdayRangeBuilder](./builder-WeekdayRangeBuilder.md)
