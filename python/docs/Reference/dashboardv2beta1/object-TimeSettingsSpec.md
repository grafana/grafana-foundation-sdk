---
title: <span class="badge object-type-class"></span> TimeSettingsSpec
---
# <span class="badge object-type-class"></span> TimeSettingsSpec

Time configuration

It defines the default time config for the time picker, the refresh picker for the specific dashboard.

## Definition

```python
class TimeSettingsSpec:
    """
    Time configuration
    It defines the default time config for the time picker, the refresh picker for the specific dashboard.
    """

    # Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    timezone: typing.Optional[str]
    # Start time range for dashboard.
    # Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    from_val: str
    # End time range for dashboard.
    # Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    to: str
    # Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    # v1: refresh
    auto_refresh: str
    # Interval options available in the refresh picker dropdown.
    # v1: timepicker.refresh_intervals
    auto_refresh_intervals: list[str]
    # Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    # v1: timepicker.quick_ranges , not exposed in the UI
    quick_ranges: typing.Optional[list[dashboardv2beta1.TimeRangeOption]]
    # Whether timepicker is visible or not.
    # v1: timepicker.hidden
    hide_timepicker: bool
    # Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    week_start: typing.Optional[typing.Literal["saturday", "monday", "sunday"]]
    # The month that the fiscal year starts on. 0 = January, 11 = December
    fiscal_year_start_month: int
    # Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    # v1: timepicker.nowDelay
    now_delay: typing.Optional[str]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [TimeSettings](./builder-TimeSettings.md)
