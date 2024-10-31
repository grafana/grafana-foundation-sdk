// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricNamesQuery] = (*MetricNamesQueryBuilder)(nil)

type MetricNamesQueryBuilder struct {
	internal *MetricNamesQuery
	errors   map[string]cog.BuildErrors
}

func NewMetricNamesQueryBuilder() *MetricNamesQueryBuilder {
	resource := &MetricNamesQuery{}
	builder := &MetricNamesQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "MetricNamesQuery"

	return builder
}

func (builder *MetricNamesQueryBuilder) Build() (MetricNamesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricNamesQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *MetricNamesQueryBuilder) RawQuery(rawQuery string) *MetricNamesQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *MetricNamesQueryBuilder) Subscription(subscription string) *MetricNamesQueryBuilder {
	builder.internal.Subscription = subscription

	return builder
}

func (builder *MetricNamesQueryBuilder) ResourceGroup(resourceGroup string) *MetricNamesQueryBuilder {
	builder.internal.ResourceGroup = resourceGroup

	return builder
}

func (builder *MetricNamesQueryBuilder) ResourceName(resourceName string) *MetricNamesQueryBuilder {
	builder.internal.ResourceName = resourceName

	return builder
}

func (builder *MetricNamesQueryBuilder) MetricNamespace(metricNamespace string) *MetricNamesQueryBuilder {
	builder.internal.MetricNamespace = metricNamespace

	return builder
}

func (builder *MetricNamesQueryBuilder) applyDefaults() {
}
