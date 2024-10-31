// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMaxSettings] = (*ElasticsearchMaxSettingsBuilder)(nil)

type ElasticsearchMaxSettingsBuilder struct {
	internal *ElasticsearchMaxSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchMaxSettingsBuilder() *ElasticsearchMaxSettingsBuilder {
	resource := &ElasticsearchMaxSettings{}
	builder := &ElasticsearchMaxSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchMaxSettingsBuilder) Build() (ElasticsearchMaxSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMaxSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMaxSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchMaxSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

	return builder
}

func (builder *ElasticsearchMaxSettingsBuilder) Missing(missing string) *ElasticsearchMaxSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchMaxSettingsBuilder) applyDefaults() {
}
