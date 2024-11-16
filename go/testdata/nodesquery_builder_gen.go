// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NodesQuery] = (*NodesQueryBuilder)(nil)

type NodesQueryBuilder struct {
	internal *NodesQuery
	errors   map[string]cog.BuildErrors
}

func NewNodesQueryBuilder() *NodesQueryBuilder {
	resource := NewNodesQuery()
	builder := &NodesQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *NodesQueryBuilder) Build() (NodesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return NodesQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *NodesQueryBuilder) Count(count int64) *NodesQueryBuilder {
	builder.internal.Count = &count

	return builder
}

func (builder *NodesQueryBuilder) Seed(seed int64) *NodesQueryBuilder {
	builder.internal.Seed = &seed

	return builder
}

// Possible enum values:
//   - `"random"`
//   - `"random edges"`
//   - `"response_medium"`
//   - `"response_small"`
//   - `"feature_showcase"`
func (builder *NodesQueryBuilder) Type(typeArg NodesQueryType) *NodesQueryBuilder {
	builder.internal.Type = &typeArg

	return builder
}
