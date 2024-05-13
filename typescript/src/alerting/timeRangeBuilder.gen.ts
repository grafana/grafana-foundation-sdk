// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

// Redefining this to avoid an import cycle
export class TimeRangeBuilder implements cog.Builder<alerting.TimeRange> {
    protected readonly internal: alerting.TimeRange;

    constructor() {
        this.internal = alerting.defaultTimeRange();
    }

    build(): alerting.TimeRange {
        return this.internal;
    }

    // Redefining this to avoid an import cycle
    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    // Redefining this to avoid an import cycle
    to(to: string): this {
        this.internal.to = to;
        return this;
    }
}
