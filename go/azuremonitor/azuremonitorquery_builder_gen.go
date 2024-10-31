// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*AzureMonitorQueryBuilder)(nil)

type AzureMonitorQueryBuilder struct {
	internal *AzureMonitorQuery
	errors   map[string]cog.BuildErrors
}

func NewAzureMonitorQueryBuilder() *AzureMonitorQueryBuilder {
	resource := &AzureMonitorQuery{}
	builder := &AzureMonitorQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureMonitorQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureMonitorQuery{}, err
	}

	return *builder.internal, nil
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *AzureMonitorQueryBuilder) RefId(refId string) *AzureMonitorQueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// Note this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *AzureMonitorQueryBuilder) Hide(hide bool) *AzureMonitorQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *AzureMonitorQueryBuilder) QueryType(queryType string) *AzureMonitorQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// Azure subscription containing the resource(s) to be queried.
func (builder *AzureMonitorQueryBuilder) Subscription(subscription string) *AzureMonitorQueryBuilder {
	builder.internal.Subscription = &subscription

	return builder
}

// Subscriptions to be queried via Azure Resource Graph.
func (builder *AzureMonitorQueryBuilder) Subscriptions(subscriptions []string) *AzureMonitorQueryBuilder {
	builder.internal.Subscriptions = subscriptions

	return builder
}

// Azure Monitor Metrics sub-query properties.
func (builder *AzureMonitorQueryBuilder) AzureMonitor(azureMonitor cog.Builder[AzureMetricQuery]) *AzureMonitorQueryBuilder {
	azureMonitorResource, err := azureMonitor.Build()
	if err != nil {
		builder.errors["azureMonitor"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.AzureMonitor = &azureMonitorResource

	return builder
}

// Azure Monitor Logs sub-query properties.
func (builder *AzureMonitorQueryBuilder) AzureLogAnalytics(azureLogAnalytics cog.Builder[AzureLogsQuery]) *AzureMonitorQueryBuilder {
	azureLogAnalyticsResource, err := azureLogAnalytics.Build()
	if err != nil {
		builder.errors["azureLogAnalytics"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.AzureLogAnalytics = &azureLogAnalyticsResource

	return builder
}

// Azure Resource Graph sub-query properties.
func (builder *AzureMonitorQueryBuilder) AzureResourceGraph(azureResourceGraph cog.Builder[AzureResourceGraphQuery]) *AzureMonitorQueryBuilder {
	azureResourceGraphResource, err := azureResourceGraph.Build()
	if err != nil {
		builder.errors["azureResourceGraph"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.AzureResourceGraph = &azureResourceGraphResource

	return builder
}

// Application Insights Traces sub-query properties.
func (builder *AzureMonitorQueryBuilder) AzureTraces(azureTraces cog.Builder[AzureTracesQuery]) *AzureMonitorQueryBuilder {
	azureTracesResource, err := azureTraces.Build()
	if err != nil {
		builder.errors["azureTraces"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.AzureTraces = &azureTracesResource

	return builder
}

// @deprecated Legacy template variable support.
func (builder *AzureMonitorQueryBuilder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn GrafanaTemplateVariableQuery) *AzureMonitorQueryBuilder {
	builder.internal.GrafanaTemplateVariableFn = &grafanaTemplateVariableFn

	return builder
}

// Template variables params. These exist for backwards compatiblity with legacy template variables.
func (builder *AzureMonitorQueryBuilder) ResourceGroup(resourceGroup string) *AzureMonitorQueryBuilder {
	builder.internal.ResourceGroup = &resourceGroup

	return builder
}

func (builder *AzureMonitorQueryBuilder) Namespace(namespace string) *AzureMonitorQueryBuilder {
	builder.internal.Namespace = &namespace

	return builder
}

func (builder *AzureMonitorQueryBuilder) Resource(resource string) *AzureMonitorQueryBuilder {
	builder.internal.Resource = &resource

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *AzureMonitorQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *AzureMonitorQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Azure Monitor query type.
// queryType: #AzureQueryType
func (builder *AzureMonitorQueryBuilder) Region(region string) *AzureMonitorQueryBuilder {
	builder.internal.Region = &region

	return builder
}

func (builder *AzureMonitorQueryBuilder) applyDefaults() {
}
