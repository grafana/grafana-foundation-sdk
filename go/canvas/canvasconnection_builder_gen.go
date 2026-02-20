// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[CanvasConnection] = (*CanvasConnectionBuilder)(nil)

type CanvasConnectionBuilder struct {
	internal *CanvasConnection
	errors   cog.BuildErrors
}

func NewCanvasConnectionBuilder() *CanvasConnectionBuilder {
	resource := NewCanvasConnection()
	builder := &CanvasConnectionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CanvasConnectionBuilder) Build() (CanvasConnection, error) {
	if err := builder.internal.Validate(); err != nil {
		return CanvasConnection{}, err
	}

	if len(builder.errors) > 0 {
		return CanvasConnection{}, cog.MakeBuildErrors("canvas.canvasConnection", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CanvasConnectionBuilder) RecordError(path string, err error) *CanvasConnectionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *CanvasConnectionBuilder) Source(source cog.Builder[ConnectionCoordinates]) *CanvasConnectionBuilder {
	sourceResource, err := source.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Source = sourceResource

	return builder
}

func (builder *CanvasConnectionBuilder) Target(target cog.Builder[ConnectionCoordinates]) *CanvasConnectionBuilder {
	targetResource, err := target.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Target = targetResource

	return builder
}

func (builder *CanvasConnectionBuilder) TargetName(targetName string) *CanvasConnectionBuilder {
	builder.internal.TargetName = &targetName

	return builder
}

func (builder *CanvasConnectionBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *CanvasConnectionBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *CanvasConnectionBuilder) Size(size cog.Builder[common.ScaleDimensionConfig]) *CanvasConnectionBuilder {
	sizeResource, err := size.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Size = &sizeResource

	return builder
}

func (builder *CanvasConnectionBuilder) Vertices(vertices []cog.Builder[ConnectionCoordinates]) *CanvasConnectionBuilder {
	verticesResources := make([]ConnectionCoordinates, 0, len(vertices))
	for _, r1 := range vertices {
		verticesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		verticesResources = append(verticesResources, verticesDepth1)
	}
	builder.internal.Vertices = verticesResources

	return builder
}

func (builder *CanvasConnectionBuilder) SourceOriginal(sourceOriginal cog.Builder[ConnectionCoordinates]) *CanvasConnectionBuilder {
	sourceOriginalResource, err := sourceOriginal.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.SourceOriginal = &sourceOriginalResource

	return builder
}

func (builder *CanvasConnectionBuilder) TargetOriginal(targetOriginal cog.Builder[ConnectionCoordinates]) *CanvasConnectionBuilder {
	targetOriginalResource, err := targetOriginal.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TargetOriginal = &targetOriginalResource

	return builder
}
