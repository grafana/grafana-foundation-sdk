// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMovingFunctionSettings] = (*ElasticsearchMovingFunctionSettingsBuilder)(nil)

type ElasticsearchMovingFunctionSettingsBuilder struct {
	internal *ElasticsearchMovingFunctionSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchMovingFunctionSettingsBuilder() *ElasticsearchMovingFunctionSettingsBuilder {
	resource := NewElasticsearchMovingFunctionSettings()
	builder := &ElasticsearchMovingFunctionSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ElasticsearchMovingFunctionSettingsBuilder) Build() (ElasticsearchMovingFunctionSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMovingFunctionSettings{}, err
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
