// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	showLabels: boolean;
	showCommonLabels: boolean;
	showTime: boolean;
	showLogContextToggle: boolean;
	wrapLogMessage: boolean;
	prettifyLogMessage: boolean;
	enableLogDetails: boolean;
	sortOrder: common.LogsSortOrder;
	dedupStrategy: common.LogsDedupStrategy;
}

export const defaultOptions = (): Options => ({
	showLabels: false,
	showCommonLabels: false,
	showTime: false,
	showLogContextToggle: false,
	wrapLogMessage: false,
	prettifyLogMessage: false,
	enableLogDetails: false,
	sortOrder: common.LogsSortOrder.Descending,
	dedupStrategy: common.LogsDedupStrategy.None,
});

