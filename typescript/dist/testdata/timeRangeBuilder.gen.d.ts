import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class TimeRangeBuilder implements cog.Builder<testdata.TimeRange> {
    protected readonly internal: testdata.TimeRange;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.TimeRange;
    from(from: string): this;
    to(to: string): this;
}
