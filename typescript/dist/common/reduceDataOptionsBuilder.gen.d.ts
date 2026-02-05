import * as cog from '../cog';
import * as common from '../common';
export declare class ReduceDataOptionsBuilder implements cog.Builder<common.ReduceDataOptions> {
    protected readonly internal: common.ReduceDataOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.ReduceDataOptions;
    values(values: boolean): this;
    limit(limit: number): this;
    calcs(calcs: string[]): this;
    fields(fields: string): this;
}
