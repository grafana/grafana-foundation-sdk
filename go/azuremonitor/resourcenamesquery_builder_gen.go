// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResourceNamesQuery] = (*ResourceNamesQueryBuilder)(nil)

type ResourceNamesQueryBuilder struct {
	internal *ResourceNamesQuery
	errors   cog.BuildErrors
}

func NewResourceNamesQueryBuilder() *ResourceNamesQueryBuilder {
	resource := NewResourceNamesQuery()
	builder := &ResourceNamesQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "ResourceNamesQuery"

	return builder
}

func (builder *ResourceNamesQueryBuilder) Build() (ResourceNamesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return ResourceNamesQuery{}, err
	}

	if len(builder.errors) > 0 {
		return ResourceNamesQuery{}, cog.MakeBuildErrors("azuremonitor.resourceNamesQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ResourceNamesQueryBuilder) RawQuery(rawQuery string) *ResourceNamesQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *ResourceNamesQueryBuilder) Subscription(subscription string) *ResourceNamesQueryBuilder {
	builder.internal.Subscription = subscription

	return builder
}

func (builder *ResourceNamesQueryBuilder) ResourceGroup(resourceGroup string) *ResourceNamesQueryBuilder {
	builder.internal.ResourceGroup = resourceGroup

	return builder
}

func (builder *ResourceNamesQueryBuilder) MetricNamespace(metricNamespace string) *ResourceNamesQueryBuilder {
	builder.internal.MetricNamespace = metricNamespace

	return builder
}
