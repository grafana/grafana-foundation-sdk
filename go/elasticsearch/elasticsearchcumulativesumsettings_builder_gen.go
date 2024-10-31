// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchCumulativeSumSettings] = (*ElasticsearchCumulativeSumSettingsBuilder)(nil)

type ElasticsearchCumulativeSumSettingsBuilder struct {
	internal *ElasticsearchCumulativeSumSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchCumulativeSumSettingsBuilder() *ElasticsearchCumulativeSumSettingsBuilder {
	resource := &ElasticsearchCumulativeSumSettings{}
	builder := &ElasticsearchCumulativeSumSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchCumulativeSumSettingsBuilder) Build() (ElasticsearchCumulativeSumSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchCumulativeSumSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchCumulativeSumSettingsBuilder) Format(format string) *ElasticsearchCumulativeSumSettingsBuilder {
	builder.internal.Format = &format

	return builder
}

func (builder *ElasticsearchCumulativeSumSettingsBuilder) applyDefaults() {
}
