import * as common from '../common';
export interface Options {
    showTime: boolean;
    wrapLogMessage: boolean;
    enableLogDetails: boolean;
    sortOrder: common.LogsSortOrder;
    dedupStrategy: common.LogsDedupStrategy;
    enableInfiniteScrolling?: boolean;
    onNewLogsReceived?: any;
}
export declare const defaultOptions: () => Options;
