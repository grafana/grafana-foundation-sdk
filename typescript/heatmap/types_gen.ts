// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


// Controls the color mode of the heatmap
export enum HeatmapColorMode {
	Opacity = "opacity",
	Scheme = "scheme",
}

export const defaultHeatmapColorMode = (): HeatmapColorMode => (HeatmapColorMode.Opacity);

// Controls the color scale of the heatmap
export enum HeatmapColorScale {
	Linear = "linear",
	Exponential = "exponential",
}

export const defaultHeatmapColorScale = (): HeatmapColorScale => (HeatmapColorScale.Linear);

// Controls various color options
export interface HeatmapColorOptions {
	// Sets the color mode
	mode?: HeatmapColorMode;
	// Controls the color scheme used
	scheme: string;
	// Controls the color fill when in opacity mode
	fill: string;
	// Controls the color scale
	scale?: HeatmapColorScale;
	// Controls the exponent when scale is set to exponential
	exponent: number;
	// Controls the number of color steps
	steps: number;
	// Reverses the color scheme
	reverse: boolean;
	// Sets the minimum value for the color scale
	min?: number;
	// Sets the maximum value for the color scale
	max?: number;
}

export const defaultHeatmapColorOptions = (): HeatmapColorOptions => ({
	scheme: "",
	fill: "",
	exponent: 0,
	steps: 0,
	reverse: false,
});

// Configuration options for the yAxis
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

export const defaultYAxisConfig = (): YAxisConfig => ({
});

// Controls cell value options
export interface CellValues {
	// Controls the cell value unit
	unit?: string;
	// Controls the number of decimals for cell values
	decimals?: number;
}

export const defaultCellValues = (): CellValues => ({
});

// Controls the value filter range
export interface FilterValueRange {
	// Sets the filter range to values less than or equal to the given value
	le?: number;
	// Sets the filter range to values greater than or equal to the given value
	ge?: number;
}

export const defaultFilterValueRange = (): FilterValueRange => ({
});

// Controls tooltip options
export interface HeatmapTooltip {
	// Controls if the tooltip is shown
	show: boolean;
	// Controls if the tooltip shows a histogram of the y-axis values
	yHistogram?: boolean;
}

export const defaultHeatmapTooltip = (): HeatmapTooltip => ({
	show: false,
});

// Controls legend options
export interface HeatmapLegend {
	// Controls if the legend is shown
	show: boolean;
}

export const defaultHeatmapLegend = (): HeatmapLegend => ({
	show: false,
});

// Controls exemplar options
export interface ExemplarConfig {
	// Sets the color of the exemplar markers
	color: string;
}

export const defaultExemplarConfig = (): ExemplarConfig => ({
	color: "",
});

// Controls frame rows options
export interface RowsHeatmapOptions {
	// Sets the name of the cell when not calculating from data
	value?: string;
	// Controls tick alignment when not calculating from data
	layout?: common.HeatmapCellLayout;
}

export const defaultRowsHeatmapOptions = (): RowsHeatmapOptions => ({
});

export interface Options {
	// Controls if the heatmap should be calculated from data
	calculate?: boolean;
	// Calculation options for the heatmap
	calculation?: common.HeatmapCalculationOptions;
	// Controls the color options
	color: HeatmapColorOptions;
	// Filters values between a given range
	filterValues?: FilterValueRange;
	// Controls tick alignment and value name when not calculating from data
	rowsFrame?: RowsHeatmapOptions;
	// | *{
	// 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
	// }
	// Controls the display of the value in the cell
	showValue: common.VisibilityMode;
	// Controls gap between cells
	cellGap?: number;
	// Controls cell radius
	cellRadius?: number;
	// Controls cell value unit
	cellValues?: CellValues;
	// Controls yAxis placement
	yAxis: YAxisConfig;
	// | *{
	// 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
	// }
	// Controls legend options
	legend: {
		show: true;
	};
	// Controls tooltip options
	tooltip: {
		show: true;
		yHistogram: false;
	};
	// Controls exemplar options
	exemplars: {
		color: "rgba(255,0,255,0.7)";
	};
}

export const defaultOptions = (): Options => ({
	calculate: false,
	color: defaultHeatmapColorOptions(),
	showValue: common.VisibilityMode.Auto,
	cellGap: 1,
	yAxis: defaultYAxisConfig(),
	legend: {
	show: true,
},
	tooltip: {
	show: true,
	yHistogram: false,
},
	exemplars: {
	color: "rgba(255,0,255,0.7)",
},
});

export interface FieldConfig {
	scaleDistribution?: common.ScaleDistributionConfig;
	hideFrom?: common.HideSeriesConfig;
}

export const defaultFieldConfig = (): FieldConfig => ({
});

