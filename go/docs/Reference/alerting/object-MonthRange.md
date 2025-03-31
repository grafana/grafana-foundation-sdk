---
title: <span class="badge object-type-struct"></span> MonthRange
---
# <span class="badge object-type-struct"></span> MonthRange

## Definition

```go
type MonthRange struct {
    Begin *int32 `json:"begin,omitempty"`
    End *int32 `json:"end,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MonthRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (monthRange *MonthRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MonthRange` objects.

```go
func (monthRange *MonthRange) Equals(other MonthRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MonthRange` fields for violations and returns them.

```go
func (monthRange *MonthRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [MonthRangeBuilder](./builder-MonthRangeBuilder.md)
