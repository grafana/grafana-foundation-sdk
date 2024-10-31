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
	resource := &ElasticsearchMovingFunctionSettings{}
	builder := &ElasticsearchMovingFunctionSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

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

func (builder *ElasticsearchMovingFunctionSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchMovingFunctionSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

	return builder
}

func (builder *ElasticsearchMovingFunctionSettingsBuilder) Shift(shift string) *ElasticsearchMovingFunctionSettingsBuilder {
	builder.internal.Shift = &shift

	return builder
}

func (builder *ElasticsearchMovingFunctionSettingsBuilder) applyDefaults() {
}
