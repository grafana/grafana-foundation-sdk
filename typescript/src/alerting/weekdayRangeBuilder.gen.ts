// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class WeekdayRangeBuilder implements cog.Builder<alerting.WeekdayRange> {
    protected readonly internal: alerting.WeekdayRange;

    constructor() {
        this.internal = alerting.defaultWeekdayRange();
    }

    /**
     * Builds the object.
     */
    build(): alerting.WeekdayRange {
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

