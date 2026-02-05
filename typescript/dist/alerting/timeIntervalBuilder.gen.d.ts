import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class TimeIntervalBuilder implements cog.Builder<alerting.TimeInterval> {
    protected readonly internal: alerting.TimeInterval;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.TimeInterval;
    times(times: cog.Builder<alerting.TimeRange>[]): this;
    weekdays(weekdays: cog.Builder<alerting.WeekdayRange>[]): this;
    daysOfMonth(daysOfMonth: cog.Builder<alerting.DayOfMonthRange>[]): this;
    months(months: cog.Builder<alerting.MonthRange>[]): this;
    years(years: cog.Builder<alerting.YearRange>[]): this;
    location(location: alerting.Location): this;
}
