// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchAverageSettings] = (*ElasticsearchAverageSettingsBuilder)(nil)

type ElasticsearchAverageSettingsBuilder struct {
	internal *ElasticsearchAverageSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchAverageSettingsBuilder() *ElasticsearchAverageSettingsBuilder {
	resource := &ElasticsearchAverageSettings{}
	builder := &ElasticsearchAverageSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchAverageSettingsBuilder) Build() (ElasticsearchAverageSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchAverageSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchAverageSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchAverageSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

	return builder
}

func (builder *ElasticsearchAverageSettingsBuilder) Missing(missing string) *ElasticsearchAverageSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchAverageSettingsBuilder) applyDefaults() {
}
