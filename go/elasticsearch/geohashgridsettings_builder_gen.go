// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GeoHashGridSettings] = (*GeoHashGridSettingsBuilder)(nil)

type GeoHashGridSettingsBuilder struct {
	internal *GeoHashGridSettings
	errors   map[string]cog.BuildErrors
}

func NewGeoHashGridSettingsBuilder() *GeoHashGridSettingsBuilder {
	resource := &GeoHashGridSettings{}
	builder := &GeoHashGridSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *GeoHashGridSettingsBuilder) Build() (GeoHashGridSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("GeoHashGridSettings", err)...)
	}

	if len(errs) != 0 {
		return GeoHashGridSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *GeoHashGridSettingsBuilder) Precision(precision string) *GeoHashGridSettingsBuilder {
	builder.internal.Precision = &precision

	return builder
}

func (builder *GeoHashGridSettingsBuilder) applyDefaults() {
}
