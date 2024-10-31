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
	axisPlacement?: common.AxisPlacement;
	axisColorMode?: common.AxisColorMode;
	axisLabel?: string;
	axisWidth?: number;
	axisSoftMin?: number;
	axisSoftMax?: number;
	axisGridShow?: boolean;
	scaleDistribution?: common.ScaleDistributionConfig;
	hideFrom?: common.HideSeriesConfig;
	// Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
	// Gradient appearance is influenced by the Fill opacity setting.
	gradientMode?: common.GraphGradientMode;
	axisCenteredZero?: boolean;
}

```
## Methods

No methods.
