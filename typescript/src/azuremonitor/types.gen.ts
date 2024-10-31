// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


export interface AzureMonitorQuery {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// Azure subscription containing the resource(s) to be queried.
	subscription?: string;
	// Subscriptions to be queried via Azure Resource Graph.
	subscriptions?: string[];
	// Azure Monitor Metrics sub-query properties.
	azureMonitor?: AzureMetricQuery;
	// Azure Monitor Logs sub-query properties.
	azureLogAnalytics?: AzureLogsQuery;
	// Azure Resource Graph sub-query properties.
	azureResourceGraph?: AzureResourceGraphQuery;
	// Application Insights Traces sub-query properties.
	azureTraces?: AzureTracesQuery;
	// @deprecated Legacy template variable support.
	grafanaTemplateVariableFn?: GrafanaTemplateVariableQuery;
	// Template variables params. These exist for backwards compatiblity with legacy template variables.
	resourceGroup?: string;
	namespace?: string;
	resource?: string;
	region?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// Used only for exemplar queries from Prometheus
	query?: string;
	_implementsDataqueryVariant(): void;
}

export const defaultAzureMonitorQuery = (): AzureMonitorQuery => ({
	refId: "",
	_implementsDataqueryVariant: () => {},
});

// Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated
export enum AzureQueryType {
	AzureMonitor = "Azure Monitor",
	LogAnalytics = "Azure Log Analytics",
	AzureResourceGraph = "Azure Resource Graph",
	AzureTraces = "Azure Traces",
	SubscriptionsQuery = "Azure Subscriptions",
	ResourceGroupsQuery = "Azure Resource Groups",
	NamespacesQuery = "Azure Namespaces",
	ResourceNamesQuery = "Azure Resource Names",
	MetricNamesQuery = "Azure Metric Names",
	WorkspacesQuery = "Azure Workspaces",
	LocationsQuery = "Azure Regions",
	GrafanaTemplateVariableFn = "Grafana Template Variable Function",
	TraceExemplar = "traceql",
}

export const defaultAzureQueryType = (): AzureQueryType => (AzureQueryType.AzureMonitor);

export interface AzureMetricQuery {
	// Array of resource URIs to be queried.
	resources?: AzureMonitorResource[];
	// metricNamespace is used as the resource type (or resource namespace).
	// It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
	// Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
	metricNamespace?: string;
	// Used as the value for the metricNamespace property when it's different from the resource namespace.
	customNamespace?: string;
	// The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
	metricName?: string;
	// The Azure region containing the resource(s).
	region?: string;
	// The granularity of data points to be queried. Defaults to auto.
	timeGrain?: string;
	// The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
	aggregation?: string;
	// Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
	dimensionFilters?: AzureMetricDimension[];
	// Maximum number of records to return. Defaults to 10.
	top?: string;
	// Time grains that are supported by the metric.
	allowedTimeGrainsMs?: number[];
	// Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
	alias?: string;
	// @deprecated
	timeGrainUnit?: string;
	// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
	dimension?: string;
	// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
	dimensionFilter?: string;
	// @deprecated Use metricNamespace instead
	metricDefinition?: string;
	// @deprecated Use resourceGroup, resourceName and metricNamespace instead
	resourceUri?: string;
	// @deprecated Use resources instead
	resourceGroup?: string;
	// @deprecated Use resources instead
	resourceName?: string;
}

export const defaultAzureMetricQuery = (): AzureMetricQuery => ({
});

// Azure Monitor Logs sub-query properties
export interface AzureLogsQuery {
	// KQL query to be executed.
	query?: string;
	// Specifies the format results should be returned as.
	resultFormat?: ResultFormat;
	// Array of resource URIs to be queried.
	resources?: string[];
	// If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
	dashboardTime?: boolean;
	// If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
	timeColumn?: string;
	// If set to true the query will be run as a basic logs query
	basicLogsQuery?: boolean;
	// Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
	workspace?: string;
	// @deprecated Use resources instead
	resource?: string;
	// @deprecated Use dashboardTime instead
	intersectTime?: boolean;
}

export const defaultAzureLogsQuery = (): AzureLogsQuery => ({
});

// Application Insights Traces sub-query properties
export interface AzureTracesQuery {
	// Specifies the format results should be returned as.
	resultFormat?: ResultFormat;
	// Array of resource URIs to be queried.
	resources?: string[];
	// Operation ID. Used only for Traces queries.
	operationId?: string;
	// Types of events to filter by.
	traceTypes?: string[];
	// Filters for property values.
	filters?: AzureTracesFilter[];
	// KQL query to be executed.
	query?: string;
}

export const defaultAzureTracesQuery = (): AzureTracesQuery => ({
});

export interface AzureTracesFilter {
	// Property name, auto-populated based on available traces.
	property: string;
	// Comparison operator to use. Either equals or not equals.
	operation: string;
	// Values to filter by.
	filters: string[];
}

export const defaultAzureTracesFilter = (): AzureTracesFilter => ({
	property: "",
	operation: "",
	filters: [],
});

export enum ResultFormat {
	Table = "table",
	TimeSeries = "time_series",
	Trace = "trace",
	Logs = "logs",
}

export const defaultResultFormat = (): ResultFormat => (ResultFormat.Table);

export interface AzureResourceGraphQuery {
	// Azure Resource Graph KQL query to be executed.
	query?: string;
	// Specifies the format results should be returned as. Defaults to table.
	resultFormat?: string;
}

export const defaultAzureResourceGraphQuery = (): AzureResourceGraphQuery => ({
});

export interface AzureMonitorResource {
	subscription?: string;
	resourceGroup?: string;
	resourceName?: string;
	metricNamespace?: string;
	region?: string;
}

export const defaultAzureMonitorResource = (): AzureMonitorResource => ({
});

export interface AzureMetricDimension {
	// Name of Dimension to be filtered on.
	dimension?: string;
	// String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
	operator?: string;
	// Values to match with the filter.
	filters?: string[];
	// @deprecated filter is deprecated in favour of filters to support multiselect.
	filter?: string;
}

export const defaultAzureMetricDimension = (): AzureMetricDimension => ({
});

export enum GrafanaTemplateVariableQueryType {
	AppInsightsMetricNameQuery = "AppInsightsMetricNameQuery",
	AppInsightsGroupByQuery = "AppInsightsGroupByQuery",
	SubscriptionsQuery = "SubscriptionsQuery",
	ResourceGroupsQuery = "ResourceGroupsQuery",
	ResourceNamesQuery = "ResourceNamesQuery",
	MetricNamespaceQuery = "MetricNamespaceQuery",
	MetricNamesQuery = "MetricNamesQuery",
	WorkspacesQuery = "WorkspacesQuery",
	UnknownQuery = "UnknownQuery",
}

export const defaultGrafanaTemplateVariableQueryType = (): GrafanaTemplateVariableQueryType => (GrafanaTemplateVariableQueryType.AppInsightsMetricNameQuery);

export interface BaseGrafanaTemplateVariableQuery {
	rawQuery?: string;
}

export const defaultBaseGrafanaTemplateVariableQuery = (): BaseGrafanaTemplateVariableQuery => ({
});

export interface UnknownQuery {
	rawQuery?: string;
	kind: "UnknownQuery";
}

export const defaultUnknownQuery = (): UnknownQuery => ({
	kind: "UnknownQuery",
});

export interface AppInsightsMetricNameQuery {
	rawQuery?: string;
	kind: "AppInsightsMetricNameQuery";
}

export const defaultAppInsightsMetricNameQuery = (): AppInsightsMetricNameQuery => ({
	kind: "AppInsightsMetricNameQuery",
});

export interface AppInsightsGroupByQuery {
	rawQuery?: string;
	kind: "AppInsightsGroupByQuery";
	metricName: string;
}

export const defaultAppInsightsGroupByQuery = (): AppInsightsGroupByQuery => ({
	kind: "AppInsightsGroupByQuery",
	metricName: "",
});

export interface SubscriptionsQuery {
	rawQuery?: string;
	kind: "SubscriptionsQuery";
}

export const defaultSubscriptionsQuery = (): SubscriptionsQuery => ({
	kind: "SubscriptionsQuery",
});

export interface ResourceGroupsQuery {
	rawQuery?: string;
	kind: "ResourceGroupsQuery";
	subscription: string;
}

export const defaultResourceGroupsQuery = (): ResourceGroupsQuery => ({
	kind: "ResourceGroupsQuery",
	subscription: "",
});

export interface ResourceNamesQuery {
	rawQuery?: string;
	kind: "ResourceNamesQuery";
	subscription: string;
	resourceGroup: string;
	metricNamespace: string;
}

export const defaultResourceNamesQuery = (): ResourceNamesQuery => ({
	kind: "ResourceNamesQuery",
	subscription: "",
	resourceGroup: "",
	metricNamespace: "",
});

export interface MetricNamespaceQuery {
	rawQuery?: string;
	kind: "MetricNamespaceQuery";
	subscription: string;
	resourceGroup: string;
	metricNamespace?: string;
	resourceName?: string;
}

export const defaultMetricNamespaceQuery = (): MetricNamespaceQuery => ({
	kind: "MetricNamespaceQuery",
	subscription: "",
	resourceGroup: "",
});

// @deprecated Use MetricNamespaceQuery instead
export interface MetricDefinitionsQuery {
	rawQuery?: string;
	kind: "MetricDefinitionsQuery";
	subscription: string;
	resourceGroup: string;
	metricNamespace?: string;
	resourceName?: string;
}

export const defaultMetricDefinitionsQuery = (): MetricDefinitionsQuery => ({
	kind: "MetricDefinitionsQuery",
	subscription: "",
	resourceGroup: "",
});

export interface MetricNamesQuery {
	rawQuery?: string;
	kind: "MetricNamesQuery";
	subscription: string;
	resourceGroup: string;
	resourceName: string;
	metricNamespace: string;
}

export const defaultMetricNamesQuery = (): MetricNamesQuery => ({
	kind: "MetricNamesQuery",
	subscription: "",
	resourceGroup: "",
	resourceName: "",
	metricNamespace: "",
});

export interface WorkspacesQuery {
	rawQuery?: string;
	kind: "WorkspacesQuery";
	subscription: string;
}

export const defaultWorkspacesQuery = (): WorkspacesQuery => ({
	kind: "WorkspacesQuery",
	subscription: "",
});

export type GrafanaTemplateVariableQuery = AppInsightsMetricNameQuery | AppInsightsGroupByQuery | SubscriptionsQuery | ResourceGroupsQuery | ResourceNamesQuery | MetricNamespaceQuery | MetricDefinitionsQuery | MetricNamesQuery | WorkspacesQuery | UnknownQuery;

export const defaultGrafanaTemplateVariableQuery = (): GrafanaTemplateVariableQuery => (defaultAppInsightsMetricNameQuery());

