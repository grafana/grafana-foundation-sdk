// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	timezone?: common.TimeZone[];
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
}

export const defaultOptions = (): Options => ({
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
});

export type FieldConfig = common.GraphFieldConfig;

export const defaultFieldConfig = (): FieldConfig => (common.defaultGraphFieldConfig());

