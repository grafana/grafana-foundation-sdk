---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Comma-separated list of values used to filter alert results
	labels: string;
	// Name of the alertmanager used as a source for alerts
	alertmanager: string;
	// Expand all alert groups by default
	expandAll: boolean;
}

```
## See also

 * <span class="badge builder"></span> [OptionsBuilder](./builder-OptionsBuilder.md)
