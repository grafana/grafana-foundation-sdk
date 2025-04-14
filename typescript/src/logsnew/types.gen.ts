// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	showControls: boolean;
	showTime: boolean;
	wrapLogMessage: boolean;
	enableLogDetails: boolean;
	syntaxHighlighting: boolean;
	sortOrder: common.LogsSortOrder;
	dedupStrategy: common.LogsDedupStrategy;
	grammar?: any;
	enableInfiniteScrolling?: boolean;
	onLogOptionsChange?: any;
	onNewLogsReceived?: any;
}

export const defaultOptions = (): Options => ({
	showControls: false,
	showTime: false,
	wrapLogMessage: false,
	enableLogDetails: false,
	syntaxHighlighting: false,
	sortOrder: common.LogsSortOrder.Descending,
	dedupStrategy: common.LogsDedupStrategy.None,
});

