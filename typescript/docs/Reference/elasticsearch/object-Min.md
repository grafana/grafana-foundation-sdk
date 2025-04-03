---
title: <span class="badge object-type-interface"></span> Min
---
# <span class="badge object-type-interface"></span> Min

## Definition

```typescript
export interface Min {
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

 * <span class="badge builder"></span> [MinBuilder](./builder-MinBuilder.md)
