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
	type: unknown;
	id: string;
	settings?: Record<string, any>;
	hide?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [MovingAverageBuilder](./builder-MovingAverageBuilder.md)
