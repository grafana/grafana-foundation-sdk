// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as logs from '../logs';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<logs.Options> {
    protected readonly internal: logs.Options;

    constructor() {
        this.internal = logs.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): logs.Options {
        return this.internal;
    }

    showLabels(showLabels: boolean): this {
        this.internal.showLabels = showLabels;
        return this;
    }

    showCommonLabels(showCommonLabels: boolean): this {
        this.internal.showCommonLabels = showCommonLabels;
        return this;
    }

    showTime(showTime: boolean): this {
        this.internal.showTime = showTime;
        return this;
    }

    showLogContextToggle(showLogContextToggle: boolean): this {
        this.internal.showLogContextToggle = showLogContextToggle;
        return this;
    }

    wrapLogMessage(wrapLogMessage: boolean): this {
        this.internal.wrapLogMessage = wrapLogMessage;
        return this;
    }

    prettifyLogMessage(prettifyLogMessage: boolean): this {
        this.internal.prettifyLogMessage = prettifyLogMessage;
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

    // TODO: figure out how to define callbacks
    onClickFilterLabel(onClickFilterLabel: any): this {
        this.internal.onClickFilterLabel = onClickFilterLabel;
        return this;
    }

    onClickFilterOutLabel(onClickFilterOutLabel: any): this {
        this.internal.onClickFilterOutLabel = onClickFilterOutLabel;
        return this;
    }

    isFilterLabelActive(isFilterLabelActive: any): this {
        this.internal.isFilterLabelActive = isFilterLabelActive;
        return this;
    }

    onClickFilterString(onClickFilterString: any): this {
        this.internal.onClickFilterString = onClickFilterString;
        return this;
    }

    onClickFilterOutString(onClickFilterOutString: any): this {
        this.internal.onClickFilterOutString = onClickFilterOutString;
        return this;
    }

    onClickShowField(onClickShowField: any): this {
        this.internal.onClickShowField = onClickShowField;
        return this;
    }

    onClickHideField(onClickHideField: any): this {
        this.internal.onClickHideField = onClickHideField;
        return this;
    }

    logRowMenuIconsBefore(logRowMenuIconsBefore: any): this {
        this.internal.logRowMenuIconsBefore = logRowMenuIconsBefore;
        return this;
    }

    logRowMenuIconsAfter(logRowMenuIconsAfter: any): this {
        this.internal.logRowMenuIconsAfter = logRowMenuIconsAfter;
        return this;
    }

    onNewLogsReceived(onNewLogsReceived: any): this {
        this.internal.onNewLogsReceived = onNewLogsReceived;
        return this;
    }

    displayedFields(displayedFields: string[]): this {
        this.internal.displayedFields = displayedFields;
        return this;
    }
}

