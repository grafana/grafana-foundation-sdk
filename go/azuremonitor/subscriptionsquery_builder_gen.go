// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SubscriptionsQuery] = (*SubscriptionsQueryBuilder)(nil)

type SubscriptionsQueryBuilder struct {
	internal *SubscriptionsQuery
	errors   cog.BuildErrors
}

func NewSubscriptionsQueryBuilder() *SubscriptionsQueryBuilder {
	resource := NewSubscriptionsQuery()
	builder := &SubscriptionsQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "SubscriptionsQuery"

	return builder
}

func (builder *SubscriptionsQueryBuilder) Build() (SubscriptionsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return SubscriptionsQuery{}, err
	}

	if len(builder.errors) > 0 {
		return SubscriptionsQuery{}, cog.MakeBuildErrors("azuremonitor.subscriptionsQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SubscriptionsQueryBuilder) RecordError(path string, err error) *SubscriptionsQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SubscriptionsQueryBuilder) RawQuery(rawQuery string) *SubscriptionsQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}
