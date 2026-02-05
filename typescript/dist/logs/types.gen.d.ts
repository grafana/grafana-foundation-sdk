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
export declare const defaultOptions: () => Options;
