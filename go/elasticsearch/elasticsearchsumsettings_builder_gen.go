// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchSumSettings] = (*ElasticsearchSumSettingsBuilder)(nil)

type ElasticsearchSumSettingsBuilder struct {
	internal *ElasticsearchSumSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchSumSettingsBuilder() *ElasticsearchSumSettingsBuilder {
	resource := &ElasticsearchSumSettings{}
	builder := &ElasticsearchSumSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchSumSettingsBuilder) Build() (ElasticsearchSumSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchSumSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchSumSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchSumSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

	return builder
}

func (builder *ElasticsearchSumSettingsBuilder) Missing(missing string) *ElasticsearchSumSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchSumSettingsBuilder) applyDefaults() {
}
