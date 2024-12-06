---
title: <span class="badge object-type-interface"></span> Percentiles
---
# <span class="badge object-type-interface"></span> Percentiles

## Definition

```typescript
export interface Percentiles {
	type: "percentiles";
	field?: string;
	id: string;
	settings?: {
		script?: elasticsearch.InlineScript;
		missing?: string;
		percents?: string[];
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [PercentilesBuilder](./builder-PercentilesBuilder.md)
