"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOrderByDirection = exports.OrderByDirection = exports.defaultQueryEditorPropertyExpression = exports.defaultQueryEditorPropertyType = exports.QueryEditorPropertyType = exports.defaultQueryEditorProperty = exports.defaultQueryEditorGroupByExpression = exports.defaultQueryEditorFunctionParameterExpression = exports.defaultQueryEditorExpressionType = exports.QueryEditorExpressionType = exports.defaultQueryEditorFunctionExpression = exports.defaultSQLExpression = exports.defaultEditorMode = exports.EditorMode = exports.defaultQueryPriority = exports.QueryPriority = exports.defaultQueryFormat = exports.QueryFormat = exports.defaultDataquery = void 0;
const defaultDataquery = () => ({
    format: QueryFormat.Timeseries,
    rawSql: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultDataquery = defaultDataquery;
var QueryFormat;
(function (QueryFormat) {
    QueryFormat[QueryFormat["Timeseries"] = 0] = "Timeseries";
    QueryFormat[QueryFormat["Table"] = 1] = "Table";
})(QueryFormat || (exports.QueryFormat = QueryFormat = {}));
const defaultQueryFormat = () => (QueryFormat.Timeseries);
exports.defaultQueryFormat = defaultQueryFormat;
var QueryPriority;
(function (QueryPriority) {
    QueryPriority["Interactive"] = "INTERACTIVE";
    QueryPriority["Batch"] = "BATCH";
})(QueryPriority || (exports.QueryPriority = QueryPriority = {}));
const defaultQueryPriority = () => (QueryPriority.Interactive);
exports.defaultQueryPriority = defaultQueryPriority;
var EditorMode;
(function (EditorMode) {
    EditorMode["Code"] = "code";
    EditorMode["Builder"] = "builder";
})(EditorMode || (exports.EditorMode = EditorMode = {}));
const defaultEditorMode = () => (EditorMode.Code);
exports.defaultEditorMode = defaultEditorMode;
const defaultSQLExpression = () => ({});
exports.defaultSQLExpression = defaultSQLExpression;
const defaultQueryEditorFunctionExpression = () => ({
    type: QueryEditorExpressionType.Function,
});
exports.defaultQueryEditorFunctionExpression = defaultQueryEditorFunctionExpression;
var QueryEditorExpressionType;
(function (QueryEditorExpressionType) {
    QueryEditorExpressionType["Property"] = "property";
    QueryEditorExpressionType["Operator"] = "operator";
    QueryEditorExpressionType["Or"] = "or";
    QueryEditorExpressionType["And"] = "and";
    QueryEditorExpressionType["GroupBy"] = "groupBy";
    QueryEditorExpressionType["Function"] = "function";
    QueryEditorExpressionType["FunctionParameter"] = "functionParameter";
})(QueryEditorExpressionType || (exports.QueryEditorExpressionType = QueryEditorExpressionType = {}));
const defaultQueryEditorExpressionType = () => (QueryEditorExpressionType.Property);
exports.defaultQueryEditorExpressionType = defaultQueryEditorExpressionType;
const defaultQueryEditorFunctionParameterExpression = () => ({
    type: QueryEditorExpressionType.FunctionParameter,
});
exports.defaultQueryEditorFunctionParameterExpression = defaultQueryEditorFunctionParameterExpression;
const defaultQueryEditorGroupByExpression = () => ({
    type: QueryEditorExpressionType.GroupBy,
    property: (0, exports.defaultQueryEditorProperty)(),
});
exports.defaultQueryEditorGroupByExpression = defaultQueryEditorGroupByExpression;
const defaultQueryEditorProperty = () => ({
    type: QueryEditorPropertyType.String,
});
exports.defaultQueryEditorProperty = defaultQueryEditorProperty;
var QueryEditorPropertyType;
(function (QueryEditorPropertyType) {
    QueryEditorPropertyType["String"] = "string";
})(QueryEditorPropertyType || (exports.QueryEditorPropertyType = QueryEditorPropertyType = {}));
const defaultQueryEditorPropertyType = () => (QueryEditorPropertyType.String);
exports.defaultQueryEditorPropertyType = defaultQueryEditorPropertyType;
const defaultQueryEditorPropertyExpression = () => ({
    type: QueryEditorExpressionType.Property,
    property: (0, exports.defaultQueryEditorProperty)(),
});
exports.defaultQueryEditorPropertyExpression = defaultQueryEditorPropertyExpression;
var OrderByDirection;
(function (OrderByDirection) {
    OrderByDirection["ASC"] = "ASC";
    OrderByDirection["DESC"] = "DESC";
})(OrderByDirection || (exports.OrderByDirection = OrderByDirection = {}));
const defaultOrderByDirection = () => (OrderByDirection.ASC);
exports.defaultOrderByDirection = defaultOrderByDirection;
//# sourceMappingURL=types.gen.js.map