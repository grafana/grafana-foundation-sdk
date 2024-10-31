---
title: <span class="badge object-type-interface"></span> AzureMetricDimension
---
# <span class="badge object-type-interface"></span> AzureMetricDimension

## Definition

```typescript
export interface AzureMetricDimension {
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
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AzureMetricDimensionBuilder](./builder-AzureMetricDimensionBuilder.md)
