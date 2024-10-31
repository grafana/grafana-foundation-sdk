---
title: <span class="badge object-type-interface"></span> SpecialValueMap
---
# <span class="badge object-type-interface"></span> SpecialValueMap

Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.

See SpecialValueMatch to see the list of special values.

For example, you can configure a special value mapping so that null values appear as N/A.

## Definition

```typescript
export interface SpecialValueMap {
	type: "special";
	options: {
		// Special value to match against
		match: dashboard.SpecialValueMatch;
		// Config to apply when the value matches the special value
		result: dashboard.ValueMappingResult;
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [SpecialValueMapBuilder](./builder-SpecialValueMapBuilder.md)
