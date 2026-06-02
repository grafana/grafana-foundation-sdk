---
title: <span class="badge object-type-interface"></span> MetricDimension
---
# <span class="badge object-type-interface"></span> MetricDimension

## Definition

```typescript
export interface MetricDimension {
	// Name of Dimension to be filtered on.
	dimension?: string;
	// String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
	operator?: string;
	// Values to match with the filter.
	filters?: string[];
	// @deprecated filter is deprecated in favour of filters to support multiselect.
	filter?: string;
}

```
## See also

 * <span class="badge builder"></span> [MetricDimensionBuilder](./builder-MetricDimensionBuilder.md)
