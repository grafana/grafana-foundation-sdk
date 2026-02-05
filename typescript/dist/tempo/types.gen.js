"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultSearchStreamingState = exports.SearchStreamingState = exports.defaultTempoQueryType = exports.TempoQueryType = exports.defaultMetricsQueryType = exports.MetricsQueryType = exports.defaultSearchTableType = exports.SearchTableType = exports.defaultTraceqlSearchScope = exports.TraceqlSearchScope = exports.defaultTraceqlFilter = exports.defaultTempoQuery = void 0;
const defaultTempoQuery = () => ({
    filters: [],
    _implementsDataqueryVariant: () => { },
});
exports.defaultTempoQuery = defaultTempoQuery;
const defaultTraceqlFilter = () => ({
    id: "",
});
exports.defaultTraceqlFilter = defaultTraceqlFilter;
// static fields are pre-set in the UI, dynamic fields are added by the user
var TraceqlSearchScope;
(function (TraceqlSearchScope) {
    TraceqlSearchScope["Intrinsic"] = "intrinsic";
    TraceqlSearchScope["Unscoped"] = "unscoped";
    TraceqlSearchScope["Event"] = "event";
    TraceqlSearchScope["Instrumentation"] = "instrumentation";
    TraceqlSearchScope["Link"] = "link";
    TraceqlSearchScope["Resource"] = "resource";
    TraceqlSearchScope["Span"] = "span";
})(TraceqlSearchScope || (exports.TraceqlSearchScope = TraceqlSearchScope = {}));
const defaultTraceqlSearchScope = () => (TraceqlSearchScope.Intrinsic);
exports.defaultTraceqlSearchScope = defaultTraceqlSearchScope;
// The type of the table that is used to display the search results
var SearchTableType;
(function (SearchTableType) {
    SearchTableType["Traces"] = "traces";
    SearchTableType["Spans"] = "spans";
    SearchTableType["Raw"] = "raw";
})(SearchTableType || (exports.SearchTableType = SearchTableType = {}));
const defaultSearchTableType = () => (SearchTableType.Traces);
exports.defaultSearchTableType = defaultSearchTableType;
var MetricsQueryType;
(function (MetricsQueryType) {
    MetricsQueryType["Range"] = "range";
    MetricsQueryType["Instant"] = "instant";
})(MetricsQueryType || (exports.MetricsQueryType = MetricsQueryType = {}));
const defaultMetricsQueryType = () => (MetricsQueryType.Range);
exports.defaultMetricsQueryType = defaultMetricsQueryType;
var TempoQueryType;
(function (TempoQueryType) {
    TempoQueryType["Traceql"] = "traceql";
    TempoQueryType["TraceqlSearch"] = "traceqlSearch";
    TempoQueryType["ServiceMap"] = "serviceMap";
    TempoQueryType["Upload"] = "upload";
    TempoQueryType["NativeSearch"] = "nativeSearch";
    TempoQueryType["TraceId"] = "traceId";
    TempoQueryType["Clear"] = "clear";
})(TempoQueryType || (exports.TempoQueryType = TempoQueryType = {}));
const defaultTempoQueryType = () => (TempoQueryType.Traceql);
exports.defaultTempoQueryType = defaultTempoQueryType;
// The state of the TraceQL streaming search query
var SearchStreamingState;
(function (SearchStreamingState) {
    SearchStreamingState["Pending"] = "pending";
    SearchStreamingState["Streaming"] = "streaming";
    SearchStreamingState["Done"] = "done";
    SearchStreamingState["Error"] = "error";
})(SearchStreamingState || (exports.SearchStreamingState = SearchStreamingState = {}));
const defaultSearchStreamingState = () => (SearchStreamingState.Pending);
exports.defaultSearchStreamingState = defaultSearchStreamingState;
//# sourceMappingURL=types.gen.js.map