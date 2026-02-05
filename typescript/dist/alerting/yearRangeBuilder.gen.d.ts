import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class YearRangeBuilder implements cog.Builder<alerting.YearRange> {
    protected readonly internal: alerting.YearRange;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.YearRange;
    begin(begin: number): this;
    end(end: number): this;
}
