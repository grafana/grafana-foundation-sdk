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
	enableInfiniteScrolling?: boolean;
	// TODO: figure out how to define callbacks
	onClickFilterLabel?: any;
	onClickFilterOutLabel?: any;
	isFilterLabelActive?: any;
	onClickFilterString?: any;
	onClickFilterOutString?: any;
	onClickShowField?: any;
	onClickHideField?: any;
	logRowMenuIconsBefore?: any;
	logRowMenuIconsAfter?: any;
	onNewLogsReceived?: any;
	displayedFields?: string[];
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

