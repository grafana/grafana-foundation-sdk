---
title: <span class="badge object-type-interface"></span> SerialDiff
---
# <span class="badge object-type-interface"></span> SerialDiff

## Definition

```typescript
export interface SerialDiff {
	pipelineAgg?: string;
	field?: string;
	type: "serial_diff";
	id: string;
	settings?: {
		lag?: string;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [SerialDiffBuilder](./builder-SerialDiffBuilder.md)
