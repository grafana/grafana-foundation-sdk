import * as cog from '../cog';
import * as expr from '../expr';
import * as common from '../common';
export declare class TypeReduceBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: expr.TypeReduce;
    constructor();
    /**
     * Builds the object.
     */
    build(): expr.TypeReduce;
    datasource(datasource: common.DataSourceRef): this;
    expression(expression: string): this;
    hide(hide: boolean): this;
    intervalMs(intervalMs: number): this;
    maxDataPoints(maxDataPoints: number): this;
    queryType(queryType: string): this;
    reducer(reducer: "sum" | "mean" | "min" | "max" | "count" | "last" | "median"): this;
    refId(refId: string): this;
    resultAssertions(resultAssertions: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    }): this;
    settings(settings: {
        mode: "dropNN" | "replaceNN";
        replaceWithValue?: number;
    }): this;
    timeRange(timeRange: {
        from: string;
        to: string;
    }): this;
}
