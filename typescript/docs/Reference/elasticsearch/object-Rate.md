---
title: <span class="badge object-type-interface"></span> Rate
---
# <span class="badge object-type-interface"></span> Rate

## Definition

```typescript
export interface Rate {
	type: unknown;
	field?: string;
	id: string;
	settings?: {
		unit?: string;
		mode?: string;
	};
	hide?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [RateBuilder](./builder-RateBuilder.md)
