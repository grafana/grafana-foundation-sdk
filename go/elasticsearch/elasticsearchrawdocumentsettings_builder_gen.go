// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchRawDocumentSettings] = (*ElasticsearchRawDocumentSettingsBuilder)(nil)

type ElasticsearchRawDocumentSettingsBuilder struct {
	internal *ElasticsearchRawDocumentSettings
	errors   cog.BuildErrors
}

func NewElasticsearchRawDocumentSettingsBuilder() *ElasticsearchRawDocumentSettingsBuilder {
	resource := NewElasticsearchRawDocumentSettings()
	builder := &ElasticsearchRawDocumentSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchRawDocumentSettingsBuilder) Build() (ElasticsearchRawDocumentSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchRawDocumentSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchRawDocumentSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchRawDocumentSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchRawDocumentSettingsBuilder) RecordError(path string, err error) *ElasticsearchRawDocumentSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchRawDocumentSettingsBuilder) Size(size string) *ElasticsearchRawDocumentSettingsBuilder {
	builder.internal.Size = &size

	return builder
}
