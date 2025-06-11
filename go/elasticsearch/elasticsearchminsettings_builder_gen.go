// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMinSettings] = (*ElasticsearchMinSettingsBuilder)(nil)

type ElasticsearchMinSettingsBuilder struct {
	internal *ElasticsearchMinSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMinSettingsBuilder() *ElasticsearchMinSettingsBuilder {
	resource := NewElasticsearchMinSettings()
	builder := &ElasticsearchMinSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMinSettingsBuilder) Build() (ElasticsearchMinSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMinSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMinSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMinSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMinSettingsBuilder) Script(script InlineScript) *ElasticsearchMinSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

func (builder *ElasticsearchMinSettingsBuilder) Missing(missing string) *ElasticsearchMinSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}
