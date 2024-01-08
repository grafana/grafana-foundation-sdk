// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	displayMode: common.BarGaugeDisplayMode;
	valueMode: common.BarGaugeValueMode;
	namePlacement: common.BarGaugeNamePlacement;
	showUnfilled: boolean;
	sizing: common.BarGaugeSizing;
	minVizWidth: number;
	minVizHeight: number;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	maxVizHeight: number;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	displayMode: common.BarGaugeDisplayMode.Gradient,
	valueMode: common.BarGaugeValueMode.Color,
	namePlacement: common.BarGaugeNamePlacement.Auto,
	showUnfilled: true,
	sizing: common.BarGaugeSizing.Auto,
	minVizWidth: 8,
	minVizHeight: 16,
	reduceOptions: common.defaultReduceDataOptions(),
	maxVizHeight: 300,
	orientation: common.VizOrientation.Auto,
});

