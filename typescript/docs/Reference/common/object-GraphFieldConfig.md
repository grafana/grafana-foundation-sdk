---
title: <span class="badge object-type-interface"></span> GraphFieldConfig
---
# <span class="badge object-type-interface"></span> GraphFieldConfig

TODO docs

## Definition

```typescript
export interface GraphFieldConfig {
	drawStyle?: common.GraphDrawStyle;
	gradientMode?: common.GraphGradientMode;
	thresholdsStyle?: common.GraphThresholdsStyleConfig;
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
	barAlignment?: common.BarAlignment;
	barWidthFactor?: number;
	stacking?: common.StackingConfig;
	hideFrom?: common.HideSeriesConfig;
	transform?: common.GraphTransform;
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	spanNulls?: boolean | number;
	fillBelowTo?: string;
	pointSymbol?: string;
	axisCenteredZero?: boolean;
	barMaxWidth?: number;
	insertNulls?: boolean | number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [GraphFieldConfigBuilder](./builder-GraphFieldConfigBuilder.md)
