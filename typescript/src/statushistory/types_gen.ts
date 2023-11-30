// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	// Set the height of the rows
	rowHeight: number;
	// Show values on the columns
	showValue: common.VisibilityMode;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	timezone?: common.TimeZone[];
	// Controls the column width
	colWidth?: number;
}

export const defaultOptions = (): Options => ({
	rowHeight: 0.9,
	showValue: common.VisibilityMode.Auto,
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
	colWidth: 0.9,
});

export interface FieldConfig {
	lineWidth?: number;
	hideFrom?: common.HideSeriesConfig;
	fillOpacity?: number;
}

export const defaultFieldConfig = (): FieldConfig => ({
	lineWidth: 1,
	fillOpacity: 70,
});

