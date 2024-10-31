---
title: <span class="badge object-type-interface"></span> CumulativeSum
---
# <span class="badge object-type-interface"></span> CumulativeSum

## Definition

```typescript
export interface CumulativeSum {
	pipelineAgg?: string;
	field?: string;
	type: "cumulative_sum";
	id: string;
	settings?: {
		format?: string;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [CumulativeSumBuilder](./builder-CumulativeSumBuilder.md)
