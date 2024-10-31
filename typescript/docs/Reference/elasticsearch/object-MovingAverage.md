---
title: <span class="badge object-type-interface"></span> MovingAverage
---
# <span class="badge object-type-interface"></span> MovingAverage

#MovingAverage's settings are overridden in types.ts

## Definition

```typescript
export interface MovingAverage {
	pipelineAgg?: string;
	field?: string;
	type: "moving_avg";
	id: string;
	settings?: Record<string, any>;
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MovingAverageBuilder](./builder-MovingAverageBuilder.md)
