---
title: <span class="badge object-type-struct"></span> TimeRange
---
# <span class="badge object-type-struct"></span> TimeRange

## Definition

```go
type TimeRange struct {
    // From is the start time of the query.
    From string `json:"from"`
    // To is the end time of the query.
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeRange *TimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeRange` objects.

```go
func (timeRange *TimeRange) Equals(other TimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeRange` fields for violations and returns them.

```go
func (timeRange *TimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeRangeBuilder](./builder-TimeRangeBuilder.md)
