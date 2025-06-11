// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ControlsOptions] = (*ControlsOptionsBuilder)(nil)

type ControlsOptionsBuilder struct {
	internal *ControlsOptions
	errors   cog.BuildErrors
}

func NewControlsOptionsBuilder() *ControlsOptionsBuilder {
	resource := NewControlsOptions()
	builder := &ControlsOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ControlsOptionsBuilder) Build() (ControlsOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ControlsOptions{}, err
	}

	if len(builder.errors) > 0 {
		return ControlsOptions{}, cog.MakeBuildErrors("geomap.controlsOptions", builder.errors)
	}

	return *builder.internal, nil
}

// Zoom (upper left)
func (builder *ControlsOptionsBuilder) ShowZoom(showZoom bool) *ControlsOptionsBuilder {
	builder.internal.ShowZoom = &showZoom

	return builder
}

// let the mouse wheel zoom
func (builder *ControlsOptionsBuilder) MouseWheelZoom(mouseWheelZoom bool) *ControlsOptionsBuilder {
	builder.internal.MouseWheelZoom = &mouseWheelZoom

	return builder
}

// Lower right
func (builder *ControlsOptionsBuilder) ShowAttribution(showAttribution bool) *ControlsOptionsBuilder {
	builder.internal.ShowAttribution = &showAttribution

	return builder
}

// Scale options
func (builder *ControlsOptionsBuilder) ShowScale(showScale bool) *ControlsOptionsBuilder {
	builder.internal.ShowScale = &showScale

	return builder
}

// Show debug
func (builder *ControlsOptionsBuilder) ShowDebug(showDebug bool) *ControlsOptionsBuilder {
	builder.internal.ShowDebug = &showDebug

	return builder
}

// Show measure
func (builder *ControlsOptionsBuilder) ShowMeasure(showMeasure bool) *ControlsOptionsBuilder {
	builder.internal.ShowMeasure = &showMeasure

	return builder
}
