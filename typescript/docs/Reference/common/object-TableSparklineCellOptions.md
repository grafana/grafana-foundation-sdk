---
title: <span class="badge object-type-interface"></span> TableSparklineCellOptions
---
# <span class="badge object-type-interface"></span> TableSparklineCellOptions

Sparkline cell options

## Definition

```typescript
export interface TableSparklineCellOptions {
	type: "sparkline";
	drawStyle?: common.GraphDrawStyle;
	gradientMode?: common.GraphGradientMode;
	thresholdsStyle?: common.GraphThresholdsStyleConfig;
	transform?: common.GraphTransform;
	lineColor?: string;
	lineWidth?: number;
	lineInterpolation?: common.LineInterpolation;
	lineStyle?: common.LineStyle;
	fillColor?: string;
	fillOpacity?: number;
	showPoints?: common.VisibilityMode;
	pointSize?: number;
	pointColor?: string;
	axisPlacement?: common.AxisPlacement;
	axisColorMode?: common.AxisColorMode;
	axisLabel?: string;
	axisWidth?: number;
	axisSoftMin?: number;
	axisSoftMax?: number;
	axisGridShow?: boolean;
	scaleDistribution?: common.ScaleDistributionConfig;
	axisCenteredZero?: boolean;
	barAlignment?: common.BarAlignment;
	barWidthFactor?: number;
	stacking?: common.StackingConfig;
	hideFrom?: common.HideSeriesConfig;
	hideValue?: boolean;
	insertNulls?: boolean | number;
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	spanNulls?: boolean | number;
	fillBelowTo?: string;
	pointSymbol?: string;
	axisBorderShow?: boolean;
	barMaxWidth?: number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TableSparklineCellOptionsBuilder](./builder-TableSparklineCellOptionsBuilder.md)
