// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricNamespaceQuery] = (*MetricNamespaceQueryBuilder)(nil)

type MetricNamespaceQueryBuilder struct {
	internal *MetricNamespaceQuery
	errors   map[string]cog.BuildErrors
}

func NewMetricNamespaceQueryBuilder() *MetricNamespaceQueryBuilder {
	resource := &MetricNamespaceQuery{}
	builder := &MetricNamespaceQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "MetricNamespaceQuery"

	return builder
}

func (builder *MetricNamespaceQueryBuilder) Build() (MetricNamespaceQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricNamespaceQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *MetricNamespaceQueryBuilder) RawQuery(rawQuery string) *MetricNamespaceQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *MetricNamespaceQueryBuilder) Subscription(subscription string) *MetricNamespaceQueryBuilder {
	builder.internal.Subscription = subscription

	return builder
}

func (builder *MetricNamespaceQueryBuilder) ResourceGroup(resourceGroup string) *MetricNamespaceQueryBuilder {
	builder.internal.ResourceGroup = resourceGroup

	return builder
}

func (builder *MetricNamespaceQueryBuilder) MetricNamespace(metricNamespace string) *MetricNamespaceQueryBuilder {
	builder.internal.MetricNamespace = &metricNamespace

	return builder
}

func (builder *MetricNamespaceQueryBuilder) ResourceName(resourceName string) *MetricNamespaceQueryBuilder {
	builder.internal.ResourceName = &resourceName

	return builder
}

func (builder *MetricNamespaceQueryBuilder) applyDefaults() {
}
