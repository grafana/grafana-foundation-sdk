---
title: <span class="badge object-type-struct"></span> TimeInterval
---
# <span class="badge object-type-struct"></span> TimeInterval

## Definition

```go
type TimeInterval struct {
    Times []alerting.TimeRange `json:"times,omitempty"`
    Weekdays []alerting.WeekdayRange `json:"weekdays,omitempty"`
    DaysOfMonth []alerting.DayOfMonthRange `json:"days_of_month,omitempty"`
    Months []alerting.MonthRange `json:"months,omitempty"`
    Years []alerting.YearRange `json:"years,omitempty"`
    Location *alerting.Location `json:"location,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeInterval` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeInterval *TimeInterval) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeInterval` objects.

```go
func (timeInterval *TimeInterval) Equals(other TimeInterval) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeInterval` fields for violations and returns them.

```go
func (timeInterval *TimeInterval) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeIntervalBuilder](./builder-TimeIntervalBuilder.md)
