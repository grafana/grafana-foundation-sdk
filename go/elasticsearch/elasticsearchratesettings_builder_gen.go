// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchRateSettings] = (*ElasticsearchRateSettingsBuilder)(nil)

type ElasticsearchRateSettingsBuilder struct {
	internal *ElasticsearchRateSettings
	errors   cog.BuildErrors
}

func NewElasticsearchRateSettingsBuilder() *ElasticsearchRateSettingsBuilder {
	resource := NewElasticsearchRateSettings()
	builder := &ElasticsearchRateSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchRateSettingsBuilder) Build() (ElasticsearchRateSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchRateSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchRateSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchRateSettings", builder.errors)
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
