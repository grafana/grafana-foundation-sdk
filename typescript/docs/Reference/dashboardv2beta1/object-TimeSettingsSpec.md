---
title: <span class="badge object-type-interface"></span> TimeSettingsSpec
---
# <span class="badge object-type-interface"></span> TimeSettingsSpec

Time configuration

It defines the default time config for the time picker, the refresh picker for the specific dashboard.

## Definition

```typescript
export interface TimeSettingsSpec {
	// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
	timezone?: string;
	// Start time range for dashboard.
	// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
	from: string;
	// End time range for dashboard.
	// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
	to: string;
	// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
	// v1: refresh
	autoRefresh: string;
	// Interval options available in the refresh picker dropdown.
	// v1: timepicker.refresh_intervals
	autoRefreshIntervals: string[];
	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	// v1: timepicker.quick_ranges , not exposed in the UI
	quickRanges?: dashboardv2beta1.TimeRangeOption[];
	// Whether timepicker is visible or not.
	// v1: timepicker.hidden
	hideTimepicker: boolean;
	// Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
	weekStart?: "saturday" | "monday" | "sunday";
	// The month that the fiscal year starts on. 0 = January, 11 = December
	fiscalYearStartMonth: number;
	// Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
	// v1: timepicker.nowDelay
	nowDelay?: string;
}

```
## See also

 * <span class="badge builder"></span> [TimeSettingsBuilder](./builder-TimeSettingsBuilder.md)
