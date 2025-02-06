---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Set the height of the rows
	rowHeight: number;
	// Show values on the columns
	showValue: common.VisibilityMode;
	// Controls the column width
	colWidth?: number;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	timezone?: common.TimeZone[];
	// Enables pagination when > 0
	perPage?: number;
}

```
## Methods

No methods.
