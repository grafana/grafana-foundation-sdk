// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
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
		return Options{}, cog.MakeBuildErrors("geomap.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) View(view cog.Builder[MapViewConfig]) *OptionsBuilder {
	viewResource, err := view.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.View = viewResource

	return builder
}

func (builder *OptionsBuilder) Controls(controls cog.Builder[ControlsOptions]) *OptionsBuilder {
	controlsResource, err := controls.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Controls = controlsResource

	return builder
}

func (builder *OptionsBuilder) Basemap(basemap cog.Builder[common.MapLayerOptions]) *OptionsBuilder {
	basemapResource, err := basemap.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Basemap = basemapResource

	return builder
}

func (builder *OptionsBuilder) Layers(layers []cog.Builder[common.MapLayerOptions]) *OptionsBuilder {
	layersResources := make([]common.MapLayerOptions, 0, len(layers))
	for _, r1 := range layers {
		layersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		layersResources = append(layersResources, layersDepth1)
	}
	builder.internal.Layers = layersResources

	return builder
}

func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[TooltipOptions]) *OptionsBuilder {
	tooltipResource, err := tooltip.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Tooltip = tooltipResource

	return builder
}
