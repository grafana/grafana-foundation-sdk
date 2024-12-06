// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FrameGeometrySource] = (*FrameGeometrySourceBuilder)(nil)

type FrameGeometrySourceBuilder struct {
	internal *FrameGeometrySource
	errors   map[string]cog.BuildErrors
}

func NewFrameGeometrySourceBuilder() *FrameGeometrySourceBuilder {
	resource := NewFrameGeometrySource()
	builder := &FrameGeometrySourceBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *FrameGeometrySourceBuilder) Build() (FrameGeometrySource, error) {
	if err := builder.internal.Validate(); err != nil {
		return FrameGeometrySource{}, err
	}

	return *builder.internal, nil
}

func (builder *FrameGeometrySourceBuilder) Mode(mode FrameGeometrySourceMode) *FrameGeometrySourceBuilder {
	builder.internal.Mode = mode

	return builder
}

// Field mappings
func (builder *FrameGeometrySourceBuilder) Geohash(geohash string) *FrameGeometrySourceBuilder {
	builder.internal.Geohash = &geohash

	return builder
}

func (builder *FrameGeometrySourceBuilder) Latitude(latitude string) *FrameGeometrySourceBuilder {
	builder.internal.Latitude = &latitude

	return builder
}

func (builder *FrameGeometrySourceBuilder) Longitude(longitude string) *FrameGeometrySourceBuilder {
	builder.internal.Longitude = &longitude

	return builder
}

func (builder *FrameGeometrySourceBuilder) Wkt(wkt string) *FrameGeometrySourceBuilder {
	builder.internal.Wkt = &wkt

	return builder
}

func (builder *FrameGeometrySourceBuilder) Lookup(lookup string) *FrameGeometrySourceBuilder {
	builder.internal.Lookup = &lookup

	return builder
}

// Path to Gazetteer
func (builder *FrameGeometrySourceBuilder) Gazetteer(gazetteer string) *FrameGeometrySourceBuilder {
	builder.internal.Gazetteer = &gazetteer

	return builder
}
