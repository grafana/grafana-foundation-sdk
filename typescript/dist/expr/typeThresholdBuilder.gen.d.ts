import * as cog from '../cog';
import * as expr from '../expr';
import * as common from '../common';
export declare class TypeThresholdBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: expr.TypeThreshold;
    constructor();
    /**
     * Builds the object.
     */
    build(): expr.TypeThreshold;
    conditions(conditions: {
        evaluator: {
            params: number[];
            type: "gt" | "lt" | "eq" | "ne" | "gte" | "lte" | "within_range" | "outside_range" | "within_range_included" | "outside_range_included";
        };
        loadedDimensions?: any;
        unloadEvaluator?: {
            params: number[];
            type: "gt" | "lt" | "eq" | "ne" | "gte" | "lte" | "within_range" | "outside_range" | "within_range_included" | "outside_range_included";
        };
    }[]): this;
    datasource(datasource: common.DataSourceRef): this;
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
}
