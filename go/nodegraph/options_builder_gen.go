// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Options] = (*OptionsBuilder)(nil)

type OptionsBuilder struct {
	internal *Options
	errors   cog.BuildErrors
}

func NewOptionsBuilder() *OptionsBuilder {
	resource := NewOptions()
	builder := &OptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsBuilder) Build() (Options, error) {
	if err := builder.internal.Validate(); err != nil {
		return Options{}, err
	}

	if len(builder.errors) > 0 {
		return Options{}, cog.MakeBuildErrors("nodegraph.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) Nodes(nodes cog.Builder[NodeOptions]) *OptionsBuilder {
	nodesResource, err := nodes.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Nodes = &nodesResource

	return builder
}

func (builder *OptionsBuilder) Edges(edges cog.Builder[EdgeOptions]) *OptionsBuilder {
	edgesResource, err := edges.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Edges = &edgesResource

	return builder
}
