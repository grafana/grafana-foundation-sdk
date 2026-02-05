import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class USAQueryBuilder implements cog.Builder<testdata.USAQuery> {
    protected readonly internal: testdata.USAQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.USAQuery;
    fields(fields: string[]): this;
    mode(mode: string): this;
    period(period: string): this;
    states(states: string[]): this;
}
