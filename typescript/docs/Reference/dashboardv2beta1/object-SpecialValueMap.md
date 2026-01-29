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
	type: dashboardv2beta1.MappingType.Special;
	options: {
		// Special value to match against
		match: dashboardv2beta1.SpecialValueMatch;
		// Config to apply when the value matches the special value
		result: dashboardv2beta1.ValueMappingResult;
	};
}

```
