// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export type BucketAggregation = DateHistogram | Histogram | Terms | Filters | GeoHashGrid | Nested;

export const defaultBucketAggregation = (): BucketAggregation => (defaultDateHistogram());

export type MetricAggregation = Count | PipelineMetricAggregation | MetricAggregationWithSettings;

export const defaultMetricAggregation = (): MetricAggregation => (defaultCount());

export enum BucketAggregationType {
	Terms = "terms",
	Filters = "filters",
	GeohashGrid = "geohash_grid",
	DateHistogram = "date_histogram",
	Histogram = "histogram",
	Nested = "nested",
}

export const defaultBucketAggregationType = (): BucketAggregationType => (BucketAggregationType.Terms);

export interface BaseBucketAggregation {
	id: string;
	type: BucketAggregationType;
	settings?: any;
}

export const defaultBaseBucketAggregation = (): BaseBucketAggregation => ({
	id: "",
	type: BucketAggregationType.Terms,
});

export interface BucketAggregationWithField {
	field?: string;
	id: string;
	type: BucketAggregationType;
	settings?: any;
}

export const defaultBucketAggregationWithField = (): BucketAggregationWithField => ({
	id: "",
	type: BucketAggregationType.Terms,
});

export interface DateHistogram {
	field?: string;
	id: string;
	type: "date_histogram";
	settings?: {
		interval?: string;
		min_doc_count?: string;
		trimEdges?: string;
		offset?: string;
		timeZone?: string;
	};
}

export const defaultDateHistogram = (): DateHistogram => ({
	id: "",
	type: "date_histogram",
});

export interface DateHistogramSettings {
	interval?: string;
	min_doc_count?: string;
	trimEdges?: string;
	offset?: string;
	timeZone?: string;
}

export const defaultDateHistogramSettings = (): DateHistogramSettings => ({
});

export interface Histogram {
	field?: string;
	id: string;
	type: "histogram";
	settings?: {
		interval?: string;
		min_doc_count?: string;
	};
}

export const defaultHistogram = (): Histogram => ({
	id: "",
	type: "histogram",
});

export interface HistogramSettings {
	interval?: string;
	min_doc_count?: string;
}

export const defaultHistogramSettings = (): HistogramSettings => ({
});

export enum TermsOrder {
	Desc = "desc",
	Asc = "asc",
}

export const defaultTermsOrder = (): TermsOrder => (TermsOrder.Desc);

export interface Nested {
	field?: string;
	id: string;
	type: "nested";
	settings?: any;
}

export const defaultNested = (): Nested => ({
	id: "",
	type: "nested",
});

export interface Terms {
	field?: string;
	id: string;
	type: "terms";
	settings?: {
		order?: TermsOrder;
		size?: string;
		min_doc_count?: string;
		orderBy?: string;
		missing?: string;
	};
}

export const defaultTerms = (): Terms => ({
	id: "",
	type: "terms",
});

export interface TermsSettings {
	order?: TermsOrder;
	size?: string;
	min_doc_count?: string;
	orderBy?: string;
	missing?: string;
}

export const defaultTermsSettings = (): TermsSettings => ({
});

export interface Filters {
	id: string;
	type: "filters";
	settings?: {
		filters?: Filter[];
	};
}

export const defaultFilters = (): Filters => ({
	id: "",
	type: "filters",
});

export interface Filter {
	query: string;
	label: string;
}

export const defaultFilter = (): Filter => ({
	query: "",
	label: "",
});

export interface FiltersSettings {
	filters?: Filter[];
}

export const defaultFiltersSettings = (): FiltersSettings => ({
});

export interface GeoHashGrid {
	field?: string;
	id: string;
	type: "geohash_grid";
	settings?: {
		precision?: string;
	};
}

export const defaultGeoHashGrid = (): GeoHashGrid => ({
	id: "",
	type: "geohash_grid",
});

export interface GeoHashGridSettings {
	precision?: string;
}

export const defaultGeoHashGridSettings = (): GeoHashGridSettings => ({
});

export enum PipelineMetricAggregationType {
	MovingAvg = "moving_avg",
	MovingFn = "moving_fn",
	Derivative = "derivative",
	SerialDiff = "serial_diff",
	CumulativeSum = "cumulative_sum",
	BucketScript = "bucket_script",
}

export const defaultPipelineMetricAggregationType = (): PipelineMetricAggregationType => (PipelineMetricAggregationType.MovingAvg);

export type MetricAggregationType = "count" | "avg" | "sum" | "min" | "max" | "extended_stats" | "percentiles" | "cardinality" | "raw_document" | "raw_data" | "logs" | "rate" | "top_metrics" | PipelineMetricAggregationType;

export const defaultMetricAggregationType = (): MetricAggregationType => ("count");

export interface BaseMetricAggregation {
	type: MetricAggregationType;
	id: string;
	hide?: boolean;
}

export const defaultBaseMetricAggregation = (): BaseMetricAggregation => ({
	type: defaultMetricAggregationType(),
	id: "",
});

export interface PipelineVariable {
	name: string;
	pipelineAgg: string;
}

export const defaultPipelineVariable = (): PipelineVariable => ({
	name: "",
	pipelineAgg: "",
});

export interface MetricAggregationWithField {
	field?: string;
	type: MetricAggregationType;
	id: string;
	hide?: boolean;
}

export const defaultMetricAggregationWithField = (): MetricAggregationWithField => ({
	type: defaultMetricAggregationType(),
	id: "",
});

export interface MetricAggregationWithMissingSupport {
	settings?: {
		missing?: string;
	};
	type: MetricAggregationType;
	id: string;
	hide?: boolean;
}

export const defaultMetricAggregationWithMissingSupport = (): MetricAggregationWithMissingSupport => ({
	type: defaultMetricAggregationType(),
	id: "",
});

export type InlineScript = string | {
	inline?: string;
};

export const defaultInlineScript = (): InlineScript => ("");

export interface MetricAggregationWithInlineScript {
	settings?: {
		script?: InlineScript;
	};
	type: MetricAggregationType;
	id: string;
	hide?: boolean;
}

export const defaultMetricAggregationWithInlineScript = (): MetricAggregationWithInlineScript => ({
	type: defaultMetricAggregationType(),
	id: "",
});

export interface Count {
	type: "count";
	id: string;
	hide?: boolean;
}

export const defaultCount = (): Count => ({
	type: "count",
	id: "",
});

export interface Average {
	type: "avg";
	field?: string;
	id: string;
	settings?: {
		script?: InlineScript;
		missing?: string;
	};
	hide?: boolean;
}

export const defaultAverage = (): Average => ({
	type: "avg",
	id: "",
});

export interface Sum {
	type: "sum";
	field?: string;
	id: string;
	settings?: {
		script?: InlineScript;
		missing?: string;
	};
	hide?: boolean;
}

export const defaultSum = (): Sum => ({
	type: "sum",
	id: "",
});

export interface Max {
	type: "max";
	field?: string;
	id: string;
	settings?: {
		script?: InlineScript;
		missing?: string;
	};
	hide?: boolean;
}

export const defaultMax = (): Max => ({
	type: "max",
	id: "",
});

export interface Min {
	type: "min";
	field?: string;
	id: string;
	settings?: {
		script?: InlineScript;
		missing?: string;
	};
	hide?: boolean;
}

export const defaultMin = (): Min => ({
	type: "min",
	id: "",
});

export enum ExtendedStatMetaType {
	Avg = "avg",
	Min = "min",
	Max = "max",
	Sum = "sum",
	Count = "count",
	StdDeviation = "std_deviation",
	StdDeviationBoundsUpper = "std_deviation_bounds_upper",
	StdDeviationBoundsLower = "std_deviation_bounds_lower",
}

export const defaultExtendedStatMetaType = (): ExtendedStatMetaType => (ExtendedStatMetaType.Avg);

export interface ExtendedStat {
	label: string;
	value: ExtendedStatMetaType;
}

export const defaultExtendedStat = (): ExtendedStat => ({
	label: "",
	value: ExtendedStatMetaType.Avg,
});

export interface ExtendedStats {
	type: "extended_stats";
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

export const defaultExtendedStats = (): ExtendedStats => ({
	type: "extended_stats",
	id: "",
});

export interface Percentiles {
	type: "percentiles";
	field?: string;
	id: string;
	settings?: {
		script?: InlineScript;
		missing?: string;
		percents?: string[];
	};
	hide?: boolean;
}

export const defaultPercentiles = (): Percentiles => ({
	type: "percentiles",
	id: "",
});

export interface UniqueCount {
	type: "cardinality";
	field?: string;
	id: string;
	settings?: {
		precision_threshold?: string;
		missing?: string;
	};
	hide?: boolean;
}

export const defaultUniqueCount = (): UniqueCount => ({
	type: "cardinality",
	id: "",
});

export interface RawDocument {
	type: "raw_document";
	id: string;
	settings?: {
		size?: string;
	};
	hide?: boolean;
}

export const defaultRawDocument = (): RawDocument => ({
	type: "raw_document",
	id: "",
});

export interface RawData {
	type: "raw_data";
	id: string;
	settings?: {
		size?: string;
	};
	hide?: boolean;
}

export const defaultRawData = (): RawData => ({
	type: "raw_data",
	id: "",
});

export interface Logs {
	type: "logs";
	id: string;
	settings?: {
		limit?: string;
	};
	hide?: boolean;
}

export const defaultLogs = (): Logs => ({
	type: "logs",
	id: "",
});

export interface Rate {
	type: "rate";
	field?: string;
	id: string;
	settings?: {
		unit?: string;
		mode?: string;
	};
	hide?: boolean;
}

export const defaultRate = (): Rate => ({
	type: "rate",
	id: "",
});

export interface BasePipelineMetricAggregation {
	pipelineAgg?: string;
	field?: string;
	type: string;
	id: string;
	hide?: boolean;
}

export const defaultBasePipelineMetricAggregation = (): BasePipelineMetricAggregation => ({
	type: "",
	id: "",
});

export interface PipelineMetricAggregationWithMultipleBucketPaths {
	pipelineVariables?: PipelineVariable[];
	type: MetricAggregationType;
	id: string;
	hide?: boolean;
}

export const defaultPipelineMetricAggregationWithMultipleBucketPaths = (): PipelineMetricAggregationWithMultipleBucketPaths => ({
	type: defaultMetricAggregationType(),
	id: "",
});

export enum MovingAverageModel {
	Simple = "simple",
	Linear = "linear",
	Ewma = "ewma",
	Holt = "holt",
	HoltWinters = "holt_winters",
}

export const defaultMovingAverageModel = (): MovingAverageModel => (MovingAverageModel.Simple);

export interface MovingAverageModelOption {
	label: string;
	value: MovingAverageModel;
}

export const defaultMovingAverageModelOption = (): MovingAverageModelOption => ({
	label: "",
	value: MovingAverageModel.Simple,
});

export interface BaseMovingAverageModelSettings {
	model: MovingAverageModel;
	window: string;
	predict: string;
}

export const defaultBaseMovingAverageModelSettings = (): BaseMovingAverageModelSettings => ({
	model: MovingAverageModel.Simple,
	window: "",
	predict: "",
});

export interface MovingAverageSimpleModelSettings {
	model: "simple";
	window: string;
	predict: string;
}

export const defaultMovingAverageSimpleModelSettings = (): MovingAverageSimpleModelSettings => ({
	model: "simple",
	window: "",
	predict: "",
});

export interface MovingAverageLinearModelSettings {
	model: "linear";
	window: string;
	predict: string;
}

export const defaultMovingAverageLinearModelSettings = (): MovingAverageLinearModelSettings => ({
	model: "linear",
	window: "",
	predict: "",
});

export interface MovingAverageEWMAModelSettings {
	model: "ewma";
	settings?: {
		alpha?: string;
	};
	window: string;
	minimize: boolean;
	predict: string;
}

export const defaultMovingAverageEWMAModelSettings = (): MovingAverageEWMAModelSettings => ({
	model: "ewma",
	window: "",
	minimize: false,
	predict: "",
});

export interface MovingAverageHoltModelSettings {
	model: "holt";
	settings: {
		alpha?: string;
		beta?: string;
	};
	window: string;
	minimize: boolean;
	predict: string;
}

export const defaultMovingAverageHoltModelSettings = (): MovingAverageHoltModelSettings => ({
	model: "holt",
	settings: {
},
	window: "",
	minimize: false,
	predict: "",
});

export interface MovingAverageHoltWintersModelSettings {
	model: "holt_winters";
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

export const defaultMovingAverageHoltWintersModelSettings = (): MovingAverageHoltWintersModelSettings => ({
	model: "holt_winters",
	settings: {
},
	window: "",
	minimize: false,
	predict: "",
});

// #MovingAverage's settings are overridden in types.ts
export interface MovingAverage {
	pipelineAgg?: string;
	field?: string;
	type: "moving_avg";
	id: string;
	settings?: Record<string, any>;
	hide?: boolean;
}

export const defaultMovingAverage = (): MovingAverage => ({
	type: "moving_avg",
	id: "",
});

export interface MovingFunction {
	pipelineAgg?: string;
	field?: string;
	type: "moving_fn";
	id: string;
	settings?: {
		window?: string;
		script?: InlineScript;
		shift?: string;
	};
	hide?: boolean;
}

export const defaultMovingFunction = (): MovingFunction => ({
	type: "moving_fn",
	id: "",
});

export interface Derivative {
	pipelineAgg?: string;
	field?: string;
	type: "derivative";
	id: string;
	settings?: {
		unit?: string;
	};
	hide?: boolean;
}

export const defaultDerivative = (): Derivative => ({
	type: "derivative",
	id: "",
});

export interface SerialDiff {
	pipelineAgg?: string;
	field?: string;
	type: "serial_diff";
	id: string;
	settings?: {
		lag?: string;
	};
	hide?: boolean;
}

export const defaultSerialDiff = (): SerialDiff => ({
	type: "serial_diff",
	id: "",
});

export interface CumulativeSum {
	pipelineAgg?: string;
	field?: string;
	type: "cumulative_sum";
	id: string;
	settings?: {
		format?: string;
	};
	hide?: boolean;
}

export const defaultCumulativeSum = (): CumulativeSum => ({
	type: "cumulative_sum",
	id: "",
});

export interface BucketScript {
	type: "bucket_script";
	pipelineVariables?: PipelineVariable[];
	id: string;
	settings?: {
		script?: InlineScript;
	};
	hide?: boolean;
}

export const defaultBucketScript = (): BucketScript => ({
	type: "bucket_script",
	id: "",
});

export interface TopMetrics {
	type: "top_metrics";
	id: string;
	settings?: {
		order?: string;
		orderBy?: string;
		metrics?: string[];
	};
	hide?: boolean;
}

export const defaultTopMetrics = (): TopMetrics => ({
	type: "top_metrics",
	id: "",
});

export type PipelineMetricAggregation = MovingAverage | Derivative | CumulativeSum | BucketScript;

export const defaultPipelineMetricAggregation = (): PipelineMetricAggregation => (defaultMovingAverage());

export type MetricAggregationWithSettings = BucketScript | CumulativeSum | Derivative | SerialDiff | RawData | RawDocument | UniqueCount | Percentiles | ExtendedStats | Min | Max | Sum | Average | MovingAverage | MovingFunction | Logs | Rate | TopMetrics;

export const defaultMetricAggregationWithSettings = (): MetricAggregationWithSettings => (defaultBucketScript());

export interface dataquery {
	// Alias pattern
	alias?: string;
	// Lucene query
	query?: string;
	// Name of time field
	timeField?: string;
	// List of bucket aggregations
	bucketAggs?: BucketAggregation[];
	// List of metric aggregations
	metrics?: MetricAggregation[];
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: any;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	refId: "",
	_implementsDataqueryVariant: () => {},
});

