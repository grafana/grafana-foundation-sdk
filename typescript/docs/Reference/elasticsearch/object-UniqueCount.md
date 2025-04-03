---
title: <span class="badge object-type-interface"></span> UniqueCount
---
# <span class="badge object-type-interface"></span> UniqueCount

## Definition

```typescript
export interface UniqueCount {
	type: unknown;
	field?: string;
	id: string;
	settings?: {
		precision_threshold?: string;
		missing?: string;
	};
	hide?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [UniqueCountBuilder](./builder-UniqueCountBuilder.md)
