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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("GeoHashGrid", err)...)
	}

	if len(errs) != 0 {
		return GeoHashGrid{}, errs
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

func (builder *GeoHashGridBuilder) Settings(settings struct {
	Precision *string `json:"precision,omitempty"`
}) *GeoHashGridBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *GeoHashGridBuilder) applyDefaults() {
}
