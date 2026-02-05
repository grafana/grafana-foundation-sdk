"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultDataquery = exports.defaultPromQueryFormat = exports.PromQueryFormat = exports.defaultQueryEditorMode = exports.QueryEditorMode = void 0;
var QueryEditorMode;
(function (QueryEditorMode) {
    QueryEditorMode["Code"] = "code";
    QueryEditorMode["Builder"] = "builder";
})(QueryEditorMode || (exports.QueryEditorMode = QueryEditorMode = {}));
const defaultQueryEditorMode = () => (QueryEditorMode.Code);
exports.defaultQueryEditorMode = defaultQueryEditorMode;
var PromQueryFormat;
(function (PromQueryFormat) {
    PromQueryFormat["TimeSeries"] = "time_series";
    PromQueryFormat["Table"] = "table";
    PromQueryFormat["Heatmap"] = "heatmap";
})(PromQueryFormat || (exports.PromQueryFormat = PromQueryFormat = {}));
const defaultPromQueryFormat = () => (PromQueryFormat.TimeSeries);
exports.defaultPromQueryFormat = defaultPromQueryFormat;
const defaultDataquery = () => ({
    expr: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultDataquery = defaultDataquery;
//# sourceMappingURL=types.gen.js.map