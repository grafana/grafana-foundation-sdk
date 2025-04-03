// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	// Set the height of the rows
	rowHeight: number;
	// Show values on the columns
	showValue: common.VisibilityMode;
	// Controls the column width
	colWidth?: number;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	timezone?: common.TimeZone[];
	// Enables pagination when > 0
	perPage?: number;
}

export const defaultOptions = (): Options => ({
	rowHeight: 0.9,
	showValue: common.VisibilityMode.Auto,
	colWidth: 0.9,
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
	perPage: 20,
});

export interface FieldConfig {
	lineWidth?: number;
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
	fillOpacity?: number;
	axisBorderShow?: boolean;
}

export const defaultFieldConfig = (): FieldConfig => ({
	lineWidth: 1,
	fillOpacity: 70,
});

