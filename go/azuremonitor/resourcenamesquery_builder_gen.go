// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResourceNamesQuery] = (*ResourceNamesQueryBuilder)(nil)

type ResourceNamesQueryBuilder struct {
	internal *ResourceNamesQuery
	errors   map[string]cog.BuildErrors
}

func NewResourceNamesQueryBuilder() *ResourceNamesQueryBuilder {
	resource := &ResourceNamesQuery{}
	builder := &ResourceNamesQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "ResourceNamesQuery"

	return builder
}

func (builder *ResourceNamesQueryBuilder) Build() (ResourceNamesQuery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ResourceNamesQuery", err)...)
	}

	if len(errs) != 0 {
		return ResourceNamesQuery{}, errs
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

func (builder *ResourceNamesQueryBuilder) applyDefaults() {
}
