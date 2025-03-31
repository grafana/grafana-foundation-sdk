// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class DayOfMonthRangeBuilder implements cog.Builder<alerting.DayOfMonthRange> {
    protected readonly internal: alerting.DayOfMonthRange;

    constructor() {
        this.internal = alerting.defaultDayOfMonthRange();
    }

    /**
     * Builds the object.
     */
    build(): alerting.DayOfMonthRange {
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
