---
title: <span class="badge object-type-interface"></span> Derivative
---
# <span class="badge object-type-interface"></span> Derivative

## Definition

```typescript
export interface Derivative {
	pipelineAgg?: string;
	field?: string;
	type: "derivative";
	id: string;
	settings?: {
		unit?: string;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [DerivativeBuilder](./builder-DerivativeBuilder.md)
