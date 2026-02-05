// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CanvasOptionsRoot] = (*CanvasOptionsRootBuilder)(nil)

type CanvasOptionsRootBuilder struct {
	internal *CanvasOptionsRoot
	errors   cog.BuildErrors
}

func NewCanvasOptionsRootBuilder() *CanvasOptionsRootBuilder {
	resource := NewCanvasOptionsRoot()
	builder := &CanvasOptionsRootBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Type = "frame"

	return builder
}

func (builder *CanvasOptionsRootBuilder) Build() (CanvasOptionsRoot, error) {
	if err := builder.internal.Validate(); err != nil {
		return CanvasOptionsRoot{}, err
	}

	if len(builder.errors) > 0 {
		return CanvasOptionsRoot{}, cog.MakeBuildErrors("canvas.canvasOptionsRoot", builder.errors)
	}

	return *builder.internal, nil
}

// Name of the root element
func (builder *CanvasOptionsRootBuilder) Name(name string) *CanvasOptionsRootBuilder {
	builder.internal.Name = name

	return builder
}

// The list of canvas elements attached to the root element
func (builder *CanvasOptionsRootBuilder) Elements(elements []cog.Builder[CanvasElementOptions]) *CanvasOptionsRootBuilder {
	elementsResources := make([]CanvasElementOptions, 0, len(elements))
	for _, r1 := range elements {
		elementsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		elementsResources = append(elementsResources, elementsDepth1)
	}
	builder.internal.Elements = elementsResources

	return builder
}
