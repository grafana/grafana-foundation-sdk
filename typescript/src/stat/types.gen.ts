// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	graphMode: common.BigValueGraphMode;
	colorMode: common.BigValueColorMode;
	justifyMode: common.BigValueJustifyMode;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	textMode: common.BigValueTextMode;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	graphMode: common.BigValueGraphMode.Area,
	colorMode: common.BigValueColorMode.Value,
	justifyMode: common.BigValueJustifyMode.Auto,
	reduceOptions: common.defaultReduceDataOptions(),
	textMode: common.BigValueTextMode.Auto,
	orientation: common.VizOrientation.Auto,
});

