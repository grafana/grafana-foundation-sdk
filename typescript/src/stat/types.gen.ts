// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	graphMode: common.BigValueGraphMode;
	colorMode: common.BigValueColorMode;
	justifyMode: common.BigValueJustifyMode;
	textMode: common.BigValueTextMode;
	wideLayout: boolean;
	showPercentChange: boolean;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	percentChangeColorMode: common.PercentChangeColorMode;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	graphMode: common.BigValueGraphMode.Area,
	colorMode: common.BigValueColorMode.Value,
	justifyMode: common.BigValueJustifyMode.Auto,
	textMode: common.BigValueTextMode.Auto,
	wideLayout: true,
	showPercentChange: false,
	reduceOptions: common.defaultReduceDataOptions(),
	percentChangeColorMode: common.PercentChangeColorMode.Standard,
	orientation: common.VizOrientation.Auto,
});

