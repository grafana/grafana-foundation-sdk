// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	// Size of each bucket
	bucketSize?: number;
	// Offset buckets by this amount
	bucketOffset?: number;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	// Combines multiple series into a single histogram
	combine?: boolean;
}

export const defaultOptions = (): Options => ({
	bucketOffset: 0,
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
});

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

export const defaultFieldConfig = (): FieldConfig => ({
	lineWidth: 1,
	fillOpacity: 80,
	gradientMode: common.GraphGradientMode.None,
});

