// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricDefinitionsQuery] = (*MetricDefinitionsQueryBuilder)(nil)

// @deprecated Use MetricNamespaceQuery instead
type MetricDefinitionsQueryBuilder struct {
	internal *MetricDefinitionsQuery
	errors   map[string]cog.BuildErrors
}

func NewMetricDefinitionsQueryBuilder() *MetricDefinitionsQueryBuilder {
	resource := &MetricDefinitionsQuery{}
	builder := &MetricDefinitionsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "MetricDefinitionsQuery"

	return builder
}

func (builder *MetricDefinitionsQueryBuilder) Build() (MetricDefinitionsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricDefinitionsQuery{}, err
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

func (builder *MetricDefinitionsQueryBuilder) applyDefaults() {
}
