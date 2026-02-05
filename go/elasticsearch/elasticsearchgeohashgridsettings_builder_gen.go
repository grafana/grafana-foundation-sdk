// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchGeoHashGridSettings] = (*ElasticsearchGeoHashGridSettingsBuilder)(nil)

type ElasticsearchGeoHashGridSettingsBuilder struct {
	internal *ElasticsearchGeoHashGridSettings
	errors   cog.BuildErrors
}

func NewElasticsearchGeoHashGridSettingsBuilder() *ElasticsearchGeoHashGridSettingsBuilder {
	resource := NewElasticsearchGeoHashGridSettings()
	builder := &ElasticsearchGeoHashGridSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchGeoHashGridSettingsBuilder) Build() (ElasticsearchGeoHashGridSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchGeoHashGridSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchGeoHashGridSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchGeoHashGridSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchGeoHashGridSettingsBuilder) Precision(precision string) *ElasticsearchGeoHashGridSettingsBuilder {
	builder.internal.Precision = &precision

	return builder
}
