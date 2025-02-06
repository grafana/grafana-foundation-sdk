// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	showTime: boolean;
	wrapLogMessage: boolean;
	enableLogDetails: boolean;
	sortOrder: common.LogsSortOrder;
	dedupStrategy: common.LogsDedupStrategy;
	enableInfiniteScrolling?: boolean;
}

export const defaultOptions = (): Options => ({
	showTime: false,
	wrapLogMessage: false,
	enableLogDetails: false,
	sortOrder: common.LogsSortOrder.Descending,
	dedupStrategy: common.LogsDedupStrategy.None,
});

