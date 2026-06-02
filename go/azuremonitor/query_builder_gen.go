// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

var _ cog.Builder[dashboardv2beta1.DataQueryKind] = (*QueryBuilder)(nil)

type QueryBuilder struct {
	internal *dashboardv2beta1.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryBuilder() *QueryBuilder {
	resource := dashboardv2beta1.NewDataQueryKind()
	builder := &QueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "grafana-azure-monitor-datasource"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("azuremonitor.query", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) RecordError(path string, err error) *QueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryBuilder) Version(version string) *QueryBuilder {
	builder.internal.Version = version

	return builder
}

func (builder *QueryBuilder) Labels(labels map[string]string) *QueryBuilder {
	builder.internal.Labels = labels

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).QueryType = &queryType

	return builder
}

// Azure subscription containing the resource(s) to be queried.
// Also used for template variable queries
func (builder *QueryBuilder) Subscription(subscription string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).Subscription = &subscription

	return builder
}

// Subscriptions to be queried via Azure Resource Graph.
func (builder *QueryBuilder) Subscriptions(subscriptions []string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).Subscriptions = subscriptions

	return builder
}

// Azure Monitor Metrics sub-query properties.
func (builder *QueryBuilder) AzureMonitor(azureMonitor cog.Builder[MetricQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	azureMonitorResource, err := azureMonitor.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*MonitorQuery).AzureMonitor = &azureMonitorResource

	return builder
}

// Azure Monitor Logs sub-query properties.
func (builder *QueryBuilder) AzureLogAnalytics(azureLogAnalytics cog.Builder[LogsQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	azureLogAnalyticsResource, err := azureLogAnalytics.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*MonitorQuery).AzureLogAnalytics = &azureLogAnalyticsResource

	return builder
}

// Azure Resource Graph sub-query properties.
func (builder *QueryBuilder) AzureResourceGraph(azureResourceGraph cog.Builder[ResourceGraphQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	azureResourceGraphResource, err := azureResourceGraph.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*MonitorQuery).AzureResourceGraph = &azureResourceGraphResource

	return builder
}

// Application Insights Traces sub-query properties.
func (builder *QueryBuilder) AzureTraces(azureTraces cog.Builder[TracesQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	azureTracesResource, err := azureTraces.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*MonitorQuery).AzureTraces = &azureTracesResource

	return builder
}

// @deprecated Legacy template variable support.
func (builder *QueryBuilder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn GrafanaTemplateVariableQuery) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).GrafanaTemplateVariableFn = &grafanaTemplateVariableFn

	return builder
}

// Resource group used in template variable queries
func (builder *QueryBuilder) ResourceGroup(resourceGroup string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).ResourceGroup = &resourceGroup

	return builder
}

// Namespace used in template variable queries
func (builder *QueryBuilder) Namespace(namespace string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).Namespace = &namespace

	return builder
}

// Resource used in template variable queries
func (builder *QueryBuilder) Resource(resource string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).Resource = &resource

	return builder
}

// Region used in template variable queries
func (builder *QueryBuilder) Region(region string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).Region = &region

	return builder
}

// Custom namespace used in template variable queries
func (builder *QueryBuilder) CustomNamespace(customNamespace string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).CustomNamespace = &customNamespace

	return builder
}

// Used only for exemplar queries from Prometheus
func (builder *QueryBuilder) Query(query string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewMonitorQuery()
	}
	builder.internal.Spec.(*MonitorQuery).Query = &query

	return builder
}
