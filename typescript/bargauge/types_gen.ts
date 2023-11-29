// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	displayMode: common.BarGaugeDisplayMode;
	valueMode: common.BarGaugeValueMode;
	namePlacement: common.BarGaugeNamePlacement;
	showUnfilled: boolean;
	minVizWidth: number;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	minVizHeight: number;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	displayMode: common.BarGaugeDisplayMode.Gradient,
	valueMode: common.BarGaugeValueMode.Color,
	namePlacement: common.BarGaugeNamePlacement.Auto,
	showUnfilled: true,
	minVizWidth: 0,
	reduceOptions: common.defaultReduceDataOptions(),
	minVizHeight: 10,
	orientation: common.VizOrientation.Auto,
});

