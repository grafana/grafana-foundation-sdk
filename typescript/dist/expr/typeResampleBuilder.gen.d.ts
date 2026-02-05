import * as cog from '../cog';
import * as expr from '../expr';
import * as common from '../common';
export declare class TypeResampleBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: expr.TypeResample;
    constructor();
    /**
     * Builds the object.
     */
    build(): expr.TypeResample;
    datasource(datasource: common.DataSourceRef): this;
    downsampler(downsampler: "sum" | "mean" | "min" | "max" | "count" | "last" | "median"): this;
    expression(expression: string): this;
    hide(hide: boolean): this;
    intervalMs(intervalMs: number): this;
    maxDataPoints(maxDataPoints: number): this;
    queryType(queryType: string): this;
    refId(refId: string): this;
    resultAssertions(resultAssertions: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    }): this;
    timeRange(timeRange: {
        from: string;
        to: string;
    }): this;
    upsampler(upsampler: "pad" | "backfilling" | "fillna"): this;
    window(window: string): this;
}
