// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class MonthRangeBuilder implements cog.Builder<alerting.MonthRange> {
    protected readonly internal: alerting.MonthRange;

    constructor() {
        this.internal = alerting.defaultMonthRange();
    }

    /**
     * Builds the object.
     */
    build(): alerting.MonthRange {
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
