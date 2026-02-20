// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CanvasElementOptions] = (*CanvasElementOptionsBuilder)(nil)

type CanvasElementOptionsBuilder struct {
	internal *CanvasElementOptions
	errors   cog.BuildErrors
}

func NewCanvasElementOptionsBuilder() *CanvasElementOptionsBuilder {
	resource := NewCanvasElementOptions()
	builder := &CanvasElementOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CanvasElementOptionsBuilder) Build() (CanvasElementOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return CanvasElementOptions{}, err
	}

	if len(builder.errors) > 0 {
		return CanvasElementOptions{}, cog.MakeBuildErrors("canvas.canvasElementOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CanvasElementOptionsBuilder) RecordError(path string, err error) *CanvasElementOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *CanvasElementOptionsBuilder) Name(name string) *CanvasElementOptionsBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *CanvasElementOptionsBuilder) Type(typeArg string) *CanvasElementOptionsBuilder {
	builder.internal.Type = typeArg

	return builder
}

// TODO: figure out how to define this (element config(s))
func (builder *CanvasElementOptionsBuilder) Config(config any) *CanvasElementOptionsBuilder {
	builder.internal.Config = &config

	return builder
}

func (builder *CanvasElementOptionsBuilder) Constraint(constraint cog.Builder[Constraint]) *CanvasElementOptionsBuilder {
	constraintResource, err := constraint.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Constraint = &constraintResource

	return builder
}

func (builder *CanvasElementOptionsBuilder) Placement(placement cog.Builder[Placement]) *CanvasElementOptionsBuilder {
	placementResource, err := placement.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Placement = &placementResource

	return builder
}

func (builder *CanvasElementOptionsBuilder) Background(background cog.Builder[BackgroundConfig]) *CanvasElementOptionsBuilder {
	backgroundResource, err := background.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Background = &backgroundResource

	return builder
}

func (builder *CanvasElementOptionsBuilder) Border(border cog.Builder[LineConfig]) *CanvasElementOptionsBuilder {
	borderResource, err := border.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Border = &borderResource

	return builder
}

func (builder *CanvasElementOptionsBuilder) Connections(connections []cog.Builder[CanvasConnection]) *CanvasElementOptionsBuilder {
	connectionsResources := make([]CanvasConnection, 0, len(connections))
	for _, r1 := range connections {
		connectionsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		connectionsResources = append(connectionsResources, connectionsDepth1)
	}
	builder.internal.Connections = connectionsResources

	return builder
}
