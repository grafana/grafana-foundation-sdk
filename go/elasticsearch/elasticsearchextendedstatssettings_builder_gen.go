// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchExtendedStatsSettings] = (*ElasticsearchExtendedStatsSettingsBuilder)(nil)

type ElasticsearchExtendedStatsSettingsBuilder struct {
	internal *ElasticsearchExtendedStatsSettings
	errors   cog.BuildErrors
}

func NewElasticsearchExtendedStatsSettingsBuilder() *ElasticsearchExtendedStatsSettingsBuilder {
	resource := NewElasticsearchExtendedStatsSettings()
	builder := &ElasticsearchExtendedStatsSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchExtendedStatsSettingsBuilder) Build() (ElasticsearchExtendedStatsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchExtendedStatsSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchExtendedStatsSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchExtendedStatsSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchExtendedStatsSettingsBuilder) Script(script InlineScript) *ElasticsearchExtendedStatsSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

func (builder *ElasticsearchExtendedStatsSettingsBuilder) Missing(missing string) *ElasticsearchExtendedStatsSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchExtendedStatsSettingsBuilder) Sigma(sigma string) *ElasticsearchExtendedStatsSettingsBuilder {
	builder.internal.Sigma = &sigma

	return builder
}
