---
title: <span class="badge builder"></span> TimeSettingsBuilder
---
# <span class="badge builder"></span> TimeSettingsBuilder

## Constructor

```typescript
new TimeSettingsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> autoRefresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

v1: refresh

```typescript
autoRefresh(autoRefresh: string)
```

### <span class="badge object-method"></span> autoRefreshIntervals

Interval options available in the refresh picker dropdown.

v1: timepicker.refresh_intervals

```typescript
autoRefreshIntervals(autoRefreshIntervals: string[])
```

### <span class="badge object-method"></span> fiscalYearStartMonth

The month that the fiscal year starts on. 0 = January, 11 = December

```typescript
fiscalYearStartMonth(fiscalYearStartMonth: number)
```

### <span class="badge object-method"></span> from

Start time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```typescript
from(from: string)
```

### <span class="badge object-method"></span> hideTimepicker

Whether timepicker is visible or not.

v1: timepicker.hidden

```typescript
hideTimepicker(hideTimepicker: boolean)
```

### <span class="badge object-method"></span> nowDelay

Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.

v1: timepicker.nowDelay

```typescript
nowDelay(nowDelay: string)
```

### <span class="badge object-method"></span> quickRanges

Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.

v1: timepicker.quick_ranges , not exposed in the UI

```typescript
quickRanges(quickRanges: cog.Builder<dashboardv2beta1.TimeRangeOption>[])
```

### <span class="badge object-method"></span> timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```typescript
timezone(timezone: string)
```

### <span class="badge object-method"></span> to

End time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```typescript
to(to: string)
```

### <span class="badge object-method"></span> weekStart

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```typescript
weekStart(weekStart: "saturday" | "monday" | "sunday")
```

## See also

 * <span class="badge object-type-interface"></span> [TimeSettingsSpec](./object-TimeSettingsSpec.md)
