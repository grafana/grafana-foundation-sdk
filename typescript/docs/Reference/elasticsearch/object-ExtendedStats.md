---
title: <span class="badge object-type-interface"></span> ExtendedStats
---
# <span class="badge object-type-interface"></span> ExtendedStats

## Definition

```typescript
export interface ExtendedStats {
	type: unknown;
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
## See also

 * <span class="badge builder"></span> [ExtendedStatsBuilder](./builder-ExtendedStatsBuilder.md)
