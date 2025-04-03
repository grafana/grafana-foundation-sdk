---
title: <span class="badge object-type-interface"></span> TimeInterval
---
# <span class="badge object-type-interface"></span> TimeInterval

## Definition

```typescript
export interface TimeInterval {
	times?: alerting.TimeRange[];
	weekdays?: alerting.WeekdayRange[];
	days_of_month?: alerting.DayOfMonthRange[];
	months?: alerting.MonthRange[];
	years?: alerting.YearRange[];
	location?: alerting.Location;
}

```
## See also

 * <span class="badge builder"></span> [TimeIntervalBuilder](./builder-TimeIntervalBuilder.md)
