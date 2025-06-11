// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchSerialDiffSettings] = (*ElasticsearchSerialDiffSettingsBuilder)(nil)

type ElasticsearchSerialDiffSettingsBuilder struct {
	internal *ElasticsearchSerialDiffSettings
	errors   cog.BuildErrors
}

func NewElasticsearchSerialDiffSettingsBuilder() *ElasticsearchSerialDiffSettingsBuilder {
	resource := NewElasticsearchSerialDiffSettings()
	builder := &ElasticsearchSerialDiffSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchSerialDiffSettingsBuilder) Build() (ElasticsearchSerialDiffSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchSerialDiffSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchSerialDiffSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchSerialDiffSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchSerialDiffSettingsBuilder) Lag(lag string) *ElasticsearchSerialDiffSettingsBuilder {
	builder.internal.Lag = &lag

	return builder
}
