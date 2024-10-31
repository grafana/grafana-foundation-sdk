// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchExtendedStatsSettings] = (*ElasticsearchExtendedStatsSettingsBuilder)(nil)

type ElasticsearchExtendedStatsSettingsBuilder struct {
	internal *ElasticsearchExtendedStatsSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchExtendedStatsSettingsBuilder() *ElasticsearchExtendedStatsSettingsBuilder {
	resource := &ElasticsearchExtendedStatsSettings{}
	builder := &ElasticsearchExtendedStatsSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchExtendedStatsSettingsBuilder) Build() (ElasticsearchExtendedStatsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchExtendedStatsSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchExtendedStatsSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchExtendedStatsSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

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

func (builder *ElasticsearchExtendedStatsSettingsBuilder) applyDefaults() {
}
