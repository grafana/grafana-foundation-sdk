---
title: <span class="badge object-type-interface"></span> ValueMappingResult
---
# <span class="badge object-type-interface"></span> ValueMappingResult

Result used as replacement with text and color when the value matches

## Definition

```typescript
export interface ValueMappingResult {
	// Text to display when the value matches
	text?: string;
	// Text to use when the value matches
	color?: string;
	// Icon to display when the value matches. Only specific visualizations.
	icon?: string;
	// Position in the mapping array. Only used internally.
	index?: number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ValueMappingResultBuilder](./builder-ValueMappingResultBuilder.md)
