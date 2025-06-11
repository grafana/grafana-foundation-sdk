// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchLogsSettings] = (*ElasticsearchLogsSettingsBuilder)(nil)

type ElasticsearchLogsSettingsBuilder struct {
	internal *ElasticsearchLogsSettings
	errors   cog.BuildErrors
}

func NewElasticsearchLogsSettingsBuilder() *ElasticsearchLogsSettingsBuilder {
	resource := NewElasticsearchLogsSettings()
	builder := &ElasticsearchLogsSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchLogsSettingsBuilder) Build() (ElasticsearchLogsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchLogsSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchLogsSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchLogsSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchLogsSettingsBuilder) Limit(limit string) *ElasticsearchLogsSettingsBuilder {
	builder.internal.Limit = &limit

	return builder
}
