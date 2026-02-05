import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class WeekdayRangeBuilder implements cog.Builder<alerting.WeekdayRange> {
    protected readonly internal: alerting.WeekdayRange;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.WeekdayRange;
    begin(begin: number): this;
    end(end: number): this;
}
