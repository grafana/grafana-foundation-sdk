---
title: <span class="badge object-type-interface"></span> Max
---
# <span class="badge object-type-interface"></span> Max

## Definition

```typescript
export interface Max {
	type: "max";
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

 * <span class="badge builder"></span> [MaxBuilder](./builder-MaxBuilder.md)
