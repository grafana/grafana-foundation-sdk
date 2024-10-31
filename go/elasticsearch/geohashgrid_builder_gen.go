// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GeoHashGrid] = (*GeoHashGridBuilder)(nil)

type GeoHashGridBuilder struct {
	internal *GeoHashGrid
	errors   map[string]cog.BuildErrors
}

func NewGeoHashGridBuilder() *GeoHashGridBuilder {
	resource := &GeoHashGrid{}
	builder := &GeoHashGridBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "geohash_grid"

	return builder
}

func (builder *GeoHashGridBuilder) Build() (GeoHashGrid, error) {
	if err := builder.internal.Validate(); err != nil {
		return GeoHashGrid{}, err
	}

	return *builder.internal, nil
}

func (builder *GeoHashGridBuilder) Field(field string) *GeoHashGridBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *GeoHashGridBuilder) Id(id string) *GeoHashGridBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *GeoHashGridBuilder) Settings(settings cog.Builder[ElasticsearchGeoHashGridSettings]) *GeoHashGridBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *GeoHashGridBuilder) applyDefaults() {
}
