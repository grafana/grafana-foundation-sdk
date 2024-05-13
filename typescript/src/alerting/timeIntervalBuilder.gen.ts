// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

// TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
// within the interval.
export class TimeIntervalBuilder implements cog.Builder<alerting.TimeInterval> {
    protected readonly internal: alerting.TimeInterval;

    constructor() {
        this.internal = alerting.defaultTimeInterval();
    }

    build(): alerting.TimeInterval {
        return this.internal;
    }

    // TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    // within the interval.
    daysOfMonth(daysOfMonth: string[]): this {
        this.internal.days_of_month = daysOfMonth;
        return this;
    }

    // TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    // within the interval.
    location(location: string): this {
        this.internal.location = location;
        return this;
    }

    // TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    // within the interval.
    months(months: string[]): this {
        this.internal.months = months;
        return this;
    }

    // TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    // within the interval.
    times(times: cog.Builder<alerting.TimeRange>[]): this {
        const timesResources = times.map(builder1 => builder1.build());
        this.internal.times = timesResources;
        return this;
    }

    // TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    // within the interval.
    weekdays(weekdays: string[]): this {
        this.internal.weekdays = weekdays;
        return this;
    }

    // TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
    // within the interval.
    years(years: string[]): this {
        this.internal.years = years;
        return this;
    }
}
