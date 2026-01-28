---
title: <span class="badge object-type-interface"></span> FieldConfig
---
# <span class="badge object-type-interface"></span> FieldConfig

## Definition

```typescript
export interface FieldConfig {
	show?: xychart.XYShowMode;
	pointSize?: {
		fixed?: number;
		min?: number;
		max?: number;
	};
	pointShape?: xychart.PointShape;
	pointStrokeWidth?: number;
	fillOpacity?: number;
	lineWidth?: number;
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
	lineStyle?: common.LineStyle;
	axisBorderShow?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [FieldConfigBuilder](./builder-FieldConfigBuilder.md)
