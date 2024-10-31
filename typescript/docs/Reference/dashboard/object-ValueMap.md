---
title: <span class="badge object-type-interface"></span> ValueMap
---
# <span class="badge object-type-interface"></span> ValueMap

Maps text values to a color or different display text and color.

For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.

## Definition

```typescript
export interface ValueMap {
	type: "value";
	// Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
	options: Record<string, dashboard.ValueMappingResult>;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ValueMapBuilder](./builder-ValueMapBuilder.md)
