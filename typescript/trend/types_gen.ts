// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


// Identical to timeseries... except it does not have timezone settings
export interface Options {
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	// Name of the x field to use (defaults to first number)
	xField?: string;
}

export const defaultOptions = (): Options => ({
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
});

export type FieldConfig = common.GraphFieldConfig;

export const defaultFieldConfig = (): FieldConfig => (common.defaultGraphFieldConfig());

