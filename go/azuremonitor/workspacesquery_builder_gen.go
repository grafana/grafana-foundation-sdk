// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[WorkspacesQuery] = (*WorkspacesQueryBuilder)(nil)

type WorkspacesQueryBuilder struct {
	internal *WorkspacesQuery
	errors   cog.BuildErrors
}

func NewWorkspacesQueryBuilder() *WorkspacesQueryBuilder {
	resource := NewWorkspacesQuery()
	builder := &WorkspacesQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "WorkspacesQuery"

	return builder
}

func (builder *WorkspacesQueryBuilder) Build() (WorkspacesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return WorkspacesQuery{}, err
	}

	if len(builder.errors) > 0 {
		return WorkspacesQuery{}, cog.MakeBuildErrors("azuremonitor.workspacesQuery", builder.errors)
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
