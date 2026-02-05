// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricDefinitionsQuery] = (*MetricDefinitionsQueryBuilder)(nil)

// @deprecated Use MetricNamespaceQuery instead
type MetricDefinitionsQueryBuilder struct {
	internal *MetricDefinitionsQuery
	errors   cog.BuildErrors
}

func NewMetricDefinitionsQueryBuilder() *MetricDefinitionsQueryBuilder {
	resource := NewMetricDefinitionsQuery()
	builder := &MetricDefinitionsQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "MetricDefinitionsQuery"

	return builder
}

func (builder *MetricDefinitionsQueryBuilder) Build() (MetricDefinitionsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricDefinitionsQuery{}, err
	}

	if len(builder.errors) > 0 {
		return MetricDefinitionsQuery{}, cog.MakeBuildErrors("azuremonitor.metricDefinitionsQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricDefinitionsQueryBuilder) RawQuery(rawQuery string) *MetricDefinitionsQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *MetricDefinitionsQueryBuilder) Subscription(subscription string) *MetricDefinitionsQueryBuilder {
	builder.internal.Subscription = subscription

	return builder
}

func (builder *MetricDefinitionsQueryBuilder) ResourceGroup(resourceGroup string) *MetricDefinitionsQueryBuilder {
	builder.internal.ResourceGroup = resourceGroup

	return builder
}

func (builder *MetricDefinitionsQueryBuilder) MetricNamespace(metricNamespace string) *MetricDefinitionsQueryBuilder {
	builder.internal.MetricNamespace = &metricNamespace

	return builder
}

func (builder *MetricDefinitionsQueryBuilder) ResourceName(resourceName string) *MetricDefinitionsQueryBuilder {
	builder.internal.ResourceName = &resourceName

	return builder
}
