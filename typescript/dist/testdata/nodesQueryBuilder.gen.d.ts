import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class NodesQueryBuilder implements cog.Builder<testdata.NodesQuery> {
    protected readonly internal: testdata.NodesQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.NodesQuery;
    count(count: number): this;
    seed(seed: number): this;
    type(type: "random" | "random edges" | "response_medium" | "response_small" | "feature_showcase"): this;
}
