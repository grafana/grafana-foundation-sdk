// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CanvasOptionsRoot] = (*CanvasOptionsRootBuilder)(nil)

type CanvasOptionsRootBuilder struct {
	internal *CanvasOptionsRoot
	errors   map[string]cog.BuildErrors
}

func NewCanvasOptionsRootBuilder() *CanvasOptionsRootBuilder {
	resource := &CanvasOptionsRoot{}
	builder := &CanvasOptionsRootBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "frame"

	return builder
}

func (builder *CanvasOptionsRootBuilder) Build() (CanvasOptionsRoot, error) {
	if err := builder.internal.Validate(); err != nil {
		return CanvasOptionsRoot{}, err
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
			builder.errors["elements"] = err.(cog.BuildErrors)
			return builder
		}
		elementsResources = append(elementsResources, elementsDepth1)
	}
	builder.internal.Elements = elementsResources

	return builder
}

func (builder *CanvasOptionsRootBuilder) applyDefaults() {
}
