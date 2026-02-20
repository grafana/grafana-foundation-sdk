// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMaxSettings] = (*ElasticsearchMaxSettingsBuilder)(nil)

type ElasticsearchMaxSettingsBuilder struct {
	internal *ElasticsearchMaxSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMaxSettingsBuilder() *ElasticsearchMaxSettingsBuilder {
	resource := NewElasticsearchMaxSettings()
	builder := &ElasticsearchMaxSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMaxSettingsBuilder) Build() (ElasticsearchMaxSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMaxSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMaxSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMaxSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMaxSettingsBuilder) RecordError(path string, err error) *ElasticsearchMaxSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchMaxSettingsBuilder) Script(script InlineScript) *ElasticsearchMaxSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

func (builder *ElasticsearchMaxSettingsBuilder) Missing(missing string) *ElasticsearchMaxSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}
