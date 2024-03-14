// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	// Manually select which field from the dataset to represent the x field.
	xField?: string;
	// Use the color value for a sibling field to color each bar value.
	colorByField?: string;
	// Controls the orientation of the bar chart, either vertical or horizontal.
	orientation: common.VizOrientation;
	// Controls the radius of each bar.
	barRadius?: number;
	// Controls the rotation of the x axis labels.
	xTickLabelRotation: number;
	// Sets the max length that a label can have before it is truncated.
	xTickLabelMaxLength: number;
	// Controls the spacing between x axis labels.
	// negative values indicate backwards skipping behavior
	xTickLabelSpacing?: number;
	// Controls whether bars are stacked or not, either normally or in percent mode.
	stacking: common.StackingMode;
	// This controls whether values are shown on top or to the left of bars.
	showValue: common.VisibilityMode;
	// Controls the width of bars. 1 = Max width, 0 = Min width.
	barWidth: number;
	// Controls the width of groups. 1 = max with, 0 = min width.
	groupWidth: number;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	text?: common.VizTextDisplayOptions;
	// Enables mode which highlights the entire bar area and shows tooltip when cursor
	// hovers over highlighted area
	fullHighlight: boolean;
}

export const defaultOptions = (): Options => ({
	orientation: common.VizOrientation.Auto,
	barRadius: 0,
	xTickLabelRotation: 0,
	xTickLabelMaxLength: 0,
	xTickLabelSpacing: 0,
	stacking: common.StackingMode.None,
	showValue: common.VisibilityMode.Auto,
	barWidth: 0.97,
	groupWidth: 0.7,
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
	fullHighlight: false,
});

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
	hideFrom?: common.HideSeriesConfig;
	// Threshold rendering
	thresholdsStyle?: common.GraphThresholdsStyleConfig;
	axisCenteredZero?: boolean;
}

export const defaultFieldConfig = (): FieldConfig => ({
	lineWidth: 1,
	fillOpacity: 80,
	gradientMode: common.GraphGradientMode.None,
});

