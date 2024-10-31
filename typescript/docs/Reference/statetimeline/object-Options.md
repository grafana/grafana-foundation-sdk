---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Show timeline values on chart
	showValue: common.VisibilityMode;
	// Controls the row height
	rowHeight: number;
	// Merge equal consecutive values
	mergeValues?: boolean;
	// Controls value alignment on the timelines
	alignValue?: common.TimelineValueAlignment;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	timezone?: common.TimeZone[];
	// Enables pagination when > 0
	perPage?: number;
}

```
## Methods

No methods.
