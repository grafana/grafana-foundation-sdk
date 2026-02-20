// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchRawDataSettings] = (*ElasticsearchRawDataSettingsBuilder)(nil)

type ElasticsearchRawDataSettingsBuilder struct {
	internal *ElasticsearchRawDataSettings
	errors   cog.BuildErrors
}

func NewElasticsearchRawDataSettingsBuilder() *ElasticsearchRawDataSettingsBuilder {
	resource := NewElasticsearchRawDataSettings()
	builder := &ElasticsearchRawDataSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchRawDataSettingsBuilder) Build() (ElasticsearchRawDataSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchRawDataSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchRawDataSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchRawDataSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchRawDataSettingsBuilder) RecordError(path string, err error) *ElasticsearchRawDataSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchRawDataSettingsBuilder) Size(size string) *ElasticsearchRawDataSettingsBuilder {
	builder.internal.Size = &size

	return builder
}
