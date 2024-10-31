// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class TimeRangeBuilder implements cog.Builder<testdata.TimeRange> {
    protected readonly internal: testdata.TimeRange;

    constructor() {
        this.internal = testdata.defaultTimeRange();
    }

    /**
     * Builds the object.
     */
    build(): testdata.TimeRange {
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
