// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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

// VariantConfig returns the configuration related to grafana-azure-monitor-datasource dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &AzureMonitorQuery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMonitorQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureMonitorQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}

		}
		delete(fields, "subscription")

	}
	// Field "subscriptions"
	if fields["subscriptions"] != nil {
		if string(fields["subscriptions"]) != "null" {

			if err := json.Unmarshal(fields["subscriptions"], &resource.Subscriptions); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscriptions", err)...)
			}

		}
		delete(fields, "subscriptions")

	}
	// Field "azureMonitor"
	if fields["azureMonitor"] != nil {
		if string(fields["azureMonitor"]) != "null" {

			resource.AzureMonitor = &AzureMetricQuery{}
			if err := resource.AzureMonitor.UnmarshalJSONStrict(fields["azureMonitor"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("azureMonitor", err)...)
			}

		}
		delete(fields, "azureMonitor")

	}
	// Field "azureLogAnalytics"
	if fields["azureLogAnalytics"] != nil {
		if string(fields["azureLogAnalytics"]) != "null" {

			resource.AzureLogAnalytics = &AzureLogsQuery{}
			if err := resource.AzureLogAnalytics.UnmarshalJSONStrict(fields["azureLogAnalytics"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("azureLogAnalytics", err)...)
			}

		}
		delete(fields, "azureLogAnalytics")

	}
	// Field "azureResourceGraph"
	if fields["azureResourceGraph"] != nil {
		if string(fields["azureResourceGraph"]) != "null" {

			resource.AzureResourceGraph = &AzureResourceGraphQuery{}
			if err := resource.AzureResourceGraph.UnmarshalJSONStrict(fields["azureResourceGraph"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("azureResourceGraph", err)...)
			}

		}
		delete(fields, "azureResourceGraph")

	}
	// Field "azureTraces"
	if fields["azureTraces"] != nil {
		if string(fields["azureTraces"]) != "null" {

			resource.AzureTraces = &AzureTracesQuery{}
			if err := resource.AzureTraces.UnmarshalJSONStrict(fields["azureTraces"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("azureTraces", err)...)
			}

		}
		delete(fields, "azureTraces")

	}
	// Field "grafanaTemplateVariableFn"
	if fields["grafanaTemplateVariableFn"] != nil {
		if string(fields["grafanaTemplateVariableFn"]) != "null" {

			resource.GrafanaTemplateVariableFn = &GrafanaTemplateVariableQuery{}
			if err := resource.GrafanaTemplateVariableFn.UnmarshalJSONStrict(fields["grafanaTemplateVariableFn"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("grafanaTemplateVariableFn", err)...)
			}

		}
		delete(fields, "grafanaTemplateVariableFn")

	}
	// Field "resourceGroup"
	if fields["resourceGroup"] != nil {
		if string(fields["resourceGroup"]) != "null" {
			if err := json.Unmarshal(fields["resourceGroup"], &resource.ResourceGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceGroup", err)...)
			}

		}
		delete(fields, "resourceGroup")

	}
	// Field "namespace"
	if fields["namespace"] != nil {
		if string(fields["namespace"]) != "null" {
			if err := json.Unmarshal(fields["namespace"], &resource.Namespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("namespace", err)...)
			}

		}
		delete(fields, "namespace")

	}
	// Field "resource"
	if fields["resource"] != nil {
		if string(fields["resource"]) != "null" {
			if err := json.Unmarshal(fields["resource"], &resource.Resource); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resource", err)...)
			}

		}
		delete(fields, "resource")

	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}

		}
		delete(fields, "region")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}

		}
		delete(fields, "query")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureMonitorQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `AzureMonitorQuery` fields for violations and returns them.
func (resource AzureMonitorQuery) Validate() error {
	var errs cog.BuildErrors
	if resource.AzureMonitor != nil {
		if err := resource.AzureMonitor.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("azureMonitor", err)...)
		}
	}
	if resource.AzureLogAnalytics != nil {
		if err := resource.AzureLogAnalytics.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("azureLogAnalytics", err)...)
		}
	}
	if resource.AzureResourceGraph != nil {
		if err := resource.AzureResourceGraph.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("azureResourceGraph", err)...)
		}
	}
	if resource.AzureTraces != nil {
		if err := resource.AzureTraces.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("azureTraces", err)...)
		}
	}
	if resource.GrafanaTemplateVariableFn != nil {
		if err := resource.GrafanaTemplateVariableFn.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("grafanaTemplateVariableFn", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMetricQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureMetricQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "resources"
	if fields["resources"] != nil {
		if string(fields["resources"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["resources"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AzureMonitorResource

				result1 = AzureMonitorResource{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("resources["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Resources = append(resource.Resources, result1)
			}

		}
		delete(fields, "resources")

	}
	// Field "metricNamespace"
	if fields["metricNamespace"] != nil {
		if string(fields["metricNamespace"]) != "null" {
			if err := json.Unmarshal(fields["metricNamespace"], &resource.MetricNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricNamespace", err)...)
			}

		}
		delete(fields, "metricNamespace")

	}
	// Field "customNamespace"
	if fields["customNamespace"] != nil {
		if string(fields["customNamespace"]) != "null" {
			if err := json.Unmarshal(fields["customNamespace"], &resource.CustomNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("customNamespace", err)...)
			}

		}
		delete(fields, "customNamespace")

	}
	// Field "metricName"
	if fields["metricName"] != nil {
		if string(fields["metricName"]) != "null" {
			if err := json.Unmarshal(fields["metricName"], &resource.MetricName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricName", err)...)
			}

		}
		delete(fields, "metricName")

	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}

		}
		delete(fields, "region")

	}
	// Field "timeGrain"
	if fields["timeGrain"] != nil {
		if string(fields["timeGrain"]) != "null" {
			if err := json.Unmarshal(fields["timeGrain"], &resource.TimeGrain); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeGrain", err)...)
			}

		}
		delete(fields, "timeGrain")

	}
	// Field "aggregation"
	if fields["aggregation"] != nil {
		if string(fields["aggregation"]) != "null" {
			if err := json.Unmarshal(fields["aggregation"], &resource.Aggregation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("aggregation", err)...)
			}

		}
		delete(fields, "aggregation")

	}
	// Field "dimensionFilters"
	if fields["dimensionFilters"] != nil {
		if string(fields["dimensionFilters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["dimensionFilters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AzureMetricDimension

				result1 = AzureMetricDimension{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("dimensionFilters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.DimensionFilters = append(resource.DimensionFilters, result1)
			}

		}
		delete(fields, "dimensionFilters")

	}
	// Field "top"
	if fields["top"] != nil {
		if string(fields["top"]) != "null" {
			if err := json.Unmarshal(fields["top"], &resource.Top); err != nil {
				errs = append(errs, cog.MakeBuildErrors("top", err)...)
			}

		}
		delete(fields, "top")

	}
	// Field "allowedTimeGrainsMs"
	if fields["allowedTimeGrainsMs"] != nil {
		if string(fields["allowedTimeGrainsMs"]) != "null" {

			if err := json.Unmarshal(fields["allowedTimeGrainsMs"], &resource.AllowedTimeGrainsMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allowedTimeGrainsMs", err)...)
			}

		}
		delete(fields, "allowedTimeGrainsMs")

	}
	// Field "alias"
	if fields["alias"] != nil {
		if string(fields["alias"]) != "null" {
			if err := json.Unmarshal(fields["alias"], &resource.Alias); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alias", err)...)
			}

		}
		delete(fields, "alias")

	}
	// Field "timeGrainUnit"
	if fields["timeGrainUnit"] != nil {
		if string(fields["timeGrainUnit"]) != "null" {
			if err := json.Unmarshal(fields["timeGrainUnit"], &resource.TimeGrainUnit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeGrainUnit", err)...)
			}

		}
		delete(fields, "timeGrainUnit")

	}
	// Field "dimension"
	if fields["dimension"] != nil {
		if string(fields["dimension"]) != "null" {
			if err := json.Unmarshal(fields["dimension"], &resource.Dimension); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dimension", err)...)
			}

		}
		delete(fields, "dimension")

	}
	// Field "dimensionFilter"
	if fields["dimensionFilter"] != nil {
		if string(fields["dimensionFilter"]) != "null" {
			if err := json.Unmarshal(fields["dimensionFilter"], &resource.DimensionFilter); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dimensionFilter", err)...)
			}

		}
		delete(fields, "dimensionFilter")

	}
	// Field "metricDefinition"
	if fields["metricDefinition"] != nil {
		if string(fields["metricDefinition"]) != "null" {
			if err := json.Unmarshal(fields["metricDefinition"], &resource.MetricDefinition); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricDefinition", err)...)
			}

		}
		delete(fields, "metricDefinition")

	}
	// Field "resourceUri"
	if fields["resourceUri"] != nil {
		if string(fields["resourceUri"]) != "null" {
			if err := json.Unmarshal(fields["resourceUri"], &resource.ResourceUri); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceUri", err)...)
			}

		}
		delete(fields, "resourceUri")

	}
	// Field "resourceGroup"
	if fields["resourceGroup"] != nil {
		if string(fields["resourceGroup"]) != "null" {
			if err := json.Unmarshal(fields["resourceGroup"], &resource.ResourceGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceGroup", err)...)
			}

		}
		delete(fields, "resourceGroup")

	}
	// Field "resourceName"
	if fields["resourceName"] != nil {
		if string(fields["resourceName"]) != "null" {
			if err := json.Unmarshal(fields["resourceName"], &resource.ResourceName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceName", err)...)
			}

		}
		delete(fields, "resourceName")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureMetricQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AzureMetricQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `AzureMetricQuery` fields for violations and returns them.
func (resource AzureMetricQuery) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Resources {
		if err := resource.Resources[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resources["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.DimensionFilters {
		if err := resource.DimensionFilters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("dimensionFilters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureLogsQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureLogsQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}

		}
		delete(fields, "query")

	}
	// Field "resultFormat"
	if fields["resultFormat"] != nil {
		if string(fields["resultFormat"]) != "null" {
			if err := json.Unmarshal(fields["resultFormat"], &resource.ResultFormat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultFormat", err)...)
			}

		}
		delete(fields, "resultFormat")

	}
	// Field "resources"
	if fields["resources"] != nil {
		if string(fields["resources"]) != "null" {

			if err := json.Unmarshal(fields["resources"], &resource.Resources); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resources", err)...)
			}

		}
		delete(fields, "resources")

	}
	// Field "dashboardTime"
	if fields["dashboardTime"] != nil {
		if string(fields["dashboardTime"]) != "null" {
			if err := json.Unmarshal(fields["dashboardTime"], &resource.DashboardTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dashboardTime", err)...)
			}

		}
		delete(fields, "dashboardTime")

	}
	// Field "timeColumn"
	if fields["timeColumn"] != nil {
		if string(fields["timeColumn"]) != "null" {
			if err := json.Unmarshal(fields["timeColumn"], &resource.TimeColumn); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeColumn", err)...)
			}

		}
		delete(fields, "timeColumn")

	}
	// Field "basicLogsQuery"
	if fields["basicLogsQuery"] != nil {
		if string(fields["basicLogsQuery"]) != "null" {
			if err := json.Unmarshal(fields["basicLogsQuery"], &resource.BasicLogsQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("basicLogsQuery", err)...)
			}

		}
		delete(fields, "basicLogsQuery")

	}
	// Field "workspace"
	if fields["workspace"] != nil {
		if string(fields["workspace"]) != "null" {
			if err := json.Unmarshal(fields["workspace"], &resource.Workspace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("workspace", err)...)
			}

		}
		delete(fields, "workspace")

	}
	// Field "resource"
	if fields["resource"] != nil {
		if string(fields["resource"]) != "null" {
			if err := json.Unmarshal(fields["resource"], &resource.Resource); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resource", err)...)
			}

		}
		delete(fields, "resource")

	}
	// Field "intersectTime"
	if fields["intersectTime"] != nil {
		if string(fields["intersectTime"]) != "null" {
			if err := json.Unmarshal(fields["intersectTime"], &resource.IntersectTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intersectTime", err)...)
			}

		}
		delete(fields, "intersectTime")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureLogsQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AzureLogsQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `AzureLogsQuery` fields for violations and returns them.
func (resource AzureLogsQuery) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureTracesQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureTracesQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "resultFormat"
	if fields["resultFormat"] != nil {
		if string(fields["resultFormat"]) != "null" {
			if err := json.Unmarshal(fields["resultFormat"], &resource.ResultFormat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultFormat", err)...)
			}

		}
		delete(fields, "resultFormat")

	}
	// Field "resources"
	if fields["resources"] != nil {
		if string(fields["resources"]) != "null" {

			if err := json.Unmarshal(fields["resources"], &resource.Resources); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resources", err)...)
			}

		}
		delete(fields, "resources")

	}
	// Field "operationId"
	if fields["operationId"] != nil {
		if string(fields["operationId"]) != "null" {
			if err := json.Unmarshal(fields["operationId"], &resource.OperationId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operationId", err)...)
			}

		}
		delete(fields, "operationId")

	}
	// Field "traceTypes"
	if fields["traceTypes"] != nil {
		if string(fields["traceTypes"]) != "null" {

			if err := json.Unmarshal(fields["traceTypes"], &resource.TraceTypes); err != nil {
				errs = append(errs, cog.MakeBuildErrors("traceTypes", err)...)
			}

		}
		delete(fields, "traceTypes")

	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["filters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AzureTracesFilter

				result1 = AzureTracesFilter{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Filters = append(resource.Filters, result1)
			}

		}
		delete(fields, "filters")

	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}

		}
		delete(fields, "query")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureTracesQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AzureTracesQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `AzureTracesQuery` fields for violations and returns them.
func (resource AzureTracesQuery) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Filters {
		if err := resource.Filters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AzureTracesFilter struct {
	// Property name, auto-populated based on available traces.
	Property string `json:"property"`
	// Comparison operator to use. Either equals or not equals.
	Operation string `json:"operation"`
	// Values to filter by.
	Filters []string `json:"filters"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureTracesFilter` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureTracesFilter) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "property"
	if fields["property"] != nil {
		if string(fields["property"]) != "null" {
			if err := json.Unmarshal(fields["property"], &resource.Property); err != nil {
				errs = append(errs, cog.MakeBuildErrors("property", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is null"))...)

		}
		delete(fields, "property")
	} else {
		errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is missing from input"))...)
	}
	// Field "operation"
	if fields["operation"] != nil {
		if string(fields["operation"]) != "null" {
			if err := json.Unmarshal(fields["operation"], &resource.Operation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operation", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operation", errors.New("required field is null"))...)

		}
		delete(fields, "operation")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operation", errors.New("required field is missing from input"))...)
	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			if err := json.Unmarshal(fields["filters"], &resource.Filters); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filters", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is null"))...)

		}
		delete(fields, "filters")
	} else {
		errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureTracesFilter", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AzureTracesFilter` objects.
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

// Validate checks all the validation constraints that may be defined on `AzureTracesFilter` fields for violations and returns them.
func (resource AzureTracesFilter) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureResourceGraphQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureResourceGraphQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}

		}
		delete(fields, "query")

	}
	// Field "resultFormat"
	if fields["resultFormat"] != nil {
		if string(fields["resultFormat"]) != "null" {
			if err := json.Unmarshal(fields["resultFormat"], &resource.ResultFormat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultFormat", err)...)
			}

		}
		delete(fields, "resultFormat")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureResourceGraphQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AzureResourceGraphQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `AzureResourceGraphQuery` fields for violations and returns them.
func (resource AzureResourceGraphQuery) Validate() error {
	return nil
}

type AzureMonitorResource struct {
	Subscription    *string `json:"subscription,omitempty"`
	ResourceGroup   *string `json:"resourceGroup,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty"`
	MetricNamespace *string `json:"metricNamespace,omitempty"`
	Region          *string `json:"region,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMonitorResource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureMonitorResource) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}

		}
		delete(fields, "subscription")

	}
	// Field "resourceGroup"
	if fields["resourceGroup"] != nil {
		if string(fields["resourceGroup"]) != "null" {
			if err := json.Unmarshal(fields["resourceGroup"], &resource.ResourceGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceGroup", err)...)
			}

		}
		delete(fields, "resourceGroup")

	}
	// Field "resourceName"
	if fields["resourceName"] != nil {
		if string(fields["resourceName"]) != "null" {
			if err := json.Unmarshal(fields["resourceName"], &resource.ResourceName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceName", err)...)
			}

		}
		delete(fields, "resourceName")

	}
	// Field "metricNamespace"
	if fields["metricNamespace"] != nil {
		if string(fields["metricNamespace"]) != "null" {
			if err := json.Unmarshal(fields["metricNamespace"], &resource.MetricNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricNamespace", err)...)
			}

		}
		delete(fields, "metricNamespace")

	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}

		}
		delete(fields, "region")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureMonitorResource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AzureMonitorResource` objects.
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

// Validate checks all the validation constraints that may be defined on `AzureMonitorResource` fields for violations and returns them.
func (resource AzureMonitorResource) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMetricDimension` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AzureMetricDimension) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "dimension"
	if fields["dimension"] != nil {
		if string(fields["dimension"]) != "null" {
			if err := json.Unmarshal(fields["dimension"], &resource.Dimension); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dimension", err)...)
			}

		}
		delete(fields, "dimension")

	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {
			if err := json.Unmarshal(fields["operator"], &resource.Operator); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}

		}
		delete(fields, "operator")

	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			if err := json.Unmarshal(fields["filters"], &resource.Filters); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filters", err)...)
			}

		}
		delete(fields, "filters")

	}
	// Field "filter"
	if fields["filter"] != nil {
		if string(fields["filter"]) != "null" {
			if err := json.Unmarshal(fields["filter"], &resource.Filter); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filter", err)...)
			}

		}
		delete(fields, "filter")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AzureMetricDimension", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AzureMetricDimension` objects.
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

// Validate checks all the validation constraints that may be defined on `AzureMetricDimension` fields for violations and returns them.
func (resource AzureMetricDimension) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseGrafanaTemplateVariableQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BaseGrafanaTemplateVariableQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BaseGrafanaTemplateVariableQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BaseGrafanaTemplateVariableQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `BaseGrafanaTemplateVariableQuery` fields for violations and returns them.
func (resource BaseGrafanaTemplateVariableQuery) Validate() error {
	return nil
}

type UnknownQuery struct {
	RawQuery *string `json:"rawQuery,omitempty"`
	Kind     string  `json:"kind"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `UnknownQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *UnknownQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("UnknownQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `UnknownQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `UnknownQuery` fields for violations and returns them.
func (resource UnknownQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "UnknownQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == UnknownQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AppInsightsMetricNameQuery struct {
	RawQuery *string `json:"rawQuery,omitempty"`
	Kind     string  `json:"kind"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AppInsightsMetricNameQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AppInsightsMetricNameQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AppInsightsMetricNameQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AppInsightsMetricNameQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `AppInsightsMetricNameQuery` fields for violations and returns them.
func (resource AppInsightsMetricNameQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "AppInsightsMetricNameQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == AppInsightsMetricNameQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type AppInsightsGroupByQuery struct {
	RawQuery   *string `json:"rawQuery,omitempty"`
	Kind       string  `json:"kind"`
	MetricName string  `json:"metricName"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AppInsightsGroupByQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AppInsightsGroupByQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "metricName"
	if fields["metricName"] != nil {
		if string(fields["metricName"]) != "null" {
			if err := json.Unmarshal(fields["metricName"], &resource.MetricName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metricName", errors.New("required field is null"))...)

		}
		delete(fields, "metricName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metricName", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AppInsightsGroupByQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AppInsightsGroupByQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `AppInsightsGroupByQuery` fields for violations and returns them.
func (resource AppInsightsGroupByQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "AppInsightsGroupByQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == AppInsightsGroupByQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type SubscriptionsQuery struct {
	RawQuery *string `json:"rawQuery,omitempty"`
	Kind     string  `json:"kind"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SubscriptionsQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SubscriptionsQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SubscriptionsQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SubscriptionsQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `SubscriptionsQuery` fields for violations and returns them.
func (resource SubscriptionsQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "SubscriptionsQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == SubscriptionsQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ResourceGroupsQuery struct {
	RawQuery     *string `json:"rawQuery,omitempty"`
	Kind         string  `json:"kind"`
	Subscription string  `json:"subscription"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceGroupsQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ResourceGroupsQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is null"))...)

		}
		delete(fields, "subscription")
	} else {
		errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ResourceGroupsQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ResourceGroupsQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `ResourceGroupsQuery` fields for violations and returns them.
func (resource ResourceGroupsQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "ResourceGroupsQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == ResourceGroupsQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ResourceNamesQuery struct {
	RawQuery        *string `json:"rawQuery,omitempty"`
	Kind            string  `json:"kind"`
	Subscription    string  `json:"subscription"`
	ResourceGroup   string  `json:"resourceGroup"`
	MetricNamespace string  `json:"metricNamespace"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceNamesQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ResourceNamesQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is null"))...)

		}
		delete(fields, "subscription")
	} else {
		errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is missing from input"))...)
	}
	// Field "resourceGroup"
	if fields["resourceGroup"] != nil {
		if string(fields["resourceGroup"]) != "null" {
			if err := json.Unmarshal(fields["resourceGroup"], &resource.ResourceGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceGroup", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is null"))...)

		}
		delete(fields, "resourceGroup")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is missing from input"))...)
	}
	// Field "metricNamespace"
	if fields["metricNamespace"] != nil {
		if string(fields["metricNamespace"]) != "null" {
			if err := json.Unmarshal(fields["metricNamespace"], &resource.MetricNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricNamespace", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metricNamespace", errors.New("required field is null"))...)

		}
		delete(fields, "metricNamespace")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metricNamespace", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ResourceNamesQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ResourceNamesQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `ResourceNamesQuery` fields for violations and returns them.
func (resource ResourceNamesQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "ResourceNamesQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == ResourceNamesQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MetricNamespaceQuery struct {
	RawQuery        *string `json:"rawQuery,omitempty"`
	Kind            string  `json:"kind"`
	Subscription    string  `json:"subscription"`
	ResourceGroup   string  `json:"resourceGroup"`
	MetricNamespace *string `json:"metricNamespace,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricNamespaceQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricNamespaceQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is null"))...)

		}
		delete(fields, "subscription")
	} else {
		errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is missing from input"))...)
	}
	// Field "resourceGroup"
	if fields["resourceGroup"] != nil {
		if string(fields["resourceGroup"]) != "null" {
			if err := json.Unmarshal(fields["resourceGroup"], &resource.ResourceGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceGroup", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is null"))...)

		}
		delete(fields, "resourceGroup")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is missing from input"))...)
	}
	// Field "metricNamespace"
	if fields["metricNamespace"] != nil {
		if string(fields["metricNamespace"]) != "null" {
			if err := json.Unmarshal(fields["metricNamespace"], &resource.MetricNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricNamespace", err)...)
			}

		}
		delete(fields, "metricNamespace")

	}
	// Field "resourceName"
	if fields["resourceName"] != nil {
		if string(fields["resourceName"]) != "null" {
			if err := json.Unmarshal(fields["resourceName"], &resource.ResourceName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceName", err)...)
			}

		}
		delete(fields, "resourceName")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricNamespaceQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricNamespaceQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `MetricNamespaceQuery` fields for violations and returns them.
func (resource MetricNamespaceQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "MetricNamespaceQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == MetricNamespaceQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricDefinitionsQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricDefinitionsQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is null"))...)

		}
		delete(fields, "subscription")
	} else {
		errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is missing from input"))...)
	}
	// Field "resourceGroup"
	if fields["resourceGroup"] != nil {
		if string(fields["resourceGroup"]) != "null" {
			if err := json.Unmarshal(fields["resourceGroup"], &resource.ResourceGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceGroup", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is null"))...)

		}
		delete(fields, "resourceGroup")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is missing from input"))...)
	}
	// Field "metricNamespace"
	if fields["metricNamespace"] != nil {
		if string(fields["metricNamespace"]) != "null" {
			if err := json.Unmarshal(fields["metricNamespace"], &resource.MetricNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricNamespace", err)...)
			}

		}
		delete(fields, "metricNamespace")

	}
	// Field "resourceName"
	if fields["resourceName"] != nil {
		if string(fields["resourceName"]) != "null" {
			if err := json.Unmarshal(fields["resourceName"], &resource.ResourceName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceName", err)...)
			}

		}
		delete(fields, "resourceName")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricDefinitionsQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricDefinitionsQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `MetricDefinitionsQuery` fields for violations and returns them.
func (resource MetricDefinitionsQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "MetricDefinitionsQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == MetricDefinitionsQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MetricNamesQuery struct {
	RawQuery        *string `json:"rawQuery,omitempty"`
	Kind            string  `json:"kind"`
	Subscription    string  `json:"subscription"`
	ResourceGroup   string  `json:"resourceGroup"`
	ResourceName    string  `json:"resourceName"`
	MetricNamespace string  `json:"metricNamespace"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricNamesQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricNamesQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is null"))...)

		}
		delete(fields, "subscription")
	} else {
		errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is missing from input"))...)
	}
	// Field "resourceGroup"
	if fields["resourceGroup"] != nil {
		if string(fields["resourceGroup"]) != "null" {
			if err := json.Unmarshal(fields["resourceGroup"], &resource.ResourceGroup); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceGroup", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is null"))...)

		}
		delete(fields, "resourceGroup")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resourceGroup", errors.New("required field is missing from input"))...)
	}
	// Field "resourceName"
	if fields["resourceName"] != nil {
		if string(fields["resourceName"]) != "null" {
			if err := json.Unmarshal(fields["resourceName"], &resource.ResourceName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resourceName", errors.New("required field is null"))...)

		}
		delete(fields, "resourceName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resourceName", errors.New("required field is missing from input"))...)
	}
	// Field "metricNamespace"
	if fields["metricNamespace"] != nil {
		if string(fields["metricNamespace"]) != "null" {
			if err := json.Unmarshal(fields["metricNamespace"], &resource.MetricNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricNamespace", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metricNamespace", errors.New("required field is null"))...)

		}
		delete(fields, "metricNamespace")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metricNamespace", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricNamesQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricNamesQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `MetricNamesQuery` fields for violations and returns them.
func (resource MetricNamesQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "MetricNamesQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == MetricNamesQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type WorkspacesQuery struct {
	RawQuery     *string `json:"rawQuery,omitempty"`
	Kind         string  `json:"kind"`
	Subscription string  `json:"subscription"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `WorkspacesQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *WorkspacesQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "rawQuery"
	if fields["rawQuery"] != nil {
		if string(fields["rawQuery"]) != "null" {
			if err := json.Unmarshal(fields["rawQuery"], &resource.RawQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawQuery", err)...)
			}

		}
		delete(fields, "rawQuery")

	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "subscription"
	if fields["subscription"] != nil {
		if string(fields["subscription"]) != "null" {
			if err := json.Unmarshal(fields["subscription"], &resource.Subscription); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subscription", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is null"))...)

		}
		delete(fields, "subscription")
	} else {
		errs = append(errs, cog.MakeBuildErrors("subscription", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("WorkspacesQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `WorkspacesQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `WorkspacesQuery` fields for violations and returns them.
func (resource WorkspacesQuery) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Kind == "WorkspacesQuery") {
		errs = append(errs, cog.MakeBuildErrors(
			"kind",
			errors.New("must be == WorkspacesQuery"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// MarshalJSON implements a custom JSON marshalling logic to encode `AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery` as JSON.
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

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) UnmarshalJSONStrict(raw []byte) error {
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
		return fmt.Errorf("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "AppInsightsGroupByQuery":
		appInsightsGroupByQuery := &AppInsightsGroupByQuery{}
		if err := appInsightsGroupByQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.AppInsightsGroupByQuery = appInsightsGroupByQuery
		return nil
	case "AppInsightsMetricNameQuery":
		appInsightsMetricNameQuery := &AppInsightsMetricNameQuery{}
		if err := appInsightsMetricNameQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.AppInsightsMetricNameQuery = appInsightsMetricNameQuery
		return nil
	case "MetricDefinitionsQuery":
		metricDefinitionsQuery := &MetricDefinitionsQuery{}
		if err := metricDefinitionsQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MetricDefinitionsQuery = metricDefinitionsQuery
		return nil
	case "MetricNamesQuery":
		metricNamesQuery := &MetricNamesQuery{}
		if err := metricNamesQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MetricNamesQuery = metricNamesQuery
		return nil
	case "MetricNamespaceQuery":
		metricNamespaceQuery := &MetricNamespaceQuery{}
		if err := metricNamespaceQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MetricNamespaceQuery = metricNamespaceQuery
		return nil
	case "ResourceGroupsQuery":
		resourceGroupsQuery := &ResourceGroupsQuery{}
		if err := resourceGroupsQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ResourceGroupsQuery = resourceGroupsQuery
		return nil
	case "ResourceNamesQuery":
		resourceNamesQuery := &ResourceNamesQuery{}
		if err := resourceNamesQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ResourceNamesQuery = resourceNamesQuery
		return nil
	case "SubscriptionsQuery":
		subscriptionsQuery := &SubscriptionsQuery{}
		if err := subscriptionsQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.SubscriptionsQuery = subscriptionsQuery
		return nil
	case "UnknownQuery":
		unknownQuery := &UnknownQuery{}
		if err := unknownQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.UnknownQuery = unknownQuery
		return nil
	case "WorkspacesQuery":
		workspacesQuery := &WorkspacesQuery{}
		if err := workspacesQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.WorkspacesQuery = workspacesQuery
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// Equals tests the equality of two `AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery` fields for violations and returns them.
func (resource AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) Validate() error {
	var errs cog.BuildErrors
	if resource.AppInsightsMetricNameQuery != nil {
		if err := resource.AppInsightsMetricNameQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("AppInsightsMetricNameQuery", err)...)
		}
	}
	if resource.AppInsightsGroupByQuery != nil {
		if err := resource.AppInsightsGroupByQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("AppInsightsGroupByQuery", err)...)
		}
	}
	if resource.SubscriptionsQuery != nil {
		if err := resource.SubscriptionsQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("SubscriptionsQuery", err)...)
		}
	}
	if resource.ResourceGroupsQuery != nil {
		if err := resource.ResourceGroupsQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ResourceGroupsQuery", err)...)
		}
	}
	if resource.ResourceNamesQuery != nil {
		if err := resource.ResourceNamesQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ResourceNamesQuery", err)...)
		}
	}
	if resource.MetricNamespaceQuery != nil {
		if err := resource.MetricNamespaceQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MetricNamespaceQuery", err)...)
		}
	}
	if resource.MetricDefinitionsQuery != nil {
		if err := resource.MetricDefinitionsQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MetricDefinitionsQuery", err)...)
		}
	}
	if resource.MetricNamesQuery != nil {
		if err := resource.MetricNamesQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MetricNamesQuery", err)...)
		}
	}
	if resource.WorkspacesQuery != nil {
		if err := resource.WorkspacesQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("WorkspacesQuery", err)...)
		}
	}
	if resource.UnknownQuery != nil {
		if err := resource.UnknownQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("UnknownQuery", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
