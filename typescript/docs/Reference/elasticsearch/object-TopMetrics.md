---
title: <span class="badge object-type-interface"></span> TopMetrics
---
# <span class="badge object-type-interface"></span> TopMetrics

## Definition

```typescript
export interface TopMetrics {
	type: "top_metrics";
	id: string;
	settings?: {
		order?: string;
		orderBy?: string;
		metrics?: string[];
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TopMetricsBuilder](./builder-TopMetricsBuilder.md)
