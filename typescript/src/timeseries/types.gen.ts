// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	timezone?: common.TimeZone[];
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	orientation?: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	legend: { displayMode: common.LegendDisplayMode.List, placement: common.LegendPlacement.Bottom, showLegend: false, calcs: [
], },
	tooltip: common.defaultVizTooltipOptions(),
});

export type FieldConfig = common.GraphFieldConfig;

export const defaultFieldConfig = (): FieldConfig => (common.defaultGraphFieldConfig());

