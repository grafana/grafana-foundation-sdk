// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NodesQuery] = (*NodesQueryBuilder)(nil)

type NodesQueryBuilder struct {
	internal *NodesQuery
	errors   cog.BuildErrors
}

func NewNodesQueryBuilder() *NodesQueryBuilder {
	resource := NewNodesQuery()
	builder := &NodesQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NodesQueryBuilder) Build() (NodesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return NodesQuery{}, err
	}

	if len(builder.errors) > 0 {
		return NodesQuery{}, cog.MakeBuildErrors("testdata.nodesQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NodesQueryBuilder) Type(typeArg NodesQueryType) *NodesQueryBuilder {
	builder.internal.Type = &typeArg

	return builder
}

func (builder *NodesQueryBuilder) Count(count int64) *NodesQueryBuilder {
	builder.internal.Count = &count

	return builder
}
