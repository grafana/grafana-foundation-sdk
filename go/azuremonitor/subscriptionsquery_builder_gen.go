// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SubscriptionsQuery] = (*SubscriptionsQueryBuilder)(nil)

type SubscriptionsQueryBuilder struct {
	internal *SubscriptionsQuery
	errors   map[string]cog.BuildErrors
}

func NewSubscriptionsQueryBuilder() *SubscriptionsQueryBuilder {
	resource := &SubscriptionsQuery{}
	builder := &SubscriptionsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "SubscriptionsQuery"

	return builder
}

func (builder *SubscriptionsQueryBuilder) Build() (SubscriptionsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return SubscriptionsQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *SubscriptionsQueryBuilder) RawQuery(rawQuery string) *SubscriptionsQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *SubscriptionsQueryBuilder) applyDefaults() {
}
