---
title: <span class="badge object-type-struct"></span> TimeIntervalItem
---
# <span class="badge object-type-struct"></span> TimeIntervalItem

## Definition

```go
type TimeIntervalItem struct {
    DaysOfMonth []string `json:"days_of_month,omitempty"`
    Location *string `json:"location,omitempty"`
    Months []string `json:"months,omitempty"`
    Times []alerting.TimeIntervalTimeRange `json:"times,omitempty"`
    Weekdays []string `json:"weekdays,omitempty"`
    Years []string `json:"years,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeIntervalItem` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeIntervalItem *TimeIntervalItem) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeIntervalItem` objects.

```go
func (timeIntervalItem *TimeIntervalItem) Equals(other TimeIntervalItem) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeIntervalItem` fields for violations and returns them.

```go
func (timeIntervalItem *TimeIntervalItem) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeIntervalItemBuilder](./builder-TimeIntervalItemBuilder.md)
