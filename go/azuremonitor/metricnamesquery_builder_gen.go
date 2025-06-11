// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricNamesQuery] = (*MetricNamesQueryBuilder)(nil)

type MetricNamesQueryBuilder struct {
	internal *MetricNamesQuery
	errors   cog.BuildErrors
}

func NewMetricNamesQueryBuilder() *MetricNamesQueryBuilder {
	resource := NewMetricNamesQuery()
	builder := &MetricNamesQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "MetricNamesQuery"

	return builder
}

func (builder *MetricNamesQueryBuilder) Build() (MetricNamesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricNamesQuery{}, err
	}

	if len(builder.errors) > 0 {
		return MetricNamesQuery{}, cog.MakeBuildErrors("azuremonitor.metricNamesQuery", builder.errors)
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
