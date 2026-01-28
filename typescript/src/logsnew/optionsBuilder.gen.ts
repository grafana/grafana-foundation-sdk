// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as logsnew from '../logsnew';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<logsnew.Options> {
    protected readonly internal: logsnew.Options;

    constructor() {
        this.internal = logsnew.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): logsnew.Options {
        return this.internal;
    }

    showTime(showTime: boolean): this {
        this.internal.showTime = showTime;
        return this;
    }

    wrapLogMessage(wrapLogMessage: boolean): this {
        this.internal.wrapLogMessage = wrapLogMessage;
        return this;
    }

    enableLogDetails(enableLogDetails: boolean): this {
        this.internal.enableLogDetails = enableLogDetails;
        return this;
    }

    sortOrder(sortOrder: common.LogsSortOrder): this {
        this.internal.sortOrder = sortOrder;
        return this;
    }

    dedupStrategy(dedupStrategy: common.LogsDedupStrategy): this {
        this.internal.dedupStrategy = dedupStrategy;
        return this;
    }

    enableInfiniteScrolling(enableInfiniteScrolling: boolean): this {
        this.internal.enableInfiniteScrolling = enableInfiniteScrolling;
        return this;
    }

    onNewLogsReceived(onNewLogsReceived: any): this {
        this.internal.onNewLogsReceived = onNewLogsReceived;
        return this;
    }
}

