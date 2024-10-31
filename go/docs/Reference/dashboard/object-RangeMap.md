---
title: <span class="badge object-type-struct"></span> RangeMap
---
# <span class="badge object-type-struct"></span> RangeMap

Maps numerical ranges to a display text and color.

For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.

## Definition

```go
type RangeMap struct {
    Type string `json:"type"`
    // Range to match against and the result to apply when the value is within the range
    Options dashboard.DashboardRangeMapOptions `json:"options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RangeMap` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rangeMap *RangeMap) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RangeMap` objects.

```go
func (rangeMap *RangeMap) Equals(other RangeMap) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RangeMap` fields for violations and returns them.

```go
func (rangeMap *RangeMap) Validate() error
```

## See also

 * <span class="badge builder"></span> [RangeMapBuilder](./builder-RangeMapBuilder.md)
