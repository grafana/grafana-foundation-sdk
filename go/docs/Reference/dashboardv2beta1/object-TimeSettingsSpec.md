---
title: <span class="badge object-type-struct"></span> TimeSettingsSpec
---
# <span class="badge object-type-struct"></span> TimeSettingsSpec

Time configuration

It defines the default time config for the time picker, the refresh picker for the specific dashboard.

## Definition

```go
type TimeSettingsSpec struct {
    // Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    Timezone *string `json:"timezone,omitempty"`
    // Start time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    From string `json:"from"`
    // End time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    To string `json:"to"`
    // Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    // v1: refresh
    AutoRefresh string `json:"autoRefresh"`
    // Interval options available in the refresh picker dropdown.
    // v1: timepicker.refresh_intervals
    AutoRefreshIntervals []string `json:"autoRefreshIntervals"`
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    // v1: timepicker.quick_ranges , not exposed in the UI
    QuickRanges []dashboardv2beta1.TimeRangeOption `json:"quickRanges,omitempty"`
    // Whether timepicker is visible or not.
    // v1: timepicker.hidden
    HideTimepicker bool `json:"hideTimepicker"`
    // Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    WeekStart *dashboardv2beta1.TimeSettingsSpecWeekStart `json:"weekStart,omitempty"`
    // The month that the fiscal year starts on. 0 = January, 11 = December
    FiscalYearStartMonth int64 `json:"fiscalYearStartMonth"`
    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    // v1: timepicker.nowDelay
    NowDelay *string `json:"nowDelay,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeSettingsSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeSettingsSpec *TimeSettingsSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeSettingsSpec` objects.

```go
func (timeSettingsSpec *TimeSettingsSpec) Equals(other TimeSettingsSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeSettingsSpec` fields for violations and returns them.

```go
func (timeSettingsSpec *TimeSettingsSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeSettingsBuilder](./builder-TimeSettingsBuilder.md)
