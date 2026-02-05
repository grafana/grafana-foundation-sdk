---
title: <span class="badge builder"></span> TimeSettingsBuilder
---
# <span class="badge builder"></span> TimeSettingsBuilder

## Constructor

```php
new TimeSettingsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> autoRefresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

v1: refresh

```php
autoRefresh(string $autoRefresh)
```

### <span class="badge object-method"></span> autoRefreshIntervals

Interval options available in the refresh picker dropdown.

v1: timepicker.refresh_intervals

@param array<string> $autoRefreshIntervals

```php
autoRefreshIntervals(array $autoRefreshIntervals)
```

### <span class="badge object-method"></span> fiscalYearStartMonth

The month that the fiscal year starts on. 0 = January, 11 = December

```php
fiscalYearStartMonth(int $fiscalYearStartMonth)
```

### <span class="badge object-method"></span> from

Start time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```php
from(string $from)
```

### <span class="badge object-method"></span> hideTimepicker

Whether timepicker is visible or not.

v1: timepicker.hidden

```php
hideTimepicker(bool $hideTimepicker)
```

### <span class="badge object-method"></span> nowDelay

Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.

v1: timepicker.nowDelay

```php
nowDelay(string $nowDelay)
```

### <span class="badge object-method"></span> quickRanges

Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.

v1: timepicker.quick_ranges , not exposed in the UI

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeRangeOption>> $quickRanges

```php
quickRanges(array $quickRanges)
```

### <span class="badge object-method"></span> timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```php
timezone(string $timezone)
```

### <span class="badge object-method"></span> to

End time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```php
to(string $to)
```

### <span class="badge object-method"></span> weekStart

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```php
weekStart(\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart $weekStart)
```

## See also

 * <span class="badge object-type-class"></span> [TimeSettingsSpec](./object-TimeSettingsSpec.md)
