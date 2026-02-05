"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultMetricFindQueryTypes = exports.MetricFindQueryTypes = exports.defaultFilter = exports.defaultLegacyCloudMonitoringAnnotationQuery = exports.defaultAlignmentTypes = exports.AlignmentTypes = exports.defaultValueTypes = exports.ValueTypes = exports.defaultMetricKind = exports.MetricKind = exports.defaultMetricQuery = exports.defaultQueryType = exports.QueryType = exports.defaultPromQLQuery = exports.defaultSLOQuery = exports.defaultTimeSeriesQuery = exports.defaultPreprocessorType = exports.PreprocessorType = exports.defaultTimeSeriesList = exports.defaultCloudMonitoringQuery = void 0;
const defaultCloudMonitoringQuery = () => ({
    _implementsDataqueryVariant: () => { },
});
exports.defaultCloudMonitoringQuery = defaultCloudMonitoringQuery;
const defaultTimeSeriesList = () => ({
    projectName: "",
    crossSeriesReducer: "",
});
exports.defaultTimeSeriesList = defaultTimeSeriesList;
// Types of pre-processor available. Defined by the metric.
var PreprocessorType;
(function (PreprocessorType) {
    PreprocessorType["None"] = "none";
    PreprocessorType["Rate"] = "rate";
    PreprocessorType["Delta"] = "delta";
})(PreprocessorType || (exports.PreprocessorType = PreprocessorType = {}));
const defaultPreprocessorType = () => (PreprocessorType.None);
exports.defaultPreprocessorType = defaultPreprocessorType;
const defaultTimeSeriesQuery = () => ({
    projectName: "",
    query: "",
    graphPeriod: "disabled",
});
exports.defaultTimeSeriesQuery = defaultTimeSeriesQuery;
const defaultSLOQuery = () => ({
    projectName: "",
    selectorName: "",
    serviceId: "",
    serviceName: "",
    sloId: "",
    sloName: "",
});
exports.defaultSLOQuery = defaultSLOQuery;
const defaultPromQLQuery = () => ({
    projectName: "",
    expr: "",
    step: "",
});
exports.defaultPromQLQuery = defaultPromQLQuery;
// Defines the supported queryTypes.
var QueryType;
(function (QueryType) {
    QueryType["TIMESERIESLIST"] = "timeSeriesList";
    QueryType["TIMESERIESQUERY"] = "timeSeriesQuery";
    QueryType["SLO"] = "slo";
    QueryType["ANNOTATION"] = "annotation";
    QueryType["PROMQL"] = "promQL";
})(QueryType || (exports.QueryType = QueryType = {}));
const defaultQueryType = () => (QueryType.TIMESERIESLIST);
exports.defaultQueryType = defaultQueryType;
const defaultMetricQuery = () => ({
    projectName: "",
    editorMode: "",
    metricType: "",
    crossSeriesReducer: "",
    query: "",
    graphPeriod: "disabled",
});
exports.defaultMetricQuery = defaultMetricQuery;
var MetricKind;
(function (MetricKind) {
    MetricKind["METRICKINDUNSPECIFIED"] = "METRIC_KIND_UNSPECIFIED";
    MetricKind["GAUGE"] = "GAUGE";
    MetricKind["DELTA"] = "DELTA";
    MetricKind["CUMULATIVE"] = "CUMULATIVE";
})(MetricKind || (exports.MetricKind = MetricKind = {}));
const defaultMetricKind = () => (MetricKind.METRICKINDUNSPECIFIED);
exports.defaultMetricKind = defaultMetricKind;
var ValueTypes;
(function (ValueTypes) {
    ValueTypes["VALUETYPEUNSPECIFIED"] = "VALUE_TYPE_UNSPECIFIED";
    ValueTypes["BOOL"] = "BOOL";
    ValueTypes["INT64"] = "INT64";
    ValueTypes["DOUBLE"] = "DOUBLE";
    ValueTypes["STRING"] = "STRING";
    ValueTypes["DISTRIBUTION"] = "DISTRIBUTION";
    ValueTypes["MONEY"] = "MONEY";
})(ValueTypes || (exports.ValueTypes = ValueTypes = {}));
const defaultValueTypes = () => (ValueTypes.VALUETYPEUNSPECIFIED);
exports.defaultValueTypes = defaultValueTypes;
var AlignmentTypes;
(function (AlignmentTypes) {
    AlignmentTypes["ALIGNDELTA"] = "ALIGN_DELTA";
    AlignmentTypes["ALIGNRATE"] = "ALIGN_RATE";
    AlignmentTypes["ALIGNINTERPOLATE"] = "ALIGN_INTERPOLATE";
    AlignmentTypes["ALIGNNEXTOLDER"] = "ALIGN_NEXT_OLDER";
    AlignmentTypes["ALIGNMIN"] = "ALIGN_MIN";
    AlignmentTypes["ALIGNMAX"] = "ALIGN_MAX";
    AlignmentTypes["ALIGNMEAN"] = "ALIGN_MEAN";
    AlignmentTypes["ALIGNCOUNT"] = "ALIGN_COUNT";
    AlignmentTypes["ALIGNSUM"] = "ALIGN_SUM";
    AlignmentTypes["ALIGNSTDDEV"] = "ALIGN_STDDEV";
    AlignmentTypes["ALIGNCOUNTTRUE"] = "ALIGN_COUNT_TRUE";
    AlignmentTypes["ALIGNCOUNTFALSE"] = "ALIGN_COUNT_FALSE";
    AlignmentTypes["ALIGNFRACTIONTRUE"] = "ALIGN_FRACTION_TRUE";
    AlignmentTypes["ALIGNPERCENTILE99"] = "ALIGN_PERCENTILE_99";
    AlignmentTypes["ALIGNPERCENTILE95"] = "ALIGN_PERCENTILE_95";
    AlignmentTypes["ALIGNPERCENTILE50"] = "ALIGN_PERCENTILE_50";
    AlignmentTypes["ALIGNPERCENTILE05"] = "ALIGN_PERCENTILE_05";
    AlignmentTypes["ALIGNPERCENTCHANGE"] = "ALIGN_PERCENT_CHANGE";
    AlignmentTypes["ALIGNNONE"] = "ALIGN_NONE";
})(AlignmentTypes || (exports.AlignmentTypes = AlignmentTypes = {}));
const defaultAlignmentTypes = () => (AlignmentTypes.ALIGNDELTA);
exports.defaultAlignmentTypes = defaultAlignmentTypes;
const defaultLegacyCloudMonitoringAnnotationQuery = () => ({
    projectName: "",
    metricType: "",
    refId: "",
    filters: [],
    metricKind: MetricKind.METRICKINDUNSPECIFIED,
    valueType: "",
    title: "",
    text: "",
});
exports.defaultLegacyCloudMonitoringAnnotationQuery = defaultLegacyCloudMonitoringAnnotationQuery;
const defaultFilter = () => ({
    key: "",
    operator: "",
    value: "",
});
exports.defaultFilter = defaultFilter;
var MetricFindQueryTypes;
(function (MetricFindQueryTypes) {
    MetricFindQueryTypes["Projects"] = "projects";
    MetricFindQueryTypes["Services"] = "services";
    MetricFindQueryTypes["DefaultProject"] = "defaultProject";
    MetricFindQueryTypes["MetricTypes"] = "metricTypes";
    MetricFindQueryTypes["LabelKeys"] = "labelKeys";
    MetricFindQueryTypes["LabelValues"] = "labelValues";
    MetricFindQueryTypes["ResourceTypes"] = "resourceTypes";
    MetricFindQueryTypes["Aggregations"] = "aggregations";
    MetricFindQueryTypes["Aligners"] = "aligners";
    MetricFindQueryTypes["AlignmentPeriods"] = "alignmentPeriods";
    MetricFindQueryTypes["Selectors"] = "selectors";
    MetricFindQueryTypes["SLOServices"] = "sloServices";
    MetricFindQueryTypes["SLO"] = "slo";
})(MetricFindQueryTypes || (exports.MetricFindQueryTypes = MetricFindQueryTypes = {}));
const defaultMetricFindQueryTypes = () => (MetricFindQueryTypes.Projects);
exports.defaultMetricFindQueryTypes = defaultMetricFindQueryTypes;
//# sourceMappingURL=types.gen.js.map