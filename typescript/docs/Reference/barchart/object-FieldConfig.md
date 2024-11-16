---
title: <span class="badge object-type-interface"></span> FieldConfig
---
# <span class="badge object-type-interface"></span> FieldConfig

## Definition

```typescript
export interface FieldConfig {
	// Controls line width of the bars.
	lineWidth?: number;
	// Controls the fill opacity of the bars.
	fillOpacity?: number;
	// Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
	// Gradient appearance is influenced by the Fill opacity setting.
	gradientMode?: common.GraphGradientMode;
	axisPlacement?: common.AxisPlacement;
	axisColorMode?: common.AxisColorMode;
	axisLabel?: string;
	axisWidth?: number;
	axisSoftMin?: number;
	axisSoftMax?: number;
	axisGridShow?: boolean;
	scaleDistribution?: common.ScaleDistributionConfig;
	axisCenteredZero?: boolean;
	hideFrom?: common.HideSeriesConfig;
	// Threshold rendering
	thresholdsStyle?: common.GraphThresholdsStyleConfig;
	axisBorderShow?: boolean;
}

```
## Methods

No methods.