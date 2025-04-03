// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchPercentilesSettings] = (*ElasticsearchPercentilesSettingsBuilder)(nil)

type ElasticsearchPercentilesSettingsBuilder struct {
	internal *ElasticsearchPercentilesSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchPercentilesSettingsBuilder() *ElasticsearchPercentilesSettingsBuilder {
	resource := NewElasticsearchPercentilesSettings()
	builder := &ElasticsearchPercentilesSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ElasticsearchPercentilesSettingsBuilder) Build() (ElasticsearchPercentilesSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchPercentilesSettings{}, err
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
