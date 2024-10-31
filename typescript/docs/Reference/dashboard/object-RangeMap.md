---
title: <span class="badge object-type-interface"></span> RangeMap
---
# <span class="badge object-type-interface"></span> RangeMap

Maps numerical ranges to a display text and color.

For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.

## Definition

```typescript
export interface RangeMap {
	type: "range";
	// Range to match against and the result to apply when the value is within the range
	options: {
		// Min value of the range. It can be null which means -Infinity
		from: number | null;
		// Max value of the range. It can be null which means +Infinity
		to: number | null;
		// Config to apply when the value is within the range
		result: dashboard.ValueMappingResult;
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RangeMapBuilder](./builder-RangeMapBuilder.md)
