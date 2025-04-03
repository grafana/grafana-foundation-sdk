---
title: <span class="badge object-type-interface"></span> Average
---
# <span class="badge object-type-interface"></span> Average

## Definition

```typescript
export interface Average {
	type: unknown;
	field?: string;
	id: string;
	settings?: {
		script?: elasticsearch.InlineScript;
		missing?: string;
	};
	hide?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [AverageBuilder](./builder-AverageBuilder.md)
