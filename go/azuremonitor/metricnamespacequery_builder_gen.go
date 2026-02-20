// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricNamespaceQuery] = (*MetricNamespaceQueryBuilder)(nil)

type MetricNamespaceQueryBuilder struct {
	internal *MetricNamespaceQuery
	errors   cog.BuildErrors
}

func NewMetricNamespaceQueryBuilder() *MetricNamespaceQueryBuilder {
	resource := NewMetricNamespaceQuery()
	builder := &MetricNamespaceQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "MetricNamespaceQuery"

	return builder
}

func (builder *MetricNamespaceQueryBuilder) Build() (MetricNamespaceQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricNamespaceQuery{}, err
	}

	if len(builder.errors) > 0 {
		return MetricNamespaceQuery{}, cog.MakeBuildErrors("azuremonitor.metricNamespaceQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricNamespaceQueryBuilder) RecordError(path string, err error) *MetricNamespaceQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
