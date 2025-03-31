// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class YearRangeBuilder implements cog.Builder<alerting.YearRange> {
    protected readonly internal: alerting.YearRange;

    constructor() {
        this.internal = alerting.defaultYearRange();
    }

    /**
     * Builds the object.
     */
    build(): alerting.YearRange {
        return this.internal;
    }

    begin(begin: number): this {
        this.internal.begin = begin;
        return this;
    }

    end(end: number): this {
        this.internal.end = end;
        return this;
    }
}
