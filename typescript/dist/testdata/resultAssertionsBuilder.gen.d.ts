import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class ResultAssertionsBuilder implements cog.Builder<testdata.ResultAssertions> {
    protected readonly internal: testdata.ResultAssertions;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.ResultAssertions;
    maxFrames(maxFrames: number): this;
    type(type: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines"): this;
    typeVersion(typeVersion: number[]): this;
}
