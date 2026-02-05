import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class DayOfMonthRangeBuilder implements cog.Builder<alerting.DayOfMonthRange> {
    protected readonly internal: alerting.DayOfMonthRange;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.DayOfMonthRange;
    begin(begin: number): this;
    end(end: number): this;
}
