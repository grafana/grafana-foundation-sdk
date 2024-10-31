---
title: <span class="badge object-type-interface"></span> YAxisConfig
---
# <span class="badge object-type-interface"></span> YAxisConfig

Configuration options for the yAxis

## Definition

```typescript
export interface YAxisConfig {
	// Sets the yAxis unit
	unit?: string;
	// Reverses the yAxis
	reverse?: boolean;
	// Controls the number of decimals for yAxis values
	decimals?: number;
	// Sets the minimum value for the yAxis
	min?: number;
	axisPlacement?: common.AxisPlacement;
	axisColorMode?: common.AxisColorMode;
	axisLabel?: string;
	axisWidth?: number;
	axisSoftMin?: number;
	axisSoftMax?: number;
	axisGridShow?: boolean;
	scaleDistribution?: common.ScaleDistributionConfig;
	axisCenteredZero?: boolean;
	// Sets the maximum value for the yAxis
	max?: number;
	axisBorderShow?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [YAxisConfigBuilder](./builder-YAxisConfigBuilder.md)
