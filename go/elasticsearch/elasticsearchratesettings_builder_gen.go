// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchRateSettings] = (*ElasticsearchRateSettingsBuilder)(nil)

type ElasticsearchRateSettingsBuilder struct {
	internal *ElasticsearchRateSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchRateSettingsBuilder() *ElasticsearchRateSettingsBuilder {
	resource := &ElasticsearchRateSettings{}
	builder := &ElasticsearchRateSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchRateSettingsBuilder) Build() (ElasticsearchRateSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchRateSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchRateSettingsBuilder) Unit(unit string) *ElasticsearchRateSettingsBuilder {
	builder.internal.Unit = &unit

	return builder
}

func (builder *ElasticsearchRateSettingsBuilder) Mode(mode string) *ElasticsearchRateSettingsBuilder {
	builder.internal.Mode = &mode

	return builder
}

func (builder *ElasticsearchRateSettingsBuilder) applyDefaults() {
}
