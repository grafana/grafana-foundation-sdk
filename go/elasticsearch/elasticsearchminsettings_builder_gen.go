// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMinSettings] = (*ElasticsearchMinSettingsBuilder)(nil)

type ElasticsearchMinSettingsBuilder struct {
	internal *ElasticsearchMinSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchMinSettingsBuilder() *ElasticsearchMinSettingsBuilder {
	resource := &ElasticsearchMinSettings{}
	builder := &ElasticsearchMinSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchMinSettingsBuilder) Build() (ElasticsearchMinSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMinSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMinSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchMinSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

	return builder
}

func (builder *ElasticsearchMinSettingsBuilder) Missing(missing string) *ElasticsearchMinSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchMinSettingsBuilder) applyDefaults() {
}
