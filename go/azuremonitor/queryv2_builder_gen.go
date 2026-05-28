// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

var _ cog.Builder[dashboardv2.DataQueryKind] = (*QueryV2Builder)(nil)

type QueryV2Builder struct {
	internal *dashboardv2.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryV2Builder() *QueryV2Builder {
	resource := dashboardv2.NewDataQueryKind()
	builder := &QueryV2Builder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "grafana-azure-monitor-datasource"

	return builder
}

func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.DataQueryKind{}, cog.MakeBuildErrors("azuremonitor.queryV2", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryV2Builder) RecordError(path string, err error) *QueryV2Builder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryV2Builder) Version(version string) *QueryV2Builder {
	builder.internal.Version = version

	return builder
}

func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder {
	builder.internal.Labels = labels

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder {
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
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).QueryType = &queryType

	return builder
}

// Azure subscription containing the resource(s) to be queried.
// Also used for template variable queries
func (builder *QueryV2Builder) Subscription(subscription string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).Subscription = &subscription

	return builder
}

// Subscriptions to be queried via Azure Resource Graph.
func (builder *QueryV2Builder) Subscriptions(subscriptions []string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).Subscriptions = subscriptions

	return builder
}

// Azure Monitor Metrics sub-query properties.
func (builder *QueryV2Builder) AzureMonitor(azureMonitor cog.Builder[AzureMetricQuery]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	azureMonitorResource, err := azureMonitor.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*AzureMonitorQuery).AzureMonitor = &azureMonitorResource

	return builder
}

// Azure Monitor Logs sub-query properties.
func (builder *QueryV2Builder) AzureLogAnalytics(azureLogAnalytics cog.Builder[AzureLogsQuery]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	azureLogAnalyticsResource, err := azureLogAnalytics.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*AzureMonitorQuery).AzureLogAnalytics = &azureLogAnalyticsResource

	return builder
}

// Azure Resource Graph sub-query properties.
func (builder *QueryV2Builder) AzureResourceGraph(azureResourceGraph cog.Builder[AzureResourceGraphQuery]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	azureResourceGraphResource, err := azureResourceGraph.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*AzureMonitorQuery).AzureResourceGraph = &azureResourceGraphResource

	return builder
}

// Application Insights Traces sub-query properties.
func (builder *QueryV2Builder) AzureTraces(azureTraces cog.Builder[AzureTracesQuery]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	azureTracesResource, err := azureTraces.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*AzureMonitorQuery).AzureTraces = &azureTracesResource

	return builder
}

// @deprecated Legacy template variable support.
func (builder *QueryV2Builder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn GrafanaTemplateVariableQuery) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).GrafanaTemplateVariableFn = &grafanaTemplateVariableFn

	return builder
}

// Resource group used in template variable queries
func (builder *QueryV2Builder) ResourceGroup(resourceGroup string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).ResourceGroup = &resourceGroup

	return builder
}

// Namespace used in template variable queries
func (builder *QueryV2Builder) Namespace(namespace string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).Namespace = &namespace

	return builder
}

// Resource used in template variable queries
func (builder *QueryV2Builder) Resource(resource string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).Resource = &resource

	return builder
}

// Region used in template variable queries
func (builder *QueryV2Builder) Region(region string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).Region = &region

	return builder
}

// Custom namespace used in template variable queries
func (builder *QueryV2Builder) CustomNamespace(customNamespace string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).CustomNamespace = &customNamespace

	return builder
}

// Used only for exemplar queries from Prometheus
func (builder *QueryV2Builder) Query(query string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewAzureMonitorQuery()
	}
	builder.internal.Spec.(*AzureMonitorQuery).Query = &query

	return builder
}
