---
title: <span class="badge object-type-interface"></span> FieldConfig
---
# <span class="badge object-type-interface"></span> FieldConfig

## Definition

```typescript
export interface FieldConfig {
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
	labelValue?: common.TextDimensionConfig;
	axisCenteredZero?: boolean;
}

```
## Methods

No methods.
