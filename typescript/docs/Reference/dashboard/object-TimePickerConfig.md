---
title: <span class="badge object-type-interface"></span> TimePickerConfig
---
# <span class="badge object-type-interface"></span> TimePickerConfig

Time picker configuration

It defines the default config for the time picker and the refresh picker for the specific dashboard.

## Definition

```typescript
export interface TimePickerConfig {
	// Whether timepicker is visible or not.
	hidden: boolean;
	// Interval options available in the refresh picker dropdown.
	refresh_intervals: string[];
	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	time_options: string[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TimePickerBuilder](./builder-TimePickerBuilder.md)
