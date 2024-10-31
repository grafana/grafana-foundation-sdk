---
title: <span class="badge object-type-interface"></span> Sum
---
# <span class="badge object-type-interface"></span> Sum

## Definition

```typescript
export interface Sum {
	type: "sum";
	field?: string;
	id: string;
	settings?: {
		script?: elasticsearch.InlineScript;
		missing?: string;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [SumBuilder](./builder-SumBuilder.md)
