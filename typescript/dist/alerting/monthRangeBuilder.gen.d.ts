import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class MonthRangeBuilder implements cog.Builder<alerting.MonthRange> {
    protected readonly internal: alerting.MonthRange;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.MonthRange;
    begin(begin: number): this;
    end(end: number): this;
}
