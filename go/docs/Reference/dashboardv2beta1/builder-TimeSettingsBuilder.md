---
title: <span class="badge builder"></span> TimeSettingsBuilder
---
# <span class="badge builder"></span> TimeSettingsBuilder

## Constructor

```go
func NewTimeSettingsBuilder() *TimeSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TimeSettingsBuilder) Build() (TimeSettingsSpec, error)
```

### <span class="badge object-method"></span> AutoRefresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

v1: refresh

```go
func (builder *TimeSettingsBuilder) AutoRefresh(autoRefresh string) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> AutoRefreshIntervals

Interval options available in the refresh picker dropdown.

v1: timepicker.refresh_intervals

```go
func (builder *TimeSettingsBuilder) AutoRefreshIntervals(autoRefreshIntervals []string) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> FiscalYearStartMonth

The month that the fiscal year starts on. 0 = January, 11 = December

```go
func (builder *TimeSettingsBuilder) FiscalYearStartMonth(fiscalYearStartMonth int64) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> From

Start time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```go
func (builder *TimeSettingsBuilder) From(from string) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> HideTimepicker

Whether timepicker is visible or not.

v1: timepicker.hidden

```go
func (builder *TimeSettingsBuilder) HideTimepicker(hideTimepicker bool) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> NowDelay

Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.

v1: timepicker.nowDelay

```go
func (builder *TimeSettingsBuilder) NowDelay(nowDelay string) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> QuickRanges

Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.

v1: timepicker.quick_ranges , not exposed in the UI

```go
func (builder *TimeSettingsBuilder) QuickRanges(quickRanges []cog.Builder[dashboardv2beta1.TimeRangeOption]) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> Timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```go
func (builder *TimeSettingsBuilder) Timezone(timezone string) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> To

End time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```go
func (builder *TimeSettingsBuilder) To(to string) *TimeSettingsBuilder
```

### <span class="badge object-method"></span> WeekStart

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```go
func (builder *TimeSettingsBuilder) WeekStart(weekStart dashboardv2beta1.TimeSettingsSpecWeekStart) *TimeSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeSettingsSpec](./object-TimeSettingsSpec.md)
