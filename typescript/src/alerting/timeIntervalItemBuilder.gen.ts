// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class TimeIntervalItemBuilder implements cog.Builder<alerting.TimeIntervalItem> {
    protected readonly internal: alerting.TimeIntervalItem;

    constructor() {
        this.internal = alerting.defaultTimeIntervalItem();
    }

    /**
     * Builds the object.
     */
    build(): alerting.TimeIntervalItem {
        return this.internal;
    }

    daysOfMonth(daysOfMonth: string[]): this {
        this.internal.days_of_month = daysOfMonth;
        return this;
    }

    location(location: string): this {
        this.internal.location = location;
        return this;
    }

    months(months: string[]): this {
        this.internal.months = months;
        return this;
    }

    times(times: cog.Builder<alerting.TimeIntervalTimeRange>[]): this {
        const timesResources = times.map(builder1 => builder1.build());
        this.internal.times = timesResources;
        return this;
    }

    weekdays(weekdays: string[]): this {
        this.internal.weekdays = weekdays;
        return this;
    }

    years(years: string[]): this {
        this.internal.years = years;
        return this;
    }
}
