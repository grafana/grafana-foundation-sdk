// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


export interface CloudMonitoringQuery {
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
	// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
	aliasBy?: string;
	// GCM query type.
	// queryType: #QueryType
	// Time Series List sub-query properties.
	timeSeriesList?: TimeSeriesList;
	// Time Series sub-query properties.
	timeSeriesQuery?: TimeSeriesQuery;
	// SLO sub-query properties.
	sloQuery?: SLOQuery;
	// PromQL sub-query properties.
	promQLQuery?: PromQLQuery;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// Time interval in milliseconds.
	intervalMs?: number;
	_implementsDataqueryVariant(): void;
}

export const defaultCloudMonitoringQuery = (): CloudMonitoringQuery => ({
	refId: "",
	_implementsDataqueryVariant: () => {},
});

// Time Series List sub-query properties.
export interface TimeSeriesList {
	// GCP project to execute the query against.
	projectName: string;
	// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	crossSeriesReducer: string;
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	alignmentPeriod?: string;
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	perSeriesAligner?: string;
	// Array of labels to group data by.
	groupBys?: string[];
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	filters?: string[];
	// Data view, defaults to FULL.
	view?: string;
	// Annotation title.
	title?: string;
	// Annotation text.
	text?: string;
	// Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	secondaryCrossSeriesReducer?: string;
	// Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	secondaryAlignmentPeriod?: string;
	// Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
	secondaryPerSeriesAligner?: string;
	// Only present if a preprocessor is selected. Array of labels to group data by.
	secondaryGroupBys?: string[];
	// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
	preprocessor?: PreprocessorType;
}

export const defaultTimeSeriesList = (): TimeSeriesList => ({
	projectName: "",
	crossSeriesReducer: "",
});

// Types of pre-processor available. Defined by the metric.
export enum PreprocessorType {
	None = "none",
	Rate = "rate",
	Delta = "delta",
}

export const defaultPreprocessorType = (): PreprocessorType => (PreprocessorType.None);

// Time Series sub-query properties.
export interface TimeSeriesQuery {
	// GCP project to execute the query against.
	projectName: string;
	// MQL query to be executed.
	query: string;
	// To disable the graphPeriod, it should explictly be set to 'disabled'.
	graphPeriod?: string;
}

export const defaultTimeSeriesQuery = (): TimeSeriesQuery => ({
	projectName: "",
	query: "",
	graphPeriod: "disabled",
});

// SLO sub-query properties.
export interface SLOQuery {
	// GCP project to execute the query against.
	projectName: string;
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	perSeriesAligner?: string;
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	alignmentPeriod?: string;
	// SLO selector.
	selectorName: string;
	// ID for the service the SLO is in.
	serviceId: string;
	// Name for the service the SLO is in.
	serviceName: string;
	// ID for the SLO.
	sloId: string;
	// Name of the SLO.
	sloName: string;
	// SLO goal value.
	goal?: number;
	// Specific lookback period for the SLO.
	lookbackPeriod?: string;
}

export const defaultSLOQuery = (): SLOQuery => ({
	projectName: "",
	selectorName: "",
	serviceId: "",
	serviceName: "",
	sloId: "",
	sloName: "",
});

// PromQL sub-query properties.
export interface PromQLQuery {
	// GCP project to execute the query against.
	projectName: string;
	// PromQL expression/query to be executed.
	expr: string;
	// PromQL min step
	step: string;
}

export const defaultPromQLQuery = (): PromQLQuery => ({
	projectName: "",
	expr: "",
	step: "",
});

// Defines the supported queryTypes.
export enum QueryType {
	TIMESERIESLIST = "timeSeriesList",
	TIMESERIESQUERY = "timeSeriesQuery",
	SLO = "slo",
	ANNOTATION = "annotation",
	PROMQL = "promQL",
}

export const defaultQueryType = (): QueryType => (QueryType.TIMESERIESLIST);

// @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
export interface MetricQuery {
	// GCP project to execute the query against.
	projectName: string;
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	perSeriesAligner?: string;
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	alignmentPeriod?: string;
	// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
	aliasBy?: string;
	editorMode: string;
	metricType: string;
	// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	crossSeriesReducer: string;
	// Array of labels to group data by.
	groupBys?: string[];
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	filters?: string[];
	metricKind?: MetricKind;
	valueType?: string;
	view?: string;
	// MQL query to be executed.
	query: string;
	// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
	preprocessor?: PreprocessorType;
	// To disable the graphPeriod, it should explictly be set to 'disabled'.
	graphPeriod?: string;
}

export const defaultMetricQuery = (): MetricQuery => ({
	projectName: "",
	editorMode: "",
	metricType: "",
	crossSeriesReducer: "",
	query: "",
	graphPeriod: "disabled",
});

export enum MetricKind {
	METRICKINDUNSPECIFIED = "METRIC_KIND_UNSPECIFIED",
	GAUGE = "GAUGE",
	DELTA = "DELTA",
	CUMULATIVE = "CUMULATIVE",
}

export const defaultMetricKind = (): MetricKind => (MetricKind.METRICKINDUNSPECIFIED);

export enum ValueTypes {
	VALUETYPEUNSPECIFIED = "VALUE_TYPE_UNSPECIFIED",
	BOOL = "BOOL",
	INT64 = "INT64",
	DOUBLE = "DOUBLE",
	STRING = "STRING",
	DISTRIBUTION = "DISTRIBUTION",
	MONEY = "MONEY",
}

export const defaultValueTypes = (): ValueTypes => (ValueTypes.VALUETYPEUNSPECIFIED);

export enum AlignmentTypes {
	ALIGNDELTA = "ALIGN_DELTA",
	ALIGNRATE = "ALIGN_RATE",
	ALIGNINTERPOLATE = "ALIGN_INTERPOLATE",
	ALIGNNEXTOLDER = "ALIGN_NEXT_OLDER",
	ALIGNMIN = "ALIGN_MIN",
	ALIGNMAX = "ALIGN_MAX",
	ALIGNMEAN = "ALIGN_MEAN",
	ALIGNCOUNT = "ALIGN_COUNT",
	ALIGNSUM = "ALIGN_SUM",
	ALIGNSTDDEV = "ALIGN_STDDEV",
	ALIGNCOUNTTRUE = "ALIGN_COUNT_TRUE",
	ALIGNCOUNTFALSE = "ALIGN_COUNT_FALSE",
	ALIGNFRACTIONTRUE = "ALIGN_FRACTION_TRUE",
	ALIGNPERCENTILE99 = "ALIGN_PERCENTILE_99",
	ALIGNPERCENTILE95 = "ALIGN_PERCENTILE_95",
	ALIGNPERCENTILE50 = "ALIGN_PERCENTILE_50",
	ALIGNPERCENTILE05 = "ALIGN_PERCENTILE_05",
	ALIGNPERCENTCHANGE = "ALIGN_PERCENT_CHANGE",
	ALIGNNONE = "ALIGN_NONE",
}

export const defaultAlignmentTypes = (): AlignmentTypes => (AlignmentTypes.ALIGNDELTA);

// @deprecated Use TimeSeriesList instead. Legacy annotation query properties for migration purposes.
export interface LegacyCloudMonitoringAnnotationQuery {
	// GCP project to execute the query against.
	projectName: string;
	metricType: string;
	// Query refId.
	refId: string;
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	filters: string[];
	metricKind: MetricKind;
	valueType: string;
	// Annotation title.
	title: string;
	// Annotation text.
	text: string;
}

export const defaultLegacyCloudMonitoringAnnotationQuery = (): LegacyCloudMonitoringAnnotationQuery => ({
	projectName: "",
	metricType: "",
	refId: "",
	filters: [],
	metricKind: MetricKind.METRICKINDUNSPECIFIED,
	valueType: "",
	title: "",
	text: "",
});

// Query filter representation.
export interface Filter {
	// Filter key.
	key: string;
	// Filter operator.
	operator: string;
	// Filter value.
	value: string;
	// Filter condition.
	condition?: string;
}

export const defaultFilter = (): Filter => ({
	key: "",
	operator: "",
	value: "",
});

export enum MetricFindQueryTypes {
	Projects = "projects",
	Services = "services",
	DefaultProject = "defaultProject",
	MetricTypes = "metricTypes",
	LabelKeys = "labelKeys",
	LabelValues = "labelValues",
	ResourceTypes = "resourceTypes",
	Aggregations = "aggregations",
	Aligners = "aligners",
	AlignmentPeriods = "alignmentPeriods",
	Selectors = "selectors",
	SLOServices = "sloServices",
	SLO = "slo",
}

export const defaultMetricFindQueryTypes = (): MetricFindQueryTypes => (MetricFindQueryTypes.Projects);

