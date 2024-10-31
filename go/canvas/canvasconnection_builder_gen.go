// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[CanvasConnection] = (*CanvasConnectionBuilder)(nil)

type CanvasConnectionBuilder struct {
	internal *CanvasConnection
	errors   map[string]cog.BuildErrors
}

func NewCanvasConnectionBuilder() *CanvasConnectionBuilder {
	resource := &CanvasConnection{}
	builder := &CanvasConnectionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CanvasConnectionBuilder) Build() (CanvasConnection, error) {
	if err := builder.internal.Validate(); err != nil {
		return CanvasConnection{}, err
	}

	return *builder.internal, nil
}

func (builder *CanvasConnectionBuilder) Source(source cog.Builder[ConnectionCoordinates]) *CanvasConnectionBuilder {
	sourceResource, err := source.Build()
	if err != nil {
		builder.errors["source"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Source = sourceResource

	return builder
}

func (builder *CanvasConnectionBuilder) Target(target cog.Builder[ConnectionCoordinates]) *CanvasConnectionBuilder {
	targetResource, err := target.Build()
	if err != nil {
		builder.errors["target"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Target = targetResource

	return builder
}

func (builder *CanvasConnectionBuilder) TargetName(targetName string) *CanvasConnectionBuilder {
	builder.internal.TargetName = &targetName

	return builder
}

func (builder *CanvasConnectionBuilder) Path(path ConnectionPath) *CanvasConnectionBuilder {
	builder.internal.Path = path

	return builder
}

func (builder *CanvasConnectionBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *CanvasConnectionBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors["color"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *CanvasConnectionBuilder) Size(size cog.Builder[common.ScaleDimensionConfig]) *CanvasConnectionBuilder {
	sizeResource, err := size.Build()
	if err != nil {
		builder.errors["size"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Size = &sizeResource

	return builder
}

func (builder *CanvasConnectionBuilder) applyDefaults() {
}
