---
title: <span class="badge object-type-class"></span> TimeSettingsSpec
---
# <span class="badge object-type-class"></span> TimeSettingsSpec

Time configuration

It defines the default time config for the time picker, the refresh picker for the specific dashboard.

## Definition

```php
class TimeSettingsSpec implements \JsonSerializable
{
    /**
     * Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
     */
    public ?string $timezone;

    /**
     * Start time range for dashboard.
     * Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
     */
    public string $from;

    /**
     * End time range for dashboard.
     * Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
     */
    public string $to;

    /**
     * Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
     * v1: refresh
     */
    public string $autoRefresh;

    /**
     * Interval options available in the refresh picker dropdown.
     * v1: timepicker.refresh_intervals
     * @var array<string>
     */
    public array $autoRefreshIntervals;

    /**
     * Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
     * v1: timepicker.quick_ranges , not exposed in the UI
     * @var array<\Grafana\Foundation\Dashboardv2beta1\TimeRangeOption>|null
     */
    public ?array $quickRanges;

    /**
     * Whether timepicker is visible or not.
     * v1: timepicker.hidden
     */
    public bool $hideTimepicker;

    /**
     * Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart $weekStart;

    /**
     * The month that the fiscal year starts on. 0 = January, 11 = December
     */
    public int $fiscalYearStartMonth;

    /**
     * Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
     * v1: timepicker.nowDelay
     */
    public ?string $nowDelay;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [TimeSettingsBuilder](./builder-TimeSettingsBuilder.md)
