// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class TimeIntervalBuilder implements cog.Builder<alerting.TimeInterval> {
    protected readonly internal: alerting.TimeInterval;

    constructor() {
        this.internal = alerting.defaultTimeInterval();
    }

    /**
     * Builds the object.
     */
    build(): alerting.TimeInterval {
        return this.internal;
    }

    times(times: cog.Builder<alerting.TimeRange>[]): this {
        const timesResources = times.map(builder1 => builder1.build());
        this.internal.times = timesResources;
        return this;
    }

    weekdays(weekdays: cog.Builder<alerting.WeekdayRange>[]): this {
        const weekdaysResources = weekdays.map(builder1 => builder1.build());
        this.internal.weekdays = weekdaysResources;
        return this;
    }

    daysOfMonth(daysOfMonth: cog.Builder<alerting.DayOfMonthRange>[]): this {
        const daysOfMonthResources = daysOfMonth.map(builder1 => builder1.build());
        this.internal.days_of_month = daysOfMonthResources;
        return this;
    }

    months(months: cog.Builder<alerting.MonthRange>[]): this {
        const monthsResources = months.map(builder1 => builder1.build());
        this.internal.months = monthsResources;
        return this;
    }

    years(years: cog.Builder<alerting.YearRange>[]): this {
        const yearsResources = years.map(builder1 => builder1.build());
        this.internal.years = yearsResources;
        return this;
    }

    location(location: alerting.Location): this {
        this.internal.location = location;
        return this;
    }
}
