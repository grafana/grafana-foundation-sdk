// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResourceGroupsQuery] = (*ResourceGroupsQueryBuilder)(nil)

type ResourceGroupsQueryBuilder struct {
	internal *ResourceGroupsQuery
	errors   map[string]cog.BuildErrors
}

func NewResourceGroupsQueryBuilder() *ResourceGroupsQueryBuilder {
	resource := &ResourceGroupsQuery{}
	builder := &ResourceGroupsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "ResourceGroupsQuery"

	return builder
}

func (builder *ResourceGroupsQueryBuilder) Build() (ResourceGroupsQuery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ResourceGroupsQuery", err)...)
	}

	if len(errs) != 0 {
		return ResourceGroupsQuery{}, errs
	}

	return *builder.internal, nil
}

func (builder *ResourceGroupsQueryBuilder) RawQuery(rawQuery string) *ResourceGroupsQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *ResourceGroupsQueryBuilder) Subscription(subscription string) *ResourceGroupsQueryBuilder {
	builder.internal.Subscription = subscription

	return builder
}

func (builder *ResourceGroupsQueryBuilder) applyDefaults() {
}
