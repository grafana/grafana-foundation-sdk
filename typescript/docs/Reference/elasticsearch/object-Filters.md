---
title: <span class="badge object-type-interface"></span> Filters
---
# <span class="badge object-type-interface"></span> Filters

## Definition

```typescript
export interface Filters {
	id: string;
	type: "filters";
	settings?: {
		filters?: elasticsearch.Filter[];
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [FiltersBuilder](./builder-FiltersBuilder.md)
