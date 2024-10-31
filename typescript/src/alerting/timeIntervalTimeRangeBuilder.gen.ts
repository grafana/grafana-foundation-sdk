// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class TimeIntervalTimeRangeBuilder implements cog.Builder<alerting.TimeIntervalTimeRange> {
    protected readonly internal: alerting.TimeIntervalTimeRange;

    constructor() {
        this.internal = alerting.defaultTimeIntervalTimeRange();
    }

    /**
     * Builds the object.
     */
    build(): alerting.TimeIntervalTimeRange {
        return this.internal;
    }

    endTime(endTime: string): this {
        this.internal.end_time = endTime;
        return this;
    }

    startTime(startTime: string): this {
        this.internal.start_time = startTime;
        return this;
    }
}
