import * as common from '../common';
export type BucketAggregation = DateHistogram | Histogram | Terms | Filters | GeoHashGrid | Nested;
export declare const defaultBucketAggregation: () => BucketAggregation;
export interface DateHistogram {
    field?: string;
    id: string;
    type: BucketAggregationType.DateHistogram;
    settings?: {
        interval?: string;
        min_doc_count?: string;
        trimEdges?: string;
        offset?: string;
        timeZone?: string;
    };
}
export declare const defaultDateHistogram: () => DateHistogram;
export declare enum BucketAggregationType {
    Terms = "terms",
    Filters = "filters",
    GeohashGrid = "geohash_grid",
    DateHistogram = "date_histogram",
    Histogram = "histogram",
    Nested = "nested"
}
export declare const defaultBucketAggregationType: () => BucketAggregationType;
export interface Histogram {
    field?: string;
    id: string;
    type: BucketAggregationType.Histogram;
    settings?: {
        interval?: string;
        min_doc_count?: string;
    };
}
export declare const defaultHistogram: () => Histogram;
export interface Terms {
    field?: string;
    id: string;
    type: BucketAggregationType.Terms;
    settings?: {
        order?: TermsOrder;
        size?: string;
        min_doc_count?: string;
        orderBy?: string;
        missing?: string;
    };
}
export declare const defaultTerms: () => Terms;
export declare enum TermsOrder {
    Desc = "desc",
    Asc = "asc"
}
export declare const defaultTermsOrder: () => TermsOrder;
export interface Filters {
    id: string;
    type: BucketAggregationType.Filters;
    settings?: {
        filters?: Filter[];
    };
}
export declare const defaultFilters: () => Filters;
export interface Filter {
    query: string;
    label: string;
}
export declare const defaultFilter: () => Filter;
export interface GeoHashGrid {
    field?: string;
    id: string;
    type: BucketAggregationType.GeohashGrid;
    settings?: {
        precision?: string;
    };
}
export declare const defaultGeoHashGrid: () => GeoHashGrid;
export interface Nested {
    field?: string;
    id: string;
    type: BucketAggregationType.Nested;
    settings?: any;
}
export declare const defaultNested: () => Nested;
export type MetricAggregation = Count | PipelineMetricAggregation | MetricAggregationWithSettings;
export declare const defaultMetricAggregation: () => MetricAggregation;
export interface Count {
    type: unknown;
    id: string;
    hide?: boolean;
}
export declare const defaultCount: () => Count;
export type MetricAggregationType = "count" | "avg" | "sum" | "min" | "max" | "extended_stats" | "percentiles" | "cardinality" | "raw_document" | "raw_data" | "logs" | "rate" | "top_metrics" | PipelineMetricAggregationType;
export declare const defaultMetricAggregationType: () => MetricAggregationType;
export declare enum PipelineMetricAggregationType {
    MovingAvg = "moving_avg",
    MovingFn = "moving_fn",
    Derivative = "derivative",
    SerialDiff = "serial_diff",
    CumulativeSum = "cumulative_sum",
    BucketScript = "bucket_script"
}
export declare const defaultPipelineMetricAggregationType: () => PipelineMetricAggregationType;
export type PipelineMetricAggregation = MovingAverage | Derivative | CumulativeSum | BucketScript;
export declare const defaultPipelineMetricAggregation: () => PipelineMetricAggregation;
export interface MovingAverage {
    pipelineAgg?: string;
    field?: string;
    type: unknown;
    id: string;
    settings?: Record<string, any>;
    hide?: boolean;
}
export declare const defaultMovingAverage: () => MovingAverage;
export interface Derivative {
    pipelineAgg?: string;
    field?: string;
    type: unknown;
    id: string;
    settings?: {
        unit?: string;
    };
    hide?: boolean;
}
export declare const defaultDerivative: () => Derivative;
export interface CumulativeSum {
    pipelineAgg?: string;
    field?: string;
    type: unknown;
    id: string;
    settings?: {
        format?: string;
    };
    hide?: boolean;
}
export declare const defaultCumulativeSum: () => CumulativeSum;
export interface BucketScript {
    type: unknown;
    pipelineVariables?: PipelineVariable[];
    id: string;
    settings?: {
        script?: InlineScript;
    };
    hide?: boolean;
}
export declare const defaultBucketScript: () => BucketScript;
export interface PipelineVariable {
    name: string;
    pipelineAgg: string;
}
export declare const defaultPipelineVariable: () => PipelineVariable;
export type InlineScript = string | {
    inline?: string;
};
export declare const defaultInlineScript: () => InlineScript;
export type MetricAggregationWithSettings = BucketScript | CumulativeSum | Derivative | SerialDiff | RawData | RawDocument | UniqueCount | Percentiles | ExtendedStats | Min | Max | Sum | Average | MovingAverage | MovingFunction | Logs | Rate | TopMetrics;
export declare const defaultMetricAggregationWithSettings: () => MetricAggregationWithSettings;
export interface SerialDiff {
    pipelineAgg?: string;
    field?: string;
    type: unknown;
    id: string;
    settings?: {
        lag?: string;
    };
    hide?: boolean;
}
export declare const defaultSerialDiff: () => SerialDiff;
export interface RawData {
    type: unknown;
    id: string;
    settings?: {
        size?: string;
    };
    hide?: boolean;
}
export declare const defaultRawData: () => RawData;
export interface RawDocument {
    type: unknown;
    id: string;
    settings?: {
        size?: string;
    };
    hide?: boolean;
}
export declare const defaultRawDocument: () => RawDocument;
export interface UniqueCount {
    type: unknown;
    field?: string;
    id: string;
    settings?: {
        precision_threshold?: string;
        missing?: string;
    };
    hide?: boolean;
}
export declare const defaultUniqueCount: () => UniqueCount;
export interface Percentiles {
    type: unknown;
    field?: string;
    id: string;
    settings?: {
        script?: InlineScript;
        missing?: string;
        percents?: string[];
    };
    hide?: boolean;
}
export declare const defaultPercentiles: () => Percentiles;
export interface ExtendedStats {
    type: unknown;
    settings?: {
        script?: InlineScript;
        missing?: string;
        sigma?: string;
    };
    field?: string;
    id: string;
    meta?: any;
    hide?: boolean;
}
export declare const defaultExtendedStats: () => ExtendedStats;
export interface Min {
    type: unknown;
    field?: string;
    id: string;
    settings?: {
        script?: InlineScript;
        missing?: string;
    };
    hide?: boolean;
}
export declare const defaultMin: () => Min;
export interface Max {
    type: unknown;
    field?: string;
    id: string;
    settings?: {
        script?: InlineScript;
        missing?: string;
    };
    hide?: boolean;
}
export declare const defaultMax: () => Max;
export interface Sum {
    type: unknown;
    field?: string;
    id: string;
    settings?: {
        script?: InlineScript;
        missing?: string;
    };
    hide?: boolean;
}
export declare const defaultSum: () => Sum;
export interface Average {
    type: unknown;
    field?: string;
    id: string;
    settings?: {
        script?: InlineScript;
        missing?: string;
    };
    hide?: boolean;
}
export declare const defaultAverage: () => Average;
export interface MovingFunction {
    pipelineAgg?: string;
    field?: string;
    type: unknown;
    id: string;
    settings?: {
        window?: string;
        script?: InlineScript;
        shift?: string;
    };
    hide?: boolean;
}
export declare const defaultMovingFunction: () => MovingFunction;
export interface Logs {
    type: unknown;
    id: string;
    settings?: {
        limit?: string;
    };
    hide?: boolean;
}
export declare const defaultLogs: () => Logs;
export interface Rate {
    type: unknown;
    field?: string;
    id: string;
    settings?: {
        unit?: string;
        mode?: string;
    };
    hide?: boolean;
}
export declare const defaultRate: () => Rate;
export interface TopMetrics {
    type: unknown;
    id: string;
    settings?: {
        order?: string;
        orderBy?: string;
        metrics?: string[];
    };
    hide?: boolean;
}
export declare const defaultTopMetrics: () => TopMetrics;
export interface BaseBucketAggregation {
    id: string;
    type: BucketAggregationType;
    settings?: any;
}
export declare const defaultBaseBucketAggregation: () => BaseBucketAggregation;
export interface BucketAggregationWithField {
    field?: string;
    id: string;
    type: BucketAggregationType;
    settings?: any;
}
export declare const defaultBucketAggregationWithField: () => BucketAggregationWithField;
export interface DateHistogramSettings {
    interval?: string;
    min_doc_count?: string;
    trimEdges?: string;
    offset?: string;
    timeZone?: string;
}
export declare const defaultDateHistogramSettings: () => DateHistogramSettings;
export interface HistogramSettings {
    interval?: string;
    min_doc_count?: string;
}
export declare const defaultHistogramSettings: () => HistogramSettings;
export interface TermsSettings {
    order?: TermsOrder;
    size?: string;
    min_doc_count?: string;
    orderBy?: string;
    missing?: string;
}
export declare const defaultTermsSettings: () => TermsSettings;
export interface FiltersSettings {
    filters?: Filter[];
}
export declare const defaultFiltersSettings: () => FiltersSettings;
export interface GeoHashGridSettings {
    precision?: string;
}
export declare const defaultGeoHashGridSettings: () => GeoHashGridSettings;
export interface BaseMetricAggregation {
    type: MetricAggregationType;
    id: string;
    hide?: boolean;
}
export declare const defaultBaseMetricAggregation: () => BaseMetricAggregation;
export interface MetricAggregationWithField {
    field?: string;
    type: MetricAggregationType;
    id: string;
    hide?: boolean;
}
export declare const defaultMetricAggregationWithField: () => MetricAggregationWithField;
export interface MetricAggregationWithMissingSupport {
    settings?: {
        missing?: string;
    };
    type: MetricAggregationType;
    id: string;
    hide?: boolean;
}
export declare const defaultMetricAggregationWithMissingSupport: () => MetricAggregationWithMissingSupport;
export interface MetricAggregationWithInlineScript {
    settings?: {
        script?: InlineScript;
    };
    type: MetricAggregationType;
    id: string;
    hide?: boolean;
}
export declare const defaultMetricAggregationWithInlineScript: () => MetricAggregationWithInlineScript;
export declare enum ExtendedStatMetaType {
    Avg = "avg",
    Min = "min",
    Max = "max",
    Sum = "sum",
    Count = "count",
    StdDeviation = "std_deviation",
    StdDeviationBoundsUpper = "std_deviation_bounds_upper",
    StdDeviationBoundsLower = "std_deviation_bounds_lower"
}
export declare const defaultExtendedStatMetaType: () => ExtendedStatMetaType;
export interface ExtendedStat {
    label: string;
    value: ExtendedStatMetaType;
}
export declare const defaultExtendedStat: () => ExtendedStat;
export interface BasePipelineMetricAggregation {
    pipelineAgg?: string;
    field?: string;
    type: string;
    id: string;
    hide?: boolean;
}
export declare const defaultBasePipelineMetricAggregation: () => BasePipelineMetricAggregation;
export interface PipelineMetricAggregationWithMultipleBucketPaths {
    pipelineVariables?: PipelineVariable[];
    type: MetricAggregationType;
    id: string;
    hide?: boolean;
}
export declare const defaultPipelineMetricAggregationWithMultipleBucketPaths: () => PipelineMetricAggregationWithMultipleBucketPaths;
export declare enum MovingAverageModel {
    Simple = "simple",
    Linear = "linear",
    Ewma = "ewma",
    Holt = "holt",
    HoltWinters = "holt_winters"
}
export declare const defaultMovingAverageModel: () => MovingAverageModel;
export interface MovingAverageModelOption {
    label: string;
    value: MovingAverageModel;
}
export declare const defaultMovingAverageModelOption: () => MovingAverageModelOption;
export interface BaseMovingAverageModelSettings {
    model: MovingAverageModel;
    window: string;
    predict: string;
}
export declare const defaultBaseMovingAverageModelSettings: () => BaseMovingAverageModelSettings;
export interface MovingAverageSimpleModelSettings {
    model: MovingAverageModel.Simple;
    window: string;
    predict: string;
}
export declare const defaultMovingAverageSimpleModelSettings: () => MovingAverageSimpleModelSettings;
export interface MovingAverageLinearModelSettings {
    model: MovingAverageModel.Linear;
    window: string;
    predict: string;
}
export declare const defaultMovingAverageLinearModelSettings: () => MovingAverageLinearModelSettings;
export interface MovingAverageEWMAModelSettings {
    model: MovingAverageModel.Ewma;
    settings?: {
        alpha?: string;
    };
    window: string;
    minimize: boolean;
    predict: string;
}
export declare const defaultMovingAverageEWMAModelSettings: () => MovingAverageEWMAModelSettings;
export interface MovingAverageHoltModelSettings {
    model: MovingAverageModel.Holt;
    settings: {
        alpha?: string;
        beta?: string;
    };
    window: string;
    minimize: boolean;
    predict: string;
}
export declare const defaultMovingAverageHoltModelSettings: () => MovingAverageHoltModelSettings;
export interface MovingAverageHoltWintersModelSettings {
    model: MovingAverageModel.HoltWinters;
    settings: {
        alpha?: string;
        beta?: string;
        gamma?: string;
        period?: string;
        pad?: boolean;
    };
    window: string;
    minimize: boolean;
    predict: string;
}
export declare const defaultMovingAverageHoltWintersModelSettings: () => MovingAverageHoltWintersModelSettings;
export interface dataquery {
    alias?: string;
    query?: string;
    timeField?: string;
    bucketAggs?: BucketAggregation[];
    metrics?: MetricAggregation[];
    refId?: string;
    hide?: boolean;
    queryType?: string;
    datasource?: common.DataSourceRef;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => dataquery;
