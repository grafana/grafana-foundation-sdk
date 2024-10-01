// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchGeoHashGridSettings] = (*ElasticsearchGeoHashGridSettingsBuilder)(nil)

type ElasticsearchGeoHashGridSettingsBuilder struct {
	internal *ElasticsearchGeoHashGridSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchGeoHashGridSettingsBuilder() *ElasticsearchGeoHashGridSettingsBuilder {
	resource := &ElasticsearchGeoHashGridSettings{}
	builder := &ElasticsearchGeoHashGridSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchGeoHashGridSettingsBuilder) Build() (ElasticsearchGeoHashGridSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchGeoHashGridSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchGeoHashGridSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchGeoHashGridSettingsBuilder) Precision(precision string) *ElasticsearchGeoHashGridSettingsBuilder {
	builder.internal.Precision = &precision

	return builder
}

func (builder *ElasticsearchGeoHashGridSettingsBuilder) applyDefaults() {
}
