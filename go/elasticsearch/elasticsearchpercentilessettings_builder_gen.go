// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchPercentilesSettings] = (*ElasticsearchPercentilesSettingsBuilder)(nil)

type ElasticsearchPercentilesSettingsBuilder struct {
	internal *ElasticsearchPercentilesSettings
	errors   cog.BuildErrors
}

func NewElasticsearchPercentilesSettingsBuilder() *ElasticsearchPercentilesSettingsBuilder {
	resource := NewElasticsearchPercentilesSettings()
	builder := &ElasticsearchPercentilesSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchPercentilesSettingsBuilder) Build() (ElasticsearchPercentilesSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchPercentilesSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchPercentilesSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchPercentilesSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchPercentilesSettingsBuilder) Script(script InlineScript) *ElasticsearchPercentilesSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

func (builder *ElasticsearchPercentilesSettingsBuilder) Missing(missing string) *ElasticsearchPercentilesSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchPercentilesSettingsBuilder) Percents(percents []string) *ElasticsearchPercentilesSettingsBuilder {
	builder.internal.Percents = percents

	return builder
}
