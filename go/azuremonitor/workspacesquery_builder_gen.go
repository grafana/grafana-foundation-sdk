// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[WorkspacesQuery] = (*WorkspacesQueryBuilder)(nil)

type WorkspacesQueryBuilder struct {
	internal *WorkspacesQuery
	errors   map[string]cog.BuildErrors
}

func NewWorkspacesQueryBuilder() *WorkspacesQueryBuilder {
	resource := &WorkspacesQuery{}
	builder := &WorkspacesQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "WorkspacesQuery"

	return builder
}

func (builder *WorkspacesQueryBuilder) Build() (WorkspacesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return WorkspacesQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *WorkspacesQueryBuilder) RawQuery(rawQuery string) *WorkspacesQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *WorkspacesQueryBuilder) Subscription(subscription string) *WorkspacesQueryBuilder {
	builder.internal.Subscription = subscription

	return builder
}

func (builder *WorkspacesQueryBuilder) applyDefaults() {
}
