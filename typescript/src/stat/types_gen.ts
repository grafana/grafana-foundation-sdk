// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	graphMode: common.BigValueGraphMode;
	colorMode: common.BigValueColorMode;
	justifyMode: common.BigValueJustifyMode;
	textMode: common.BigValueTextMode;
	wideLayout: boolean;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	showPercentChange: boolean;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	graphMode: common.BigValueGraphMode.Area,
	colorMode: common.BigValueColorMode.Value,
	justifyMode: common.BigValueJustifyMode.Auto,
	textMode: common.BigValueTextMode.Auto,
	wideLayout: true,
	reduceOptions: common.defaultReduceDataOptions(),
	showPercentChange: false,
	orientation: common.VizOrientation.Auto,
});

