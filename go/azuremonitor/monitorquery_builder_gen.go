// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[variants.Dataquery] = (*MonitorQueryBuilder)(nil)

type MonitorQueryBuilder struct {
	internal *MonitorQuery
	errors   cog.BuildErrors
}

func NewMonitorQueryBuilder() *MonitorQueryBuilder {
	resource := NewMonitorQuery()
	builder := &MonitorQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MonitorQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MonitorQuery{}, err
	}

	if len(builder.errors) > 0 {
		return MonitorQuery{}, cog.MakeBuildErrors("azuremonitor.monitorQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MonitorQueryBuilder) RecordError(path string, err error) *MonitorQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *MonitorQueryBuilder) RefId(refId string) *MonitorQueryBuilder {
	builder.internal.RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *MonitorQueryBuilder) Hide(hide bool) *MonitorQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *MonitorQueryBuilder) QueryType(queryType string) *MonitorQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// Azure subscription containing the resource(s) to be queried.
// Also used for template variable queries
func (builder *MonitorQueryBuilder) Subscription(subscription string) *MonitorQueryBuilder {
	builder.internal.Subscription = &subscription

	return builder
}

// Subscriptions to be queried via Azure Resource Graph.
func (builder *MonitorQueryBuilder) Subscriptions(subscriptions []string) *MonitorQueryBuilder {
	builder.internal.Subscriptions = subscriptions

	return builder
}

// Azure Monitor Metrics sub-query properties.
func (builder *MonitorQueryBuilder) AzureMonitor(azureMonitor cog.Builder[MetricQuery]) *MonitorQueryBuilder {
	azureMonitorResource, err := azureMonitor.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.AzureMonitor = &azureMonitorResource

	return builder
}

// Azure Monitor Logs sub-query properties.
func (builder *MonitorQueryBuilder) AzureLogAnalytics(azureLogAnalytics cog.Builder[LogsQuery]) *MonitorQueryBuilder {
	azureLogAnalyticsResource, err := azureLogAnalytics.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.AzureLogAnalytics = &azureLogAnalyticsResource

	return builder
}

// Azure Resource Graph sub-query properties.
func (builder *MonitorQueryBuilder) AzureResourceGraph(azureResourceGraph cog.Builder[ResourceGraphQuery]) *MonitorQueryBuilder {
	azureResourceGraphResource, err := azureResourceGraph.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.AzureResourceGraph = &azureResourceGraphResource

	return builder
}

// Application Insights Traces sub-query properties.
func (builder *MonitorQueryBuilder) AzureTraces(azureTraces cog.Builder[TracesQuery]) *MonitorQueryBuilder {
	azureTracesResource, err := azureTraces.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.AzureTraces = &azureTracesResource

	return builder
}

// @deprecated Legacy template variable support.
func (builder *MonitorQueryBuilder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn GrafanaTemplateVariableQuery) *MonitorQueryBuilder {
	builder.internal.GrafanaTemplateVariableFn = &grafanaTemplateVariableFn

	return builder
}

// Resource group used in template variable queries
func (builder *MonitorQueryBuilder) ResourceGroup(resourceGroup string) *MonitorQueryBuilder {
	builder.internal.ResourceGroup = &resourceGroup

	return builder
}

// Namespace used in template variable queries
func (builder *MonitorQueryBuilder) Namespace(namespace string) *MonitorQueryBuilder {
	builder.internal.Namespace = &namespace

	return builder
}

// Resource used in template variable queries
func (builder *MonitorQueryBuilder) Resource(resource string) *MonitorQueryBuilder {
	builder.internal.Resource = &resource

	return builder
}

// Region used in template variable queries
func (builder *MonitorQueryBuilder) Region(region string) *MonitorQueryBuilder {
	builder.internal.Region = &region

	return builder
}

// Custom namespace used in template variable queries
func (builder *MonitorQueryBuilder) CustomNamespace(customNamespace string) *MonitorQueryBuilder {
	builder.internal.CustomNamespace = &customNamespace

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *MonitorQueryBuilder) Datasource(datasource common.DataSourceRef) *MonitorQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Used only for exemplar queries from Prometheus
func (builder *MonitorQueryBuilder) Query(query string) *MonitorQueryBuilder {
	builder.internal.Query = &query

	return builder
}
