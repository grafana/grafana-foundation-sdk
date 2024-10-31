---
title: <span class="badge object-type-interface"></span> ExtendedStats
---
# <span class="badge object-type-interface"></span> ExtendedStats

## Definition

```typescript
export interface ExtendedStats {
	type: "extended_stats";
	settings?: {
		script?: elasticsearch.InlineScript;
		missing?: string;
		sigma?: string;
	};
	field?: string;
	id: string;
	meta?: any;
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ExtendedStatsBuilder](./builder-ExtendedStatsBuilder.md)
