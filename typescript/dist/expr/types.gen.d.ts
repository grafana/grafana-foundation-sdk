import * as common from '../common';
export type expr = TypeMath | TypeReduce | TypeResample | TypeClassicConditions | TypeThreshold | TypeSql;
export declare const defaultExpr: () => expr;
export interface TypeMath {
    datasource?: common.DataSourceRef;
    expression: string;
    hide?: boolean;
    intervalMs?: number;
    maxDataPoints?: number;
    queryType?: string;
    refId?: string;
    resultAssertions?: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    };
    timeRange?: {
        from: string;
        to: string;
    };
    type: "math";
    _implementsDataqueryVariant(): void;
}
export declare const defaultTypeMath: () => TypeMath;
export interface TypeReduce {
    datasource?: common.DataSourceRef;
    expression: string;
    hide?: boolean;
    intervalMs?: number;
    maxDataPoints?: number;
    queryType?: string;
    reducer: "sum" | "mean" | "min" | "max" | "count" | "last" | "median";
    refId?: string;
    resultAssertions?: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    };
    settings?: {
        mode: "dropNN" | "replaceNN";
        replaceWithValue?: number;
    };
    timeRange?: {
        from: string;
        to: string;
    };
    type: "reduce";
    _implementsDataqueryVariant(): void;
}
export declare const defaultTypeReduce: () => TypeReduce;
export interface TypeResample {
    datasource?: common.DataSourceRef;
    downsampler: "sum" | "mean" | "min" | "max" | "count" | "last" | "median";
    expression: string;
    hide?: boolean;
    intervalMs?: number;
    maxDataPoints?: number;
    queryType?: string;
    refId?: string;
    resultAssertions?: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    };
    timeRange?: {
        from: string;
        to: string;
    };
    type: "resample";
    upsampler: "pad" | "backfilling" | "fillna";
    window: string;
    _implementsDataqueryVariant(): void;
}
export declare const defaultTypeResample: () => TypeResample;
export interface TypeClassicConditions {
    conditions: {
        evaluator: {
            params: number[];
            type: string;
        };
        operator: {
            type: "and" | "or" | "logic-or";
        };
        query: {
            params: string[];
        };
        reducer: {
            type: string;
        };
    }[];
    datasource?: common.DataSourceRef;
    hide?: boolean;
    intervalMs?: number;
    maxDataPoints?: number;
    queryType?: string;
    refId?: string;
    resultAssertions?: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    };
    timeRange?: {
        from: string;
        to: string;
    };
    type: "classic_conditions";
    _implementsDataqueryVariant(): void;
}
export declare const defaultTypeClassicConditions: () => TypeClassicConditions;
export interface TypeThreshold {
    conditions: {
        evaluator: {
            params: number[];
            type: "gt" | "lt" | "eq" | "ne" | "gte" | "lte" | "within_range" | "outside_range" | "within_range_included" | "outside_range_included";
        };
        loadedDimensions?: any;
        unloadEvaluator?: {
            params: number[];
            type: "gt" | "lt" | "eq" | "ne" | "gte" | "lte" | "within_range" | "outside_range" | "within_range_included" | "outside_range_included";
        };
    }[];
    datasource?: common.DataSourceRef;
    expression: string;
    hide?: boolean;
    intervalMs?: number;
    maxDataPoints?: number;
    queryType?: string;
    refId?: string;
    resultAssertions?: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    };
    timeRange?: {
        from: string;
        to: string;
    };
    type: "threshold";
    _implementsDataqueryVariant(): void;
}
export declare const defaultTypeThreshold: () => TypeThreshold;
export interface TypeSql {
    datasource?: common.DataSourceRef;
    expression: string;
    format: string;
    hide?: boolean;
    intervalMs?: number;
    maxDataPoints?: number;
    queryType?: string;
    refId?: string;
    resultAssertions?: {
        maxFrames?: number;
        type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
        typeVersion: number[];
    };
    timeRange?: {
        from: string;
        to: string;
    };
    type: "sql";
    _implementsDataqueryVariant(): void;
}
export declare const defaultTypeSql: () => TypeSql;
export interface TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql {
    TypeMath?: TypeMath;
    TypeReduce?: TypeReduce;
    TypeResample?: TypeResample;
    TypeClassicConditions?: TypeClassicConditions;
    TypeThreshold?: TypeThreshold;
    TypeSql?: TypeSql;
}
export declare const defaultTypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql: () => TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql;
