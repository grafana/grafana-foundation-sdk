---
title: <span class="badge object-type-interface"></span> AzureTracesFilter
---
# <span class="badge object-type-interface"></span> AzureTracesFilter

## Definition

```typescript
export interface AzureTracesFilter {
	// Property name, auto-populated based on available traces.
	property: string;
	// Comparison operator to use. Either equals or not equals.
	operation: string;
	// Values to filter by.
	filters: string[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AzureTracesFilterBuilder](./builder-AzureTracesFilterBuilder.md)
