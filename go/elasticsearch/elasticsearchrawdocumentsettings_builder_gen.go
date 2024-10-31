// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchRawDocumentSettings] = (*ElasticsearchRawDocumentSettingsBuilder)(nil)

type ElasticsearchRawDocumentSettingsBuilder struct {
	internal *ElasticsearchRawDocumentSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchRawDocumentSettingsBuilder() *ElasticsearchRawDocumentSettingsBuilder {
	resource := &ElasticsearchRawDocumentSettings{}
	builder := &ElasticsearchRawDocumentSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchRawDocumentSettingsBuilder) Build() (ElasticsearchRawDocumentSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchRawDocumentSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchRawDocumentSettingsBuilder) Size(size string) *ElasticsearchRawDocumentSettingsBuilder {
	builder.internal.Size = &size

	return builder
}

func (builder *ElasticsearchRawDocumentSettingsBuilder) applyDefaults() {
}
