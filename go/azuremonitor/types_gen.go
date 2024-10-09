// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"encoding/json"
	"errors"
	"fmt"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type AzureMonitorQuery struct {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// Azure subscription containing the resource(s) to be queried.
	Subscription *string `json:"subscription,omitempty"`
	// Subscriptions to be queried via Azure Resource Graph.
	Subscriptions []string `json:"subscriptions,omitempty"`
	// Azure Monitor Metrics sub-query properties.
	AzureMonitor *AzureMetricQuery `json:"azureMonitor,omitempty"`
	// Azure Monitor Logs sub-query properties.
	AzureLogAnalytics *AzureLogsQuery `json:"azureLogAnalytics,omitempty"`
	// Azure Resource Graph sub-query properties.
	AzureResourceGraph *AzureResourceGraphQuery `json:"azureResourceGraph,omitempty"`
	// Application Insights Traces sub-query properties.
	AzureTraces *AzureTracesQuery `json:"azureTraces,omitempty"`
	// @deprecated Legacy template variable support.
	GrafanaTemplateVariableFn *GrafanaTemplateVariableQuery `json:"grafanaTemplateVariableFn,omitempty"`
	// Template variables params. These exist for backwards compatiblity with legacy template variables.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	Namespace     *string `json:"namespace,omitempty"`
	Resource      *string `json:"resource,omitempty"`
	Region        *string `json:"region,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// Used only for exemplar queries from Prometheus
	Query *string `json:"query,omitempty"`
}

func (resource AzureMonitorQuery) ImplementsDataqueryVariant() {}

func (resource AzureMonitorQuery) DataqueryType() string {
	return "grafana-azure-monitor-datasource"
}

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "grafana-azure-monitor-datasource",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &AzureMonitorQuery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery AzureMonitorQuery
			if cast, ok := input.(*AzureMonitorQuery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(AzureMonitorQuery)
			}
			return AzureMonitorQueryConverter(dataquery)
		},
	}
}

func (resource AzureMonitorQuery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(AzureMonitorQuery)
	if !ok {
		return false
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.Subscription == nil && other.Subscription != nil || resource.Subscription != nil && other.Subscription == nil {
		return false
	}

	if resource.Subscription != nil {
		if *resource.Subscription != *other.Subscription {
			return false
		}
	}

	if len(resource.Subscriptions) != len(other.Subscriptions) {
		return false
	}

	for i1 := range resource.Subscriptions {
		if resource.Subscriptions[i1] != other.Subscriptions[i1] {
			return false
		}
	}
	if resource.AzureMonitor == nil && other.AzureMonitor != nil || resource.AzureMonitor != nil && other.AzureMonitor == nil {
		return false
	}

	if resource.AzureMonitor != nil {
		if !resource.AzureMonitor.Equals(*other.AzureMonitor) {
			return false
		}
	}
	if resource.AzureLogAnalytics == nil && other.AzureLogAnalytics != nil || resource.AzureLogAnalytics != nil && other.AzureLogAnalytics == nil {
		return false
	}

	if resource.AzureLogAnalytics != nil {
		if !resource.AzureLogAnalytics.Equals(*other.AzureLogAnalytics) {
			return false
		}
	}
	if resource.AzureResourceGraph == nil && other.AzureResourceGraph != nil || resource.AzureResourceGraph != nil && other.AzureResourceGraph == nil {
		return false
	}

	if resource.AzureResourceGraph != nil {
		if !resource.AzureResourceGraph.Equals(*other.AzureResourceGraph) {
			return false
		}
	}
	if resource.AzureTraces == nil && other.AzureTraces != nil || resource.AzureTraces != nil && other.AzureTraces == nil {
		return false
	}

	if resource.AzureTraces != nil {
		if !resource.AzureTraces.Equals(*other.AzureTraces) {
			return false
		}
	}
	if resource.GrafanaTemplateVariableFn == nil && other.GrafanaTemplateVariableFn != nil || resource.GrafanaTemplateVariableFn != nil && other.GrafanaTemplateVariableFn == nil {
		return false
	}

	if resource.GrafanaTemplateVariableFn != nil {
		if !resource.GrafanaTemplateVariableFn.Equals(*other.GrafanaTemplateVariableFn) {
			return false
		}
	}
	if resource.ResourceGroup == nil && other.ResourceGroup != nil || resource.ResourceGroup != nil && other.ResourceGroup == nil {
		return false
	}

	if resource.ResourceGroup != nil {
		if *resource.ResourceGroup != *other.ResourceGroup {
			return false
		}
	}
	if resource.Namespace == nil && other.Namespace != nil || resource.Namespace != nil && other.Namespace == nil {
		return false
	}

	if resource.Namespace != nil {
		if *resource.Namespace != *other.Namespace {
			return false
		}
	}
	if resource.Resource == nil && other.Resource != nil || resource.Resource != nil && other.Resource == nil {
		return false
	}

	if resource.Resource != nil {
		if *resource.Resource != *other.Resource {
			return false
		}
	}
	if resource.Region == nil && other.Region != nil || resource.Region != nil && other.Region == nil {
		return false
	}

	if resource.Region != nil {
		if *resource.Region != *other.Region {
			return false
		}
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if resource.Query == nil && other.Query != nil || resource.Query != nil && other.Query == nil {
		return false
	}

	if resource.Query != nil {
		if *resource.Query != *other.Query {
			return false
		}
	}

	return true
}

// Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated
type AzureQueryType string

const (
	AzureQueryTypeAzureMonitor              AzureQueryType = "Azure Monitor"
	AzureQueryTypeLogAnalytics              AzureQueryType = "Azure Log Analytics"
	AzureQueryTypeAzureResourceGraph        AzureQueryType = "Azure Resource Graph"
	AzureQueryTypeAzureTraces               AzureQueryType = "Azure Traces"
	AzureQueryTypeSubscriptionsQuery        AzureQueryType = "Azure Subscriptions"
	AzureQueryTypeResourceGroupsQuery       AzureQueryType = "Azure Resource Groups"
	AzureQueryTypeNamespacesQuery           AzureQueryType = "Azure Namespaces"
	AzureQueryTypeResourceNamesQuery        AzureQueryType = "Azure Resource Names"
	AzureQueryTypeMetricNamesQuery          AzureQueryType = "Azure Metric Names"
	AzureQueryTypeWorkspacesQuery           AzureQueryType = "Azure Workspaces"
	AzureQueryTypeLocationsQuery            AzureQueryType = "Azure Regions"
	AzureQueryTypeGrafanaTemplateVariableFn AzureQueryType = "Grafana Template Variable Function"
	AzureQueryTypeTraceExemplar             AzureQueryType = "traceql"
)

type AzureMetricQuery struct {
	// Array of resource URIs to be queried.
	Resources []AzureMonitorResource `json:"resources,omitempty"`
	// metricNamespace is used as the resource type (or resource namespace).
	// It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
	// Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
	MetricNamespace *string `json:"metricNamespace,omitempty"`
	// Used as the value for the metricNamespace property when it's different from the resource namespace.
	CustomNamespace *string `json:"customNamespace,omitempty"`
	// The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
	MetricName *string `json:"metricName,omitempty"`
	// The Azure region containing the resource(s).
	Region *string `json:"region,omitempty"`
	// The granularity of data points to be queried. Defaults to auto.
	TimeGrain *string `json:"timeGrain,omitempty"`
	// The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
	Aggregation *string `json:"aggregation,omitempty"`
	// Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
	DimensionFilters []AzureMetricDimension `json:"dimensionFilters,omitempty"`
	// Maximum number of records to return. Defaults to 10.
	Top *string `json:"top,omitempty"`
	// Time grains that are supported by the metric.
	AllowedTimeGrainsMs []int64 `json:"allowedTimeGrainsMs,omitempty"`
	// Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
	Alias *string `json:"alias,omitempty"`
	// @deprecated
	TimeGrainUnit *string `json:"timeGrainUnit,omitempty"`
	// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
	Dimension *string `json:"dimension,omitempty"`
	// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
	DimensionFilter *string `json:"dimensionFilter,omitempty"`
	// @deprecated Use metricNamespace instead
	MetricDefinition *string `json:"metricDefinition,omitempty"`
	// @deprecated Use resourceGroup, resourceName and metricNamespace instead
	ResourceUri *string `json:"resourceUri,omitempty"`
	// @deprecated Use resources instead
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// @deprecated Use resources instead
	ResourceName *string `json:"resourceName,omitempty"`
}

func (resource AzureMetricQuery) Equals(other AzureMetricQuery) bool {

	if len(resource.Resources) != len(other.Resources) {
		return false
	}

	for i1 := range resource.Resources {
		if !resource.Resources[i1].Equals(other.Resources[i1]) {
			return false
		}
	}
	if resource.MetricNamespace == nil && other.MetricNamespace != nil || resource.MetricNamespace != nil && other.MetricNamespace == nil {
		return false
	}

	if resource.MetricNamespace != nil {
		if *resource.MetricNamespace != *other.MetricNamespace {
			return false
		}
	}
	if resource.CustomNamespace == nil && other.CustomNamespace != nil || resource.CustomNamespace != nil && other.CustomNamespace == nil {
		return false
	}

	if resource.CustomNamespace != nil {
		if *resource.CustomNamespace != *other.CustomNamespace {
			return false
		}
	}
	if resource.MetricName == nil && other.MetricName != nil || resource.MetricName != nil && other.MetricName == nil {
		return false
	}

	if resource.MetricName != nil {
		if *resource.MetricName != *other.MetricName {
			return false
		}
	}
	if resource.Region == nil && other.Region != nil || resource.Region != nil && other.Region == nil {
		return false
	}

	if resource.Region != nil {
		if *resource.Region != *other.Region {
			return false
		}
	}
	if resource.TimeGrain == nil && other.TimeGrain != nil || resource.TimeGrain != nil && other.TimeGrain == nil {
		return false
	}

	if resource.TimeGrain != nil {
		if *resource.TimeGrain != *other.TimeGrain {
			return false
		}
	}
	if resource.Aggregation == nil && other.Aggregation != nil || resource.Aggregation != nil && other.Aggregation == nil {
		return false
	}

	if resource.Aggregation != nil {
		if *resource.Aggregation != *other.Aggregation {
			return false
		}
	}

	if len(resource.DimensionFilters) != len(other.DimensionFilters) {
		return false
	}

	for i1 := range resource.DimensionFilters {
		if !resource.DimensionFilters[i1].Equals(other.DimensionFilters[i1]) {
			return false
		}
	}
	if resource.Top == nil && other.Top != nil || resource.Top != nil && other.Top == nil {
		return false
	}

	if resource.Top != nil {
		if *resource.Top != *other.Top {
			return false
		}
	}

	if len(resource.AllowedTimeGrainsMs) != len(other.AllowedTimeGrainsMs) {
		return false
	}

	for i1 := range resource.AllowedTimeGrainsMs {
		if resource.AllowedTimeGrainsMs[i1] != other.AllowedTimeGrainsMs[i1] {
			return false
		}
	}
	if resource.Alias == nil && other.Alias != nil || resource.Alias != nil && other.Alias == nil {
		return false
	}

	if resource.Alias != nil {
		if *resource.Alias != *other.Alias {
			return false
		}
	}
	if resource.TimeGrainUnit == nil && other.TimeGrainUnit != nil || resource.TimeGrainUnit != nil && other.TimeGrainUnit == nil {
		return false
	}

	if resource.TimeGrainUnit != nil {
		if *resource.TimeGrainUnit != *other.TimeGrainUnit {
			return false
		}
	}
	if resource.Dimension == nil && other.Dimension != nil || resource.Dimension != nil && other.Dimension == nil {
		return false
	}

	if resource.Dimension != nil {
		if *resource.Dimension != *other.Dimension {
			return false
		}
	}
	if resource.DimensionFilter == nil && other.DimensionFilter != nil || resource.DimensionFilter != nil && other.DimensionFilter == nil {
		return false
	}

	if resource.DimensionFilter != nil {
		if *resource.DimensionFilter != *other.DimensionFilter {
			return false
		}
	}
	if resource.MetricDefinition == nil && other.MetricDefinition != nil || resource.MetricDefinition != nil && other.MetricDefinition == nil {
		return false
	}

	if resource.MetricDefinition != nil {
		if *resource.MetricDefinition != *other.MetricDefinition {
			return false
		}
	}
	if resource.ResourceUri == nil && other.ResourceUri != nil || resource.ResourceUri != nil && other.ResourceUri == nil {
		return false
	}

	if resource.ResourceUri != nil {
		if *resource.ResourceUri != *other.ResourceUri {
			return false
		}
	}
	if resource.ResourceGroup == nil && other.ResourceGroup != nil || resource.ResourceGroup != nil && other.ResourceGroup == nil {
		return false
	}

	if resource.ResourceGroup != nil {
		if *resource.ResourceGroup != *other.ResourceGroup {
			return false
		}
	}
	if resource.ResourceName == nil && other.ResourceName != nil || resource.ResourceName != nil && other.ResourceName == nil {
		return false
	}

	if resource.ResourceName != nil {
		if *resource.ResourceName != *other.ResourceName {
			return false
		}
	}

	return true
}

// Azure Monitor Logs sub-query properties
type AzureLogsQuery struct {
	// KQL query to be executed.
	Query *string `json:"query,omitempty"`
	// Specifies the format results should be returned as.
	ResultFormat *ResultFormat `json:"resultFormat,omitempty"`
	// Array of resource URIs to be queried.
	Resources []string `json:"resources,omitempty"`
	// If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
	DashboardTime *bool `json:"dashboardTime,omitempty"`
	// If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
	TimeColumn *string `json:"timeColumn,omitempty"`
	// If set to true the query will be run as a basic logs query
	BasicLogsQuery *bool `json:"basicLogsQuery,omitempty"`
	// Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
	Workspace *string `json:"workspace,omitempty"`
	// @deprecated Use resources instead
	Resource *string `json:"resource,omitempty"`
	// @deprecated Use dashboardTime instead
	IntersectTime *bool `json:"intersectTime,omitempty"`
}

func (resource AzureLogsQuery) Equals(other AzureLogsQuery) bool {
	if resource.Query == nil && other.Query != nil || resource.Query != nil && other.Query == nil {
		return false
	}

	if resource.Query != nil {
		if *resource.Query != *other.Query {
			return false
		}
	}
	if resource.ResultFormat == nil && other.ResultFormat != nil || resource.ResultFormat != nil && other.ResultFormat == nil {
		return false
	}

	if resource.ResultFormat != nil {
		if *resource.ResultFormat != *other.ResultFormat {
			return false
		}
	}

	if len(resource.Resources) != len(other.Resources) {
		return false
	}

	for i1 := range resource.Resources {
		if resource.Resources[i1] != other.Resources[i1] {
			return false
		}
	}
	if resource.DashboardTime == nil && other.DashboardTime != nil || resource.DashboardTime != nil && other.DashboardTime == nil {
		return false
	}

	if resource.DashboardTime != nil {
		if *resource.DashboardTime != *other.DashboardTime {
			return false
		}
	}
	if resource.TimeColumn == nil && other.TimeColumn != nil || resource.TimeColumn != nil && other.TimeColumn == nil {
		return false
	}

	if resource.TimeColumn != nil {
		if *resource.TimeColumn != *other.TimeColumn {
			return false
		}
	}
	if resource.BasicLogsQuery == nil && other.BasicLogsQuery != nil || resource.BasicLogsQuery != nil && other.BasicLogsQuery == nil {
		return false
	}

	if resource.BasicLogsQuery != nil {
		if *resource.BasicLogsQuery != *other.BasicLogsQuery {
			return false
		}
	}
	if resource.Workspace == nil && other.Workspace != nil || resource.Workspace != nil && other.Workspace == nil {
		return false
	}

	if resource.Workspace != nil {
		if *resource.Workspace != *other.Workspace {
			return false
		}
	}
	if resource.Resource == nil && other.Resource != nil || resource.Resource != nil && other.Resource == nil {
		return false
	}

	if resource.Resource != nil {
		if *resource.Resource != *other.Resource {
			return false
		}
	}
	if resource.IntersectTime == nil && other.IntersectTime != nil || resource.IntersectTime != nil && other.IntersectTime == nil {
		return false
	}

	if resource.IntersectTime != nil {
		if *resource.IntersectTime != *other.IntersectTime {
			return false
		}
	}

	return true
}

// Application Insights Traces sub-query properties
type AzureTracesQuery struct {
	// Specifies the format results should be returned as.
	ResultFormat *ResultFormat `json:"resultFormat,omitempty"`
	// Array of resource URIs to be queried.
	Resources []string `json:"resources,omitempty"`
	// Operation ID. Used only for Traces queries.
	OperationId *string `json:"operationId,omitempty"`
	// Types of events to filter by.
	TraceTypes []string `json:"traceTypes,omitempty"`
	// Filters for property values.
	Filters []AzureTracesFilter `json:"filters,omitempty"`
	// KQL query to be executed.
	Query *string `json:"query,omitempty"`
}

func (resource AzureTracesQuery) Equals(other AzureTracesQuery) bool {
	if resource.ResultFormat == nil && other.ResultFormat != nil || resource.ResultFormat != nil && other.ResultFormat == nil {
		return false
	}

	if resource.ResultFormat != nil {
		if *resource.ResultFormat != *other.ResultFormat {
			return false
		}
	}

	if len(resource.Resources) != len(other.Resources) {
		return false
	}

	for i1 := range resource.Resources {
		if resource.Resources[i1] != other.Resources[i1] {
			return false
		}
	}
	if resource.OperationId == nil && other.OperationId != nil || resource.OperationId != nil && other.OperationId == nil {
		return false
	}

	if resource.OperationId != nil {
		if *resource.OperationId != *other.OperationId {
			return false
		}
	}

	if len(resource.TraceTypes) != len(other.TraceTypes) {
		return false
	}

	for i1 := range resource.TraceTypes {
		if resource.TraceTypes[i1] != other.TraceTypes[i1] {
			return false
		}
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if !resource.Filters[i1].Equals(other.Filters[i1]) {
			return false
		}
	}
	if resource.Query == nil && other.Query != nil || resource.Query != nil && other.Query == nil {
		return false
	}

	if resource.Query != nil {
		if *resource.Query != *other.Query {
			return false
		}
	}

	return true
}

type AzureTracesFilter struct {
	// Property name, auto-populated based on available traces.
	Property string `json:"property"`
	// Comparison operator to use. Either equals or not equals.
	Operation string `json:"operation"`
	// Values to filter by.
	Filters []string `json:"filters"`
}

func (resource AzureTracesFilter) Equals(other AzureTracesFilter) bool {
	if resource.Property != other.Property {
		return false
	}
	if resource.Operation != other.Operation {
		return false
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if resource.Filters[i1] != other.Filters[i1] {
			return false
		}
	}

	return true
}

type ResultFormat string

const (
	ResultFormatTable      ResultFormat = "table"
	ResultFormatTimeSeries ResultFormat = "time_series"
	ResultFormatTrace      ResultFormat = "trace"
	ResultFormatLogs       ResultFormat = "logs"
)

type AzureResourceGraphQuery struct {
	// Azure Resource Graph KQL query to be executed.
	Query *string `json:"query,omitempty"`
	// Specifies the format results should be returned as. Defaults to table.
	ResultFormat *string `json:"resultFormat,omitempty"`
}

func (resource AzureResourceGraphQuery) Equals(other AzureResourceGraphQuery) bool {
	if resource.Query == nil && other.Query != nil || resource.Query != nil && other.Query == nil {
		return false
	}

	if resource.Query != nil {
		if *resource.Query != *other.Query {
			return false
		}
	}
	if resource.ResultFormat == nil && other.ResultFormat != nil || resource.ResultFormat != nil && other.ResultFormat == nil {
		return false
	}

	if resource.ResultFormat != nil {
		if *resource.ResultFormat != *other.ResultFormat {
			return false
		}
	}

	return true
}

type AzureMonitorResource struct {
	Subscription    *string `json:"subscription,omitempty"`
	ResourceGroup   *string `json:"resourceGroup,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty"`
	MetricNamespace *string `json:"metricNamespace,omitempty"`
	Region          *string `json:"region,omitempty"`
}

func (resource AzureMonitorResource) Equals(other AzureMonitorResource) bool {
	if resource.Subscription == nil && other.Subscription != nil || resource.Subscription != nil && other.Subscription == nil {
		return false
	}

	if resource.Subscription != nil {
		if *resource.Subscription != *other.Subscription {
			return false
		}
	}
	if resource.ResourceGroup == nil && other.ResourceGroup != nil || resource.ResourceGroup != nil && other.ResourceGroup == nil {
		return false
	}

	if resource.ResourceGroup != nil {
		if *resource.ResourceGroup != *other.ResourceGroup {
			return false
		}
	}
	if resource.ResourceName == nil && other.ResourceName != nil || resource.ResourceName != nil && other.ResourceName == nil {
		return false
	}

	if resource.ResourceName != nil {
		if *resource.ResourceName != *other.ResourceName {
			return false
		}
	}
	if resource.MetricNamespace == nil && other.MetricNamespace != nil || resource.MetricNamespace != nil && other.MetricNamespace == nil {
		return false
	}

	if resource.MetricNamespace != nil {
		if *resource.MetricNamespace != *other.MetricNamespace {
			return false
		}
	}
	if resource.Region == nil && other.Region != nil || resource.Region != nil && other.Region == nil {
		return false
	}

	if resource.Region != nil {
		if *resource.Region != *other.Region {
			return false
		}
	}

	return true
}

type AzureMetricDimension struct {
	// Name of Dimension to be filtered on.
	Dimension *string `json:"dimension,omitempty"`
	// String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
	Operator *string `json:"operator,omitempty"`
	// Values to match with the filter.
	Filters []string `json:"filters,omitempty"`
	// @deprecated filter is deprecated in favour of filters to support multiselect.
	Filter *string `json:"filter,omitempty"`
}

func (resource AzureMetricDimension) Equals(other AzureMetricDimension) bool {
	if resource.Dimension == nil && other.Dimension != nil || resource.Dimension != nil && other.Dimension == nil {
		return false
	}

	if resource.Dimension != nil {
		if *resource.Dimension != *other.Dimension {
			return false
		}
	}
	if resource.Operator == nil && other.Operator != nil || resource.Operator != nil && other.Operator == nil {
		return false
	}

	if resource.Operator != nil {
		if *resource.Operator != *other.Operator {
			return false
		}
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if resource.Filters[i1] != other.Filters[i1] {
			return false
		}
	}
	if resource.Filter == nil && other.Filter != nil || resource.Filter != nil && other.Filter == nil {
		return false
	}

	if resource.Filter != nil {
		if *resource.Filter != *other.Filter {
			return false
		}
	}

	return true
}

type GrafanaTemplateVariableQueryType string

const (
	GrafanaTemplateVariableQueryTypeAppInsightsMetricNameQuery GrafanaTemplateVariableQueryType = "AppInsightsMetricNameQuery"
	GrafanaTemplateVariableQueryTypeAppInsightsGroupByQuery    GrafanaTemplateVariableQueryType = "AppInsightsGroupByQuery"
	GrafanaTemplateVariableQueryTypeSubscriptionsQuery         GrafanaTemplateVariableQueryType = "SubscriptionsQuery"
	GrafanaTemplateVariableQueryTypeResourceGroupsQuery        GrafanaTemplateVariableQueryType = "ResourceGroupsQuery"
	GrafanaTemplateVariableQueryTypeResourceNamesQuery         GrafanaTemplateVariableQueryType = "ResourceNamesQuery"
	GrafanaTemplateVariableQueryTypeMetricNamespaceQuery       GrafanaTemplateVariableQueryType = "MetricNamespaceQuery"
	GrafanaTemplateVariableQueryTypeMetricNamesQuery           GrafanaTemplateVariableQueryType = "MetricNamesQuery"
	GrafanaTemplateVariableQueryTypeWorkspacesQuery            GrafanaTemplateVariableQueryType = "WorkspacesQuery"
	GrafanaTemplateVariableQueryTypeUnknownQuery               GrafanaTemplateVariableQueryType = "UnknownQuery"
)

type BaseGrafanaTemplateVariableQuery struct {
	RawQuery *string `json:"rawQuery,omitempty"`
}

func (resource BaseGrafanaTemplateVariableQuery) Equals(other BaseGrafanaTemplateVariableQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}

	return true
}

type UnknownQuery struct {
	RawQuery *string `json:"rawQuery,omitempty"`
	Kind     string  `json:"kind"`
}

func (resource UnknownQuery) Equals(other UnknownQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}

	return true
}

type AppInsightsMetricNameQuery struct {
	RawQuery *string `json:"rawQuery,omitempty"`
	Kind     string  `json:"kind"`
}

func (resource AppInsightsMetricNameQuery) Equals(other AppInsightsMetricNameQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}

	return true
}

type AppInsightsGroupByQuery struct {
	RawQuery   *string `json:"rawQuery,omitempty"`
	Kind       string  `json:"kind"`
	MetricName string  `json:"metricName"`
}

func (resource AppInsightsGroupByQuery) Equals(other AppInsightsGroupByQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.MetricName != other.MetricName {
		return false
	}

	return true
}

type SubscriptionsQuery struct {
	RawQuery *string `json:"rawQuery,omitempty"`
	Kind     string  `json:"kind"`
}

func (resource SubscriptionsQuery) Equals(other SubscriptionsQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}

	return true
}

type ResourceGroupsQuery struct {
	RawQuery     *string `json:"rawQuery,omitempty"`
	Kind         string  `json:"kind"`
	Subscription string  `json:"subscription"`
}

func (resource ResourceGroupsQuery) Equals(other ResourceGroupsQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Subscription != other.Subscription {
		return false
	}

	return true
}

type ResourceNamesQuery struct {
	RawQuery        *string `json:"rawQuery,omitempty"`
	Kind            string  `json:"kind"`
	Subscription    string  `json:"subscription"`
	ResourceGroup   string  `json:"resourceGroup"`
	MetricNamespace string  `json:"metricNamespace"`
}

func (resource ResourceNamesQuery) Equals(other ResourceNamesQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Subscription != other.Subscription {
		return false
	}
	if resource.ResourceGroup != other.ResourceGroup {
		return false
	}
	if resource.MetricNamespace != other.MetricNamespace {
		return false
	}

	return true
}

type MetricNamespaceQuery struct {
	RawQuery        *string `json:"rawQuery,omitempty"`
	Kind            string  `json:"kind"`
	Subscription    string  `json:"subscription"`
	ResourceGroup   string  `json:"resourceGroup"`
	MetricNamespace *string `json:"metricNamespace,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty"`
}

func (resource MetricNamespaceQuery) Equals(other MetricNamespaceQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Subscription != other.Subscription {
		return false
	}
	if resource.ResourceGroup != other.ResourceGroup {
		return false
	}
	if resource.MetricNamespace == nil && other.MetricNamespace != nil || resource.MetricNamespace != nil && other.MetricNamespace == nil {
		return false
	}

	if resource.MetricNamespace != nil {
		if *resource.MetricNamespace != *other.MetricNamespace {
			return false
		}
	}
	if resource.ResourceName == nil && other.ResourceName != nil || resource.ResourceName != nil && other.ResourceName == nil {
		return false
	}

	if resource.ResourceName != nil {
		if *resource.ResourceName != *other.ResourceName {
			return false
		}
	}

	return true
}

// @deprecated Use MetricNamespaceQuery instead
type MetricDefinitionsQuery struct {
	RawQuery        *string `json:"rawQuery,omitempty"`
	Kind            string  `json:"kind"`
	Subscription    string  `json:"subscription"`
	ResourceGroup   string  `json:"resourceGroup"`
	MetricNamespace *string `json:"metricNamespace,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty"`
}

func (resource MetricDefinitionsQuery) Equals(other MetricDefinitionsQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Subscription != other.Subscription {
		return false
	}
	if resource.ResourceGroup != other.ResourceGroup {
		return false
	}
	if resource.MetricNamespace == nil && other.MetricNamespace != nil || resource.MetricNamespace != nil && other.MetricNamespace == nil {
		return false
	}

	if resource.MetricNamespace != nil {
		if *resource.MetricNamespace != *other.MetricNamespace {
			return false
		}
	}
	if resource.ResourceName == nil && other.ResourceName != nil || resource.ResourceName != nil && other.ResourceName == nil {
		return false
	}

	if resource.ResourceName != nil {
		if *resource.ResourceName != *other.ResourceName {
			return false
		}
	}

	return true
}

type MetricNamesQuery struct {
	RawQuery        *string `json:"rawQuery,omitempty"`
	Kind            string  `json:"kind"`
	Subscription    string  `json:"subscription"`
	ResourceGroup   string  `json:"resourceGroup"`
	ResourceName    string  `json:"resourceName"`
	MetricNamespace string  `json:"metricNamespace"`
}

func (resource MetricNamesQuery) Equals(other MetricNamesQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Subscription != other.Subscription {
		return false
	}
	if resource.ResourceGroup != other.ResourceGroup {
		return false
	}
	if resource.ResourceName != other.ResourceName {
		return false
	}
	if resource.MetricNamespace != other.MetricNamespace {
		return false
	}

	return true
}

type WorkspacesQuery struct {
	RawQuery     *string `json:"rawQuery,omitempty"`
	Kind         string  `json:"kind"`
	Subscription string  `json:"subscription"`
}

func (resource WorkspacesQuery) Equals(other WorkspacesQuery) bool {
	if resource.RawQuery == nil && other.RawQuery != nil || resource.RawQuery != nil && other.RawQuery == nil {
		return false
	}

	if resource.RawQuery != nil {
		if *resource.RawQuery != *other.RawQuery {
			return false
		}
	}
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Subscription != other.Subscription {
		return false
	}

	return true
}

type GrafanaTemplateVariableQuery = AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery

type AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery struct {
	AppInsightsMetricNameQuery *AppInsightsMetricNameQuery `json:"AppInsightsMetricNameQuery,omitempty"`
	AppInsightsGroupByQuery    *AppInsightsGroupByQuery    `json:"AppInsightsGroupByQuery,omitempty"`
	SubscriptionsQuery         *SubscriptionsQuery         `json:"SubscriptionsQuery,omitempty"`
	ResourceGroupsQuery        *ResourceGroupsQuery        `json:"ResourceGroupsQuery,omitempty"`
	ResourceNamesQuery         *ResourceNamesQuery         `json:"ResourceNamesQuery,omitempty"`
	MetricNamespaceQuery       *MetricNamespaceQuery       `json:"MetricNamespaceQuery,omitempty"`
	MetricDefinitionsQuery     *MetricDefinitionsQuery     `json:"MetricDefinitionsQuery,omitempty"`
	MetricNamesQuery           *MetricNamesQuery           `json:"MetricNamesQuery,omitempty"`
	WorkspacesQuery            *WorkspacesQuery            `json:"WorkspacesQuery,omitempty"`
	UnknownQuery               *UnknownQuery               `json:"UnknownQuery,omitempty"`
}

func (resource AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) MarshalJSON() ([]byte, error) {
	if resource.AppInsightsMetricNameQuery != nil {
		return json.Marshal(resource.AppInsightsMetricNameQuery)
	}
	if resource.AppInsightsGroupByQuery != nil {
		return json.Marshal(resource.AppInsightsGroupByQuery)
	}
	if resource.SubscriptionsQuery != nil {
		return json.Marshal(resource.SubscriptionsQuery)
	}
	if resource.ResourceGroupsQuery != nil {
		return json.Marshal(resource.ResourceGroupsQuery)
	}
	if resource.ResourceNamesQuery != nil {
		return json.Marshal(resource.ResourceNamesQuery)
	}
	if resource.MetricNamespaceQuery != nil {
		return json.Marshal(resource.MetricNamespaceQuery)
	}
	if resource.MetricDefinitionsQuery != nil {
		return json.Marshal(resource.MetricDefinitionsQuery)
	}
	if resource.MetricNamesQuery != nil {
		return json.Marshal(resource.MetricNamesQuery)
	}
	if resource.WorkspacesQuery != nil {
		return json.Marshal(resource.WorkspacesQuery)
	}
	if resource.UnknownQuery != nil {
		return json.Marshal(resource.UnknownQuery)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return errors.New("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "AppInsightsGroupByQuery":
		var appInsightsGroupByQuery AppInsightsGroupByQuery
		if err := json.Unmarshal(raw, &appInsightsGroupByQuery); err != nil {
			return err
		}

		resource.AppInsightsGroupByQuery = &appInsightsGroupByQuery
		return nil
	case "AppInsightsMetricNameQuery":
		var appInsightsMetricNameQuery AppInsightsMetricNameQuery
		if err := json.Unmarshal(raw, &appInsightsMetricNameQuery); err != nil {
			return err
		}

		resource.AppInsightsMetricNameQuery = &appInsightsMetricNameQuery
		return nil
	case "MetricDefinitionsQuery":
		var metricDefinitionsQuery MetricDefinitionsQuery
		if err := json.Unmarshal(raw, &metricDefinitionsQuery); err != nil {
			return err
		}

		resource.MetricDefinitionsQuery = &metricDefinitionsQuery
		return nil
	case "MetricNamesQuery":
		var metricNamesQuery MetricNamesQuery
		if err := json.Unmarshal(raw, &metricNamesQuery); err != nil {
			return err
		}

		resource.MetricNamesQuery = &metricNamesQuery
		return nil
	case "MetricNamespaceQuery":
		var metricNamespaceQuery MetricNamespaceQuery
		if err := json.Unmarshal(raw, &metricNamespaceQuery); err != nil {
			return err
		}

		resource.MetricNamespaceQuery = &metricNamespaceQuery
		return nil
	case "ResourceGroupsQuery":
		var resourceGroupsQuery ResourceGroupsQuery
		if err := json.Unmarshal(raw, &resourceGroupsQuery); err != nil {
			return err
		}

		resource.ResourceGroupsQuery = &resourceGroupsQuery
		return nil
	case "ResourceNamesQuery":
		var resourceNamesQuery ResourceNamesQuery
		if err := json.Unmarshal(raw, &resourceNamesQuery); err != nil {
			return err
		}

		resource.ResourceNamesQuery = &resourceNamesQuery
		return nil
	case "SubscriptionsQuery":
		var subscriptionsQuery SubscriptionsQuery
		if err := json.Unmarshal(raw, &subscriptionsQuery); err != nil {
			return err
		}

		resource.SubscriptionsQuery = &subscriptionsQuery
		return nil
	case "UnknownQuery":
		var unknownQuery UnknownQuery
		if err := json.Unmarshal(raw, &unknownQuery); err != nil {
			return err
		}

		resource.UnknownQuery = &unknownQuery
		return nil
	case "WorkspacesQuery":
		var workspacesQuery WorkspacesQuery
		if err := json.Unmarshal(raw, &workspacesQuery); err != nil {
			return err
		}

		resource.WorkspacesQuery = &workspacesQuery
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

func (resource AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) Equals(other AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) bool {
	if resource.AppInsightsMetricNameQuery == nil && other.AppInsightsMetricNameQuery != nil || resource.AppInsightsMetricNameQuery != nil && other.AppInsightsMetricNameQuery == nil {
		return false
	}

	if resource.AppInsightsMetricNameQuery != nil {
		if !resource.AppInsightsMetricNameQuery.Equals(*other.AppInsightsMetricNameQuery) {
			return false
		}
	}
	if resource.AppInsightsGroupByQuery == nil && other.AppInsightsGroupByQuery != nil || resource.AppInsightsGroupByQuery != nil && other.AppInsightsGroupByQuery == nil {
		return false
	}

	if resource.AppInsightsGroupByQuery != nil {
		if !resource.AppInsightsGroupByQuery.Equals(*other.AppInsightsGroupByQuery) {
			return false
		}
	}
	if resource.SubscriptionsQuery == nil && other.SubscriptionsQuery != nil || resource.SubscriptionsQuery != nil && other.SubscriptionsQuery == nil {
		return false
	}

	if resource.SubscriptionsQuery != nil {
		if !resource.SubscriptionsQuery.Equals(*other.SubscriptionsQuery) {
			return false
		}
	}
	if resource.ResourceGroupsQuery == nil && other.ResourceGroupsQuery != nil || resource.ResourceGroupsQuery != nil && other.ResourceGroupsQuery == nil {
		return false
	}

	if resource.ResourceGroupsQuery != nil {
		if !resource.ResourceGroupsQuery.Equals(*other.ResourceGroupsQuery) {
			return false
		}
	}
	if resource.ResourceNamesQuery == nil && other.ResourceNamesQuery != nil || resource.ResourceNamesQuery != nil && other.ResourceNamesQuery == nil {
		return false
	}

	if resource.ResourceNamesQuery != nil {
		if !resource.ResourceNamesQuery.Equals(*other.ResourceNamesQuery) {
			return false
		}
	}
	if resource.MetricNamespaceQuery == nil && other.MetricNamespaceQuery != nil || resource.MetricNamespaceQuery != nil && other.MetricNamespaceQuery == nil {
		return false
	}

	if resource.MetricNamespaceQuery != nil {
		if !resource.MetricNamespaceQuery.Equals(*other.MetricNamespaceQuery) {
			return false
		}
	}
	if resource.MetricDefinitionsQuery == nil && other.MetricDefinitionsQuery != nil || resource.MetricDefinitionsQuery != nil && other.MetricDefinitionsQuery == nil {
		return false
	}

	if resource.MetricDefinitionsQuery != nil {
		if !resource.MetricDefinitionsQuery.Equals(*other.MetricDefinitionsQuery) {
			return false
		}
	}
	if resource.MetricNamesQuery == nil && other.MetricNamesQuery != nil || resource.MetricNamesQuery != nil && other.MetricNamesQuery == nil {
		return false
	}

	if resource.MetricNamesQuery != nil {
		if !resource.MetricNamesQuery.Equals(*other.MetricNamesQuery) {
			return false
		}
	}
	if resource.WorkspacesQuery == nil && other.WorkspacesQuery != nil || resource.WorkspacesQuery != nil && other.WorkspacesQuery == nil {
		return false
	}

	if resource.WorkspacesQuery != nil {
		if !resource.WorkspacesQuery.Equals(*other.WorkspacesQuery) {
			return false
		}
	}
	if resource.UnknownQuery == nil && other.UnknownQuery != nil || resource.UnknownQuery != nil && other.UnknownQuery == nil {
		return false
	}

	if resource.UnknownQuery != nil {
		if !resource.UnknownQuery.Equals(*other.UnknownQuery) {
			return false
		}
	}

	return true
}
