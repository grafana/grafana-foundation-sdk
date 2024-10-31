---
title: <span class="badge object-type-interface"></span> UniqueCount
---
# <span class="badge object-type-interface"></span> UniqueCount

## Definition

```typescript
export interface UniqueCount {
	type: "cardinality";
	field?: string;
	id: string;
	settings?: {
		precision_threshold?: string;
		missing?: string;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [UniqueCountBuilder](./builder-UniqueCountBuilder.md)
