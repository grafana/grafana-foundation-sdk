---
title: <span class="badge object-type-interface"></span> TimePicker
---
# <span class="badge object-type-interface"></span> TimePicker

## Definition

```typescript
export interface TimePicker {
	// Whether timepicker is visible or not.
	hidden: boolean;
	// Interval options available in the refresh picker dropdown.
	refresh_intervals: string[];
	// Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
	collapse: boolean;
	// Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
	enable: boolean;
	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	time_options: string[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TimePickerBuilder](./builder-TimePickerBuilder.md)
