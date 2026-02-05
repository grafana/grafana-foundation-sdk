import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class TimeRangeBuilder implements cog.Builder<alerting.TimeRange> {
    protected readonly internal: alerting.TimeRange;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.TimeRange;
    from(from: string): this;
    to(to: string): this;
}
