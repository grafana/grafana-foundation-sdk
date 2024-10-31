// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MapViewConfig] = (*MapViewConfigBuilder)(nil)

type MapViewConfigBuilder struct {
	internal *MapViewConfig
	errors   map[string]cog.BuildErrors
}

func NewMapViewConfigBuilder() *MapViewConfigBuilder {
	resource := &MapViewConfig{}
	builder := &MapViewConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MapViewConfigBuilder) Build() (MapViewConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return MapViewConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *MapViewConfigBuilder) Id(id string) *MapViewConfigBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MapViewConfigBuilder) Lat(lat int64) *MapViewConfigBuilder {
	builder.internal.Lat = &lat

	return builder
}

func (builder *MapViewConfigBuilder) Lon(lon int64) *MapViewConfigBuilder {
	builder.internal.Lon = &lon

	return builder
}

func (builder *MapViewConfigBuilder) Zoom(zoom int64) *MapViewConfigBuilder {
	builder.internal.Zoom = &zoom

	return builder
}

func (builder *MapViewConfigBuilder) MinZoom(minZoom int64) *MapViewConfigBuilder {
	builder.internal.MinZoom = &minZoom

	return builder
}

func (builder *MapViewConfigBuilder) MaxZoom(maxZoom int64) *MapViewConfigBuilder {
	builder.internal.MaxZoom = &maxZoom

	return builder
}

func (builder *MapViewConfigBuilder) Padding(padding int64) *MapViewConfigBuilder {
	builder.internal.Padding = &padding

	return builder
}

func (builder *MapViewConfigBuilder) AllLayers(allLayers bool) *MapViewConfigBuilder {
	builder.internal.AllLayers = &allLayers

	return builder
}

func (builder *MapViewConfigBuilder) LastOnly(lastOnly bool) *MapViewConfigBuilder {
	builder.internal.LastOnly = &lastOnly

	return builder
}

func (builder *MapViewConfigBuilder) Layer(layer string) *MapViewConfigBuilder {
	builder.internal.Layer = &layer

	return builder
}

func (builder *MapViewConfigBuilder) Shared(shared bool) *MapViewConfigBuilder {
	builder.internal.Shared = &shared

	return builder
}

func (builder *MapViewConfigBuilder) applyDefaults() {
	builder.Id("zero")
	builder.Lat(0)
	builder.Lon(0)
	builder.Zoom(1)
	builder.AllLayers(true)
}
