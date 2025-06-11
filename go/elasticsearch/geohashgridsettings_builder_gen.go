// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GeoHashGridSettings] = (*GeoHashGridSettingsBuilder)(nil)

type GeoHashGridSettingsBuilder struct {
	internal *GeoHashGridSettings
	errors   cog.BuildErrors
}

func NewGeoHashGridSettingsBuilder() *GeoHashGridSettingsBuilder {
	resource := NewGeoHashGridSettings()
	builder := &GeoHashGridSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *GeoHashGridSettingsBuilder) Build() (GeoHashGridSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return GeoHashGridSettings{}, err
	}

	if len(builder.errors) > 0 {
		return GeoHashGridSettings{}, cog.MakeBuildErrors("elasticsearch.geoHashGridSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *GeoHashGridSettingsBuilder) Precision(precision string) *GeoHashGridSettingsBuilder {
	builder.internal.Precision = &precision

	return builder
}
