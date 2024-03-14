// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	// Show timeline values on chart
	showValue: common.VisibilityMode;
	// Controls the row height
	rowHeight: number;
	// Merge equal consecutive values
	mergeValues?: boolean;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	timezone?: common.TimeZone[];
	// Controls value alignment on the timelines
	alignValue?: common.TimelineValueAlignment;
}

export const defaultOptions = (): Options => ({
	showValue: common.VisibilityMode.Auto,
	rowHeight: 0.9,
	mergeValues: true,
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
	alignValue: common.TimelineValueAlignment.Left,
});

export interface FieldConfig {
	lineWidth?: number;
	hideFrom?: common.HideSeriesConfig;
	fillOpacity?: number;
}

export const defaultFieldConfig = (): FieldConfig => ({
	lineWidth: 0,
	fillOpacity: 70,
});

