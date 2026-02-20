// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchAverageSettings] = (*ElasticsearchAverageSettingsBuilder)(nil)

type ElasticsearchAverageSettingsBuilder struct {
	internal *ElasticsearchAverageSettings
	errors   cog.BuildErrors
}

func NewElasticsearchAverageSettingsBuilder() *ElasticsearchAverageSettingsBuilder {
	resource := NewElasticsearchAverageSettings()
	builder := &ElasticsearchAverageSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchAverageSettingsBuilder) Build() (ElasticsearchAverageSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchAverageSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchAverageSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchAverageSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchAverageSettingsBuilder) RecordError(path string, err error) *ElasticsearchAverageSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchAverageSettingsBuilder) Script(script InlineScript) *ElasticsearchAverageSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

func (builder *ElasticsearchAverageSettingsBuilder) Missing(missing string) *ElasticsearchAverageSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}
