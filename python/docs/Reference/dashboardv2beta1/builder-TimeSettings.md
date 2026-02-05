---
title: <span class="badge builder"></span> TimeSettings
---
# <span class="badge builder"></span> TimeSettings

## Constructor

```python
TimeSettings()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.TimeSettingsSpec
```

### <span class="badge object-method"></span> auto_refresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

v1: refresh

```python
def auto_refresh(auto_refresh: str) -> typing.Self
```

### <span class="badge object-method"></span> auto_refresh_intervals

Interval options available in the refresh picker dropdown.

v1: timepicker.refresh_intervals

```python
def auto_refresh_intervals(auto_refresh_intervals: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> fiscal_year_start_month

The month that the fiscal year starts on. 0 = January, 11 = December

```python
def fiscal_year_start_month(fiscal_year_start_month: int) -> typing.Self
```

### <span class="badge object-method"></span> from_val

Start time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```python
def from_val(from_val: str) -> typing.Self
```

### <span class="badge object-method"></span> hide_timepicker

Whether timepicker is visible or not.

v1: timepicker.hidden

```python
def hide_timepicker(hide_timepicker: bool) -> typing.Self
```

### <span class="badge object-method"></span> now_delay

Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.

v1: timepicker.nowDelay

```python
def now_delay(now_delay: str) -> typing.Self
```

### <span class="badge object-method"></span> quick_ranges

Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.

v1: timepicker.quick_ranges , not exposed in the UI

```python
def quick_ranges(quick_ranges: list[cogbuilder.Builder[dashboardv2beta1.TimeRangeOption]]) -> typing.Self
```

### <span class="badge object-method"></span> timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```python
def timezone(timezone: str) -> typing.Self
```

### <span class="badge object-method"></span> to

End time range for dashboard.

Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".

```python
def to(to: str) -> typing.Self
```

### <span class="badge object-method"></span> week_start

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```python
def week_start(week_start: typing.Literal["saturday", "monday", "sunday"]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TimeSettingsSpec](./object-TimeSettingsSpec.md)
