---
title: <span class="badge object-type-interface"></span> ScatterSeriesConfig
---
# <span class="badge object-type-interface"></span> ScatterSeriesConfig

## Definition

```typescript
export interface ScatterSeriesConfig {
	x?: string;
	y?: string;
	name?: string;
	show?: xychart.ScatterShow;
	pointSize?: common.ScaleDimensionConfig;
	pointColor?: common.ColorDimensionConfig;
	lineColor?: common.ColorDimensionConfig;
	lineWidth?: number;
	lineStyle?: common.LineStyle;
	label?: common.VisibilityMode;
	hideFrom?: common.HideSeriesConfig;
	axisPlacement?: common.AxisPlacement;
	axisColorMode?: common.AxisColorMode;
	axisLabel?: string;
	axisWidth?: number;
	axisSoftMin?: number;
	axisSoftMax?: number;
	axisGridShow?: boolean;
	scaleDistribution?: common.ScaleDistributionConfig;
	axisCenteredZero?: boolean;
	frame?: number;
	labelValue?: common.TextDimensionConfig;
	axisBorderShow?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ScatterSeriesConfigBuilder](./builder-ScatterSeriesConfigBuilder.md)
