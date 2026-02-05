// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMovingFunctionSettings] = (*ElasticsearchMovingFunctionSettingsBuilder)(nil)

type ElasticsearchMovingFunctionSettingsBuilder struct {
	internal *ElasticsearchMovingFunctionSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMovingFunctionSettingsBuilder() *ElasticsearchMovingFunctionSettingsBuilder {
	resource := NewElasticsearchMovingFunctionSettings()
	builder := &ElasticsearchMovingFunctionSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMovingFunctionSettingsBuilder) Build() (ElasticsearchMovingFunctionSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMovingFunctionSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMovingFunctionSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMovingFunctionSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMovingFunctionSettingsBuilder) Window(window string) *ElasticsearchMovingFunctionSettingsBuilder {
	builder.internal.Window = &window

	return builder
}

func (builder *ElasticsearchMovingFunctionSettingsBuilder) Script(script InlineScript) *ElasticsearchMovingFunctionSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

func (builder *ElasticsearchMovingFunctionSettingsBuilder) Shift(shift string) *ElasticsearchMovingFunctionSettingsBuilder {
	builder.internal.Shift = &shift

	return builder
}
