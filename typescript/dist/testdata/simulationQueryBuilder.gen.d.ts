import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class SimulationQueryBuilder implements cog.Builder<testdata.SimulationQuery> {
    protected readonly internal: testdata.SimulationQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.SimulationQuery;
    config(config: any): this;
    key(key: cog.Builder<testdata.Key>): this;
    last(last: boolean): this;
    stream(stream: boolean): this;
}
