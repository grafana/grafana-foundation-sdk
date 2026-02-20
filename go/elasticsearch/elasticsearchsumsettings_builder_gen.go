// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchSumSettings] = (*ElasticsearchSumSettingsBuilder)(nil)

type ElasticsearchSumSettingsBuilder struct {
	internal *ElasticsearchSumSettings
	errors   cog.BuildErrors
}

func NewElasticsearchSumSettingsBuilder() *ElasticsearchSumSettingsBuilder {
	resource := NewElasticsearchSumSettings()
	builder := &ElasticsearchSumSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchSumSettingsBuilder) Build() (ElasticsearchSumSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchSumSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchSumSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchSumSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchSumSettingsBuilder) RecordError(path string, err error) *ElasticsearchSumSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchSumSettingsBuilder) Script(script InlineScript) *ElasticsearchSumSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

func (builder *ElasticsearchSumSettingsBuilder) Missing(missing string) *ElasticsearchSumSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}
