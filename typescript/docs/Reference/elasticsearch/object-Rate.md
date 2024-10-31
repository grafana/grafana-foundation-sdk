---
title: <span class="badge object-type-interface"></span> Rate
---
# <span class="badge object-type-interface"></span> Rate

## Definition

```typescript
export interface Rate {
	type: "rate";
	field?: string;
	id: string;
	settings?: {
		unit?: string;
		mode?: string;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RateBuilder](./builder-RateBuilder.md)
