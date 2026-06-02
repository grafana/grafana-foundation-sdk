---
title: <span class="badge object-type-interface"></span> TracesFilter
---
# <span class="badge object-type-interface"></span> TracesFilter

## Definition

```typescript
export interface TracesFilter {
	// Property name, auto-populated based on available traces.
	property: string;
	// Comparison operator to use. Either equals or not equals.
	operation: string;
	// Values to filter by.
	filters: string[];
}

```
## See also

 * <span class="badge builder"></span> [TracesFilterBuilder](./builder-TracesFilterBuilder.md)
