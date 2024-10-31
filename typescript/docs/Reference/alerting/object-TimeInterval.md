---
title: <span class="badge object-type-interface"></span> TimeInterval
---
# <span class="badge object-type-interface"></span> TimeInterval

TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained

within the interval.

## Definition

```typescript
export interface TimeInterval {
	days_of_month?: string[];
	location?: string;
	months?: string[];
	times?: alerting.TimeRange[];
	weekdays?: string[];
	years?: string[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TimeIntervalBuilder](./builder-TimeIntervalBuilder.md)
