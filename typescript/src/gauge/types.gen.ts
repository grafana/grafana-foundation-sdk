// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	showThresholdLabels: boolean;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	showThresholdMarkers: boolean;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	showThresholdLabels: false,
	reduceOptions: common.defaultReduceDataOptions(),
	showThresholdMarkers: true,
	orientation: common.VizOrientation.Auto,
});

