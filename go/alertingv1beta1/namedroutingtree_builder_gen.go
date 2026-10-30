// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NamedRoutingTree] = (*NamedRoutingTreeBuilder)(nil)

type NamedRoutingTreeBuilder struct {
	internal *NamedRoutingTree
	errors   cog.BuildErrors
}

func NewNamedRoutingTreeBuilder() *NamedRoutingTreeBuilder {
	resource := NewNamedRoutingTree()
	builder := &NamedRoutingTreeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NamedRoutingTreeBuilder) Build() (NamedRoutingTree, error) {
	if err := builder.internal.Validate(); err != nil {
		return NamedRoutingTree{}, err
	}

	if len(builder.errors) > 0 {
		return NamedRoutingTree{}, cog.MakeBuildErrors("alertingv1beta1.namedRoutingTree", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NamedRoutingTreeBuilder) RecordError(path string, err error) *NamedRoutingTreeBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *NamedRoutingTreeBuilder) RoutingTree(routingTree string) *NamedRoutingTreeBuilder {
	builder.internal.RoutingTree = routingTree

	return builder
}
