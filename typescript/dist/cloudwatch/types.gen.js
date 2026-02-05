"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultCloudWatchQuery = exports.defaultCloudWatchAnnotationQuery = exports.defaultLogGroup = exports.defaultCloudWatchLogsQuery = exports.defaultLogsQueryLanguage = exports.LogsQueryLanguage = exports.defaultQueryEditorOperatorValueType = exports.defaultQueryEditorOperatorType = exports.defaultQueryEditorOperator = exports.defaultQueryEditorOperatorExpression = exports.defaultQueryEditorGroupByExpression = exports.defaultQueryEditorExpression = exports.defaultQueryEditorArrayExpression = exports.defaultQueryEditorPropertyType = exports.QueryEditorPropertyType = exports.defaultQueryEditorProperty = exports.defaultQueryEditorPropertyExpression = exports.defaultQueryEditorFunctionParameterExpression = exports.defaultQueryEditorExpressionType = exports.QueryEditorExpressionType = exports.defaultQueryEditorFunctionExpression = exports.defaultSQLExpression = exports.defaultMetricEditorMode = exports.MetricEditorMode = exports.defaultMetricQueryType = exports.MetricQueryType = exports.defaultCloudWatchQueryMode = exports.CloudWatchQueryMode = exports.defaultCloudWatchMetricsQuery = exports.defaultDimensions = exports.defaultMetricStat = void 0;
const defaultMetricStat = () => ({
    region: "",
    namespace: "",
});
exports.defaultMetricStat = defaultMetricStat;
const defaultDimensions = () => ({});
exports.defaultDimensions = defaultDimensions;
const defaultCloudWatchMetricsQuery = () => ({
    queryMode: CloudWatchQueryMode.Metrics,
    id: "",
    region: "",
    namespace: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultCloudWatchMetricsQuery = defaultCloudWatchMetricsQuery;
var CloudWatchQueryMode;
(function (CloudWatchQueryMode) {
    CloudWatchQueryMode["Metrics"] = "Metrics";
    CloudWatchQueryMode["Logs"] = "Logs";
    CloudWatchQueryMode["Annotations"] = "Annotations";
})(CloudWatchQueryMode || (exports.CloudWatchQueryMode = CloudWatchQueryMode = {}));
const defaultCloudWatchQueryMode = () => (CloudWatchQueryMode.Metrics);
exports.defaultCloudWatchQueryMode = defaultCloudWatchQueryMode;
var MetricQueryType;
(function (MetricQueryType) {
    MetricQueryType[MetricQueryType["Search"] = 0] = "Search";
    MetricQueryType[MetricQueryType["Insights"] = 1] = "Insights";
})(MetricQueryType || (exports.MetricQueryType = MetricQueryType = {}));
const defaultMetricQueryType = () => (MetricQueryType.Search);
exports.defaultMetricQueryType = defaultMetricQueryType;
var MetricEditorMode;
(function (MetricEditorMode) {
    MetricEditorMode[MetricEditorMode["Builder"] = 0] = "Builder";
    MetricEditorMode[MetricEditorMode["Code"] = 1] = "Code";
})(MetricEditorMode || (exports.MetricEditorMode = MetricEditorMode = {}));
const defaultMetricEditorMode = () => (MetricEditorMode.Builder);
exports.defaultMetricEditorMode = defaultMetricEditorMode;
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
const defaultQueryEditorPropertyExpression = () => ({
    type: QueryEditorExpressionType.Property,
    property: (0, exports.defaultQueryEditorProperty)(),
});
exports.defaultQueryEditorPropertyExpression = defaultQueryEditorPropertyExpression;
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
const defaultQueryEditorArrayExpression = () => ({
    type: "and",
    expressions: [],
});
exports.defaultQueryEditorArrayExpression = defaultQueryEditorArrayExpression;
const defaultQueryEditorExpression = () => ((0, exports.defaultQueryEditorArrayExpression)());
exports.defaultQueryEditorExpression = defaultQueryEditorExpression;
const defaultQueryEditorGroupByExpression = () => ({
    type: QueryEditorExpressionType.GroupBy,
    property: (0, exports.defaultQueryEditorProperty)(),
});
exports.defaultQueryEditorGroupByExpression = defaultQueryEditorGroupByExpression;
const defaultQueryEditorOperatorExpression = () => ({
    type: QueryEditorExpressionType.Operator,
    property: (0, exports.defaultQueryEditorProperty)(),
    operator: (0, exports.defaultQueryEditorOperator)(),
});
exports.defaultQueryEditorOperatorExpression = defaultQueryEditorOperatorExpression;
const defaultQueryEditorOperator = () => ({});
exports.defaultQueryEditorOperator = defaultQueryEditorOperator;
const defaultQueryEditorOperatorType = () => ("");
exports.defaultQueryEditorOperatorType = defaultQueryEditorOperatorType;
const defaultQueryEditorOperatorValueType = () => ((0, exports.defaultQueryEditorOperatorType)());
exports.defaultQueryEditorOperatorValueType = defaultQueryEditorOperatorValueType;
var LogsQueryLanguage;
(function (LogsQueryLanguage) {
    LogsQueryLanguage["CWLI"] = "CWLI";
    LogsQueryLanguage["SQL"] = "SQL";
    LogsQueryLanguage["PPL"] = "PPL";
})(LogsQueryLanguage || (exports.LogsQueryLanguage = LogsQueryLanguage = {}));
const defaultLogsQueryLanguage = () => (LogsQueryLanguage.CWLI);
exports.defaultLogsQueryLanguage = defaultLogsQueryLanguage;
const defaultCloudWatchLogsQuery = () => ({
    queryMode: CloudWatchQueryMode.Logs,
    id: "",
    region: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultCloudWatchLogsQuery = defaultCloudWatchLogsQuery;
const defaultLogGroup = () => ({
    arn: "",
    name: "",
});
exports.defaultLogGroup = defaultLogGroup;
const defaultCloudWatchAnnotationQuery = () => ({
    queryMode: CloudWatchQueryMode.Annotations,
    region: "",
    namespace: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultCloudWatchAnnotationQuery = defaultCloudWatchAnnotationQuery;
const defaultCloudWatchQuery = () => ((0, exports.defaultCloudWatchMetricsQuery)());
exports.defaultCloudWatchQuery = defaultCloudWatchQuery;
//# sourceMappingURL=types.gen.js.map