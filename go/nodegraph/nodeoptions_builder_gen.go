// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NodeOptions] = (*NodeOptionsBuilder)(nil)

type NodeOptionsBuilder struct {
	internal *NodeOptions
	errors   map[string]cog.BuildErrors
}

func NewNodeOptionsBuilder() *NodeOptionsBuilder {
	resource := &NodeOptions{}
	builder := &NodeOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *NodeOptionsBuilder) Build() (NodeOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return NodeOptions{}, err
	}

	return *builder.internal, nil
}

// Unit for the main stat to override what ever is set in the data frame.
func (builder *NodeOptionsBuilder) MainStatUnit(mainStatUnit string) *NodeOptionsBuilder {
	builder.internal.MainStatUnit = &mainStatUnit

	return builder
}

// Unit for the secondary stat to override what ever is set in the data frame.
func (builder *NodeOptionsBuilder) SecondaryStatUnit(secondaryStatUnit string) *NodeOptionsBuilder {
	builder.internal.SecondaryStatUnit = &secondaryStatUnit

	return builder
}

// Define which fields are shown as part of the node arc (colored circle around the node).
func (builder *NodeOptionsBuilder) Arcs(arcs []cog.Builder[ArcOption]) *NodeOptionsBuilder {
	arcsResources := make([]ArcOption, 0, len(arcs))
	for _, r1 := range arcs {
		arcsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["arcs"] = err.(cog.BuildErrors)
			return builder
		}
		arcsResources = append(arcsResources, arcsDepth1)
	}
	builder.internal.Arcs = arcsResources

	return builder
}

func (builder *NodeOptionsBuilder) applyDefaults() {
}
