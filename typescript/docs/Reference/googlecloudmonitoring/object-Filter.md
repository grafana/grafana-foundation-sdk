---
title: <span class="badge object-type-interface"></span> Filter
---
# <span class="badge object-type-interface"></span> Filter

Query filter representation.

## Definition

```typescript
export interface Filter {
	// Filter key.
	key: string;
	// Filter operator.
	operator: string;
	// Filter value.
	value: string;
	// Filter condition.
	condition?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [FilterBuilder](./builder-FilterBuilder.md)
