// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchDerivativeSettings] = (*ElasticsearchDerivativeSettingsBuilder)(nil)

type ElasticsearchDerivativeSettingsBuilder struct {
	internal *ElasticsearchDerivativeSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchDerivativeSettingsBuilder() *ElasticsearchDerivativeSettingsBuilder {
	resource := &ElasticsearchDerivativeSettings{}
	builder := &ElasticsearchDerivativeSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchDerivativeSettingsBuilder) Build() (ElasticsearchDerivativeSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchDerivativeSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchDerivativeSettingsBuilder) Unit(unit string) *ElasticsearchDerivativeSettingsBuilder {
	builder.internal.Unit = &unit

	return builder
}

func (builder *ElasticsearchDerivativeSettingsBuilder) applyDefaults() {
}
