---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Size of each bucket
	bucketSize?: number;
	// Offset buckets by this amount
	bucketOffset?: number;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	// Combines multiple series into a single histogram
	combine?: boolean;
}

```
## Methods

No methods.
