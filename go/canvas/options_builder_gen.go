// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

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
		return Options{}, cog.MakeBuildErrors("canvas.options", builder.errors)
	}

	return *builder.internal, nil
}

// Enable inline editing
func (builder *OptionsBuilder) InlineEditing(inlineEditing bool) *OptionsBuilder {
	builder.internal.InlineEditing = inlineEditing

	return builder
}

// Show all available element types
func (builder *OptionsBuilder) ShowAdvancedTypes(showAdvancedTypes bool) *OptionsBuilder {
	builder.internal.ShowAdvancedTypes = showAdvancedTypes

	return builder
}

// Enable pan and zoom
func (builder *OptionsBuilder) PanZoom(panZoom bool) *OptionsBuilder {
	builder.internal.PanZoom = panZoom

	return builder
}

// Enable infinite pan
func (builder *OptionsBuilder) InfinitePan(infinitePan bool) *OptionsBuilder {
	builder.internal.InfinitePan = infinitePan

	return builder
}

// The root element of canvas (frame), where all canvas elements are nested
// TODO: Figure out how to define a default value for this
func (builder *OptionsBuilder) Root(root cog.Builder[CanvasOptionsRoot]) *OptionsBuilder {
	rootResource, err := root.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Root = rootResource

	return builder
}
