"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultBaseGrafanaTemplateVariableQuery = exports.defaultGrafanaTemplateVariableQueryType = exports.GrafanaTemplateVariableQueryType = exports.defaultAzureQueryType = exports.AzureQueryType = exports.defaultUnknownQuery = exports.defaultWorkspacesQuery = exports.defaultMetricNamesQuery = exports.defaultMetricDefinitionsQuery = exports.defaultMetricNamespaceQuery = exports.defaultResourceNamesQuery = exports.defaultResourceGroupsQuery = exports.defaultSubscriptionsQuery = exports.defaultAppInsightsGroupByQuery = exports.defaultAppInsightsMetricNameQuery = exports.defaultGrafanaTemplateVariableQuery = exports.defaultAzureTracesFilter = exports.defaultAzureTracesQuery = exports.defaultAzureResourceGraphQuery = exports.defaultResultFormat = exports.ResultFormat = exports.defaultAzureLogsQuery = exports.defaultAzureMetricDimension = exports.defaultAzureMonitorResource = exports.defaultAzureMetricQuery = exports.defaultAzureMonitorQuery = void 0;
const defaultAzureMonitorQuery = () => ({
    _implementsDataqueryVariant: () => { },
});
exports.defaultAzureMonitorQuery = defaultAzureMonitorQuery;
const defaultAzureMetricQuery = () => ({});
exports.defaultAzureMetricQuery = defaultAzureMetricQuery;
const defaultAzureMonitorResource = () => ({});
exports.defaultAzureMonitorResource = defaultAzureMonitorResource;
const defaultAzureMetricDimension = () => ({});
exports.defaultAzureMetricDimension = defaultAzureMetricDimension;
const defaultAzureLogsQuery = () => ({});
exports.defaultAzureLogsQuery = defaultAzureLogsQuery;
var ResultFormat;
(function (ResultFormat) {
    ResultFormat["Table"] = "table";
    ResultFormat["TimeSeries"] = "time_series";
    ResultFormat["Trace"] = "trace";
    ResultFormat["Logs"] = "logs";
})(ResultFormat || (exports.ResultFormat = ResultFormat = {}));
const defaultResultFormat = () => (ResultFormat.Table);
exports.defaultResultFormat = defaultResultFormat;
const defaultAzureResourceGraphQuery = () => ({});
exports.defaultAzureResourceGraphQuery = defaultAzureResourceGraphQuery;
const defaultAzureTracesQuery = () => ({});
exports.defaultAzureTracesQuery = defaultAzureTracesQuery;
const defaultAzureTracesFilter = () => ({
    property: "",
    operation: "",
    filters: [],
});
exports.defaultAzureTracesFilter = defaultAzureTracesFilter;
const defaultGrafanaTemplateVariableQuery = () => ((0, exports.defaultAppInsightsMetricNameQuery)());
exports.defaultGrafanaTemplateVariableQuery = defaultGrafanaTemplateVariableQuery;
const defaultAppInsightsMetricNameQuery = () => ({
    kind: "AppInsightsMetricNameQuery",
});
exports.defaultAppInsightsMetricNameQuery = defaultAppInsightsMetricNameQuery;
const defaultAppInsightsGroupByQuery = () => ({
    kind: "AppInsightsGroupByQuery",
    metricName: "",
});
exports.defaultAppInsightsGroupByQuery = defaultAppInsightsGroupByQuery;
const defaultSubscriptionsQuery = () => ({
    kind: "SubscriptionsQuery",
});
exports.defaultSubscriptionsQuery = defaultSubscriptionsQuery;
const defaultResourceGroupsQuery = () => ({
    kind: "ResourceGroupsQuery",
    subscription: "",
});
exports.defaultResourceGroupsQuery = defaultResourceGroupsQuery;
const defaultResourceNamesQuery = () => ({
    kind: "ResourceNamesQuery",
    subscription: "",
    resourceGroup: "",
    metricNamespace: "",
});
exports.defaultResourceNamesQuery = defaultResourceNamesQuery;
const defaultMetricNamespaceQuery = () => ({
    kind: "MetricNamespaceQuery",
    subscription: "",
    resourceGroup: "",
});
exports.defaultMetricNamespaceQuery = defaultMetricNamespaceQuery;
const defaultMetricDefinitionsQuery = () => ({
    kind: "MetricDefinitionsQuery",
    subscription: "",
    resourceGroup: "",
});
exports.defaultMetricDefinitionsQuery = defaultMetricDefinitionsQuery;
const defaultMetricNamesQuery = () => ({
    kind: "MetricNamesQuery",
    subscription: "",
    resourceGroup: "",
    resourceName: "",
    metricNamespace: "",
});
exports.defaultMetricNamesQuery = defaultMetricNamesQuery;
const defaultWorkspacesQuery = () => ({
    kind: "WorkspacesQuery",
    subscription: "",
});
exports.defaultWorkspacesQuery = defaultWorkspacesQuery;
const defaultUnknownQuery = () => ({
    kind: "UnknownQuery",
});
exports.defaultUnknownQuery = defaultUnknownQuery;
// Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated
var AzureQueryType;
(function (AzureQueryType) {
    AzureQueryType["AzureMonitor"] = "Azure Monitor";
    AzureQueryType["LogAnalytics"] = "Azure Log Analytics";
    AzureQueryType["AzureResourceGraph"] = "Azure Resource Graph";
    AzureQueryType["AzureTraces"] = "Azure Traces";
    AzureQueryType["SubscriptionsQuery"] = "Azure Subscriptions";
    AzureQueryType["ResourceGroupsQuery"] = "Azure Resource Groups";
    AzureQueryType["NamespacesQuery"] = "Azure Namespaces";
    AzureQueryType["ResourceNamesQuery"] = "Azure Resource Names";
    AzureQueryType["MetricNamesQuery"] = "Azure Metric Names";
    AzureQueryType["WorkspacesQuery"] = "Azure Workspaces";
    AzureQueryType["LocationsQuery"] = "Azure Regions";
    AzureQueryType["GrafanaTemplateVariableFn"] = "Grafana Template Variable Function";
    AzureQueryType["TraceExemplar"] = "traceql";
    AzureQueryType["CustomNamespacesQuery"] = "Azure Custom Namespaces";
    AzureQueryType["CustomMetricNamesQuery"] = "Azure Custom Metric Names";
})(AzureQueryType || (exports.AzureQueryType = AzureQueryType = {}));
const defaultAzureQueryType = () => (AzureQueryType.AzureMonitor);
exports.defaultAzureQueryType = defaultAzureQueryType;
var GrafanaTemplateVariableQueryType;
(function (GrafanaTemplateVariableQueryType) {
    GrafanaTemplateVariableQueryType["AppInsightsMetricNameQuery"] = "AppInsightsMetricNameQuery";
    GrafanaTemplateVariableQueryType["AppInsightsGroupByQuery"] = "AppInsightsGroupByQuery";
    GrafanaTemplateVariableQueryType["SubscriptionsQuery"] = "SubscriptionsQuery";
    GrafanaTemplateVariableQueryType["ResourceGroupsQuery"] = "ResourceGroupsQuery";
    GrafanaTemplateVariableQueryType["ResourceNamesQuery"] = "ResourceNamesQuery";
    GrafanaTemplateVariableQueryType["MetricNamespaceQuery"] = "MetricNamespaceQuery";
    GrafanaTemplateVariableQueryType["MetricNamesQuery"] = "MetricNamesQuery";
    GrafanaTemplateVariableQueryType["WorkspacesQuery"] = "WorkspacesQuery";
    GrafanaTemplateVariableQueryType["UnknownQuery"] = "UnknownQuery";
})(GrafanaTemplateVariableQueryType || (exports.GrafanaTemplateVariableQueryType = GrafanaTemplateVariableQueryType = {}));
const defaultGrafanaTemplateVariableQueryType = () => (GrafanaTemplateVariableQueryType.AppInsightsMetricNameQuery);
exports.defaultGrafanaTemplateVariableQueryType = defaultGrafanaTemplateVariableQueryType;
const defaultBaseGrafanaTemplateVariableQuery = () => ({});
exports.defaultBaseGrafanaTemplateVariableQuery = defaultBaseGrafanaTemplateVariableQuery;
//# sourceMappingURL=types.gen.js.map