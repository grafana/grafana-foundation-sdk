"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultDataquery = exports.defaultLokiQueryDirection = exports.LokiQueryDirection = exports.defaultSupportingQueryType = exports.SupportingQueryType = exports.defaultLokiQueryType = exports.LokiQueryType = exports.defaultQueryEditorMode = exports.QueryEditorMode = void 0;
var QueryEditorMode;
(function (QueryEditorMode) {
    QueryEditorMode["Code"] = "code";
    QueryEditorMode["Builder"] = "builder";
})(QueryEditorMode || (exports.QueryEditorMode = QueryEditorMode = {}));
const defaultQueryEditorMode = () => (QueryEditorMode.Code);
exports.defaultQueryEditorMode = defaultQueryEditorMode;
var LokiQueryType;
(function (LokiQueryType) {
    LokiQueryType["Range"] = "range";
    LokiQueryType["Instant"] = "instant";
    LokiQueryType["Stream"] = "stream";
})(LokiQueryType || (exports.LokiQueryType = LokiQueryType = {}));
const defaultLokiQueryType = () => (LokiQueryType.Range);
exports.defaultLokiQueryType = defaultLokiQueryType;
var SupportingQueryType;
(function (SupportingQueryType) {
    SupportingQueryType["LogsVolume"] = "logsVolume";
    SupportingQueryType["LogsSample"] = "logsSample";
    SupportingQueryType["DataSample"] = "dataSample";
    SupportingQueryType["InfiniteScroll"] = "infiniteScroll";
})(SupportingQueryType || (exports.SupportingQueryType = SupportingQueryType = {}));
const defaultSupportingQueryType = () => (SupportingQueryType.LogsVolume);
exports.defaultSupportingQueryType = defaultSupportingQueryType;
var LokiQueryDirection;
(function (LokiQueryDirection) {
    LokiQueryDirection["Forward"] = "forward";
    LokiQueryDirection["Backward"] = "backward";
    LokiQueryDirection["Scan"] = "scan";
})(LokiQueryDirection || (exports.LokiQueryDirection = LokiQueryDirection = {}));
const defaultLokiQueryDirection = () => (LokiQueryDirection.Forward);
exports.defaultLokiQueryDirection = defaultLokiQueryDirection;
const defaultDataquery = () => ({
    expr: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultDataquery = defaultDataquery;
//# sourceMappingURL=types.gen.js.map