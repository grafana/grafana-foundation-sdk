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
}

