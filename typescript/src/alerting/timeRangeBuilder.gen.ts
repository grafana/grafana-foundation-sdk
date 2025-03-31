// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

// Redefining this to avoid an import cycle
export class TimeRangeBuilder implements cog.Builder<alerting.TimeRange> {
    protected readonly internal: alerting.TimeRange;

    constructor() {
        this.internal = alerting.defaultTimeRange();
    }

    /**
     * Builds the object.
     */
    build(): alerting.TimeRange {
        return this.internal;
    }

    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    to(to: string): this {
        this.internal.to = to;
        return this;
    }
}
