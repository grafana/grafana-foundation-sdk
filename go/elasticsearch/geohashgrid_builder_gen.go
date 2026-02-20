// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GeoHashGrid] = (*GeoHashGridBuilder)(nil)

type GeoHashGridBuilder struct {
	internal *GeoHashGrid
	errors   cog.BuildErrors
}

func NewGeoHashGridBuilder() *GeoHashGridBuilder {
	resource := NewGeoHashGrid()
	builder := &GeoHashGridBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *GeoHashGridBuilder) Build() (GeoHashGrid, error) {
	if err := builder.internal.Validate(); err != nil {
		return GeoHashGrid{}, err
	}

	if len(builder.errors) > 0 {
		return GeoHashGrid{}, cog.MakeBuildErrors("elasticsearch.geoHashGrid", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *GeoHashGridBuilder) RecordError(path string, err error) *GeoHashGridBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}
