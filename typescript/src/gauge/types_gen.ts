// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	showThresholdLabels: boolean;
	showThresholdMarkers: boolean;
	sizing: common.BarGaugeSizing;
	minVizWidth: number;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	minVizHeight: number;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	showThresholdLabels: false,
	showThresholdMarkers: true,
	sizing: common.BarGaugeSizing.Auto,
	minVizWidth: 75,
	reduceOptions: common.defaultReduceDataOptions(),
	minVizHeight: 75,
	orientation: common.VizOrientation.Auto,
});

