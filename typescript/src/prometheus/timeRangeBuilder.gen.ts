// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';

export class TimeRangeBuilder implements cog.Builder<prometheus.TimeRange> {
    protected readonly internal: prometheus.TimeRange;

    constructor() {
        this.internal = prometheus.defaultTimeRange();
    }

    /**
     * Builds the object.
     */
    build(): prometheus.TimeRange {
        return this.internal;
    }

    // From is the start time of the query.
    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    // To is the end time of the query.
    to(to: string): this {
        this.internal.to = to;
        return this;
    }
}

