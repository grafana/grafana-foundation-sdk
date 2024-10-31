---
title: <span class="badge object-type-struct"></span> TimeIntervalTimeRange
---
# <span class="badge object-type-struct"></span> TimeIntervalTimeRange

## Definition

```go
type TimeIntervalTimeRange struct {
    EndTime *string `json:"end_time,omitempty"`
    StartTime *string `json:"start_time,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeIntervalTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeIntervalTimeRange *TimeIntervalTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeIntervalTimeRange` objects.

```go
func (timeIntervalTimeRange *TimeIntervalTimeRange) Equals(other TimeIntervalTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeIntervalTimeRange` fields for violations and returns them.

```go
func (timeIntervalTimeRange *TimeIntervalTimeRange) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeIntervalTimeRangeBuilder](./builder-TimeIntervalTimeRangeBuilder.md)
